package openldap

import (
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

// Client gitlab Client
type Client struct {
	Host, Scheme string
	Port         int
}

// NewSearch queries baseDN using the searchFilter and returns searchAttribute provided
func (openldap *Client) NewSearch(baseDN, searchFilter, searchAttribute string) ([]string, error) {
	ldapClient, err := ldap.Dial(openldap.Scheme, fmt.Sprintf("%s:%d", openldap.Host, openldap.Port))
	if err != nil {
		return nil, fmt.Errorf("failed creating a ldap connection: %v", err)
	}
	defer ldapClient.Close()

	searchRequest := ldap.NewSearchRequest(
		baseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		searchFilter,              // The filter to apply
		[]string{searchAttribute}, // A list attributes to retrieve
		nil,
	)
	search, err := ldapClient.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("failed searching: %v", err)
	}

	var searchResult []string
	for _, entry := range search.Entries {
		cn := entry.GetAttributeValues(searchAttribute)
		for i := 0; i < len(cn); i++ {
			searchResult = append(searchResult, cn[i])
		}
	}
	return searchResult, nil

}
