package configuration

import (
	"fcas_server/utils"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	val := "clickhouse-server 36847 36855, etl 39491"
	fields := strings.Split(val, ",")

	result := make(map[string]int, len(fields))

	for _, item := range fields {
		parts := strings.Fields(item)
		switch {
		case len(parts) == 1:
			result[parts[0]] = 0
		case len(parts) >= 2:
			v, _ := strconv.Atoi(parts[1])
			result[parts[0]] = v
		}
	}

	fmt.Println(result)
}

func TestDistUse(t *testing.T) {
	val := map[string]string{
		".1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.1":  "/dev/mapper/cl-root    52399108    19201148    33197960  37% / ",
		".1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.10": "tmpfs                  26386040           0    26386040   0% /run/user/996 ",
		".1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.2":  "devtmpfs              131918704           0   131918704   0% /dev ",
		".1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.3":  "tmpfs                 131930184           0   131930184   0% /dev/shm ",
		".1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.4":  "tmpfs                 131930184     3319084   128611100   3% /run ",
		".1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.5":  "tmpfs                 131930184           0   131930184   0% /sys/fs/cgroup ",
		".1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.6":  "/dev/sda2               1040108      132312      907796  13% /boot ",
		".1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.7":  "/dev/sda1                204580       11464      193116   6% /boot/efi ",
		".1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.8":  "/dev/mapper/cl-home 54634506240 39085873436 15548632804  72% /home ",
		".1.3.6.1.4.1.2021.4321.4.1.2.8.100.105.115.107.95.117.115.101.9":  "tmpfs                  26386040           0    26386040   0% /run/user/0",
	}

	result := make(map[string][]int64, len(val))
	for _, item := range val {
		// 去除item中的多余空格, 并按空格分割
		fields := strings.Fields(item)
		if len(fields) < 6 {
			fmt.Println("Invalid format:", item)
			return
		}

		diskName := fields[len(fields)-1]
		field1, _ := strconv.ParseInt(fields[1], 10, 64)
		field2, _ := strconv.ParseInt(fields[2], 10, 64)
		usage := []int64{field1, field2}
		result[diskName] = usage
	}

	fmt.Println(result)

}

func TestGetProcess(t *testing.T) {

	d := DimDeviceInfo{
		DeviceIp: "192.168.4.146",
		SnmpName: "sinoSh",
		UdpPort:  161,
	}

	client, err := utils.NewSNMPClient(d.DeviceIp, d.UdpPort, d.SnmpName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Conn.Close()

	// 获取进程信息
	if err := d.setProcessMap(client); err != nil {
		fmt.Println(err)
		return
	}

}
