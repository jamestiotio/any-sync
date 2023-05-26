//go:generate mockgen -destination mock_headsync/mock_headsync.go github.com/anyproto/any-sync/commonspace/headsync DiffSyncer
package headsync

import (
	"context"
	"github.com/anyproto/any-sync/app/ldiff"
	"github.com/anyproto/any-sync/app/logger"
	"github.com/anyproto/any-sync/commonspace/credentialprovider"
	"github.com/anyproto/any-sync/commonspace/object/treemanager"
	"github.com/anyproto/any-sync/commonspace/peermanager"
	"github.com/anyproto/any-sync/commonspace/settings/settingsstate"
	"github.com/anyproto/any-sync/commonspace/spacestorage"
	"github.com/anyproto/any-sync/commonspace/spacesyncproto"
	"github.com/anyproto/any-sync/commonspace/syncstatus"
	"github.com/anyproto/any-sync/net/peer"
	"github.com/anyproto/any-sync/nodeconf"
	"github.com/anyproto/any-sync/util/periodicsync"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"strings"
	"sync/atomic"
	"time"
)

type TreeHeads struct {
	Id    string
	Heads []string
}

type HeadSync interface {
	Init(objectIds []string, deletionState settingsstate.ObjectDeletionState)

	UpdateHeads(id string, heads []string)
	HandleRangeRequest(ctx context.Context, req *spacesyncproto.HeadSyncRequest) (resp *spacesyncproto.HeadSyncResponse, err error)
	RemoveObjects(ids []string)
	AllIds() []string
	DebugAllHeads() (res []TreeHeads)

	Close() (err error)
}

type headSync struct {
	spaceId        string
	periodicSync   periodicsync.PeriodicSync
	storage        spacestorage.SpaceStorage
	diff           ldiff.Diff
	log            logger.CtxLogger
	syncer         DiffSyncer
	configuration  nodeconf.NodeConf
	spaceIsDeleted *atomic.Bool

	syncPeriod int
}

func NewHeadSync(
	spaceId string,
	spaceIsDeleted *atomic.Bool,
	syncPeriod int,
	configuration nodeconf.NodeConf,
	storage spacestorage.SpaceStorage,
	peerManager peermanager.PeerManager,
	cache treemanager.TreeManager,
	syncStatus syncstatus.StatusUpdater,
	credentialProvider credentialprovider.CredentialProvider,
	log logger.CtxLogger) HeadSync {

	diff := ldiff.New(16, 16)
	l := log.With(zap.String("spaceId", spaceId))
	factory := spacesyncproto.ClientFactoryFunc(spacesyncproto.NewDRPCSpaceSyncClient)
	syncer := newDiffSyncer(spaceId, diff, peerManager, cache, storage, factory, syncStatus, credentialProvider, l)
	sync := func(ctx context.Context) (err error) {
		// for clients cancelling the sync process
		if spaceIsDeleted.Load() && !configuration.IsResponsible(spaceId) {
			return spacesyncproto.ErrSpaceIsDeleted
		}
		return syncer.Sync(ctx)
	}
	periodicSync := periodicsync.NewPeriodicSync(syncPeriod, time.Minute, sync, l)

	return &headSync{
		spaceId:        spaceId,
		storage:        storage,
		syncer:         syncer,
		periodicSync:   periodicSync,
		diff:           diff,
		log:            log,
		syncPeriod:     syncPeriod,
		configuration:  configuration,
		spaceIsDeleted: spaceIsDeleted,
	}
}

func (d *headSync) Init(objectIds []string, deletionState settingsstate.ObjectDeletionState) {
	d.fillDiff(objectIds)
	d.syncer.Init(deletionState)
	d.periodicSync.Run()
}

func (d *headSync) HandleRangeRequest(ctx context.Context, req *spacesyncproto.HeadSyncRequest) (resp *spacesyncproto.HeadSyncResponse, err error) {
	if d.spaceIsDeleted.Load() {
		peerId, err := peer.CtxPeerId(ctx)
		if err != nil {
			return nil, err
		}
		// stop receiving all request for sync from clients
		if !slices.Contains(d.configuration.NodeIds(d.spaceId), peerId) {
			return nil, spacesyncproto.ErrSpaceIsDeleted
		}
	}
	return HandleRangeRequest(ctx, d.diff, req)
}

func (d *headSync) UpdateHeads(id string, heads []string) {
	d.syncer.UpdateHeads(id, heads)
}

func (d *headSync) AllIds() []string {
	return d.diff.Ids()
}

func (d *headSync) DebugAllHeads() (res []TreeHeads) {
	els := d.diff.Elements()
	for _, el := range els {
		idHead := TreeHeads{
			Id:    el.Id,
			Heads: splitString(el.Head),
		}
		res = append(res, idHead)
	}
	return
}

func (d *headSync) RemoveObjects(ids []string) {
	d.syncer.RemoveObjects(ids)
}

func (d *headSync) Close() (err error) {
	d.periodicSync.Close()
	return d.syncer.Close()
}

func (d *headSync) fillDiff(objectIds []string) {
	var els = make([]ldiff.Element, 0, len(objectIds))
	for _, id := range objectIds {
		st, err := d.storage.TreeStorage(id)
		if err != nil {
			continue
		}
		heads, err := st.Heads()
		if err != nil {
			continue
		}
		els = append(els, ldiff.Element{
			Id:   id,
			Head: concatStrings(heads),
		})
	}
	d.diff.Set(els...)
	if err := d.storage.WriteSpaceHash(d.diff.Hash()); err != nil {
		d.log.Error("can't write space hash", zap.Error(err))
	}
}

func concatStrings(strs []string) string {
	var (
		b        strings.Builder
		totalLen int
	)
	for _, s := range strs {
		totalLen += len(s)
	}

	b.Grow(totalLen)
	for _, s := range strs {
		b.WriteString(s)
	}
	return b.String()
}

func splitString(str string) (res []string) {
	const cidLen = 59
	for i := 0; i < len(str); i += cidLen {
		res = append(res, str[i:i+cidLen])
	}
	return
}
