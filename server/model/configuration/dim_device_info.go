package configuration

import (
	"fcas_server/utils"
	"fmt"
	"github.com/gosnmp/gosnmp"
	"strconv"
	"strings"
)

type DimDeviceInfo struct {
	ID            int    `json:"id" gorm:"column:id"`                        // 设备ID
	DeviceName    string `json:"deviceName" gorm:"column:device_name"`       // 设备名称
	DeviceIp      string `json:"deviceIp" gorm:"column:device_ip"`           // 设备IP
	SnmpName      string `json:"snmpName" gorm:"column:snmp_name"`           // SNMP读写团体名,community或username
	UdpPort       int    `json:"udpPort" gorm:"column:udp_port"`             // UDP端口号
	IsSnmpTrap    int    `json:"isSnmpTrap" gorm:"column:is_snmp_trap"`      // 开启SNMP Trap上报(0:关闭，1:开启)
	ReportAddress string `json:"reportAddress" gorm:"column:report_address"` // 上报服务器地址
	Remark        string `json:"remark" gorm:"column:remark"`                // 设备备注

	LinkState  int                `json:"linkState" gorm:"-"`  // 连接状态: 0-未连接, 1-已连接
	ProcessMap map[string]int     `json:"processMap" gorm:"-"` // 监控服务pid
	DistUseMap map[string][]int64 `json:"diskUseMap" gorm:"-"` // 磁盘使用情况
}

func (r *DimDeviceInfo) TableName() string {
	return "dim_device_info"
}

func (r *DimDeviceInfo) SetSnmpFields() error {
	client, err := utils.NewSNMPClient(r.DeviceIp, r.UdpPort, r.SnmpName)
	if err != nil {
		return fmt.Errorf("建立snmp连接失败，请检查snmpd服务是否正常（snmpwalk -v2c -c sinoSh %s .1.3.6.1.2.1.1）: %w", r.DeviceIp, err)
	}
	defer client.Conn.Close()

	r.LinkState = 1 // 连接状态: 1-已连接, 0-未连接

	// 获取进程信息
	if err = r.setProcessMap(client); err != nil {
		return fmt.Errorf("获取进程信息失败，请检查snmp配置是否正确（snmpget -v2c -c sinoSh %s 1.3.6.1.4.1.2021.3210.4.1.2.13.99.104.101.99.107.95.112.114.111.99.101.115.115.1）: %w", r.DeviceIp, err)
	}

	// 获取磁盘使用情况
	if err := r.setDistUseMap(client); err != nil {
		return fmt.Errorf("获取磁盘使用情况失败，请检查snmp配置是否正确（snmpwalk -v2c -c sinoSh %s 1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101）: %w", r.DeviceIp, err)
	}

	return nil
}

var processNameMap = map[string]string{
	"clickhouse-server": "大数据数据库",
	"etl":               "大数据采集程序",
	"updpi":             "流量采集程序",
	"upload":            "数据上报程序",
}

func (r *DimDeviceInfo) setProcessMap(snmpClient *gosnmp.GoSNMP) error {
	/*
		snmpget -v2c -c sinoSh localhost 1.3.6.1.4.1.2021.3210.4.1.2.13.99.104.101.99.107.95.112.114.111.99.101.115.115.1

		result:
		process walkAll result: map[.1.3.6.1.4.1.2021.3210.4.1.2.13.99.104.101.99.107.95.112.114.111.99.101.115.115.1:clickhouse-server 36847 36855, etl 39491,]
		process get result: clickhouse-server 36847 36855, etl 39491
	*/

	val, err := utils.SNMPGet[string](snmpClient, "1.3.6.1.4.1.2021.3210.4.1.2.13.99.104.101.99.107.95.112.114.111.99.101.115.115.1")
	if err != nil {
		return err
	}

	fields := strings.Split(val, ",")
	result := make(map[string]int, len(fields))
	for _, item := range fields {
		parts := strings.Fields(item)
		if len(parts) == 0 {
			continue
		}

		process := processNameMap[parts[0]]
		switch {
		case len(parts) == 1:
			result[process] = 0
		case len(parts) >= 2:
			v, _ := strconv.Atoi(parts[1])
			result[process] = v
		}
	}

	r.ProcessMap = result

	return nil
}

func (r *DimDeviceInfo) setDistUseMap(snmpClient *gosnmp.GoSNMP) error {
	/*
		snmpwalk -v2c -c sinoSh localhost 1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101

		result:
		disk walkAll result: map[
		    .1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.1:/dev/mapper/cl-root    52399108    19201148    33197960  37% /
		    .1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.10:tmpfs                  26386040           0    26386040   0% /run/user/996
		    .1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.2:devtmpfs              131918704           0   131918704   0% /dev
		    .1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.3:tmpfs                 131930184           0   131930184   0% /dev/shm
		    .1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.4:tmpfs                 131930184     3319084   128611100   3% /run
		    .1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.5:tmpfs                 131930184           0   131930184   0% /sys/fs/cgroup
		    .1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.6:/dev/sda2               1040108      132312      907796  13% /boot
		    .1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.7:/dev/sda1                204580       11464      193116   6% /boot/efi
		    .1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.8:/dev/mapper/cl-home 54634506240 39085873436 15548632804  72% /home
		    .1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.9:tmpfs                  26386040           0    26386040   0% /run/user/0
		    ]

		disk get result: nil
	*/

	val, err := utils.SNMPWalkAll[string](snmpClient, "1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101")
	if err != nil {
		return err
	}

	result := make(map[string][]int64, len(val))
	for _, item := range val {
		// 去除item中的多余空格, 并按空格分割
		fields := strings.Fields(item)
		if len(fields) < 6 {
			return fmt.Errorf("invalid disk info: %s", item)
		}

		diskName := fields[len(fields)-1]
		field1, _ := strconv.ParseInt(fields[1], 10, 64)
		field2, _ := strconv.ParseInt(fields[2], 10, 64)
		usage := []int64{field1, field2}
		result[diskName] = usage
	}

	r.DistUseMap = result

	return nil
}
