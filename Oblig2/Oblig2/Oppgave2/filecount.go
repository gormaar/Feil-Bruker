package main

import (
		"fmt"
		"os"
		"bufio"
		"io/ioutil"
		"bytes"
		"log"
)

func main() {

	file, _ := os.Open("file.txt")
	scanner := bufio.NewScanner(file)
	reader, err := ioutil.ReadFile("file.txt")	//Leser filen
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[int]int)	//Oppretter et map og en slice
	var runeFound []byte

	for i := 0x00; i <= 0x80; i++ {	//Loop for alle tegn
		for j := 0; j < len(reader); j++ {	//Loop for alle tegn i filen
			if int(reader[j]) == i {
				runeFound = append(runeFound, reader[j])	//Legger runes fra fil inn i slice
				j = len(reader)
			}
		}
	}

	var countResult= []int{}

	for i := 0; i < len(runeFound); i++ {	//Loop som teller runes, og legger det inn i map
		runeSep := []byte{runeFound[i]}
		count := bytes.Count(reader, runeSep)
		countResult = append(countResult, count)
		//fmt.Println(runeFound[i])
		m[int(runeFound[i])] = count
	}

	lineCount := 0
	for scanner.Scan() {	//Teller hvor mange linjer filen har
		lineCount++
	}

	fmt.Println("Information about file.txt:")
	fmt.Println("Number of lines in file:", lineCount)

	sortedSlice := bubbleSort(countResult)		//Finner de fem runes som har blitt brukt mest ut fra en sortert liste
	sortedSliceLenght := len(sortedSlice) - 5
	largest5 := sortedSlice[sortedSliceLenght:]

	fmt.Println("Most common runes:")

	//Sjekker verdier fra map opp mot de fem mest brukte runes (en loop for hver av de 5 mest brukte)
	//Loopen breaker når verdien har blitt funnet, og begynner deretter på neste loop
Loop1:
	for key, value := range m {
		if value == largest5[4] {
			fmt.Printf("1. Rune: %c", key)
			fmt.Printf(", Counts: %d \n", value)
			delete(m, key) //Fjerner key for at ikke samme rune skal komme flere ganger
			break Loop1
		}
	}

Loop2:
	for key, value := range m {
		if value == largest5[3] {
			fmt.Printf("2. Rune: %c", key)
			fmt.Printf(", Counts: %d \n", value)
			delete(m, key)
			break Loop2
		}
	}

Loop3:
	for key, value := range m {
		if value == largest5[2] {
			fmt.Printf("3. Rune: %c", key)
			fmt.Printf(", Counts: %d \n", value)
			delete(m, key)
			break Loop3
		}
	}

Loop4:
	for key, value := range m {
		if value == largest5[1] {
			fmt.Printf("4. Rune: %c", key)
			fmt.Printf(", Counts: %d \n", value)
			delete(m, key)
			break Loop4
		}
	}

Loop5:
	for key, value := range m {
		if value == largest5[0] {
			fmt.Printf("5. Rune: %c", key)
			fmt.Printf(", Counts: %d \n", largest5[0])
			delete(m, key)
			break Loop5
		}
	}
}

func bubbleSort(countResult []int) []int {
	// find the length of list n
	n := len(countResult)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if countResult[j] > countResult[j+1] {
				temp := countResult[j+1]
				countResult[j+1] = countResult[j]
				countResult[j] = temp
			}
		}
	}
	return countResult
}
