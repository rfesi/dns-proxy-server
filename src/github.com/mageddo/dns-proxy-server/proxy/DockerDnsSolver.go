package proxy

import (
	"github.com/miekg/dns"
	"github.com/mageddo/dns-proxy-server/events/docker"
	"github.com/mageddo/log"
	"net"
	"strings"
	"strconv"
	"errors"
)

type DockerDnsSolver struct {

}

func (DockerDnsSolver) Solve(question dns.Question) (*dns.Msg, error) {

	key := question.Name[:len(question.Name)-1]
	if docker.ContainsKey(key) {

		ip := docker.Get(key)
		ipArr := strings.Split(ip, ".")
		i1, _ := strconv.Atoi(ipArr[0])
		i2, _ := strconv.Atoi(ipArr[1])
		i3, _ := strconv.Atoi(ipArr[2])
		i4, _ := strconv.Atoi(ipArr[3])

		rr := &dns.A{
			Hdr: dns.RR_Header{Name: question.Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 0},
			A: net.IPv4(byte(i1), byte(i2), byte(i3), byte(i4)),
		}

		m := new(dns.Msg)
		m.Answer = append(m.Answer, rr)
		log.Logger.Infof("m=solve, status=success, solver=docker")
		return m, nil
	}
	return nil, errors.New("hostname not found")
}