package addr_set

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

const length = 120

// TestAddMultipleAddr 测试添加地址
func TestAddMultipleAddr(t *testing.T) {
	as := NewAddrSet()

	// 添加单个地址：192.168.1.3, 192.168.1.8
	addrs := []string{"192.168.1.3", "192.168.1.8"}
	fmt.Printf("AddAddr: %v\nResult: %v\nTreeSetSize: %d\nTreeSetValues: %v\n%s\n", addrs, as.AddMultipleAddr(addrs), as.set.Size(), as.set.Values(), strings.Repeat("*", length))

	// 添加范围段：192.168.1.1-192.168.1.15，验证分割成多个子段
	addrs = []string{"192.168.1.1,192.168.1.15", "192.168.0.0/16"}
	fmt.Printf("AddAddr: %v\nResult: %v\nTreeSetSize: %d\nTreeSetValues: %v\n%s\n", addrs, as.AddMultipleAddr(addrs), as.set.Size(), as.set.Values(), strings.Repeat("*", length))

	// 验证IPv6类型
	addrs = []string{"2001:db8::1", "200:db8::1", "2001:db8::5,2001:db8::10", "2001:db8::/120"}
	fmt.Printf("AddAddr: %v\nResult: %v\nTreeSetSize: %d\nTreeSetValues: %v\n%s\n", addrs, as.AddMultipleAddr(addrs), as.set.Size(), as.set.Values(), strings.Repeat("*", length))

}

func TestRemove(t *testing.T) {
	as := NewAddrSet()

	addrs := []string{"1.1.1.9,1.1.1.15"}
	result1 := as.AddMultipleAddr(addrs)
	fmt.Printf("AddAddr: %v\nResult: %v\nTreeSetSize: %d\nTreeSetValues: %v\n%s\n", addrs, result1, as.set.Size(), as.set.Values(), strings.Repeat("*", length))

	addrs = []string{"1.1.1.4,1.1.1.15"}
	fmt.Printf("AddAddr: %v\nResult: %v\nTreeSetSize: %d\nTreeSetValues: %v\n%s\n", addrs, as.AddMultipleAddr(addrs), as.set.Size(), as.set.Values(), strings.Repeat("*", length))

	addrs = []string{"1.1.1.40"}
	fmt.Printf("AddAddr: %v\nResult: %v\nTreeSetSize: %d\nTreeSetValues: %v\n%s\n", addrs, as.AddMultipleAddr(addrs), as.set.Size(), as.set.Values(), strings.Repeat("*", length))

	addrs = []string{"1.1.1.1,1.1.1.20"}
	fmt.Printf("AddAddr: %v\nResult: %v\nTreeSetSize: %d\nTreeSetValues: %v\n%s\n", addrs, as.AddMultipleAddr(addrs), as.set.Size(), as.set.Values(), strings.Repeat("*", length))

	fmt.Println(as.set.Values())

	removeAddr := []string{"1.1.1.1,1.1.1.3", "1.1.1.9-1.1.1.15"}
	as.Remove(removeAddr...)

	fmt.Println(as.set.Values())

}

// TestIsContainsMultipleAddr 测试是否包含地址
func TestIsContainsMultipleAddr(t *testing.T) {
	as := NewAddrSet()

	// 首先添加一些地址到 AddrSet
	initialAddrs := []string{
		"10.0.0.1,10.0.0.100",
		"172.16.0.0/16",
		"2001:db8::/32",
		"2001:db8:1::1,2001:db8:1::ffff",
	}
	as.AddMultipleAddr(initialAddrs)

	fmt.Println(as.IsContainsMultipleAddr([]string{"10.0.0.50", "10.0.0.150"}))
	fmt.Println(as.IsContainsMultipleAddr([]string{"10.0.0.20,10.0.0.30"}))
	fmt.Println(as.IsContainsMultipleAddr([]string{"172.16.5.5"}))
	fmt.Println(as.IsContainsMultipleAddr([]string{"172.17.0.1"}))
	fmt.Println(as.IsContainsMultipleAddr([]string{"2001:db8::1"}))
	fmt.Println(as.IsContainsMultipleAddr([]string{"2001:db8:1::100"}))
	fmt.Println(as.IsContainsMultipleAddr([]string{"8.8.8.8,9.9.9.9"}))

}

