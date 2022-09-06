package utils

import (
	"github.com/dongri/phonenumber"
)

// ParseMobile ...
func ParseMobile(mobile string) string {
	return phonenumber.Parse(mobile, "KE")
}
