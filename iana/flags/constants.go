// Package flags provides Kerberos 5 flag assigned numbers.
package flags

const (
	Reserved               = 0
	Forwardable            = 1
	Forwarded              = 2
	Proxiable              = 3
	Proxy                  = 4
	AllowPostDate          = 5
	MayPostDate            = 5
	PostDated              = 6
	Invalid                = 7
	Renewable              = 8
	Initial                = 9
	PreAuthent             = 10
	HWAuthent              = 11
	OptHardwareAuth        = 11
	RequestAnonymous       = 12
	TransitedPolicyChecked = 12
	OKAsDelegate           = 13
	EncPARep               = 15
	Canonicalize           = 15
	DisableTransitedCheck  = 26
	RenewableOK            = 27
	EncTktInSkey           = 28
	Renew                  = 30
	Validate               = 31
)
