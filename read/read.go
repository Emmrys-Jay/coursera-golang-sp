package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname, lname string
}

type names []name

func main() {
	nameStruct := name{}
	var nameSlice names
	var filename string

	fmt.Print("Enter the file name: ")
	fmt.Scan(&filename)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Could not open file: ", err)
	}
	defer f.Close()
	fs := bufio.NewScanner(f)
	for fs.Scan() {
		line := fs.Text()
		nameStruct.fname = strings.Split(line, " ")[0]
		nameStruct.lname = strings.Split(line, " ")[1]

		nameSlice = append(nameSlice, nameStruct)
	}

	for _, val := range nameSlice {
		fmt.Println(val.fname, val.lname)
	}

}
