package main

import (
	"io"
	"log"
	"os"
)

func RunAssignment1() {

	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal("Error in opening file", err.Error())
	}
	io.Copy(os.Stdout, file)

	// We could also use byte slices to print the contents
	// bs := make([]byte, 32*1024)
	// _, er := file.Read(bs)
	// if er != nil {
	// 	log.Fatal("Error in reading file", er.Error())
	// }
	// fmt.Println(string(bs))

}
