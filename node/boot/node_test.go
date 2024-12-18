package bootnode

import (
	"testing"

	harmonyConfigs "github.com/harmony-one/harmony/cmd/config"
	"github.com/harmony-one/harmony/crypto/bls"
	nodeconfig "github.com/harmony-one/harmony/internal/configs/node"
	"github.com/harmony-one/harmony/internal/utils"
	"github.com/harmony-one/harmony/p2p"
)

func TestNewBootNode(t *testing.T) {
	blsKey := bls.RandPrivateKey()
	pubKey := blsKey.GetPublicKey()
	leader := p2p.Peer{IP: "127.0.0.1", Port: "8882", ConsensusPubKey: pubKey}
	priKey, _, _ := utils.GenKeyP2P("127.0.0.1", "9902")
	host, err := p2p.NewHost(p2p.HostConfig{
		Self:      &leader,
		BLSKey:    priKey,
		BootNodes: nil,
	})
	if err != nil {
		t.Fatalf("new boot node host failure: %v", err)
	}

	hc := harmonyConfigs.GetDefaultHmyConfigCopy(nodeconfig.NetworkType(nodeconfig.Devnet))
	node := New(host, &hc)

	if node == nil {
		t.Error("boot node creation failed")
	}
}
