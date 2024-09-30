package twitterscraper

import (
	"fmt"
	"net/http"
	"net/url"
)

type UserSettings struct {
	Protected                                       bool        `json:"protected"`
	ScreenName                                      string      `json:"screen_name"`
	AlwaysUseHTTPS                                  bool        `json:"always_use_https"`
	UseCookiePersonalization                        bool        `json:"use_cookie_personalization"`
	SleepTime                                       interface{} `json:"sleep_time"`
	GeoEnabled                                      bool        `json:"geo_enabled"`
	Language                                        string      `json:"language"`
	DiscoverableByEmail                             bool        `json:"discoverable_by_email"`
	DiscoverableByMobilePhone                       bool        `json:"discoverable_by_mobile_phone"`
	DisplaySensitiveMedia                           bool        `json:"display_sensitive_media"`
	PersonalizedTrends                              bool        `json:"personalized_trends"`
	AllowMediaTagging                               string      `json:"allow_media_tagging"`
	AllowContributorRequest                         string      `json:"allow_contributor_request"`
	AllowAdsPersonalization                         bool        `json:"allow_ads_personalization"`
	AllowLoggedOutDevicePersonalization             bool        `json:"allow_logged_out_device_personalization"`
	AllowLocationHistoryPersonalization             bool        `json:"allow_location_history_personalization"`
	AllowSharingDataForThirdPartyPersonalization    bool        `json:"allow_sharing_data_for_third_party_personalization"`
	AllowDMsFrom                                    string      `json:"allow_dms_from"`
	AlwaysAllowDMsFromSubscribers                   interface{} `json:"always_allow_dms_from_subscribers"`
	AllowDMGroupsFrom                               string      `json:"allow_dm_groups_from"`
	TranslatorType                                  string      `json:"translator_type"`
	CountryCode                                     string      `json:"country_code"`
	NsfwUser                                        bool        `json:"nsfw_user"`
	NsfwAdmin                                       bool        `json:"nsfw_admin"`
	RankedTimelineSetting                           interface{} `json:"ranked_timeline_setting"`
	RankedTimelineEligible                          interface{} `json:"ranked_timeline_eligible"`
	AddressBookLiveSyncEnabled                      bool        `json:"address_book_live_sync_enabled"`
	UniversalQualityFilteringEnabled                string      `json:"universal_quality_filtering_enabled"`
	DMReceiptSetting                                string      `json:"dm_receipt_setting"`
	AltTextComposeEnabled                           interface{} `json:"alt_text_compose_enabled"`
	MentionFilter                                   string      `json:"mention_filter"`
	AllowAuthenticatedPeriscopeRequests             bool        `json:"allow_authenticated_periscope_requests"`
	ProtectPasswordReset                            bool        `json:"protect_password_reset"`
	RequirePasswordLogin                            bool        `json:"require_password_login"`
	RequiresLoginVerification                       bool        `json:"requires_login_verification"`
	ExtSharingAudioSpacesListeningDataWithFollowers bool        `json:"ext_sharing_audiospaces_listening_data_with_followers"`
	Ext                                             interface{} `json:"ext"`
	DMQualityFilter                                 string      `json:"dm_quality_filter"`
	AutoplayDisabled                                bool        `json:"autoplay_disabled"`
	SettingsMetadata                                struct{}    `json:"settings_metadata"` // Empty struct since the original value is an empty object
}

