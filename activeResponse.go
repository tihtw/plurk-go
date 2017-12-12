package plurk

type ActiveResponse struct {
	FromUser   User        `json:"from_user"`
	NumOthers  int         `json:"num_others"`
	PlurkID    int         `json:"plurk_id"`
	Posted     string      `json:"posted"`
	ResponseID interface{} `json:"response_id"`
	Type       string      `json:"type"`
}
