package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter your name: ")
	n := bufio.NewScanner(os.Stdin)
	n.Scan()
	name := n.Text()

	fmt.Print("Enter your address: ")
	a := bufio.NewScanner(os.Stdin)
	a.Scan()
	addr := a.Text()

	det := make(map[string]string)
	det["name"] = name
	det["address"] = addr

	jsonval, err := json.Marshal(det)
	if err != nil {
		log.Fatalln("Could not convert to json: ", err)
	}
	fmt.Println(string(jsonval))

}
