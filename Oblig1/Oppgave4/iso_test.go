package Oppg4


import ("fmt"
		"testing"
)
	func extendedASCIITest (t *testing.T) {

		input := "Å"

		test := ExtendedASCIIText(input)

		for _, c := range test {
			if c < 0x80 {
				t.Error("This string has characters from ASCII ")
			} else {
				fmt.Println("No characters from ASCII")
			}

		}
	}
