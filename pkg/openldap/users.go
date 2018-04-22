package openldap

import (
	ldap "gopkg.in/ldap.v2"
)

// User CN is recommended to be equal to SN
type User struct {
	CN, displayName, givenName, mail, SN, UID, userPassword string
}

// CreateUser baseDN and ldapuser User
func (openldap *Client) CreateUser(baseDN string, ldapUser User) error {

	addRequest := ldap.NewAddRequest("cn=" + ldapUser.CN + ",ou=people," + baseDN)

	// default attributes for adop user
	addRequest.Attribute("objectClass", []string{"inetOrgPerson", "organizationalPerson", "person", "top"})

	// assign values
	addRequest.Attribute("displayName", []string{ldapUser.displayName})
	addRequest.Attribute("givenName", []string{ldapUser.givenName})
	addRequest.Attribute("mail", []string{ldapUser.mail})
	addRequest.Attribute("sn", []string{ldapUser.SN})
	addRequest.Attribute("uid", []string{ldapUser.UID})
	addRequest.Attribute("userPassword", []string{ldapUser.userPassword})

	// Add user
	return openldap.AddEntry(addRequest)

}

// DeleteUser baseDN and ldapuser User
func (openldap *Client) DeleteUser(baseDN string, ldapUser User) error {

	// Create Delete Request
	deleteRequest := ldap.NewDelRequest("cn="+ldapUser.CN+",ou=people,"+baseDN, nil)

	// Delete User
	return openldap.DeleteEntry(deleteRequest)

}

// AddUserToGroup ka allergy yung warnings
func (openldap *Client) AddUserToGroup(baseDN string, ldapUser User, ldapGroup Group) error {

	// Modify UniqueMember
	modRequest := ldap.NewModifyRequest("cn=" + ldapGroup.CN + ",ou=groups," + baseDN)

	// append user to list of unique Members
	ldapGroup.uniqueMember = append(ldapGroup.uniqueMember, "cn="+ldapUser.CN+",ou=people,"+baseDN)

	modRequest.Replace("uniqueMember", ldapGroup.uniqueMember)

	return openldap.ModifyEntry(modRequest)

}
