package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//func File2Array(file string) (file_str []string, err error) {
//	if fp, err := os.OpenFile(file, os.O_RDONLY, 0600); err != nil {
//		return nil, err
//	} else {
//		defer fp.Close()
//		reader := bufio.NewReader(fp)
//		for {
//			str, err := reader.ReadString('\n')
//			if err == nil {
//				//fmt.Printf("read string is %s: \n", str)
//				file_str = append(file_str, str)
//			} else if err == io.EOF {
//				break
//			} else {
//				//fmt.Println("read string failed, err: ", err)
//				return nil, err
//			}
//		}
//	}
//	return
//}

func File2Array(file string) (file_str []rune, err error) {
	if fp, err := os.OpenFile(file, os.O_RDONLY, 0600); err != nil {
		return nil, err
	} else {
		defer fp.Close()
		reader := bufio.NewReader(fp)
		for {
			str, err := reader.ReadString('\n')
			if err == nil {
				//fmt.Printf("read string is %s: \n", str)
				str_r := []rune(str)
				str_r = append(str_r, rune('\n'))
				file_str = append(file_str, str_r...)
			} else if err == io.EOF {
				break
			} else {
				//fmt.Println("read string failed, err: ", err)
				return nil, err
			}
		}
	}
	return
}

func PrintFileDiff(str1 []rune, str2 []rune, s_max int) {
	//for i := size; i > 0; i-- {
	//	//if routex[i] == routey[i] || routey[i] == rune(' ') {
	//	if str1[i] == str2[i] {
	//		fmt.Printf("%c", str1[i])
	//	} else if str1[i] == rune(' ') {
	//		color.BgRed.Printf("%c", str1[i])
	//	} else {
	//		color.Red.Printf("%c", str1[i])
	//	}
	//}
	//fmt.Printf("\n")

	s2 := s_max
	s1 := s2
	//new_s2 := s2
	is_diff := false
	for i := s_max; i > 0; i-- {
		if str1[i] == rune('\n') {
			//new_s2 = i - 1
			s1 = i - 1
			if is_diff {
				fmt.Printf("file1: ")
				PrintDiff(str1[s1+1:s2+1], str2[s1+1:s2+1], s2-s1-1)
				fmt.Printf("file2: ")
				PrintDiff(str2[s1+1:s2+1], str1[s1+1:s2+1], s2-s1-1)
				is_diff = false
			}
			s2 = s1
		} else if str1[i] != str2[i] { //str1 '\n' == str2 ' '
			is_diff = true
		}
	}

	//PrintDiff(str1[s1:s1+s2+1], str2[s1:s1+s2+1], s2)
	//PrintDiff(str2[s1:s1+s2+1], str1[s1:s1+s2+1], s2)

}

func DiffFile(file1 string, file2 string) (file1_str []rune, file2_str []rune) {
	file1_str, _ = File2Array(file1)
	file2_str, _ = File2Array(file2)

	//fmt.Printf("%v\n%v\n", file1_str, file2_str)
	return
}
