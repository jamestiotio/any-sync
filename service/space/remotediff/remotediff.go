package remotediff

import (
	"context"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/pkg/ldiff"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/service/space/spacesync"
)

type Client interface {
	HeadSync(ctx context.Context, in *spacesync.HeadSyncRequest) (*spacesync.HeadSyncResponse, error)
}

func NewRemoteDiff(spaceId string, client Client) ldiff.Remote {
	return remote{
		spaceId: spaceId,
		client:  client,
	}
}

type remote struct {
	spaceId string
	client  Client
}

func (r remote) Ranges(ctx context.Context, ranges []ldiff.Range, resBuf []ldiff.RangeResult) (results []ldiff.RangeResult, err error) {
	results = resBuf[:0]
	pbRanges := make([]*spacesync.HeadSyncRange, 0, len(ranges))
	for _, rg := range ranges {
		pbRanges = append(pbRanges, &spacesync.HeadSyncRange{
			From:  rg.From,
			To:    rg.To,
			Limit: uint32(rg.Limit),
		})
	}
	req := &spacesync.HeadSyncRequest{
		SpaceId: r.spaceId,
		Ranges:  pbRanges,
	}
	resp, err := r.client.HeadSync(ctx, req)
	if err != nil {
		return
	}
	for _, rr := range resp.Results {
		var elms []ldiff.Element
		if len(rr.Elements) > 0 {
			elms = make([]ldiff.Element, 0, len(rr.Elements))
		}
		for _, e := range rr.Elements {
			elms = append(elms, ldiff.Element{
				Id:   e.Id,
				Head: e.Head,
			})
		}
		results = append(results, ldiff.RangeResult{
			Hash:     rr.Hash,
			Elements: elms,
			Count:    int(rr.Count),
		})
	}
	return
}

func HandlerRangeRequest(ctx context.Context, d ldiff.Diff, req *spacesync.HeadSyncRequest) (resp *spacesync.HeadSyncResponse, err error) {
	ranges := make([]ldiff.Range, 0, len(req.Ranges))
	for _, reqRange := range req.Ranges {
		ranges = append(ranges, ldiff.Range{
			From:  reqRange.From,
			To:    reqRange.To,
			Limit: int(reqRange.Limit),
		})
	}
	res, err := d.Ranges(ctx, ranges, nil)
	if err != nil {
		return
	}

	var rangeResp = &spacesync.HeadSyncResponse{
		Results: make([]*spacesync.HeadSyncResult, 0, len(res)),
	}
	for _, rangeRes := range res {
		var elements []*spacesync.HeadSyncResultElement
		if len(rangeRes.Elements) > 0 {
			elements = make([]*spacesync.HeadSyncResultElement, 0, len(rangeRes.Elements))
			for _, el := range rangeRes.Elements {
				elements = append(elements, &spacesync.HeadSyncResultElement{
					Id:   el.Id,
					Head: el.Head,
				})
			}
		}
		rangeResp.Results = append(rangeResp.Results, &spacesync.HeadSyncResult{
			Hash:     rangeRes.Hash,
			Elements: elements,
			Count:    uint32(rangeRes.Count),
		})
	}
	return rangeResp, nil
}
