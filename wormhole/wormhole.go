//go:build !js

package wormhole

import (
	webrtc "github.com/pion/webrtc/v3"
	"golang.org/x/net/proxy"
)

func (c *Wormhole) getStats() webrtc.StatsReport {
	return c.pc.GetStats()
}

func (c *Wormhole) onError(f func(error)) {
	c.d.OnError(f)
}

func setICEProxyDialer(s *webrtc.SettingEngine, d proxy.Dialer) {
	s.SetICEProxyDialer(d)
}
