package cns

import (
	"errors"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"strconv"
)

type Client struct {
	*common.Client
}

//
func NewClient(secretId, secretKey string) *Client {
	cpf := profile.NewClientProfile()
	//cpf.Debug = true
	client, _ := common.NewClientWithSecretId(secretId, secretKey, "")
	client.WithProfile(cpf)
	client.WithSignatureMethod("HmacSHA256")
	return &Client{client}
}

func (c *Client) doRequest(req *Request, data interface{}) error {
	res := &Response{Data: data}
	if err := c.Send(req, res); err != nil {
		return err
	}
	if res.Code != 0 {
		return errors.New(fmt.Sprintf("request failed with code %d(%s)", res.Code, res.CodeDesc))
	}
	return nil
}

// https://cloud.tencent.com/document/api/302/8516
func (c *Client) RecordCreate(domain, subDomain, recordType, value string) (*RecordCreateResponse, error) {
	req := NewRequest("RecordCreate")
	params := req.GetParams()
	params["domain"] = domain
	params["subDomain"] = subDomain
	params["recordType"] = recordType
	params["recordLine"] = "默认"
	params["value"] = value
	//if ttl > 0 {
	//	params["ttl"] = strconv.Itoa(ttl)
	//}
	//if mx > 0 {
	//	params["ms"] = strconv.Itoa(mx)
	//}

	data := &RecordCreateResponse{}
	if err := c.doRequest(req, data); err != nil {
		return nil, err
	}
	return data, nil
}

// https://cloud.tencent.com/document/api/302/8519
func (c *Client) RecordStatus(domain string, recordId int, status string) (*RecordStatusResponse, error) {
	req := NewRequest("RecordStatus")
	params := req.GetParams()
	params["domain"] = domain
	params["recordId"] = strconv.Itoa(recordId)
	params["status"] = status

	data := &RecordStatusResponse{}
	if err := c.doRequest(req, data); err != nil {
		return nil, err
	}
	return data, nil
}

// https://cloud.tencent.com/document/api/302/8511
func (c *Client) RecordModify(domain string, recordId int, subDomain, recordType, value string) (*RecordModifyResponse, error) {
	req := NewRequest("RecordModify")
	params := req.GetParams()
	params["domain"] = domain
	params["recordId"] = strconv.Itoa(recordId)
	params["subDomain"] = subDomain
	params["recordType"] = recordType
	params["recordLine"] = "默认"
	params["value"] = value
	//if ttl > 0 {
	//	params["ttl"] = strconv.Itoa(ttl)
	//}
	//if mx > 0 {
	//	params["ms"] = strconv.Itoa(mx)
	//}

	data := &RecordModifyResponse{}
	if err := c.doRequest(req, data); err != nil {
		return nil, err
	}
	return data, nil
}

// https://cloud.tencent.com/document/api/302/8517
func (c *Client) RecordList(domain string, subDomain, recordType string) (*RecordListResponse, error) {
	req := NewRequest("RecordList")
	params := req.GetParams()
	params["domain"] = domain
	//if offset > 0 {
	//	params["offset"] = strconv.Itoa(offset)
	//}
	//if length > 0 {
	//	params["length"] = strconv.Itoa(length)
	//}
	if len(subDomain) > 0 {
		params["subDomain"] = subDomain
	}
	if len(recordType) > 0 {
		params["recordType"] = recordType
	}
	//if qProjectId > 0 {
	//	params["qProjectId"] = strconv.Itoa(qProjectId)
	//}

	data := &RecordListResponse{}
	if err := c.doRequest(req, data); err != nil {
		return nil, err
	}
	return data, nil
}
