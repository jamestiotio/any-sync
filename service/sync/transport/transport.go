package transport

import (
	"context"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/app"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/app/logger"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/service/account"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/sec"
	libp2ptls "github.com/libp2p/go-libp2p/p2p/security/tls"
	"go.uber.org/zap"
	"net"
)

type HandshakeError error

var log = logger.NewNamed("transport")

const CName = "transport"

func New() Service {
	return &service{}
}

type Service interface {
	TLSListener(lis net.Listener) ContextListener
	TLSConn(ctx context.Context, conn net.Conn) (sec.SecureConn, error)
	app.Component
}

type service struct {
	key crypto.PrivKey
}

func (s *service) Init(ctx context.Context, a *app.App) (err error) {
	acc := a.MustComponent(account.CName).(account.Service)
	rawKey, err := acc.Account().SignKey.Raw()
	if err != nil {
		return err
	}

	// converting into libp2p crypto structure
	s.key, err = crypto.UnmarshalEd25519PrivateKey(rawKey)
	if err != nil {
		return err
	}

	pubKeyRaw, _ := s.key.GetPublic().Raw()
	log.Info("transport keys generated", zap.Binary("pubKey", pubKeyRaw))
	return nil
}

func (s *service) Name() (name string) {
	return CName
}

func (s *service) TLSListener(lis net.Listener) ContextListener {
	return newTLSListener(s.key, lis)
}

func (s *service) TLSConn(ctx context.Context, conn net.Conn) (sec.SecureConn, error) {
	tr, err := libp2ptls.New(s.key)
	if err != nil {
		return nil, err
	}
	return tr.SecureOutbound(ctx, conn, "")
}