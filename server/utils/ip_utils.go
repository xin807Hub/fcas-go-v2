package utils

import (
	"strconv"
	"strings"
)

func GetFullIPv6(ipv6 string) string {
	if ipv6 == "::" {
		return "0000:0000:0000:0000:0000:0000:0000:0000"
	}
	if strings.HasSuffix(ipv6, "::") {
		ipv6 += "0"
	}
	arrs := strings.Split(ipv6, ":")
	symbol := "::"
	arrleng := len(arrs)
	for arrleng < 8 {
		symbol += ":"
		arrleng++
	}
	ipv6 = strings.Replace(ipv6, "::", symbol, -1)
	fullip := ""
	for _, ip := range strings.Split(ipv6, ":") {
		for len(ip) < 4 {
			ip = "0" + ip
		}
		fullip += ip + ":"
	}
	return fullip[0 : len(fullip)-1]
}

/**
* 根据网段计算起始IP 网段格式:x.x.x.x/x
* 一个网段0一般为网络地址,255一般为广播地址.
* 起始IP计算:网段与掩码相与之后加一的IP地址
 */
func GetStartIp(segment string) string {
	startIp := strings.Builder{}
	if segment == "" {
		return segment
	}
	arr := strings.Split(segment, "/")
	ip := arr[0]
	maskIndex := arr[1]
	mask := getNetMask(maskIndex)
	if len(strings.Split(ip, ".")) != 4 || len(mask) == 0 {
		return ""
	}
	ipArray := [4]int{}
	netMaskArray := [4]int{}
	for i := 0; i < 4; i++ {
		ipNum, _ := strconv.ParseInt(strings.Split(ip, ".")[i], 0, 64)
		ipArray[i] = int(ipNum)

		maskNum, _ := strconv.ParseInt(strings.Split(mask, ".")[i], 0, 64)
		netMaskArray[i] = int(maskNum)
		if ipArray[i] > 255 || ipArray[i] < 0 || netMaskArray[i] > 255 || netMaskArray[i] < 0 {
			return ""
		}
		ipArray[i] = ipArray[i] & netMaskArray[i]
		if i == 3 {
			startIp.WriteString(strconv.Itoa(ipArray[i]))
		} else {
			startIp.WriteString(strconv.Itoa(ipArray[i]) + ".")
		}
	}
	return startIp.String()
}

func getNetMask(maskIndex string) string {
	mask := strings.Builder{}
	var inetMask int64
	inetMask, _ = strconv.ParseInt(maskIndex, 0, 64)
	if inetMask > 32 {
		return ""
	}
	// 子网掩码为1占了几个字节
	num1 := inetMask / 8
	// 子网掩码的补位位数
	num2 := inetMask % 8
	array := [4]int{}
	for i := 0; i < int(num1); i++ {
		array[i] = 255
	}
	for i := num1; i < 4; i++ {
		array[i] = 0
	}
	for i := 0; i < int(num2); num2-- {
		array[int(num1)] += 1 << (8 - int(num2))
	}

	for i := 0; i < 4; i++ {
		if i == 3 {
			mask.WriteString(strconv.Itoa(array[i]))
		} else {
			mask.WriteString(strconv.Itoa(array[i]) + ".")
		}
	}
	return mask.String()
}

/**
 * 根据网段计算结束IP
 *
 * @param segment 网段
 * @return 结束IP
 */
func GetEndIp(segment string) string {
	endIp := strings.Builder{}
	startIp := GetStartIp(segment)
	if segment == "" {
		return ""
	}
	arr := strings.Split(segment, "/")
	maskIndex := arr[1]
	hostNumber := 0
	startIpArray := [4]int{}
	maskIndexNum, _ := strconv.ParseInt(maskIndex, 0, 64)

	hostNumber = int(1 << (32 - maskIndexNum))
	for i := 0; i < 4; i++ {
		arrTemp := strings.Split(startIp, ".")
		iNum, _ := strconv.ParseInt(arrTemp[i], 0, 64)
		startIpArray[i] = int(iNum)
		if i == 3 {
			startIpArray[i] = startIpArray[i]
			break
		}
	}
	startIpArray[3] = startIpArray[3] + (hostNumber - 1)
	if startIpArray[3] > 255 {
		k := startIpArray[3] / 256
		startIpArray[3] = startIpArray[3] % 256
		startIpArray[2] = startIpArray[2] + k
	}
	if startIpArray[2] > 255 {
		j := startIpArray[2] / 256
		startIpArray[2] = startIpArray[2] % 256
		startIpArray[1] = startIpArray[1] + j
		if startIpArray[1] > 255 {
			k := startIpArray[1] / 256
			startIpArray[1] = startIpArray[1] % 256
			startIpArray[0] = startIpArray[0] + k
		}
	}
	for i := 0; i < 4; i++ {
		if i == 3 {
			startIpArray[i] = startIpArray[i]
		}
		if "" == endIp.String() || len(endIp.String()) == 0 {
			endIp.WriteString(strconv.Itoa(startIpArray[i]))
		} else {
			endIp.WriteString("." + strconv.Itoa(startIpArray[i]))
		}
	}
	return endIp.String()
}
