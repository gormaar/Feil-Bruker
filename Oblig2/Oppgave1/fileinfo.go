package main

import (
	"fmt"
	"os"

	"log"
)
var(
	fileInfo os.FileInfo
	err error
)
func main() {

	fileInfo, err = os.Stat("file.txt")
	if err != nil	{
		log.Fatal(err)
	}

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

	if mode&os.ModeAppend == 0	{		//Sjekker append
		fmt.Println("Is not append only")

	}	else {
		fmt.Println("Is append only")
	}

	if mode&os.ModeDevice == 0	{		//Sjekker om det er device fil eller ikke
		fmt.Println("Is not device file")
	}	else	{
		fmt.Println("Is device file")
	}

	if mode&os.ModeCharDevice == 0	{		//Sjekker om det et Unix character device
		fmt.Println("Is not Unix character device")
	}	else {
		fmt.Println("Is Unix character device")
	}

	if mode&os.ModeSymlink == 0	{		//Sjekker om det er en symbolisk link
		fmt.Println("Is not a symbolic link")
	}	else {
		fmt.Println("Is a symbolic link")
	}
}