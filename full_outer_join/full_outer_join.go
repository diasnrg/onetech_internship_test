package full_outer_join

import (
	"os"
	"log"
	"sort"
	"strings"
	"fmt"
)

func FullOuterJoin(f1Path, f2Path, resultPath string) {

	//read the content of two files
	f1, err := os.ReadFile(f1Path)
	if err != nil {
		log.Fatal(err)
	}

	f2, err := os.ReadFile(f2Path)
	if err != nil {
		log.Fatal(err)
	}

	//converting []byte to string and returning it like a []string
	s1 := strings.Fields(string(f1))
	s2 := strings.Fields(string(f2))

	//mapping strings
	mp1 := make(map[string]bool)
	mp2 := make(map[string]bool)

	for _,v := range s1 {
		mp1[v] = true
	}
	
	for _,v := range s2 {
		mp2[v] = true
	}

	//slice with strings that belongs only to one of the files
	var xor []string
	//go through the maps searching for XOR strings
	for k := range mp1 {
		if !mp2[k] {
			xor = append(xor, k)
		}
	}

	for k := range mp2 {
		if !mp1[k] {
			xor = append(xor, k)
		}
	}

	//sort the strings in alphabetical order 
	sort.Strings(xor)

	//result string, using 'if' statement to either add new space (\n) or not
	xorStr := ""
	for i,v := range xor {
		xorStr += fmt.Sprintf("%s", v)

		if i < len(xor)-1 {
			xorStr += "\n"
		}
	}

	//write the result to the file
	if err := os.WriteFile(resultPath, []byte(xorStr), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}