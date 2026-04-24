//go:build libbox_minimal

package daemon

import "github.com/sagernet/sing-box/log"

// Under libbox_minimal we intentionally drop the PlatformLogWriter so sing-box
// core skips Clash API creation (box.go:151). Log / clash-mode / connection
// subscription RPCs become no-ops / panicing — caller must not invoke them.
func platformLogWriter(s *StartedService) log.PlatformWriter {
	return nil
}
