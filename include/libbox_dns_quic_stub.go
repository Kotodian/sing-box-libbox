//go:build libbox_minimal && !with_dns_quic

package include

import "github.com/sagernet/sing-box/dns"

func registerLibboxDNSQUIC(registry *dns.TransportRegistry) {}
