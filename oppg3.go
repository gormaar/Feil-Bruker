package main

import "fmt"


func main() {

		for i := 1; i > 0; i++ {
			func() {
				fmt.Println(i)
			}()

		}

	}

