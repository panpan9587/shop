package common

// This file is auto-generated, don't edit it. Thanks.

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"math/rand"
	gourl "net/url"
	"time"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

// 封装 md5加密函数
func Md5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func calcAuthorization(source string, secretId string, secretKey string) (auth string, datetime string, err error) {
	timeLocation, _ := time.LoadLocation("Etc/GMT")
	datetime = time.Now().In(timeLocation).Format("Mon, 02 Jan 2006 15:04:05 GMT")
	signStr := fmt.Sprintf("x-date: %s\nx-source: %s", datetime, source)

	// hmac-sha1
	mac := hmac.New(sha1.New, []byte(secretKey))
	mac.Write([]byte(signStr))
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	auth = fmt.Sprintf("hmac id=\"%s\", algorithm=\"hmac-sha1\", headers=\"x-date x-source\", signature=\"%s\"",
		secretId, sign)

	return auth, datetime, nil
}

func urlencode(params map[string]string) string {
	var p = gourl.Values{}
	for k, v := range params {
		p.Add(k, v)
	}
	return p.Encode()
}

// 生成订单唯一标识
func GetOrderSn() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	randomPart := rand.Intn(1000)
	orderNumber := fmt.Sprintf("%d%03d", timestamp, randomPart)
	return orderNumber
}
