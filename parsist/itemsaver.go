package parsist

import (
	"context"
	"crawler/engine"
	"log"

	"github.com/olivere/elastic/v7"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item saver #%d: %v", itemCount, item)
			itemCount++

			err := save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {

	indexSaver := client.Index().Index(index).BodyJson(item)

	if item.Id != "" {
		indexSaver.Id(item.Id)
	}

	_, err := indexSaver.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
