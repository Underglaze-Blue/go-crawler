package parser

import (
	"crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]*)</a>`)

var cityUrlRe = regexp.MustCompile(`href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/baotou/[^"]+)"`)

func ParseCity(contents []byte, _ string) engine.ParseResult {

	result := engine.ParseResult{}
	matches := profileRe.FindAllSubmatch(contents, -1)
	for _, v := range matches {
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(v[1]),
			ParserFunc: ProfileParse(string(v[2])),
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, v := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(v[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
