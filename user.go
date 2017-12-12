package plurk

type User struct {
	Avatar          int         `json:"avatar"`
	BackgroundID    int         `json:"background_id"`
	BdayPrivacy     int         `json:"bday_privacy"`
	DateOfBirth     string      `json:"date_of_birth"`
	Dateformat      int         `json:"dateformat"`
	DefaultLang     string      `json:"default_lang"`
	DisplayName     string      `json:"display_name"`
	FullName        string      `json:"full_name"`
	Gender          int         `json:"gender"`
	HasProfileImage int         `json:"has_profile_image"`
	ID              int         `json:"id"`
	Karma           float64     `json:"karma"`
	Location        string      `json:"location"`
	NameColor       interface{} `json:"name_color"`
	NickName        string      `json:"nick_name"`
	PinnedPlurkID   interface{} `json:"pinned_plurk_id"`
	Premium         bool        `json:"premium"`
	ShowAds         bool        `json:"show_ads"`
	ShowLocation    int         `json:"show_location"`
	TimelinePrivacy int         `json:"timeline_privacy"`
	Timezone        string      `json:"timezone"`
	VerifiedAccount bool        `json:"verified_account"`
}
