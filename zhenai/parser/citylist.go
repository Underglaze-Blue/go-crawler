package parser

import (
	"crawler/engine"
	"regexp"
)

const urlPrefix = "http://localhost:8080/mock/www.zhenai.com/zhenghun/"

const cityListRe = `<a href="` + urlPrefix + `([0-9a-zA-Z]+)"[^>]*>([^<]*)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)

	result := engine.ParseResult{}
	matches := re.FindAllSubmatch(contents, -1)
	for _, v := range matches {
		//result.Items = append(result.Items, "City "+string(v[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        urlPrefix + string(v[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
