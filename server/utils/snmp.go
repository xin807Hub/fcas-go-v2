package utils

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"time"
)

const (
	version   = gosnmp.Version2c
	timeout   = time.Duration(1 * time.Second)
	systemOID = ".1.3.6.1.2.1.1"
)

// NewSNMPClient 用于创建一个新的 SNMP 客户端
func NewSNMPClient(target string, port int, community string) (*gosnmp.GoSNMP, error) {
	client := &gosnmp.GoSNMP{
		Target:    target,
		Port:      uint16(port),
		Community: community,
		Version:   version,
		Timeout:   timeout,
		Retries:   1,
		Transport: "udp",
	}

	if err := client.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect to SNMP target %s: %v", target, err)
	}

	if _, err := client.WalkAll(systemOID); err != nil {
		return nil, fmt.Errorf("failed to connect, get system OID from target %s failed: %v", target, err)
	}

	return client, nil
}

// SNMPGet 根据OID获取值
func SNMPGet[T any](client *gosnmp.GoSNMP, oid string) (val T, err error) {

	result, err := client.Get([]string{oid})
	if err != nil {
		return val, fmt.Errorf("failed to get value of OID %s: %v", oid, err)
	}

	if len(result.Variables) == 0 {
		return val, fmt.Errorf("no value returned for OID %s", oid)
	}

	// 类型转换
	val, err = convertPduVal[T](result.Variables[0])
	if err != nil {
		return val, fmt.Errorf("failed to convert snmpPdu value of OID %s: %v", oid, err)
	}

	return val, nil
}

// SNMPWalkAll 遍历所有OID
func SNMPWalkAll[T any](client *gosnmp.GoSNMP, oid string) (map[string]T, error) {

	results, err := client.WalkAll(oid)
	if err != nil {
		return nil, fmt.Errorf("failed to walk OID %s: %v", oid, err)
	}

	resultMap := make(map[string]T, len(results))
	for _, pdu := range results {
		// 类型转换
		val, err := convertPduVal[T](pdu)
		if err != nil {
			return nil, fmt.Errorf("failed to convert snmpPdu value of OID %s: %v", pdu.Name, err)
		}
		resultMap[pdu.Name] = val
	}

	return resultMap, nil
}

// 转换值类型
func convertPduVal[T any](pdu gosnmp.SnmpPDU) (result T, err error) {
	var val any

	switch pdu.Type {
	case gosnmp.Integer, gosnmp.Counter32, gosnmp.Counter64, gosnmp.Gauge32, gosnmp.TimeTicks:
		val = pdu.Value
	case gosnmp.OctetString:
		val = string(pdu.Value.([]byte))
	default:
		return result, fmt.Errorf("unsupported value type %d for OID %s", pdu.Type, pdu.Name)
	}

	result, ok := val.(T)
	if !ok {
		return result, fmt.Errorf("value type %T is not convertible with %T", val, result)
	}

	return result, nil
}
