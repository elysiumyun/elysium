package version

import (
	"fmt"
	"strings"
)

// framework version
type version struct {
	ver string
	upt string
}

func (v *version) ToString() string {
	ap := strings.Join([]string{v.ver, v.upt}, ".")
	return ap
}

func (v *version) Print() {
	fmt.Printf("Version: %s \n", v.ver)
	fmt.Printf("Update Time: %s \n", v.upt)
}

func GetVersion() *version {
	return &version{
		ver: verStr,
		upt: uptStr,
	}
}
