package main

import (
	"fmt"
	_ "github.com/cweill/gotests"
	"github.com/jefferyjob/go_unittest_demo/unit_base"
)

func main() {
	ret := unit_base.Factorial(5)
	fmt.Println(ret)
}
