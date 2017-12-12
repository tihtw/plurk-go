package plurk

import (
	"encoding/json"
	"github.com/garyburd/go-oauth/oauth"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Config struct {
	AppKey      string
	AppSecret   string
	TokenToken  string
	TokenSecret string
}

type Client struct {
	Config
	Host string
}

var COMMENT_HEAD_LENGTH = len("CometChannel.scriptCallback(")
var COMMENT_TAIL_LENGTH = len(");")

func NewClient(config *Config) *Client {
	return &Client{
		Host:   "https://www.plurk.com",
		Config: *config,
	}

}

func (c *Client) GetMe() (*User, error) {
	data, err := c.Call("/APP/Users/me", nil)
	if err != nil {
		return nil, err
	}
	u := &User{}
	err = json.Unmarshal(data, u)
	if err != nil {
		return nil, err
	}
	return u, nil

}

func (c *Client) GetActive() ([]*ActiveResponse, error) {
	data, err := c.Call("/APP/Alerts/getActive", nil)
	if err != nil {
		return nil, err
	}
	log.Println("get active", string(data))
	rl := []*ActiveResponse{}
	err = json.Unmarshal(data, &rl)
	if err != nil {
		return nil, err
	}
	return rl, nil
}

func (c *Client) AddAsFriend(userId int) (bool, error) {
	data, err := c.Call("/APP/Alerts/addAsFriend", map[string]string{"user_id": strconv.Itoa(userId)})
	if err != nil {
		return false, err
	}
	rl := map[string]string{}
	err = json.Unmarshal(data, &rl)
	if err != nil {
		return false, err
	}
	if rl["success_text"] == "ok" {
		return true, nil
	}

	return false, nil
}

func (c *Client) CheckToken() {

}

func (c *Client) PlurkAdd(content, qualifier string, limitedTo []int, excluded []int) (*Plurk, error) {

	limitedToByte, _ := json.Marshal(limitedTo)

	param := map[string]string{
		"content":    content,
		"qualifier":  qualifier,
		"limited_to": string(limitedToByte),
		"lang":       "tr_ch",
	}

	data, err := c.Call("/APP/Timeline/plurkAdd", param)
	if err != nil {
		return nil, err
	}
	p := Plurk{}
	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil

} // 你好，我是Maidwhite，是可以幫你開燈關燈的Maidwhite喔。

func (c *Client) Listen(handlerFunc func(responseData *ResponseData)) error {
	data, err := c.Call("/APP/Realtime/getUserChannel", nil)
	if err != nil {
		return err
	}
	// Should be
	//  {
	//     "channel_name": "generic-xxx",
	//     "comet_server": "https://comet08.plurk.com/comet?channel=generic-14411453-3c7236176e98450d72d416b81eae7fcc0aaec1fc&offset=0"
	// }
	r := &GetUserChannelResponse{}
	err = json.Unmarshal(data, r)
	if err != nil {
		return err
	}
	log.Println("comet server: " + r.CometServer)

	req, err := http.NewRequest(http.MethodGet, r.CometServer, nil)
	if err != nil {
		return err
	}
	for {
		log.Println("comet start")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			log.Println("http error:", err)
			return err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("get response error:", err)
			return err
		}
		body = body[COMMENT_HEAD_LENGTH : len(body)-COMMENT_TAIL_LENGTH]
		log.Println("response ", string(body))
		resp.Body.Close()
		cr := CometResponse{}
		err = json.Unmarshal(body, &cr)
		if err != nil {
			log.Println("get response error:", err)
			return err
		}

		log.Println("new offset:", cr.NewOffset, "data length:", len(cr.Data))

		for _, data := range cr.Data {
			go handlerFunc(data)
		}

	}

}

func (c *Client) ResponseAdd(plurkId int, content, qualifier string) (*Response, error) {

	param := map[string]string{
		"plurk_id":  strconv.Itoa(plurkId),
		"content":   content,
		"qualifier": qualifier,
	}

	data, err := c.Call("/APP/Responses/responseAdd", param)
	if err != nil {
		return nil, err
	}
	r := Response{}
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *Client) Call(urlStr string, param map[string]string) ([]byte, error) {

	values := url.Values{}
	for k, v := range param {
		values.Set(k, v)
	}

	cred := &oauth.Credentials{
		Token:  c.TokenToken,
		Secret: c.TokenSecret,
	}
	oauthClient := &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  c.AppKey,
			Secret: c.AppSecret,
		},
	}
	urlStr = c.Host + urlStr

	// oauthClient.SignParam(token, "GET", apiURL, values)
	// req ,err := http.NewRequest(http.MethodGet, urlStr, nil)
	res, err := oauthClient.Get(nil, cred, urlStr, values)
	if err != nil {
		log.Println("network error:", err, urlStr, values)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println("http error: code:", res.StatusCode, err, urlStr, values)
		if res.StatusCode != 404 {
			return nil, err
		}
	}
	body, err := ioutil.ReadAll(res.Body)
	log.Println("body", string(body))
	if err != nil {
		log.Println("get response error:", err)
		return nil, err
	}
	return body, nil

}
