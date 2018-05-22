package main

import (
	"fmt"
	"os"
)

func main() {
	if (len(os.Args)) == 1 {
		//Print help & exit
		fmt.Println("Use: ")
		fmt.Println(os.Args[0], "MP3FILE")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil { fmt.Println(err); return }

	
	
	err = file.Close()
	if err != nil {	fmt.Println(err) }
	
	fmt.Println(os.Args)
}