// TestCountAvailableAddr 测试计算可用的地址数量
func TestCountAvailableAddr(t *testing.T) {
	fmt.Println("CountAvailableAddr:", CountAvailableAddr([]string{"192.168.1.10"}))                              // 1
	fmt.Println("CountAvailableAddr:", CountAvailableAddr([]string{"192.168.1.30,192.168.1.40"}))                 // 11
	fmt.Println("CountAvailableAddr:", CountAvailableAddr([]string{"192.168.1.10", "192.168.1.30,192.168.1.40"})) // 12
	fmt.Println("CountAvailableAddr:", CountAvailableAddr([]string{"2401:800:9000:21::/64"}))                     // 18446744073709551616
	fmt.Println("CountAvailableAddr:", CountAvailableAddr([]string{"2001:db8:44::,2001:db8:65::"}))               // 39894552047282762765303809
}

// TestMergeMultipleAddr 测试合并多个地址段
func TestMergeMultipleAddr(t *testing.T) {
	fmt.Println(MergeMultipleAddr([]string{
		"192.168.1.1",
		"192.168.1.10,192.168.1.20",
		"192.168.1.15",
		"192.168.1.20",
	}))

	fmt.Println(MergeMultipleAddr([]string{
		"2001:db8::1",
		"200:db8::1",
		"2001:db8::5-2001:db8::10",
		"2001:db8::/120",
	}))

	fmt.Println(MergeMultipleAddr([]string{
		"2001:db8::/64",
		"2001:db8:0:1::/64",
	}))
}

func generateRandomIPv4() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func TestSplitRangePerformance(t *testing.T) {
	rand.NewSource(time.Now().UnixNano())

	testCases := []struct {
		name    string
		setSize int
	}{
		{"小集合", 100000},
	}

	for _, tc := range testCases {
		as := NewAddrSet()
		var originalTotalDuration time.Duration

		// 填充集合
		for i := 0; i < tc.setSize; i++ {
			startIp := generateRandomIPv4()
			endIp := generateRandomIPv4()
			start := time.Now()
			as.AddMultipleAddr([]string{fmt.Sprintf("%s,%s", startIp, endIp)})
			originalTotalDuration += time.Since(start)
		}

		originalAvgDuration := originalTotalDuration / time.Duration(tc.setSize)

		t.Logf("集合大小: %d", tc.setSize)
		t.Logf("平均耗时: %v", originalAvgDuration)
		t.Logf("总耗时: %v", originalTotalDuration)
	}
}

func TestUpdate(t *testing.T) {
	as := NewAddrSet()

	// 首先添加一些地址到 AddrSet
	initialAddrs := []string{
		"1.1.1.1-1.1.1.10",
		"1.1.1.15",
		"10.0.0.1,10.0.0.100",
		"172.16.0.0/16",
		"2001:db8::/32",
		"2001:db8:1::1,2001:db8:1::ffff",
	}
	if _, _, err := as.Add(initialAddrs); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(as.set.Values())

	if _, _, err := as.Update([]string{"1.1.1.1-1.1.1.10"}, []string{"1.1.1.15"}); err != nil {
		fmt.Println(err)
	}

	fmt.Println(as.set.Values())

}

func TestIsValidAddr(t *testing.T) {
	fmt.Println(parseAddr("192.168.1.1s3"))
	fmt.Println(parseAddr("1"))
	fmt.Println(parseAddr("192.168.1.f1f,192.x168.1.10"))
	fmt.Println(parseAddr("192.168.1.1-192.168.1.10"))
	fmt.Println(parseAddr("2001:dbfdf8::1"))
	//orRange := parseAddr("2001:db8::1,2001:db8::2")
	//orRange.GetCount()

	//fmt.Println(ipaddr.NewIPAddressString("192.168.1.1s3").GetAddress())
	//fmt.Println(ipaddr.NewIPAddressString("192.168.1.f1f").GetAddress())
	//fmt.Println(ipaddr.NewIPAddressString("192.168.1.f1f/24").GetAddress())
}
