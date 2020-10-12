package cns

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Response struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	CodeDesc string      `json:"codeDesc"`
	Data     interface{} `json:"data"`
}

func (r *Response) ParseErrorFromHTTPResponse(body []byte) error {
	err := json.Unmarshal(body, r)
	if err != nil {
		return errors.New(fmt.Sprintf("parse response failed, err:%s, body:%s", err.Error(), string(body)))
	}
	return nil
}

type RecordCreateResponse struct {
	Record struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
		Weight int    `json:"weight"`
	} `json:"record"`
}

type RecordStatusResponse struct {
	Record struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
		Weight int    `json:"weight"`
	} `json:"record"`
}

type RecordModifyResponse struct {
	Record struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Value  string `json:"value"`
		Status string `json:"status"`
		Weight int    `json:"weight"`
	}
}

type RecordListResponse struct {
	Domain struct {
		Id         string   `json:"id"`
		Name       string   `json:"name"`
		PunyCode   string   `json:"puny_code"`
		Grade      string   `json:"grade"`
		Owner      string   `json:"owner"`
		ExtStatus  string   `json:"ext_status"`
		TTL        int      `json:"ttl"`
		MinTTL     int      `json:"min_ttl"`
		DnspodNS   []string `json:"dnspod_ns"`
		Status     string   `json:"status"`
		QProjectId int      `json:"q_project_id"`
	} `json:"domain"`
	Info struct {
		SubDomains  string `json:"sub_domains"`
		RecordTotal string `json:"record_total"`
	} `json:"info"`
	Records []struct {
		Id         int    `json:"id"`
		TTL        int    `json:"ttl"`
		Value      string `json:"value"`
		Enabled    int    `json:"enabled"`
		Status     string `json:"status"`
		UpdatedOn  string `json:"updated_on"`
		QProjectId int    `json:"q_project_id"`
		Name       string `json:"name"`
		Line       string `json:"line"`
		Type       string `json:"type"`
		Remark     string `json:"remark"`
		MX         int    `json:"mx"`
		Hold       string `json:"hold"`
	} `json:"records"`
}
