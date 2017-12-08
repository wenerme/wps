package ipin

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/mholt/caddy"
)

func init() {
	caddy.RegisterPlugin("ipin", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	c.Next() // ipin
	if c.NextArg() {
		return plugin.Error("ipin", c.ArgErr())
	}

	dnsserver.
		GetConfig(c).
		AddPlugin(func(next plugin.Handler) plugin.Handler {
			return IpInName{}
		})

	return nil
}
