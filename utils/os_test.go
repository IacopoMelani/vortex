package utils

import "testing"

func TestGetPrimaryIp(t *testing.T) {

	ip, err := GetPrimaryIP()
	if err != nil {
		t.Fatal(err)
	}

	if ip.String() == "" {
		t.Fatal("Invalid ip addr")
	}
}
