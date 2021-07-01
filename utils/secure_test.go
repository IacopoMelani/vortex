package utils

import "testing"

func TestSecureRandomString(t *testing.T) {

	length := uint(64)

	first, err := SecureRandomString(length)
	if err != nil {
		t.Fatal(err)
	}

	if first == "" {
		t.Fatalf("Empty string, length %d", length)
	}

	second, err := SecureRandomString(length)
	if err != nil {
		t.Fatal(err)
	}

	if second == "" {
		t.Fatalf("Empty string, length %d", length)
	}

	if first == second {
		t.Fatal("Equals values!!!")
	}

	third, err := SecureRandomString(0)
	if err != nil {
		t.Fatal(err)
	}

	if third != "" {
		t.Fatal("Not empty string")
	}
}
