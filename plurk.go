package plurk

type Plurk struct {
	Anonymous           bool          `json:"anonymous"`
	Bookmark            bool          `json:"bookmark"`
	Coins               int           `json:"coins"`
	Content             string        `json:"content"`
	ContentRaw          string        `json:"content_raw"`
	Excluded            interface{}   `json:"excluded"`
	Favorers            []interface{} `json:"favorers"`
	Favorite            bool          `json:"favorite"`
	FavoriteCount       int           `json:"favorite_count"`
	HasGift             bool          `json:"has_gift"`
	IsUnread            int           `json:"is_unread"`
	Lang                string        `json:"lang"`
	LastEdited          interface{}   `json:"last_edited"`
	LimitedTo           interface{}   `json:"limited_to"`
	Mentioned           int           `json:"mentioned"`
	NoComments          int           `json:"no_comments"`
	OwnerID             int           `json:"owner_id"`
	PlurkID             int           `json:"plurk_id"`
	PlurkType           int           `json:"plurk_type"`
	Porn                bool          `json:"porn"`
	Posted              string        `json:"posted"`
	Qualifier           string        `json:"qualifier"`
	QualifierTranslated string        `json:"qualifier_translated"`
	Replurkable         bool          `json:"replurkable"`
	Replurked           bool          `json:"replurked"`
	ReplurkerID         interface{}   `json:"replurker_id"`
	Replurkers          []interface{} `json:"replurkers"`
	ReplurkersCount     int           `json:"replurkers_count"`
	Responded           int           `json:"responded"`
	ResponseCount       int           `json:"response_count"`
	ResponsesSeen       int           `json:"responses_seen"`
	UserID              int           `json:"user_id"`
}
