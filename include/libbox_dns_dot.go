//go:build libbox_minimal && with_dns_dot

package include

import (
	"github.com/sagernet/sing-box/dns"
	"github.com/sagernet/sing-box/dns/transport"
)

func registerLibboxDoT(registry *dns.TransportRegistry) {
	transport.RegisterTLS(registry)
}
