package ipsearch

import (
	"fmt"
	"testing"
)

/**
 * @author xiao.luo
 * @description This is the unit test for IpSearch
 */

func TestLoad(t *testing.T) {
	fmt.Println("Test Load IP Dat ...")
	p, err := New()
	if len(p.data) <= 0 || err != nil {
		t.Fatal("the IP Dat did not loaded successfully!")
	}
}

func TestGet(t *testing.T) {
	fmt.Println("Test Get IP ...")
	p, _ := New()
	ip := "210.51.200.123"
	ipstr := p.Get(ip)
	fmt.Println(ipstr)
	if ipstr != `亚洲|中国|江苏|镇江||联通/基站WiFi|321100|China|CN|119.452753|32.204402` {
		t.Fatal("the IP convert by ipSearch component is not correct!")
	}
}

func TestGetLocation(t *testing.T) {
	fmt.Println("Test Get IP ...")
	p, _ := New()
	ip := "210.51.200.123"
	location := p.GetLocation(ip)
	fmt.Printf("%+v\n", location)
	ipstr := fmt.Sprintf("%s|%s|%s|%s",
		location.Country,
		location.Province,
		location.City,
		location.CountryCode,
	)
	fmt.Println(ipstr)
	if ipstr != `中国|江苏|镇江|CN` {
		t.Fatal("the IP convert by ipSearch component is not correct!")
	}

}
