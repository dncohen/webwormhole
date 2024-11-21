//go:build !js

package wormhole

import webrtc "github.com/pion/webrtc/v3"

func (c *Wormhole) getStats() webrtc.StatsReport {
	return c.pc.GetStats()
}

func (c *Wormhole) onError(f func(error)) {
	c.d.OnError(f)
}
