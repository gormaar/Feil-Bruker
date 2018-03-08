package main

import (
	"fmt"
	"os"

)
var(
	fileInfo os.FileInfo
	err error
)
func main() {

	fileInfo, err = os.Stat("file.txt")

	byteInt := fileInfo.Size()		//Konverterer bytes fra filen til KB, MB og GB
	bytes := float64(byteInt)
	KB := bytes / 1024
	MB := KB / 1024
	GB := MB / 1024

	fmt.Println("Information about file <file.txt>:")
	fmt.Println("Size:", fileInfo.Size(), "in bytes,", KB, "KB", MB, "MB", GB, "GB")

	mode := fileInfo.Mode()
	dir := mode.IsDir()		//Sjekker om det er directory eller ikke
	if dir == true {
		fmt.Println("Is a directory")
	} else {
		fmt.Println("Is not a directory")
	}
	file := mode.IsRegular()		//Sjekker om det er fil eller ikke
	if file == true {
		fmt.Println("Is a file")
	} else {
		fmt.Println("Is not a file")
	}

	fileMode := fileInfo.Mode()
	fmt.Println("Has Unix permission bits:", fileMode) //Printer permission bits

	append := fileInfo.ModeAppend
	if append == true{
		
	}

}
