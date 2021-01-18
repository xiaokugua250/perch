package helpers

type MetaDataKey struct {
	MetaUserName  string `json:"meta_user_name"`
	MetaAPIKey    string `json:"meta_api_key"`
	MetaUUID      string `json:"meta_uuid"`
	MetadaProNo   string `json:"metda_pro_no"`
	MetaTimeZone  int    `json:"meta_time_zone"`
	MetaLoginHref string `json:"meta_login_href"`
	MetaUserAgent string `json:"meta_user_agent"`
	//MetaLoginType string  `json:"meta_login_type"`

}

func MetaDataLoginKey(key MetaDataKey) (string, error) {

	return "", nil

}
