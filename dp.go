package main

//import (
//lev "github.com/texttheater/golang-levenshtein/levenshtein"
//"os"
//)

//func CreateDf(str_x []rune, str_y []rune) (df [][]int) {
//	options := lev.Options{
//		InsCost: 1,
//		DelCost: 1,
//		SubCost: 1,
//		Matches: lev.IdenticalRunes,
//	}
//
//	df = lev.MatrixForStrings(str_x, str_y, options)
//
//	debug := false
//	if debug == true {
//		lev.WriteMatrix(str_x, str_y, df, os.Stdout)
//	}
//	return
//}

func min(a int, b int) int {
	if b < a {
		return b
	}
	return a
}

func CreateDp(source []rune, target []rune) (df [][]int) {
	//  http://www.let.rug.nl/~kleiweg/lev/levenshtein.html
	height := len(source) + 1
	width := len(target) + 1
	df = make([][]int, height)

	for i := 0; i < height; i++ {
		df[i] = make([]int, width)
		df[i][0] = i * 1
	}
	for j := 1; j < width; j++ {
		df[0][j] = j * 1
	}

	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			delCost := df[i-1][j] + 1
			matchSubCost := df[i-1][j-1]
			if source[i-1] != target[j-1] {
				matchSubCost += 1
			}
			insCost := df[i][j-1] + 1
			df[i][j] = min(delCost, min(matchSubCost,
				insCost))
		}
	}
	return
}
