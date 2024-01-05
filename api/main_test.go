package api_test

import (
	"fmt"
	"os"
	"testing"
)

var (
	CTFD_URL = ""
)

func TestMain(m *testing.M) {
	u, ok := os.LookupEnv("CTFD_URL")
	if !ok {
		fmt.Println("Environment variable CTFD is not set, please indicate the domain name/IP address to reach out the cluster.")
		os.Exit(1)
	}
	CTFD_URL = u

	os.Exit(m.Run())
}

func ptr[T any](t T) *T {
	return &t
}
