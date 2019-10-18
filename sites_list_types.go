package touchoffice

type SitesList struct {
	Status int   `json:"status"`
	Sites  Sites `json:"sites"`
}

type Sites []Site

type Site struct {
	SiteNumber    string `json:"site_number"`
	SiteName      string `json:"site_name"`
	SiteGroup     string `json:"site_group"`
	SiteGroupName string `json:"site_group_name"`
}
