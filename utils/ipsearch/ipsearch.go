package ipsearch

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

/**
 * @author xiao.luo
 * @description This is the go version for IpSearch
 */

type ipIndex struct {
	startip, endip           uint32
	localOffset, localLength uint32
}

type prefixIndex struct {
	startIndex, endIndex uint32
}

type ipSearch struct {
	data               []byte
	prefixMap          map[uint32]prefixIndex
	firstStartIpOffset uint32
	prefixStartOffset  uint32
	prefixEndOffset    uint32
	prefixCount        uint32
}

type Location struct {
	Continent      string  `json:"continent,omitempty"`
	Country        string  `json:"country,omitempty"`
	Province       string  `json:"province,omitempty"`
	City           string  `json:"city,omitempty"`
	District       string  `json:"district,omitempty"`
	Isp            string  `json:"isp,omitempty"`
	AreaCode       string  `json:"area_code,omitempty"`
	CountryEnglish string  `json:"country_english,omitempty"`
	CountryCode    string  `json:"country_code,omitempty"`
	Longitude      float64 `json:"longitude,omitempty"`
	Latitude       float64 `json:"latitude,omitempty"`
}

var ips *ipSearch = nil

func New() (ipSearch, error) {
	if ips == nil {
		var err error
		ips, err = loadIpDat()
		if err != nil {
			log.Fatal("the IP Dat loaded failed!")
			return *ips, err
		}
	}
	return *ips, nil
}

func loadIpDat() (*ipSearch, error) {

	p := ipSearch{}
	//加载ip地址库信息
	pwd, _ := os.Getwd()
	data, err := ioutil.ReadFile(pwd + "/utils/ipsearch/qqzeng-ip-utf8.dat")
	if err != nil {
		log.Fatal(err)
	}
	p.data = data
	p.prefixMap = make(map[uint32]prefixIndex)

	p.firstStartIpOffset = bytesToLong(data[0], data[1], data[2], data[3])
	p.prefixStartOffset = bytesToLong(data[8], data[9], data[10], data[11])
	p.prefixEndOffset = bytesToLong(data[12], data[13], data[14], data[15])
	p.prefixCount = (p.prefixEndOffset-p.prefixStartOffset)/9 + 1 // 前缀区块每组

	// 初始化前缀对应索引区区间
	indexBuffer := p.data[p.prefixStartOffset:(p.prefixEndOffset + 9)]
	for k := uint32(0); k < p.prefixCount; k++ {
		i := k * 9
		prefix := uint32(indexBuffer[i] & 0xFF)

		pf := prefixIndex{}
		pf.startIndex = bytesToLong(indexBuffer[i+1], indexBuffer[i+2], indexBuffer[i+3], indexBuffer[i+4])
		pf.endIndex = bytesToLong(indexBuffer[i+5], indexBuffer[i+6], indexBuffer[i+7], indexBuffer[i+8])
		p.prefixMap[prefix] = pf

	}
	return &p, nil
}

func (p ipSearch) Get(ip string) string {
	ips := strings.Split(ip, ".")
	x, _ := strconv.Atoi(ips[0])
	prefix := uint32(x)
	intIP := ipToLong(ip)

	var high uint32 = 0
	var low uint32 = 0

	if _, ok := p.prefixMap[prefix]; ok {
		low = p.prefixMap[prefix].startIndex
		high = p.prefixMap[prefix].endIndex
	} else {
		return ""
	}

	var myIndex uint32
	if low == high {
		myIndex = low
	} else {
		myIndex = p.binarySearch(low, high, intIP)
	}

	ipindex := ipIndex{}
	ipindex.getIndex(myIndex, &p)

	if ipindex.startip <= intIP && ipindex.endip >= intIP {
		return ipindex.getLocal(&p)
	} else {
		return ""
	}
}

func (p ipSearch) GetLocation(ip string) Location {
	ipstr := p.Get(ip)
	locationSclie := strings.Split(ipstr, "|")
	location := Location{
		Continent:      locationSclie[0],
		Country:        locationSclie[1],
		Province:       locationSclie[2],
		City:           locationSclie[3],
		District:       locationSclie[4],
		Isp:            locationSclie[5],
		AreaCode:       locationSclie[6],
		CountryEnglish: locationSclie[7],
		CountryCode:    locationSclie[8],
	}
	if locationSclie[9] != "" {
		location.Longitude, _ = strconv.ParseFloat(locationSclie[9], 64)
	}
	if locationSclie[10] != "" {
		location.Latitude, _ = strconv.ParseFloat(locationSclie[10], 64)
	}

	return location

}

// 二分逼近算法
func (p ipSearch) binarySearch(low uint32, high uint32, k uint32) uint32 {
	var M uint32 = 0
	for low <= high {
		mid := (low + high) / 2

		endipNum := p.getEndIp(mid)
		if endipNum >= k {
			M = mid
			if mid == 0 {
				break // 防止溢出
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return M
}

// 只获取结束ip的数值
// 索引区第left个索引
// 返回结束ip的数值
func (p ipSearch) getEndIp(left uint32) uint32 {
	leftOffset := p.firstStartIpOffset + left*12
	return bytesToLong(p.data[4+leftOffset], p.data[5+leftOffset], p.data[6+leftOffset], p.data[7+leftOffset])

}

func (p *ipIndex) getIndex(left uint32, ips *ipSearch) {
	leftOffset := ips.firstStartIpOffset + left*12
	p.startip = bytesToLong(ips.data[leftOffset], ips.data[1+leftOffset], ips.data[2+leftOffset], ips.data[3+leftOffset])
	p.endip = bytesToLong(ips.data[4+leftOffset], ips.data[5+leftOffset], ips.data[6+leftOffset], ips.data[7+leftOffset])
	p.localOffset = bytesToLong3(ips.data[8+leftOffset], ips.data[9+leftOffset], ips.data[10+leftOffset])
	p.localLength = uint32(ips.data[11+leftOffset])
}

// / 返回地址信息
// / 地址信息的流位置
// / 地址信息的流长度
func (p *ipIndex) getLocal(ips *ipSearch) string {
	bytes := ips.data[p.localOffset : p.localOffset+p.localLength]
	return string(bytes)

}

func ipToLong(ip string) uint32 {
	quads := strings.Split(ip, ".")
	var result uint32 = 0
	a, _ := strconv.Atoi(quads[3])
	result += uint32(a)
	b, _ := strconv.Atoi(quads[2])
	result += uint32(b) << 8
	c, _ := strconv.Atoi(quads[1])
	result += uint32(c) << 16
	d, _ := strconv.Atoi(quads[0])
	result += uint32(d) << 24
	return result
}

//字节转整形
func bytesToLong(a, b, c, d byte) uint32 {
	a1 := uint32(a)
	b1 := uint32(b)
	c1 := uint32(c)
	d1 := uint32(d)
	return (a1 & 0xFF) | ((b1 << 8) & 0xFF00) | ((c1 << 16) & 0xFF0000) | ((d1 << 24) & 0xFF000000)
}

func bytesToLong3(a, b, c byte) uint32 {
	a1 := uint32(a)
	b1 := uint32(b)
	c1 := uint32(c)
	return (a1 & 0xFF) | ((b1 << 8) & 0xFF00) | ((c1 << 16) & 0xFF0000)

}
