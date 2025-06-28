package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to Files")

	content := "Aydin is ready to go to school"

	file, err :=os.Create("./myAydinFile.txt")
	genericError(err)

	length, err := io.WriteString(file, content)
	genericError(err)
	fmt.Println("Length of file is : ",length)
	defer file.Close();
	fileReader("./myAydinFile.txt")

}

func fileReader(fileName string)  {
	databyte, err := os.ReadFile(fileName)
	genericError(err)

	fmt.Println("Text data inside the file is : ",string(databyte))
}

func genericError(err error){
	if err != nil {
		panic(err)
	}
}