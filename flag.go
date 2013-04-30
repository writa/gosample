package gosample

import (
	"flag"
	"fmt"
)

var intCode *int = flag.Int("int", 0, "int code")
var strCode *string = flag.String("str", "default", "string code")
var boolCode *bool = flag.Bool("bool", false, "bool code")

func Flag() {
	flag.Parse()
	fmt.Println("intCode has value", *intCode)
	fmt.Println("strCode has value", *strCode)
	fmt.Println("boolCode has value", *boolCode)
}
