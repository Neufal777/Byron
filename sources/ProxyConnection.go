package sources

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

func ProxyScraping(url string) (string, []error) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	proxies := ProxiesCleaner()

	randomElement := r1.Intn(len(proxies))

	fmt.Println("Using:", proxies[randomElement])

	//proxyInfo := strings.Split(proxies[randomElement], ":")

	//request := gorequest.New().Proxy("http://" + proxyInfo[0] + ":" + proxyInfo[1])
	request := gorequest.New().Proxy("")
	_, body, _ := request.Get(url).End()

	return body, nil
}

func ProxiesCleaner() []string {
	proxies, _ := ioutil.ReadFile("sources/proxies.txt")
	proxiesList := strings.Split(string(proxies), "\n")

	return proxiesList
}