// parseUserSettings parses the JSON response into the UserSettings struct.
func parseUserSettings(jsn map[string]interface{}) UserSettings {
	return UserSettings{
		Protected:                           jsn["protected"].(bool),
		ScreenName:                          jsn["screen_name"].(string),
		AlwaysUseHTTPS:                      jsn["always_use_https"].(bool),
		UseCookiePersonalization:            jsn["use_cookie_personalization"].(bool),
		SleepTime:                           jsn["sleep_time"],
		GeoEnabled:                          jsn["geo_enabled"].(bool),
		Language:                            jsn["language"].(string),
		DiscoverableByEmail:                 jsn["discoverable_by_email"].(bool),
		DiscoverableByMobilePhone:           jsn["discoverable_by_mobile_phone"].(bool),
		DisplaySensitiveMedia:               jsn["display_sensitive_media"].(bool),
		PersonalizedTrends:                  jsn["personalized_trends"].(bool),
		AllowMediaTagging:                   jsn["allow_media_tagging"].(string),
		AllowContributorRequest:             jsn["allow_contributor_request"].(string),
		AllowAdsPersonalization:             jsn["allow_ads_personalization"].(bool),
		AllowLoggedOutDevicePersonalization: jsn["allow_logged_out_device_personalization"].(bool),
		AllowLocationHistoryPersonalization: jsn["allow_location_history_personalization"].(bool),
		AllowSharingDataForThirdPartyPersonalization: jsn["allow_sharing_data_for_third_party_personalization"].(bool),
		AllowDMsFrom:                        jsn["allow_dms_from"].(string),
		AlwaysAllowDMsFromSubscribers:       jsn["always_allow_dms_from_subscribers"],
		AllowDMGroupsFrom:                   jsn["allow_dm_groups_from"].(string),
		TranslatorType:                      jsn["translator_type"].(string),
		CountryCode:                         jsn["country_code"].(string),
		NsfwUser:                            jsn["nsfw_user"].(bool),
		NsfwAdmin:                           jsn["nsfw_admin"].(bool),
		RankedTimelineSetting:               jsn["ranked_timeline_setting"],
		RankedTimelineEligible:              jsn["ranked_timeline_eligible"],
		AddressBookLiveSyncEnabled:          jsn["address_book_live_sync_enabled"].(bool),
		UniversalQualityFilteringEnabled:    jsn["universal_quality_filtering_enabled"].(string),
		DMReceiptSetting:                    jsn["dm_receipt_setting"].(string),
		AltTextComposeEnabled:               jsn["alt_text_compose_enabled"],
		MentionFilter:                       jsn["mention_filter"].(string),
		AllowAuthenticatedPeriscopeRequests: jsn["allow_authenticated_periscope_requests"].(bool),
		ProtectPasswordReset:                jsn["protect_password_reset"].(bool),
		RequirePasswordLogin:                jsn["require_password_login"].(bool),
		RequiresLoginVerification:           jsn["requires_login_verification"].(bool),
		ExtSharingAudioSpacesListeningDataWithFollowers: jsn["ext_sharing_audiospaces_listening_data_with_followers"].(bool),
		Ext:              jsn["ext"],
		DMQualityFilter:  jsn["dm_quality_filter"].(string),
		AutoplayDisabled: jsn["autoplay_disabled"].(bool),
		SettingsMetadata: struct{}{}, // Empty struct
	}
}

// GetProfile return parsed user profile.
func (s *Scraper) GetUserSettings(authToken AuthToken) (UserSettings, error) {
	var jsn map[string]interface{}
	req, err := http.NewRequest("GET", "https://api.x.com/1.1/account/settings.json", nil)

	if err != nil {
		return UserSettings{}, err
	}

	query := url.Values{}
	query.Set("include_ext_sharing_audiospaces_listening_data_with_followers", "true")
	query.Set("include_mention_filter", "true")
	query.Set("include_nsfw_user_flag", "true")
	query.Set("include_nsfw_admin_flag", "true")
	query.Set("include_ranked_timeline", "true")
	query.Set("include_alt_text_compose", "true")
	query.Set("ext", "ssoConnections")
	query.Set("include_country_code", "true")
	query.Set("include_ext_dm_nsfw_media_filter", "true")
	req.URL.RawQuery = query.Encode()

	req.Header.Set("Cookie", fmt.Sprintf("auth_token=%s; ct0=%s", authToken.Token, authToken.CSRFToken))
	req.Header.Set("X-CSRF-Token", authToken.CSRFToken)

	err = s.RequestAPI(req, &jsn)

	if err != nil {
		return UserSettings{}, err
	}

	return parseUserSettings(jsn), nil
}
