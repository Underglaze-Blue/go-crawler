package engine

import (
	"crawler/fetcher"

	"go.uber.org/zap"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	logger, _ := zap.NewProduction()
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		logger.Info("Fetching", zap.String("url", r.Url))
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			logger.Info("", zap.String("Fetcher: error fetching url", r.Url), zap.Error(err))
			//log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			logger.Info("Got item", zap.Any("item", item))
			//log.Printf("Got item %v", item)
		}
	}
}
