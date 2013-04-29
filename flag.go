package gosample

import (
	"flag"
	"fmt"
)

var code *int = flag.Int("testcode", 716, "give me your codes")

func Flag() {
	flag.Parse()
	fmt.Println("testcode has value", *code)
}
