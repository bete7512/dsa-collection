package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var result strings.Builder
	for _, str := range strs {
		result.WriteString(fmt.Sprintf("%d#%s", len(str), str))
	}

	return result.String()
}

func (s *Solution) Decode(str string) []string {
	var results []string
	i := 0
	for i < len(str) {
		if str[i] >= '0' && str[i] <= '9' {
			j := i
			for j < len(str) && str[j] >= '0' && str[j] <= '9' {
				j++
			}

			if j < len(str) && str[j] == '#' {
				length, _ := strconv.Atoi(str[i:j])
				if j+1+length <= len(str) {
					word := str[j+1 : j+1+length]
					results = append(results, word)
					i = j + 1 + length
					continue
				}
			}
		}
		i++
	}
	return results
}

func main() {
	s := Solution{}
	// encoded := s.Encode([]string{""})
	// fmt.Println("Encoded:", encoded)
	// decoded := s.Decode(encoded)
	// fmt.Println("Decoded:", decoded)
	// str2 := []string{"EmojiTest 😊", "🌟✨🌟✨🌟", "🤖👽🤖👽"}
	str2 := []string{"ddddddddddddddddddddddddddddwe", "say", ":", "yes", "!@#$%^&*()"}
	log.Println("STR1:", str2)
	log.Println("STR1LENNN:", len(str2))
	encoded2 := s.Encode(str2)
	fmt.Println("Encoded2:", encoded2)
	decoded2 := s.Decode(encoded2)
	fmt.Println("Decoded2:", decoded2)
	fmt.Println("Decoded2LENNN:", len(decoded2))
}

