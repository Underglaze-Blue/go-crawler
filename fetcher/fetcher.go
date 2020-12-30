package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var timeTick = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-timeTick
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code： %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
