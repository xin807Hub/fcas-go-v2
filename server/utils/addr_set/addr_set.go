package addr_set

import (
	"fmt"
	"github.com/emirpasic/gods/sets/treeset"
	"github.com/seancfoley/ipaddress-go/ipaddr"
	"math/big"
	"sort"
	"strings"
)

type AddrSet struct {
	set *treeset.Set
}

func NewAddrSet() *AddrSet {

	treeset.NewWithStringComparator()

	return &AddrSet{
		set: treeset.NewWith(func(a, b interface{}) int {
			rangeA := a.(*ipaddr.IPAddressSeqRange)
			rangeB := b.(*ipaddr.IPAddressSeqRange)
			return rangeA.GetLower().Compare(rangeB.GetLower())
		}),
	}
}

func (as *AddrSet) Add(addrs []string) ([]string, string, error) {

	// 合并相同子网的CIDR表达式和处理包含关系
	mergedAddrs := MergeMultipleAddr(addrs)

	// 判断是否已存在
	existedAddr, is := as.IsContainsMultipleAddr(mergedAddrs)
	if is {
		return nil, "", fmt.Errorf("地址 %s 已存在", existedAddr)
	}

	// 添加IP地址
	resultAddrs := as.AddMultipleAddr(mergedAddrs)

	// 计算可用的IP地址数量
	countAvailableAddr := CountAvailableAddr(resultAddrs)

	return resultAddrs, countAvailableAddr, nil
}

// AddMultipleAddr 添加多个IP地址
func (as *AddrSet) AddMultipleAddr(addrs []string) []string {
	ranges := make([]*ipaddr.IPAddressSeqRange, 0, len(addrs))

	for _, addr := range addrs {
		seqRange := parseAddr(addr)
		splitRanges := as.splitRange(seqRange)
		as.updateAddedRanges(splitRanges)
		ranges = append(ranges, splitRanges...)
	}

	resultSplitAddr := make([]string, 0, len(ranges))
	for _, rng := range ranges {
		resultSplitAddr = append(resultSplitAddr, customString(rng))
	}

	return resultSplitAddr
}

// 分割范围并跳过已添加的地址
func (as *AddrSet) splitRange(seqRange *ipaddr.IPAddressSeqRange) []*ipaddr.IPAddressSeqRange {
	var result []*ipaddr.IPAddressSeqRange
	currentRngLower := seqRange.GetLower()

	it := as.set.Iterator()
	for it.Next() {
		existedRange, ok := it.Value().(*ipaddr.IPAddressSeqRange)
		if !ok {
			continue
		}

		if existedRange.Overlaps(seqRange) {
			lower := existedRange.GetLower()
			upper := existedRange.GetUpper()

			if currentRngLower.Compare(lower) < 0 {
				result = as.addRangeOrSingleAddr(result, currentRngLower, lower.Increment(-1))
			}

			currentRngLower = upper.Increment(1)
		}
	}

	if currentRngLower.Compare(seqRange.GetUpper()) <= 0 {
		result = as.addRangeOrSingleAddr(result, currentRngLower, seqRange.GetUpper())
	}

	return result
}

// 添加单个IP地址或范围段到结果列表中
func (as *AddrSet) addRangeOrSingleAddr(result []*ipaddr.IPAddressSeqRange, start, end *ipaddr.IPAddress) []*ipaddr.IPAddressSeqRange {
	return append(result, ipaddr.NewSequentialRange(start, end))
}

// 更新已添加的地址范围
func (as *AddrSet) updateAddedRanges(seqRanges []*ipaddr.IPAddressSeqRange) {
	for _, item := range seqRanges {
		as.set.Add(item)
	}
}

// IsContainsMultipleAddr 判断是否包含指定IP地址或范围段或CIDR格式的IP地址范围段
func (as *AddrSet) IsContainsMultipleAddr(addrs []string) (string, bool) {
	for _, addr := range addrs {
		if as.isContainsSingleAddr(addr) {
			return addr, true
		}
	}
	return "", false
}

// 判断是否包含指定IP地址或范围段或CIDR格式的IP地址范围段
func (as *AddrSet) isContainsSingleAddr(addr string) bool {
	it := as.set.Iterator()
	for it.Next() {
		existedRange, ok := it.Value().(*ipaddr.IPAddressSeqRange)
		if !ok {
			continue
		}
		seqRange := parseAddr(addr)

		if rangeContains(existedRange, seqRange) {
			return true
		}
	}
	return false
}

// Remove 移除IP地址或范围段或CIDR格式的IP地址范围段
func (as *AddrSet) Remove(addrs ...string) {
	for _, addr := range addrs {
		seqRng := parseAddr(addr)
		as.set.Remove(seqRng)
	}
}

// Duplicate 复制AddrSet
func (as *AddrSet) Duplicate() *AddrSet {
	newSet := NewAddrSet()
	it := as.set.Iterator()
	for it.Next() {
		newSet.set.Add(it.Value())
	}
	return newSet
}

