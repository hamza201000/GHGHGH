package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Clean(f string) string {
	sn := false
	str := ""
	s := []rune(f)
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			str += string(s[i])
			sn = true
		} else if s[i] == ' ' && sn {
			str += " "
			sn = false
		}
	}
	str = strings.TrimSpace(str)
	
	
	return str
}
func Mark(f string) string {
	line := []rune(Clean(f))
	ismark := false
	indexmark := 0
	chngstr := []rune{}
	for i := 0; i < len(line); i++ {
		if i != 0 && i != len(line)-1 && line[i] == '\'' && unicode.IsLetter(line[i+1]) && unicode.IsLetter(line[i-1]) {
			continue
		}
		if line[i] == '\'' && !ismark {
			indexmark = i
			ismark = true
		} else if line[i] == '\'' && ismark {
			chngstr = append(chngstr, line[indexmark])
			chng := []rune(Clean(string(line[indexmark+1 : i])))
			//fmt.Println(string(line[indexmark+1 : i]))
			chngstr = append(chngstr, chng...)
			chngstr = append(chngstr, line[i])
			ismark = false
		} else if !ismark {
			chngstr = append(chngstr, line[i])
		}
	}
	if ismark {
		chngstr = append(chngstr, line[indexmark:]...)
	}
	return string(chngstr)
}
func main() {

	fmt.Println(Mark("ghg '  hgjhg    hbv ' "))
	fmt.Println(Mark("    "))
	fmt.Println(Mark("ghg '  hgjh' g' '   hbv "))
	fmt.Println(Mark("hhf '     '     '   '   '   '   '"))

}
