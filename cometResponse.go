package plurk

type CometResponse struct {
	Data      []*ResponseData `json:"data"`
	NewOffset int             `json:"new_offset"`
}

const (
	EVENT_TYPE_UPDATE_NOTIFICATION = "update_notification"
	EVENT_TYPE_MENTIONED           = "mentioned"
	EVENT_TYPE_FRIENDSHIP_REQUEST  = "friendship_request"
	EVENT_TYPE_NEW_RESPONSE        = "new_response"
	EVENT_TYPE_PRIVATE_PLURK       = "private_plurk"
	EVENT_TYPE_PLURK_LIKE          = "plurk_liked"
)

type ResponseData struct {
	Plurk         Plurk            `json:"plurk"`
	PlurkID       int              `json:"plurk_id"`
	Response      Response         `json:"response"`
	ResponseCount int              `json:"response_count"`
	Type          string           `json:"type"`
	User          map[string]*User `json:"user"`
}
