package parser

import (
	"crawler/engine"
	"regexp"
)

const cityPrefix = "http://localhost:8080/mock/album.zhenai.com/u/"

const cityRe = `<a href="` + cityPrefix + `([0-9]+)"[^>]*>([^<]*)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)

	result := engine.ParseResult{}
	matches := re.FindAllSubmatch(contents, -1)
	for _, v := range matches {
		name := string(v[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: cityPrefix + string(v[1]),
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, name)
			},
		})
	}
	return result
}
