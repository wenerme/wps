package ipin

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/caddyserver/caddy"
)

func init() {
	caddy.RegisterPlugin(Name, caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	ipin := IpInName{}

	c.Next() // ipin
	for c.NextBlock() {
		x := c.Val()
		switch x {
		case "fallback":
			ipin.Fallback = true
		default:
			return plugin.Error(Name, c.Errf("unexpected '%v' command", x))
		}
	}
	if c.NextArg() {
		return plugin.Error(Name, c.ArgErr())
	}

	dnsserver.
		GetConfig(c).
		AddPlugin(func(next plugin.Handler) plugin.Handler {
			ipin.Next = next
			return ipin
		})

	return nil
}
