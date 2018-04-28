package openldap

var ldapDomain = "dc=ldap,dc=adop,dc=com"
var openldap = &Client{
	Host:         "localhost",
	Scheme:       "tcp",
	Port:         389,
	bindUser:     "cn=admin," + ldapDomain,
	bindPassword: "123qwe123",
}
