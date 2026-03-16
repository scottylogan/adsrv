package main

import (
	"fmt"
	"os"

	"code.stanford.edu/SLAC-IAM/adsrv"
)

func main() {
	if len(os.Args) < 2 {
		panic("Usage: adsrv realm [site]")
	}

	realm := os.Args[1]

	domain, err := adsrv.GetDomainPDC(realm)
	if err != nil {
		panic(err)
	}

	fmt.Println("PDC")
	for _, srv := range domain.SRV {
		fmt.Printf("  %s %d %d %d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}

	domain, err = adsrv.GetDomain(realm)
	if err != nil {
		panic(err)
	}
	fmt.Println("ALL")
	for _, srv := range domain.SRV {
		fmt.Printf("  %s %d %d %d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}

	for _, site := range os.Args[2:] {
		domain, err = adsrv.GetDomainSite(realm, site)

		if err != nil {
			panic(err)
		}
		fmt.Println("SITE:", site)
		for _, srv := range domain.SRV {
			fmt.Printf("  %s %d %d %d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
		}
	}

}
