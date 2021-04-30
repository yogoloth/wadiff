package main

import (
	"flag"
	"fmt"
	"os"
)

var args_is_str bool
var is_print_all bool
var is_print_change bool
var no_print_diff bool

func main() {
	flag.BoolVar(&args_is_str, "s", false, "args as string")
	flag.BoolVar(&is_print_all, "a", false, "print all common data")
	flag.BoolVar(&is_print_change, "c", false, "print changes")
	flag.BoolVar(&no_print_diff, "n", false, "donnot print diff")
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
		str1 := []rune(str_x)
		str2 := []rune(str_y)

		routex, routey, route_max := GetDiff(str1, str2)

		PrintDiff(routex, routey, route_max)
		PrintDiff(routey, routex, route_max)
	} else {
		file1_str, file2_str := DiffFile(str_x, str_y)
		routex, routey, route_max := GetDiff(file1_str, file2_str)

		config := FileDiffConfig{is_print_all: is_print_all}
		PrintFileDiff(routex, routey, route_max, config)
		//PrintDiff(routey, routex, route_size)
	}
}
