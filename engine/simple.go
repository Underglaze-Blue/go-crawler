package engine

import (
	"crawler/fetcher"

	"go.uber.org/zap"
)

type SimpleEngine struct {
}

func (SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	logger, _ := zap.NewProduction()
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			logger.Info("Got item", zap.Any("item", item))
		}
	}
}

func worker(r Request) (ParseResult, error) {
	logger, _ := zap.NewProduction()
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		logger.Info("", zap.String("Fetcher: error fetching url", r.Url), zap.Error(err))
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
