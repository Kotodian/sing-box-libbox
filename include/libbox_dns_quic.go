//go:build libbox_minimal && with_dns_quic

package include

import "github.com/sagernet/sing-box/dns"

// Registers DoQ + DoH3. Requires with_quic (already forced on for Hysteria2).
func registerLibboxDNSQUIC(registry *dns.TransportRegistry) {
	registerQUICTransports(registry)
}
