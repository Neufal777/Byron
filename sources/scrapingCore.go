package sources

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

func ProxyScraping(url string) (string, []error) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	proxies := []string{
		"201.220.140.30:8181",
		"136.226.33.115:80",
		//"94.139.254.232:3128",
	}

	randomElement := r1.Intn(len(proxies))

	fmt.Println("Using:", proxies[randomElement])

	proxyInfo := strings.Split(proxies[randomElement], ":")

	request := gorequest.New().Proxy("http://" + proxyInfo[0] + ":" + proxyInfo[1] + "")
	_, html, errs := request.Get(url).End()

	return html, errs
}
