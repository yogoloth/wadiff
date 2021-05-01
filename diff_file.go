package main

import (
	"fmt"
	"io/ioutil"
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

type FileDiffConfig struct {
	is_print_all bool
}

func File2Array(file string) (file_str []rune, err error) {

	if data, err := ioutil.ReadFile(file); err == nil {
		//fmt.Println(data)
		file_str = []rune(string(data))
	} else {
		return nil, err
	}
	return
}

func PrintFileDiff(str1 []rune, str2 []rune, s_max int, config FileDiffConfig) {
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
		if str1[i] == rune('\n') || str2[i] == rune('\n') {
			s1 = i - 1
			if is_diff {
				if !no_print_diff {
					fmt.Printf("file1:  ")
					PrintDiff(str1[s1+1:s2+1], str2[s1+1:s2+1], s2-s1-1)
					fmt.Printf("file2:  ")
					PrintDiff(str2[s1+1:s2+1], str1[s1+1:s2+1], s2-s1-1)
				}
				if is_print_change {
					fmt.Printf("change: ")
					PrintChange(str1[s1+1:s2+1], str2[s1+1:s2+1], s2-s1-1)
					//PrintDiff(str2[s1+1:s2+1], str1[s1+1:s2+1], s2-s1-1)
				}
				is_diff = false
			} else {
				if config.is_print_all && s2 > s1+1 {
					//for i := s2; i > s1; i-- {
					//	fmt.Printf("%c", str1[i])
					//}
					//fmt.Printf("\n")
					fmt.Printf("same:   ")
					PrintRunes(str2[s1+1:s2+1], s2-s1-1)
				}
			}
			s2 = s1
		} else if str1[i] != str2[i] { //str1 '\n' == str2 ' '
			is_diff = true
		}
	}

}

func DiffFile(file1 string, file2 string) (file1_str []rune, file2_str []rune) {
	var err error
	if file1_str, err = File2Array(file1); err != nil {
		fmt.Printf("read %s error :%v\n", file1, err)
	}
	if file2_str, err = File2Array(file2); err != nil {
		file2_str, err = File2Array(file2)
	}

	//fmt.Printf("%v\n%v\n", file1_str, file2_str)
	return
}
