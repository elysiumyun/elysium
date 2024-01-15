package usage

import (
	"fmt"
	"strings"

	"github.com/elysiumyun/elysium/pkg/info"
)

// app usage menual
func Usage(extraUsage string) {
	fmt.Printf(`Usage: %s-%s -[hv]

     ------- < Commands Arguments > -------
argument:
%s
options:
  -h, help          Show help message. 
  -v, version       Show the app version.
For more help, you can use '%s help' for the detailed information
or you can check the docs: https://github.com/elysiumyun/elysium.git  
`, info.Prefix, info.MicroService, extraUsage, strings.Join([]string{info.Prefix, info.MicroService}, "-"))
}
