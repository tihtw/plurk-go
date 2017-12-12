package plurk

type Response struct {
	Content             string      `json:"content"`
	ContentRaw          string      `json:"content_raw"`
	Editability         int         `json:"editability"`
	ID                  int         `json:"id"`
	Lang                string      `json:"lang"`
	LastEdited          interface{} `json:"last_edited"`
	PlurkID             int         `json:"plurk_id"`
	Posted              string      `json:"posted"`
	Qualifier           string      `json:"qualifier"`
	QualifierTranslated string      `json:"qualifier_translated"`
	UserID              int         `json:"user_id"`
	WithRandomEmos      bool        `json:"with_random_emos"`
}
