package model

import (
	"log"
	"reflect"

	"github.com/h-yamada/fb-bot-sample/config"

	"gopkg.in/olivere/elastic.v2"
)

type Movikuma struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
	Key         string `json:"key"`
	PlayCount   int    `json:"play_count"`
	ShareCount  int    `json:"share_count"`
}

func (movikuma *Movikuma) Search(keyword string) ([]Movikuma, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(config.MovikumaEsUrl),
		elastic.SetScheme("https"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	termQuery := elastic.NewTermQuery("description", keyword)
	searchResult, err := client.Search().
		Index("media").
		Type("medium").
		Query(&termQuery).
		Pretty(true).
		Do()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//var movikumaType Movikuma
	movList := []Movikuma{}
	for _, result := range searchResult.Each(reflect.TypeOf(movikuma)) {
		if m, ok := result.(Movikuma); ok {
			movList = append(movList, m)
		}
	}

	return movList, nil
}
