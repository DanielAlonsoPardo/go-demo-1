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
	fmt.Println(os.Args)
}
