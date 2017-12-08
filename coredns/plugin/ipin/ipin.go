// Package whoami implements a plugin that returns details about the resolving
// querying it.
package ipin

import (
	"github.com/coredns/coredns/request"
	"net"

	"github.com/miekg/dns"
	"golang.org/x/net/context"
	"regexp"
	"strconv"
	"strings"
)

type IpInName struct {
}

var regIpDash = regexp.MustCompile(`^(\d{1,3}-\d{1,3}-\d{1,3}-\d{1,3})(-\d+)?\.`)

func (self IpInName) Name() string { return "ipin" }
func (self IpInName) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}

	a := new(dns.Msg)
	a.SetReply(r)
	a.Compress = true
	a.Authoritative = true

	matches := regIpDash.FindStringSubmatch(state.QName())
	if len(matches) > 1 {
		ip := matches[1]
		ip = strings.Replace(ip, "-", ".", -1)

		var rr dns.RR
		rr = new(dns.A)
		rr.(*dns.A).Hdr = dns.RR_Header{Name: state.QName(), Rrtype: dns.TypeA, Class: state.QClass()}
		rr.(*dns.A).A = net.ParseIP(ip).To4()

		a.Extra = []dns.RR{rr}

		if len(matches[2]) > 0 {
			srv := new(dns.SRV)
			srv.Hdr = dns.RR_Header{Name: "_port." + state.QName(), Rrtype: dns.TypeSRV, Class: state.QClass()}
			if state.QName() == "." {
				srv.Hdr.Name = "_port." + state.QName()
			}
			port, _ := strconv.Atoi(matches[2][1:])
			srv.Port = uint16(port)
			srv.Target = "."

			a.Extra = append(a.Extra, srv)
		}
	} else {
		// return empty
	}

	w.WriteMsg(a)
	return 0, nil
}
