//go:build !libbox_minimal

package daemon

import "github.com/sagernet/sing-box/log"

// platformLogWriter returns the log sink that forwards sing-box core logs to
// the libbox client. Triggers sing-box to instantiate its Clash API internally
// (see box.go:151) because the Observable log factory is shared with Clash.
// Under libbox_minimal the stub variant returns nil so no Clash server is
// created — caller must not use log / clash-mode / connection subscription
// RPCs in that build.
func platformLogWriter(s *StartedService) log.PlatformWriter {
	return s
}
