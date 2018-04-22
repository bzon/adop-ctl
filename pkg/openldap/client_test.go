package openldap

import "os"

var ldapDomain = os.Getenv("LDAP_FULL_DOMAIN")
var openldap = &Client{
	Host:         "localhost",
	Scheme:       "tcp",
	Port:         389,
	bindUser:     "cn=admin," + ldapDomain,
	bindPassword: os.Getenv("SLAPD_PASSWORD"),
}
