package cns

import (
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Request struct {
	*tchttp.BaseRequest
}

func (r *Request) GetPath() string {
	return "/v2/index.php"
}

func (r *Request) GetUrl() string {
	if r.GetHttpMethod() == "GET" {
		return r.GetScheme() + "://" + r.GetDomain() + r.GetPath() + "?" + tchttp.GetUrlQueriesEncoded(r.GetParams())
	} else if r.GetHttpMethod() == "POST" {
		return r.GetScheme() + "://" + r.GetDomain() + r.GetPath()
	} else {
		return ""
	}
}

func NewRequest(action string) *Request {
	req := &Request{&tchttp.BaseRequest{}}
	req.Init().WithApiInfo("cns", "", action)
	req.SetRootDomain("api.qcloud.com")
	return req
}
