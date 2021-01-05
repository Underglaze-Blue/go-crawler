package engine

import (
	"crawler/fetcher"

	"go.uber.org/zap"
)

func worker(r Request) (ParseResult, error) {
	logger, _ := zap.NewProduction()
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		logger.Info("", zap.String("Fetcher: error fetching url", r.Url), zap.Error(err))
		return ParseResult{}, err
	}

	return r.ParserFunc(body, r.Url), nil
}
