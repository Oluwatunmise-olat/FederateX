package main

import (
	"fmt"
	"log"
)

func main() {
	tun, err := CreateTUN("federatex0")

	if err != nil {
		log.Fatal(err)
	}
	
	defer tun.Close()

	fmt.Println("Created TUN interface: federatex0")
}
