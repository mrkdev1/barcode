package main

import (
	"os"
	"fmt"
	"strings"
	"log"
    "io/ioutil"

)

func main() {
    bc := string(os.Args[1])
	
    bar, err := ioutil.ReadFile(bc)
    if err != nil {
        log.Fatal(err)
    }
    lines := strings.Split(string(bar), "\n")

    file := string(os.Args[2])
	
    content, err := ioutil.ReadFile(file)
    if err != nil {
        log.Fatal(err)
    }
    a := strings.Split(string(content), "\n")
	
	s := ""
	
	for i := 1; i < (len(a)-1); i++ {
	   s = s + a[i]
    }
	
//	fmt.Println(lines[0] + " (input)\n" + s)

	for i := 1; i < (len(lines)-1); i++ {
	found := strings.Contains(s, lines[i])
	fmt.Println (found)
	}
}
