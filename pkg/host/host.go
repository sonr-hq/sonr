package host

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/routing"
	cmgr "github.com/libp2p/go-libp2p/p2p/net/connmgr"

	ps "github.com/libp2p/go-libp2p-pubsub"
)

// A P2PHost is a host.Host with a private key, a channel of mDNS peers, a channel of DHT peers, a
// context, a map of topics, and a DHT and PubSub.
// @property host - The host is the main object of the libp2p library. It represents the local node and
// provides all the functionality to interact with the network.
// @property {string} accAddr - The address of the account that is being used to connect to the
// network.
// @property privKey - The private key of the host.
// @property mdnsPeerChan - This is a channel that will receive peer.AddrInfo objects from the mdns
// service.
// @property dhtPeerChan - A channel that will receive peer.AddrInfo objects when a peer is found via
// the DHT.
// @property ctx - The context of the P2PHost.
// @property topics - A map of topic names to the PubSub topic object.
// @property  - `host`: The libp2p host.
// @property  - `host`: The libp2p host.
type P2PHost struct {
	// Standard Node Implementation
	host    host.Host
	accAddr string

	// Host and context
	privKey      crypto.PrivKey
	mdnsPeerChan chan peer.AddrInfo
	dhtPeerChan  <-chan peer.AddrInfo

	// Properties
	ctx    context.Context
	topics map[string]*ps.Topic

	*dht.IpfsDHT
	*ps.PubSub
}

// New Creates a Sonr libp2p Host with the given config
func New(ctx context.Context) (*P2PHost, error) {
	var err error
	// Create the host.
	hn := &P2PHost{
		ctx:          ctx,
		mdnsPeerChan: make(chan peer.AddrInfo),
		topics:       make(map[string]*ps.Topic),
	}
	// findPrivKey returns the private key for the host.
	findPrivKey := func() (crypto.PrivKey, error) {
		privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
		if err == nil {

			return privKey, nil
		}
		return nil, err
	}
	// Fetch the private key.
	hn.privKey, err = findPrivKey()
	if err != nil {
		return nil, err
	}

	// Create Connection Manager
	cnnmgr, err := cmgr.NewConnManager(10, 40)
	if err != nil {
		return nil, err
	}

	// Start Host
	hn.host, err = libp2p.New(
		libp2p.Identity(hn.privKey),
		libp2p.ConnectionManager(cnnmgr),
		libp2p.DefaultListenAddrs,
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			hn.IpfsDHT, err = dht.New(ctx, h)
			if err != nil {
				return nil, err
			}
			return hn.IpfsDHT, nil
		}),
	)
	if err != nil {
		return nil, err
	}

	// Bootstrap DHT
	if err := hn.Bootstrap(ctx); err != nil {
		return nil, err
	}

	// Connect to Bootstrap Nodes
	for _, pistr := range defaultBootstrapMultiaddrs {
		if err := hn.Connect(pistr); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := hn.createDHTDiscovery(); err != nil {
		return nil, err
	}
	return hn, nil
}

// Host returns the host of the node
func (hn *P2PHost) Host() host.Host {
	return hn.host
}

// PeerID returns the ID of the Host
func (n *P2PHost) PeerID() peer.ID {
	return n.host.ID()
}

// Connect connects with `peer.AddrInfo` if underlying Host is ready
func (hn *P2PHost) Connect(pi interface{}) error {
	// Check if type is String or AddrInfo
	switch pi.(type) {
	case string:
		pi, err := peer.AddrInfoFromString(pi.(string))
		if err != nil {
			return err
		}
		return hn.host.Connect(hn.ctx, *pi)
	case peer.AddrInfo:
		return hn.host.Connect(hn.ctx, pi.(peer.AddrInfo))
	default:
		return fmt.Errorf("Connect: Invalid type for peer.AddrInfo")
	}
}

// HandlePeerFound is to be called when new  peer is found
func (hn *P2PHost) HandlePeerFound(pi peer.AddrInfo) {
	hn.mdnsPeerChan <- pi
}

// MultiAddrs returns the MultiAddresses of the Host
func (hn *P2PHost) MultiAddrs() []ma.Multiaddr {
	return hn.host.Addrs()
}

// NewStream opens a new stream to the peer with given peer id
func (n *P2PHost) NewStream(ctx context.Context, pid peer.ID, pids ...protocol.ID) (network.Stream, error) {
	return n.host.NewStream(ctx, pid, pids...)
}

// JoinTopic creates a new topic
func (n *P2PHost) Publish(topic string, message []byte, opts ...ps.TopicOpt) error {
	ctx, cancel := context.WithTimeout(n.ctx, 10*time.Second)
	defer cancel()
	// Check if PubSub is Set
	if n.PubSub == nil {
		return nil
	}

	// Check if topic is valid
	t, ok := n.topics[topic]
	if ok {
		return t.Publish(ctx, message)
	}

	// Call Underlying Pubsub to Connect
	t, err := n.PubSub.Join(topic, opts...)
	if err != nil {
		return err
	}

	// Create Subscriber
	n.topics[topic] = t
	return t.Publish(ctx, message)
}

// SetStreamHandler sets the handler for a given protocol
func (n *P2PHost) SetStreamHandler(protocol protocol.ID, handler network.StreamHandler) {
	n.host.SetStreamHandler(protocol, handler)
}

// Join wraps around PubSub.Join and returns topic. Checks wether the host is ready before joining.
func (hn *P2PHost) Subscribe(topic string, handlers ...func(msg *ps.Message)) (*ps.Subscription, error) {
	// Check if PubSub is Set
	if hn.PubSub == nil {
		return nil, errors.New("Join: Pubsub has not been set on SNRHost")
	}

	// Check if topic is already joined
	if t, ok := hn.topics[topic]; ok {
		return t.Subscribe()
	}

	// Call Underlying Pubsub to Connect
	t, err := hn.PubSub.Join(topic)
	if err != nil {
		return nil, err
	}
	hn.topics[topic] = t

	// Subscribe to Topic
	sub, err := t.Subscribe()
	if err != nil {
		return nil, err
	}

	// Handle Subscription
	if len(handlers) > 0 {
		go hn.handleSubscription(sub, handlers[0])
	}
	return sub, nil
}

// handleSubscription handles the subscription to a topic
func (hn *P2PHost) handleSubscription(sub *ps.Subscription, handler func(msg *ps.Message)) {
	for {
		msg, err := sub.Next(hn.ctx)
		if err != nil {
			return
		}
		handler(msg)

		select {
		case <-hn.ctx.Done():
			return
		default:
		}
	}
}