package utils

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	fmt.Println("Test redeemcode Get ...")
	r := &CreateRedeemCode{BatchId: "100"}
	code := r.Get()
	fmt.Println("code is :", code)
	if len(code) <= 0 {
		t.Fatal("兑换码生成有问题！")
	}
}

func TestCheck(t *testing.T) {
	fmt.Println("Test redeemcode Check ...")
	code := "165FAE84766D236"
	fcode := "165FAE84766D235"
	r := &CreateRedeemCode{BatchId: "100"}
	if !r.Check(code) {
		t.Fatal("正确兑换码验证出错！")
	}

	if r.Check(fcode) {
		t.Fatal("错误兑换码验证出错！")
	}
}
