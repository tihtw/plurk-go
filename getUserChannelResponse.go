package plurk

type GetUserChannelResponse struct {
	ChannelName string `json:"channel_name"`
	CometServer string `json:"comet_server"`
}
