package main

import (
	"bufio"
	"os"
	// "io/ioutil"
	"strings"
	"fmt"

)

type Names struct{
	fname string
	lname string
}

func main(){
slice := make([]Names, 0, 3)
fmt.Println("Enter the file name :")

var fileName string
fmt.Scan(&fileName)

file, e := os.Open(fileName)
if e != nil {
	fmt.Println("Error is = ",e)
}

scanner := bufio.NewScanner(file)
for scanner.Scan() {

	s := strings.Split(scanner.Text(), " ")
	var namee Names
	namee.fname, namee.lname = s[0], s[1]
	slice = append(slice, namee)

}

file.Close()
fmt.Println("\n********************\n")

for _, v := range slice {
	fmt.Println(v.fname, v.lname)
}
}
