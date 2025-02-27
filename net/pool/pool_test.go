package pool

import (
	"context"
	"errors"
	"fmt"
	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/net"
	"github.com/anyproto/any-sync/net/peer"
	"github.com/anyproto/any-sync/net/secureservice/handshake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	net2 "net"
	"storj.io/drpc"
	"testing"
	"time"
)

var ctx = context.Background()

func TestPool_Get(t *testing.T) {
	t.Run("dial error", func(t *testing.T) {
		fx := newFixture(t)
		defer fx.Finish()
		var expErr = errors.New("dial error")
		fx.Dialer.dial = func(ctx context.Context, peerId string) (peer peer.Peer, err error) {
			return nil, expErr
		}
		p, err := fx.Get(ctx, "1")
		assert.Nil(t, p)
		assert.EqualError(t, err, expErr.Error())
	})
	t.Run("dial and cached", func(t *testing.T) {
		fx := newFixture(t)
		defer fx.Finish()
		fx.Dialer.dial = func(ctx context.Context, peerId string) (peer peer.Peer, err error) {
			return newTestPeer("1"), nil
		}
		p, err := fx.Get(ctx, "1")
		assert.NoError(t, err)
		assert.NotNil(t, p)
		fx.Dialer.dial = nil
		p, err = fx.Get(ctx, "1")
		assert.NoError(t, err)
		assert.NotNil(t, p)
	})
	t.Run("retry for closed", func(t *testing.T) {
		fx := newFixture(t)
		defer fx.Finish()
		tp := newTestPeer("1")
		fx.Dialer.dial = func(ctx context.Context, peerId string) (peer peer.Peer, err error) {
			return tp, nil
		}
		p, err := fx.Get(ctx, "1")
		assert.NoError(t, err)
		assert.NotNil(t, p)
		p.Close()
		tp2 := newTestPeer("1")
		fx.Dialer.dial = func(ctx context.Context, peerId string) (peer peer.Peer, err error) {
			return tp2, nil
		}
		p, err = fx.Get(ctx, "1")
		assert.NoError(t, err)
		assert.Equal(t, p, tp2)
	})
}

func TestPool_GetOneOf(t *testing.T) {
	addToCache := func(t *testing.T, fx *fixture, tp *testPeer) {
		fx.Dialer.dial = func(ctx context.Context, peerId string) (peer peer.Peer, err error) {
			return tp, nil
		}
		gp, err := fx.Get(ctx, tp.Id())
		require.NoError(t, err)
		require.Equal(t, gp, tp)
	}

	t.Run("from cache", func(t *testing.T) {
		fx := newFixture(t)
		defer fx.Finish()
		tp1 := newTestPeer("1")
		addToCache(t, fx, tp1)
		p, err := fx.GetOneOf(ctx, []string{"3", "2", "1"})
		require.NoError(t, err)
		assert.Equal(t, tp1, p)
	})
	t.Run("from cache - skip closed", func(t *testing.T) {
		fx := newFixture(t)
		defer fx.Finish()
		tp2 := newTestPeer("2")
		addToCache(t, fx, tp2)
		tp2.Close()
		tp1 := newTestPeer("1")
		addToCache(t, fx, tp1)
		p, err := fx.GetOneOf(ctx, []string{"3", "2", "1"})
		require.NoError(t, err)
		assert.Equal(t, tp1, p)
	})
	t.Run("dial", func(t *testing.T) {
		fx := newFixture(t)
		defer fx.Finish()
		var called bool
		fx.Dialer.dial = func(ctx context.Context, peerId string) (peer peer.Peer, err error) {
			if called {
				return nil, fmt.Errorf("not expected call")
			}
			called = true
			return newTestPeer(peerId), nil
		}
		p, err := fx.GetOneOf(ctx, []string{"3", "2", "1"})
		require.NoError(t, err)
		assert.NotNil(t, p)
	})
	t.Run("unable to connect", func(t *testing.T) {
		fx := newFixture(t)
		defer fx.Finish()
		fx.Dialer.dial = func(ctx context.Context, peerId string) (peer peer.Peer, err error) {
			return nil, fmt.Errorf("persistent error")
		}
		p, err := fx.GetOneOf(ctx, []string{"3", "2", "1"})
		assert.Equal(t, net.ErrUnableToConnect, err)
		assert.Nil(t, p)
	})
	t.Run("handshake error", func(t *testing.T) {
		fx := newFixture(t)
		defer fx.Finish()
		fx.Dialer.dial = func(ctx context.Context, peerId string) (peer peer.Peer, err error) {
			return nil, handshake.ErrIncompatibleVersion
		}
		p, err := fx.GetOneOf(ctx, []string{"3", "2", "1"})
		assert.Equal(t, handshake.ErrIncompatibleVersion, err)
		assert.Nil(t, p)
	})
}

