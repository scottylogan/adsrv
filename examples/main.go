package main

import (
	"fmt"
	"os"

	"code.stanford.edu/SLAC-IAM/adsrv"
)

func main() {
	if len(os.Args) != 2 {
		panic("Usage: adsrv realm")
	}

	domain, err := adsrv.GetDomain(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Printf("cname: %s\n", domain.CName)
	for _, srv := range domain.SRV {
		fmt.Printf("%s %d %d %d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}
}
