package main

import (
	"fmt"
	"github.com/gookit/color"
)

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

	routex = make([]rune, len(matrix[0])+len(matrix))
	routey = make([]rune, len(matrix[0])+len(matrix))

	i = 0
	for {
		i++
		if x < 1 && y < 1 {
			break
		}
		px, py := GetPrev(matrix, x, y)
		if px == x-1 && py == y-1 {
			routey[i] = target[y-1]
			routex[i] = source[x-1]
			x = px
			y = py
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

func PrintRunes(str []rune, size int) {
	//fmt.Printf("%d", size)

	for i := size; i > 0; i-- {
		fmt.Printf("%c", str[i])
	}
	fmt.Printf("\n")
}

func PrintDiff(str1 []rune, str2 []rune, size int) {
	for i := size; i > 0; i-- {
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

func PrintChange(str1 []rune, str2 []rune, size int) {
	max := size
	min := 0

	for i := size; i > 0; i-- {
		if str1[i] != str2[i] {
			if max == size {
				max = i
				//fmt.Printf("max %d\n", max)
			}
			min = i
			//fmt.Printf("min %d\n", min)
		}
	}
	min = min - 1
	if max == size {
		max = max + 1
	}
	//fmt.Println(min, max, size)
	for i := max; i > min; i-- {
		if str1[i] == str2[i] {
			fmt.Printf("%c", str1[i])
		} else if str1[i] == rune(' ') {
			//color.BgRed.Printf("%c", str1[i])
		} else {
			color.Red.Printf("%c", str1[i])
		}
	}
	fmt.Printf("  -->  ")
	for i := max; i > min; i-- {
		if str1[i] == str2[i] {
			fmt.Printf("%c", str1[i])
		} else if str2[i] == rune(' ') {
			//color.BgRed.Printf("%c", str2[i])
		} else {
			color.Red.Printf("%c", str2[i])
		}
	}
	fmt.Printf("\n")
}

func GetDiff(str_1 []rune, str_2 []rune) (routex []rune, routey []rune, route_size int) {

	matrix := CreateDp(str_1, str_2)

	routex, routey, route_size = GetRoute(str_1, str_2, matrix)
	return
}
