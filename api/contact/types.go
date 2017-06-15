package contact

import "github.com/qwentic/qcrm/api/site"

//CSInfo struct to hold post request for contact
type CSInfo struct {
	*Contact
	Site []site.Site
}
