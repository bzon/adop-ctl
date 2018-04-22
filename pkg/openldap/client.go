package openldap

import (
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

// Client gitlab Client
type Client struct {
	Host, Scheme, bindUser, bindPassword string
	Port                                 int
}

// NewSearch queries baseDN using the searchFilter and returns searchAttribute provided
func (openldap *Client) NewSearch(baseDN, searchFilter, searchAttribute string) ([]string, error) {

	// Initialize Ldap Client
	ldapClient, err := ldap.Dial(openldap.Scheme, fmt.Sprintf("%s:%d", openldap.Host, openldap.Port))
	if err != nil {
		return nil, fmt.Errorf("failed creating a ldap connection: %v", err)
	}
	defer ldapClient.Close()

	// Bind is optional for search will work without bind
	// err = ldapClient.Bind(openldap.bindUser, openldap.bindPassword)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed creating a ldap connection binding: %v", err)

	// }

	// Create Search Request
	searchRequest := ldap.NewSearchRequest(
		baseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		searchFilter,              // The filter to apply
		[]string{searchAttribute}, // A list attributes to retrieve
		nil,
	)

	// Run search
	search, err := ldapClient.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("failed searching: %v", err)
	}

	// Get Result
	var searchResult []string
	for _, entry := range search.Entries {
		cn := entry.GetAttributeValues(searchAttribute)
		for i := 0; i < len(cn); i++ {
			searchResult = append(searchResult, cn[i])
		}
	}
	return searchResult, nil

}

// AddEntry addRequest *ldap.AddRequest
func (openldap *Client) AddEntry(addRequest *ldap.AddRequest) error {

	// Initialize Ldap Client
	ldapClient, err := ldap.Dial(openldap.Scheme, fmt.Sprintf("%s:%d", openldap.Host, openldap.Port))
	if err != nil {
		return fmt.Errorf("failed creating a ldap connection: %v", err)
	}
	defer ldapClient.Close()

	// Bind Credentials
	err = ldapClient.Bind(openldap.bindUser, openldap.bindPassword)
	if err != nil {
		return fmt.Errorf("failed creating a ldap connection binding: %v", err)

	}

	// Run Add Request
	err = ldapClient.Add(addRequest)
	if err != nil {
		return fmt.Errorf("Entry %v not added: %v", addRequest.DN, err)
	}
	return nil

}

// DeleteEntry delete *ldap.deleteRequest
func (openldap *Client) DeleteEntry(delRequest *ldap.DelRequest) error {

	// Initialize Ldap Client
	ldapClient, err := ldap.Dial(openldap.Scheme, fmt.Sprintf("%s:%d", openldap.Host, openldap.Port))
	if err != nil {
		return fmt.Errorf("failed creating a ldap connection: %v", err)
	}
	defer ldapClient.Close()

	// Bind Credentials
	err = ldapClient.Bind(openldap.bindUser, openldap.bindPassword)
	if err != nil {
		return fmt.Errorf("failed creating a ldap connection binding: %v", err)

	}

	// Run Delete Request
	err = ldapClient.Del(delRequest)
	if err != nil {
		return fmt.Errorf("Entry %v not deleted: %v", delRequest.DN, err)
	}
	return nil

}

// ModifyEntry delete *ldap.deleteRequest
func (openldap *Client) ModifyEntry(modRequest *ldap.ModifyRequest) error {

	// Initialize Ldap Client
	ldapClient, err := ldap.Dial(openldap.Scheme, fmt.Sprintf("%s:%d", openldap.Host, openldap.Port))
	if err != nil {
		return fmt.Errorf("failed creating a ldap connection: %v", err)
	}
	defer ldapClient.Close()

	// Bind Credentials
	err = ldapClient.Bind(openldap.bindUser, openldap.bindPassword)
	if err != nil {
		return fmt.Errorf("failed creating a ldap connection binding: %v", err)

	}

	// Run Modify Request
	err = ldapClient.Modify(modRequest)
	if err != nil {
		return fmt.Errorf("Entry %v not modified: %v", modRequest.DN, err)
	}
	return nil

}
