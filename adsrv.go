package adsrv

import (
	"fmt"
	"net"
	"net/url"
	"sort"

	"github.com/go-ldap/ldap/v3"
)

const (
	proto   string = "tcp"
	service string = "ldap"
)

type SRVSet []*net.SRV

type Domain struct {
	Name  string
	CName string
	SRV   SRVSet
}

func (s SRVSet) Len() int {
	return len(s)
}

func (s SRVSet) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SRVSet) Less(i, j int) bool {
	if s[i].Priority == s[j].Priority {
		return s[i].Weight > s[j].Weight
	}
	return s[i].Priority < s[j].Priority
}

func new(domain string, cname string, srv SRVSet) (*Domain) {
	sort.Sort(SRVSet(srv))
	return &Domain{domain, cname, srv}
}

func GetDomain(domain string) (*Domain, error) {
	cname, srv, err := net.LookupSRV(service, proto, domain)
	if err != nil {
		return nil, err
	}
	return new(domain, cname, srv), nil
}

func makeSiteDomain(domain, site string) string {
	return site + "._sites." + domain
}

func GetDomainSite(domain, site string) (*Domain, error) {
	cname, srv, err := net.LookupSRV(service, proto, makeSiteDomain(domain, site))
	if err != nil {
		return nil, err
	}
	return new(domain, cname, srv), nil
}

func (d *Domain) Dial(opts ...ldap.DialOpt) (*ldap.Conn, error) {

	ldapUrl := &url.URL{Scheme: "ldap"}

	for _, srv := range d.SRV {
		ldapUrl.Host = fmt.Sprintf("%s:%d", srv.Target, srv.Port)
		conn, err := ldap.DialURL(ldapUrl.String(), opts...)
		if err == nil {
			return conn, nil
		}
	}
	return nil, ldap.NewError(ldap.ErrorNetwork, fmt.Errorf("failed to connect to any LDAP server"))
}
