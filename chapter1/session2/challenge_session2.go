package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Printf("Nilai i = %v\n", i+1)

		if i == 3 {
			for j := 0; j < 10; j++ {
				if j == 4 {
					s := "САШАРВО"

					for i, runeChar := range s {
						fmt.Printf("character %#U starts at byte position %d\n", runeChar, i)
					}
				} else {
					fmt.Printf("Nilai j = %v\n", j+1)
				}

			}
		}
	}

	// fmt.Printf("Glyph:             %q\n", s)
	// fmt.Printf("UTF-8:             [% x]\n", []byte(s))
	// fmt.Printf("Unicode codepoint: %U\n", []rune(s))

	// for i, runeChar := range b {
	// 	fmt.Printf("byte position %d: %#U\n", i, runeChar)
	// }

	//b := []byte{0xd0, 0xa1, 0xd0, 0x90, 0xd0, 0xa8, 0xd0, 0x90, 0xd0, 0xa0, 0xd0, 0x92, 0xd0, 0x9e}
}
