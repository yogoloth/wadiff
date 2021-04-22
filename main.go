package main

import (
	"flag"
	"github.com/gookit/color"
	lev "github.com/texttheater/golang-levenshtein/levenshtein"

	"fmt"
	"os"
)

//func getPosition(matrix [][]int, x int, y int) int {
//	return matrix[x][y]
//}

func GetPrev(matrix [][]int, x int, y int) (px int, py int) {
	if x > 1 {
		px = x - 1
	}
	if y > 1 {
		py = y - 1
	}
	if matrix[px][py] > matrix[px][y] {
		py = y
	}
	if matrix[px][py] > matrix[x][py] {
		px = x
	}
	return
}

func GetRoute(source []rune, target []rune, matrix [][]int) (routex []rune, routey []rune, i int) {

	x := len(matrix) - 1
	y := len(matrix[0]) - 1

	//fmt.Println("hello", x)
	//var xp, yq int

	routex = make([]rune, len(matrix[0])+len(matrix))
	routey = make([]rune, len(matrix[0])+len(matrix))

	i = 0
	for {
		i++
		//fmt.Println(x, y)
		if x < 1 && y < 1 {
			break
		}
		px, py := GetPrev(matrix, x, y)
		//pv := getPosition(matrix, px, py)
		//v := getPosition(matrix, x, y)
		if px == x-1 && py == y-1 {
			//if pv == v {
			routey[i] = target[y-1]
			routex[i] = source[x-1]
			x = px
			y = py
			//} else {
			//}
		} else if px == x-1 {
			routey[i] = rune(' ')
			routex[i] = source[x-1]
			x = px
			y = py
		} else if py == y-1 {
			routey[i] = target[y-1]
			routex[i] = rune(' ')
			x = px
			y = py
		}
	}
	return
}

func PrintDiff(str1 []rune, str2 []rune, size int) {
	for i := size; i > 0; i-- {
		//if routex[i] == routey[i] || routey[i] == rune(' ') {
		if str1[i] == str2[i] {
			fmt.Printf("%c", str1[i])
		} else if str1[i] == rune(' ') {
			color.BgRed.Printf("%c", str1[i])
		} else {
			color.Red.Printf("%c", str1[i])
		}
	}
	fmt.Printf("\n")
}

var args_is_str bool

func main() {
	flag.BoolVar(&args_is_str, "s", false, "args as string")
	flag.Parse()

	fmt.Println("args_is_str:", args_is_str)

	others := flag.Args()
	fmt.Println("others:", others)

	str_x := []rune(others[0])
	str_y := []rune(others[1])

	//str_x := []rune("neighbouxxr")
	//str_y := []rune("Niijeighbour")

	//str_x := []rune("int-uat.faw")
	//str_y := []rune("int.faw")

	options := lev.Options{
		InsCost: 1,
		DelCost: 1,
		SubCost: 1,
		Matches: lev.IdenticalRunes,
	}

	matrix := lev.MatrixForStrings(str_x, str_y, options)

	debug := false
	if debug == true {
		lev.WriteMatrix(str_x, str_y, matrix, os.Stdout)
	}

	routex, routey, route_size := GetRoute(str_x, str_y, matrix)

	PrintDiff(routex, routey, route_size)
	PrintDiff(routey, routex, route_size)

}
