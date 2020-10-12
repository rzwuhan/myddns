package utils

import (
	"errors"
	"fmt"
	"io"
	"myddns/log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var ipPattern = regexp.MustCompile("\\d{1,3}(\\.\\d{1,3}){3}")

func httpGet(url string) (string, error) {
	http.DefaultClient.Timeout = 3 * time.Second
	if res, err := http.Get(url); err != nil {
		return "", err
	} else if res.StatusCode < 200 || res.StatusCode > 399 {
		return "", errors.New(fmt.Sprintf("got status code %d", res.StatusCode))
	} else {
		var buffer []byte
		if res.ContentLength > 0 {
			buffer = make([]byte, res.ContentLength)
		} else {
			buffer = make([]byte, 1024)
		}
		nRead, err := res.Body.Read(buffer)
		if err != nil && err != io.EOF {
			return "", errors.New("read response body failed:" + err.Error())
		}
		if nRead <= 0 {
			return "", errors.New("got empty body")
		}
		return strings.TrimSpace(string(buffer[:nRead])), nil
	}
}

func GetIp() string {
	urls := []string{
		"http://www.3322.org/dyndns/getip",     // 111.172.28.136
		"http://myexternalip.com/raw",          // 111.172.28.136
		"http://pv.sohu.com/cityjson?ie=utf-8", // var returnCitySN = {"cip": "111.172.28.136", "cid": "420100", "cname": "湖北省武汉市"};
	}
	result := ""
	for _, url := range urls {
		if body, err := httpGet(url); err != nil {
			log.E("get ip failed(", url, "):", err)
			continue
		} else if ip := ipPattern.FindString(body); len(ip) == 0 {
			log.E("get ip failed(", url, "):", body)
			continue
		} else {
			log.I("get ip success(", url, "):", ip)
			result = ip
			break
		}
	}
	return result
}
