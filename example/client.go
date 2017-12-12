package main

import (
	"fmt"
	"github.com/tihtw/plurk-go"
	"log"
	"time"
)

const (
	ConsumerToken  = ""
	ConsumerSecret = ""
	AccessToken    = ""
	AccessSecret   = ""
)

func main() {
	fmt.Println("start")

	client := plurk.NewClient(&plurk.Config{
		AppKey:      ConsumerToken,
		AppSecret:   ConsumerSecret,
		TokenToken:  AccessToken,
		TokenSecret: AccessSecret,
	})

	// res, err := client.Call("/APP/Users/me", nil)
	// fmt.Println(string(res), err)

	user, err := client.GetMe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("got user: " + user.DisplayName)

	client.Listen(func(p *plurk.ResponseData) {
		log.Println("event", p.Type)
		switch p.Type {
		case plurk.EVENT_TYPE_UPDATE_NOTIFICATION:
			al, err := client.GetActive()
			if err != nil {
				log.Println("error:", err)
				return
			}
			for _, activ := range al {
				log.Println("active type:", activ.Type)
				switch activ.Type {
				case plurk.EVENT_TYPE_NEW_RESPONSE:
				case plurk.EVENT_TYPE_FRIENDSHIP_REQUEST:
					// 有新朋友了耶
					go func() {
						client.AddAsFriend(activ.FromUser.ID)
						time.Sleep(500 * time.Millisecond)
						p, _ := client.PlurkAdd("Hello, There are Maidwhite, I can help you for control your light~ 中文",
							"says",
							[]int{activ.FromUser.ID},
							[]int{})
						log.Println("create plurk success, plurk id: ", p.PlurkID)
						time.Sleep(500 * time.Millisecond)
						client.ResponseAdd(p.PlurkID, "If have any problem, be free to response here", "say")
						time.Sleep(500 * time.Millisecond)
						client.ResponseAdd(p.PlurkID, "Or tag @maidwhite.", "says")
						time.Sleep(500 * time.Millisecond)
						client.ResponseAdd(p.PlurkID, "There are the link for account binding: https://www.tih.tw", "says")
					}()

				}

			}

		}

	})

}
