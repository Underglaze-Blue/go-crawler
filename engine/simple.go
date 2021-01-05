package engine

import (
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
