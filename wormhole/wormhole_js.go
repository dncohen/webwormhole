//go:build js && wasm

package wormhole

import (
	"log"

	webrtc "github.com/pion/webrtc/v3"
	"golang.org/x/net/proxy"
)

func (c *Wormhole) onError(f func(error)) {
	log.Printf("FIXME: wasm build of Wormhole does not support onError")
}

func (c *Wormhole) getStats() (zero webrtc.StatsReport) {
	log.Printf("FIXME: wasm build of Wormhole does not support getStats")
	return zero
}

func setICEProxyDialer(s *webrtc.SettingEngine, d proxy.Dialer) {
	log.Printf("FIXME: wasm build does not support ICE proxy dialer")
}
