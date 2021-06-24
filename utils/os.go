package utils

import (
	"net"
)

// GetPrimaryIP - Returns primary host IP
func GetPrimaryIP() (net.IP, error) {

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}
