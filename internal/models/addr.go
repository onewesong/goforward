package models

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Addr struct {
	IP   net.IP
	Port int
}

func (addr Addr) IsIPv4() bool {
	return len(addr.IP.To4()) == net.IPv4len
}

func (addr Addr) Network() string {
	if addr.IsIPv4() {
		return "tcp4"
	}
	return "tcp6"
}

func (addr Addr) String() string {
	if addr.IsIPv4() {
		return fmt.Sprintf("%s:%d", addr.IP, addr.Port)
	}
	return fmt.Sprintf("[%s]:%d", addr.IP, addr.Port)
}

func ParseAddr(address string) (addr Addr, err error) {
	index := strings.LastIndex(address, ":")
	if index == -1 {
		err = fmt.Errorf("invalid addr: %s", address)
		return
	}
	addr.IP = net.ParseIP(strings.Trim(address[:index], "[]"))
	if addr.IP == nil {
		err = fmt.Errorf("invalid ip: %s", address[:index])
		return
	}
	addr.Port, err = strconv.Atoi(address[index+1:])
	if err != nil {
		err = fmt.Errorf("invalid port: %s", address[index+1:])
		return
	}
	if addr.Port < 0 || addr.Port > 65535 {
		err = fmt.Errorf("invalid port range: %s", address[index+1:])
		return
	}
	return
}
