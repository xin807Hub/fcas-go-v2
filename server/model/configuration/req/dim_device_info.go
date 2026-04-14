package req

type DimDeviceInfoSaveRequest struct {
	ID         int    `json:"id"`
	DeviceIp   string `json:"deviceIp"`
	DeviceName string `json:"deviceName"`
	Remark     string `json:"remark"`
	SnmpName   string `json:"snmpName"`
	UdpPort    int    `json:"udpPort"`
}
