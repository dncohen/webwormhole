//go:build js && wasm

package wormhole

import (
	"log"

	webrtc "github.com/pion/webrtc/v3"
)

func (c *Wormhole) onError(f func(error)) {
	log.Printf("FIXME: wasm build of Wormhole does not support onError")
}

func (c *Wormhole) getStats() (zero webrtc.StatsReport) {
	log.Printf("FIXME: wasm build of Wormhole does not support getStats")
	return zero
}