func TestPool_AddPeer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fx := newFixture(t)
		defer fx.Finish()
		require.NoError(t, fx.AddPeer(ctx, newTestPeer("p1")))
	})
	t.Run("two peers", func(t *testing.T) {
		fx := newFixture(t)
		defer fx.Finish()
		p1, p2 := newTestPeer("p1"), newTestPeer("p1")
		require.NoError(t, fx.AddPeer(ctx, p1))
		require.NoError(t, fx.AddPeer(ctx, p2))
		select {
		case <-p1.closed:
		default:
			assert.Truef(t, false, "peer not closed")
		}
	})

}

func newFixture(t *testing.T) *fixture {
	fx := &fixture{
		Service: New(),
		Dialer:  &dialerMock{},
	}
	a := new(app.App)
	a.Register(fx.Service)
	a.Register(fx.Dialer)
	require.NoError(t, a.Start(context.Background()))
	fx.a = a
	fx.t = t
	return fx
}

func (fx *fixture) Finish() {
	require.NoError(fx.t, fx.a.Close(context.Background()))
}

type fixture struct {
	Service
	Dialer *dialerMock
	a      *app.App
	t      *testing.T
}

var _ dialer = (*dialerMock)(nil)

type dialerMock struct {
	dial func(ctx context.Context, peerId string) (peer peer.Peer, err error)
}

func (d *dialerMock) Dial(ctx context.Context, peerId string) (peer peer.Peer, err error) {
	return d.dial(ctx, peerId)
}

func (d *dialerMock) UpdateAddrs(addrs map[string][]string) {
	return
}

func (d *dialerMock) SetPeerAddrs(peerId string, addrs []string) {
	return
}

func (d *dialerMock) Init(a *app.App) (err error) {
	return
}

func (d *dialerMock) Name() (name string) {
	return "net.peerservice"
}

func newTestPeer(id string) *testPeer {
	return &testPeer{
		id:     id,
		closed: make(chan struct{}),
	}
}

type testPeer struct {
	id     string
	closed chan struct{}
}

func (t *testPeer) SetTTL(ttl time.Duration) {
	return
}

func (t *testPeer) DoDrpc(ctx context.Context, do func(conn drpc.Conn) error) error {
	return fmt.Errorf("not implemented")
}

func (t *testPeer) AcquireDrpcConn(ctx context.Context) (drpc.Conn, error) {
	return nil, fmt.Errorf("not implemented")
}

func (t *testPeer) ReleaseDrpcConn(conn drpc.Conn) {}

func (t *testPeer) Context() context.Context {
	//TODO implement me
	panic("implement me")
}

func (t *testPeer) Accept() (conn net2.Conn, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *testPeer) Open(ctx context.Context) (conn net2.Conn, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *testPeer) Addr() string {
	return ""
}

func (t *testPeer) Id() string {
	return t.id
}

func (t *testPeer) TryClose(objectTTL time.Duration) (res bool, err error) {
	return true, t.Close()
}

func (t *testPeer) Close() error {
	select {
	case <-t.closed:
		return fmt.Errorf("already closed")
	default:
		close(t.closed)
	}
	return nil
}

func (t *testPeer) IsClosed() bool {
	select {
	case <-t.closed:
		return true
	default:
		return false
	}
}
