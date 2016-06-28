package handler

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/h-yamada/fb-bot-sample/config"

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
		{"img", "http://ohtm.cocolog-nifty.com/.shared/image.html?/photos/uncategorized/2013/04/26/ec98ac6b.jpg"},
		{"img", "http://ks.c.yimg.jp/res/chie-que-10136/10/136/469/560/i320"},
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

		i := rand.Intn(len(myMessages))

		var m interface{}
		switch myMessages[i].kind {
		case "text":
			m = NewTextMessage(messaging.Sender.ID, myMessages[i].message)
		case "image":
			m = NewImageMessage(messaging.Sender.ID, myMessages[i].message)
		default:
			m = NewTextMessage(messaging.Sender.ID, "orz")
		}

		if err := fb.SendMessage(m); err != nil {
			log.Println(err.Error())
		}
	}

	c.String(http.StatusOK, "OK")
}