// Update 更新IP地址或范围段或CIDR格式的IP地址范围段
// 先移除旧的IP地址或范围段或CIDR格式的IP地址范围段，再添加新的IP地址或范围段或CIDR格式的IP地址范围段
func (as *AddrSet) Update(oldAddrs, newAddrs []string) ([]string, string, error) {
	oldAs := as.Duplicate()

	as.Remove(oldAddrs...)
	addrs, addrNum, err := as.Add(newAddrs)
	if err != nil {
		as.set = oldAs.set // 回滚
		return nil, "", err
	}

	return addrs, addrNum, nil
}

func (as *AddrSet) Size() int {
	return as.set.Size()
}

func (as *AddrSet) String() string {

	result := make([]string, 0, as.set.Size())
	it := as.set.Iterator()
	for it.Next() {
		result = append(result, customString(it.Value().(*ipaddr.IPAddressSeqRange)))
	}

	return strings.Join(result, ", ")
}

// CountAvailableAddr 计算可用的IP地址数量
func CountAvailableAddr(addrs []string) string {
	count := big.NewInt(0)
	for _, addr := range addrs {
		seqRange := parseAddr(addr)
		count.Add(count, seqRange.GetCount())
	}
	return count.String()
}

// MergeMultipleAddr 合并多个IP地址段
func MergeMultipleAddr(addrs []string) []string {
	ranges := make([]*ipaddr.IPAddressSeqRange, 0, len(addrs))
	// 解析输入列表
	for _, src := range addrs {
		addressSeqRange := parseAddr(src)
		ranges = append(ranges, addressSeqRange)
	}

	// 排序ranges
	//fmt.Println("排序前", ranges)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].GetLower().Compare(ranges[j].GetLower()) < 0
	})
	//fmt.Println("排序后", ranges)

	// 归并相同子网的CIDR表达式和处理包含关系
	mergedRanges := mergeRanges(ranges)

	// 输出结果
	result := make([]string, 0, len(mergedRanges))
	for _, r := range mergedRanges {
		result = append(result, customString(r))
	}
	return result
}

// 归并用户传入的多个IP地址段（如果有包含关系）
func mergeRanges(ranges []*ipaddr.IPAddressSeqRange) []*ipaddr.IPAddressSeqRange {
	mergedRanges := make([]*ipaddr.IPAddressSeqRange, 0)

	for _, currentRange := range ranges {
		merged := false
		for i, existingRange := range mergedRanges {
			if rangeContains(existingRange, currentRange) {
				merged = true
				break
			} else if rangeContains(currentRange, existingRange) {
				mergedRanges[i] = currentRange
				merged = true
				break
			} else if existingRange.Overlaps(currentRange) {
				// 合并重叠的范围
				lower := existingRange.GetLower()
				if currentRange.GetLower().Compare(lower) < 0 {
					lower = currentRange.GetLower()
				}
				upper := existingRange.GetUpper()
				if currentRange.GetUpper().Compare(upper) > 0 {
					upper = currentRange.GetUpper()
				}

				mergedRanges[i] = ipaddr.NewSequentialRange(lower, upper)
				merged = true
				break
			}
		}
		if !merged {
			mergedRanges = append(mergedRanges, currentRange)
		}
	}

	return mergedRanges
}

// 自定义字符串输出
func customString(rng *ipaddr.IPAddressSeqRange) string {
	l := rng.GetLower()
	u := rng.GetUpper()
	if l == u {
		return l.String()
	}
	return fmt.Sprintf("%s-%s", l.String(), u.String())
}

// parseAddr 解析地址字符串并返回IP地址范围段, 若解析失败则返回nil
func parseAddr(addr string) *ipaddr.IPAddressSeqRange {

	switch {
	case strings.Contains(addr, "/"):
		// 处理CIDR表达式
		address := ipaddr.NewIPAddressString(addr).GetAddress().ToPrefixBlock()
		if address == nil {
			return nil
		}
		return address.ToSequentialRange()

	case strings.Contains(addr, ","), strings.Contains(addr, "-"):
		// 处理IPv6地址范围段
		parts := strings.FieldsFunc(addr, func(r rune) bool {
			return r == ',' || r == '-'
		})
		start := ipaddr.NewIPAddressString(strings.TrimSpace(parts[0])).GetAddress()
		end := ipaddr.NewIPAddressString(strings.TrimSpace(parts[1])).GetAddress()
		if start == nil || end == nil {
			return nil
		}
		return ipaddr.NewSequentialRange(start, end)

	default:
		// 处理单个地址
		address := ipaddr.NewIPAddressString(addr).GetAddress()
		if address == nil {
			return nil
		}
		return ipaddr.NewSequentialRange(address, address)
	}
}

// 判断range1是否包含range2
func rangeContains(rng1, rng2 *ipaddr.IPAddressSeqRange) bool {
	return rng1.GetLower().Compare(rng2.GetLower()) <= 0 &&
		rng1.GetUpper().Compare(rng2.GetUpper()) >= 0
}

// ValidateAddr 校验IP地址是否合法
func ValidateAddr(addrs ...string) error {
	for _, addr := range addrs {
		if parseAddr(addr) == nil {
			return fmt.Errorf("地址 %s 格式错误", addr)
		}
	}
	return nil
}
