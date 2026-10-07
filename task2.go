package main

import "fmt"

func count(list []string) map[string]int {
	result := make(map[string]int)

	for _, str := range list {
		count := 0
		for _, char := range str {
			if char >= '0' && char <= '9' {
				count++
			}
		}
		result[str] = count
	}
	return result
}
func main() {
	a := []string{"asdqwe123qwe", "qwertdfxvc", "qwtr", "qwetreyr"}
	b := count(a)

	for key, value := range b {
		fmt.Println(key, value)
	}

}
