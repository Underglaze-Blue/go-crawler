package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	fetch, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(fetch)

	const resultSize = 470

	if len(result.Requests) != resultSize {
		t.Errorf("result should hava %d requests; but had %d", resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should hava %d Items; but had %d", resultSize, len(result.Items))
	}

}
