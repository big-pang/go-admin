package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"math/big"
	"strconv"
	"strings"
)

// Redeemcode  兑换码
type Redeemcode struct {
	Md5    string `json:"redeemcode_md5"`
	cipher string `json:"redeemcode_cipher"`
}

// CreateRedeemCode 兑换码按批次号生成
type CreateRedeemCode struct {
	BatchId     string
	Redeemcodes []Redeemcode
}

func (crc *CreateRedeemCode) Get() string {
	//先生成随机数
	const min, max = 1000000000000, 9999999999999
	randNum, _ := rand.Int(rand.Reader, big.NewInt(max-min))
	randNum = randNum.Add(randNum, big.NewInt(min))
	verifyNum, err := convert15(randNum.Int64())
	if err != nil {
		logs.Error(err)
	}
	str := fmt.Sprintf("%s%d", crc.BatchId, verifyNum)
	code, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		logs.Error(err)
	}
	return strings.ToUpper(strconv.FormatInt(code, 16))
}

func (crc *CreateRedeemCode) Check(code string) bool {
	//先转int64
	num, err := strconv.ParseInt(code, 16, 64)
	if err != nil {
		logs.Error(err)
	}
	str := strconv.FormatInt(num, 10)
	str1, _ := strconv.ParseInt(str[len(str)-15:], 10, 64)
	str2, _ := strconv.ParseInt(str[len(str)-15:len(str)-2], 10, 64)
	str3, _ := convert15(str2)
	return str1 == str3
}
func convert15(n int64) (int64, error) {
	str := strconv.FormatInt(n, 10)
	num0, _ := strconv.ParseInt(str[0:1], 10, 64)
	num1, _ := strconv.ParseInt(str[1:2], 10, 64)
	num2, _ := strconv.ParseInt(str[2:3], 10, 64)
	num3, _ := strconv.ParseInt(str[3:4], 10, 64)
	num4, _ := strconv.ParseInt(str[4:5], 10, 64)
	num5, _ := strconv.ParseInt(str[5:6], 10, 64)
	num6, _ := strconv.ParseInt(str[6:7], 10, 64)
	num7, _ := strconv.ParseInt(str[7:8], 10, 64)
	num8, _ := strconv.ParseInt(str[8:9], 10, 64)
	num9, _ := strconv.ParseInt(str[9:10], 10, 64)
	num := ((num1+num3+num5+num7+num9)*3 + num0 + num2 + num4 + num6 + num8) % 100
	str = ""
	if num < 10 {
		str = fmt.Sprintf("%s%d", "0", num)
	}
	str = fmt.Sprintf("%d%d%s", n, num, str)
	return strconv.ParseInt(str, 10, 64)
}
