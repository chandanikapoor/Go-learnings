package main

import(
    
"encoding/json"
	"fmt"
)


func main(){
	var name string
	fmt.Printf("Enter a name:")
	fmt.Scan(&name)

	var address string
	fmt.Printf("Enter an address:")
	fmt.Scan(&address)

	mapNameAddress := map[string]string{"name": name, "address": address}

	
file, _ := json.Marshal(mapNameAddress)
fmt.Println(string(file))
}
	
