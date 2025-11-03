module adsrv-test

go 1.25.2

require code.stanford.edu/SLAC-IAM/adsrv v0.0.0-00010101000000-000000000000

require (
	github.com/Azure/go-ntlmssp v0.0.0-20221128193559-754e69321358 // indirect
	github.com/go-asn1-ber/asn1-ber v1.5.8-0.20250403174932-29230038a667 // indirect
	github.com/go-ldap/ldap/v3 v3.4.12 // indirect
	github.com/google/uuid v1.6.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
)

replace code.stanford.edu/SLAC-IAM/adsrv => ../
