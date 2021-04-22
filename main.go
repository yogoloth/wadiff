package main

import (
	"flag"

	"fmt"
	"os"
)

var args_is_str bool

func main() {
	flag.BoolVar(&args_is_str, "s", false, "args as string")
	flag.Parse()

	//fmt.Println("args_is_str:", args_is_str)

	others := flag.Args()
	//fmt.Println("others:", others)

	if len(others) != 2 {
		fmt.Println("params error: wadiff -s str1 str2 or wadiff file1 file2")
		os.Exit(1)
	}
	str_x := others[0]
	str_y := others[1]

	if args_is_str == true {

		routex, routey, route_size := GetDiff(str_x, str_y)

		PrintDiff(routex, routey, route_size)
		PrintDiff(routey, routex, route_size)
	} else {
		DiffFile(str_x, str_y)
	}
}
