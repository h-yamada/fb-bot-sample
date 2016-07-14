package handler

import (
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/h-yamada/fb-bot-sample/config"
	. "github.com/h-yamada/fb-bot-sample/model"

	. "github.com/ymd38/facebook/messenger"

	"github.com/gin-gonic/gin"
)

type MyMessage struct {
	kind    string
	message string
}

func PostWebHook(c *gin.Context) {
	myMessages := []MyMessage{
		{"text", "あぼーん"},
		{"text", "残念だな"},
		{"text", "なんだと？"},
		{"text", "...うぜえ"},
		{"img", "http://1093.up.n.seesaa.net/1093/image/takokora.jpg"},
		{"img", "http://s.eximg.jp/exnews/feed/Kotaku/Kotaku_201211_sce_new_cm_3.jpg"},
		{"img", "http://stat.ameba.jp/user_images/20160514/19/hajackass/72/cf/j/o0800060013645886097.jpg"},
	}

	receiver := &ReceivedMessage{}

	if err := c.BindJSON(&receiver); err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	log.Println(receiver)

	for _, messaging := range receiver.Entry[0].Messaging {
		log.Println(messaging.Sender.ID)
		log.Println(messaging.Message.Text)

		fb := NewFacebookMessenger(config.FbToken)
		rand.Seed(time.Now().UnixNano())

		var m interface{}

		movikuma := new(Movikuma)
		movikumaList, _ := movikuma.Search(messaging.Message.Text)

		log.Println(movikumaList)

		if verse := lime(messaging.Message.Text); verse != "" {
			m = NewTextMessage(messaging.Sender.ID, verse)
		} else {
			i := rand.Intn(len(myMessages))
			switch myMessages[i].kind {
			case "text":
				m = NewTextMessage(messaging.Sender.ID, myMessages[i].message)
			case "img":
				m = NewImageMessage(messaging.Sender.ID, myMessages[i].message)
			default:
				m = NewTextMessage(messaging.Sender.ID, "orz")
			}

		}

		if err := fb.SendMessage(m); err != nil {
			log.Println(err.Error())
		}
	}

	c.String(http.StatusOK, "OK")
}

func lime(message string) string {
	myVerse := []string{
		"かますぜ 俺はボット　最近やたらとホット　鋭いライムでおまえの喉元カット",
		"俺は検索のお色直しじゃねー F8見ただろ？可能性は半端ねー",
	}

	switch {
	case regexp.MustCompile(`おい`).Match([]byte(message)):
		return myVerse[0]
	case regexp.MustCompile(`検索`).Match([]byte(message)):
		return myVerse[1]
	}
	return ""
}
