package utils

import (
	"strconv"
	"strings"
	"time"
)

func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}

func ComputeSpeed(speed int) string {
	if speed >= 0 && speed < 1024 {
		return strconv.Itoa(speed) + "(kbps)"
	} else if speed >= 1024 && speed < 1048576 {
		return strconv.Itoa(speed/1024) + "(Mbps)"
	} else if speed >= 1048576 && speed < 1073741824 {
		return strconv.Itoa(speed/1048576) + "(Gbps)"
	} else {
		return strconv.Itoa(speed/1073741824) + "(Tbps)"
	}
}
