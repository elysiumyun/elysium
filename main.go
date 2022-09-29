package elysium

import "fmt"

//go:generate ./scripts/generate_version.sh
func New() {
	fmt.Println("init elysium")
}
