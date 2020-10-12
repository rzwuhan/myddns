package tencent

import (
	"errors"
	"myddns/log"
	"myddns/tencent/cns"
)

var (
	secretId  = ""
	secretKey = ""
)

func UseAccessKey(id, secret string) {
	secretId = id
	secretKey = secret
}

func UpdateDomainRecord(subDomain, domain, value string) error {
	cli := cns.NewClient(secretId, secretKey)

	// 1. query
	res, err := cli.RecordList(domain, subDomain, "A")
	if err != nil {
		return errors.New("record list failed:" + err.Error())
	}
	if len(res.Records) == 0 {
		// 2. add if not exists
		res, err := cli.RecordCreate(domain, subDomain, "A", value)
		if err != nil {
			return errors.New("record create failed:" + err.Error())
		}
		log.D("record create done:", *res)
	} else {
		// 3. update if exists
		record := res.Records[0]
		if record.Value == value {
			log.W("record value same, ignore update.")
			return nil
		}

		res, err := cli.RecordModify(domain, record.Id, subDomain, "A", value)
		if err != nil {
			return errors.New("record modify failed:" + err.Error())
		}
		log.D("record modify done:", *res)
	}
	return nil
}
