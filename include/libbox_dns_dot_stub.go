//go:build libbox_minimal && !with_dns_dot

package include

import "github.com/sagernet/sing-box/dns"

func registerLibboxDoT(registry *dns.TransportRegistry) {}
