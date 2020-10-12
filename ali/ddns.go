package ali

import (
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"myddns/log"
)

var (
	accessKeyId     = ""            // your accessKey
	accessKeySecret = ""            // your accessSecret
	regionId        = "cn-hangzhou" // 域名SDK请使用固定值"cn-hangzhou"
)

func UseAccessKey(id, secret string) {
	accessKeyId = id
	accessKeySecret = secret
}

func UpdateDomainRecord(rr, domain, value string) error {
	client, err := alidns.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return errors.New("create ali sdk client failed:" + err.Error())
	}

	// 1. query
	req := alidns.CreateDescribeSubDomainRecordsRequest()
	req.Scheme = "https"
	req.SubDomain = rr + "." + domain
	res, err := client.DescribeSubDomainRecords(req)
	if err != nil {
		return errors.New("query domain record failed:" + err.Error())
	}
	if res.TotalCount == 0 {
		// 2. add if not exists
		req := alidns.CreateAddDomainRecordRequest()
		req.Scheme = "https"
		req.DomainName = domain
		req.RR = rr
		req.Type = "A"
		req.Value = value
		res, err := client.AddDomainRecord(req)
		if err != nil {
			return errors.New("add domain record failed:" + err.Error())
		}
		log.D("add domain record done:\n", res.String())
	} else {
		// 3. update if exists
		record := &res.DomainRecords.Record[0]
		if record.Value == value {
			log.W("record value same, ignore update.")
			return nil
		}

		req := alidns.CreateUpdateDomainRecordRequest()
		req.Scheme = "https"
		req.RecordId = record.RecordId
		req.RR = record.RR
		req.Type = record.Type
		req.Value = value
		res, err := client.UpdateDomainRecord(req)
		if err != nil {
			return errors.New("update domain record failed:" + err.Error())
		}
		log.D("update domain record done:\n", res.String())
	}
	return nil
}
