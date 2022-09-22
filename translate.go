package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func Add() {
	var s1 string
	var s2 string
	var k int
	var b bool
	b = true
	filePtr, err := os.Open("Translate book.txt")
	if err != nil {
		filePtr, err = os.Create("Translate book.txt")
		if err != nil {
			log.Fatal(err)
		}
	} else {

		for b {
			fmt.Println("Enter a word")
			fmt.Scanf("%s", &s1)
			fmt.Println("Enter translate")
			fmt.Scanf("%s", &s2)
			b = numErrFunc(s1, s2)
			if b {

			} else {
				k = RuEng(s1)
				file_data, err := os.ReadFile(filePtr.Name())
				if err != nil {
					log.Fatal(err)
				} else {
					switch k {
					case 1:

						if string(file_data) == "" {
							os.WriteFile(filePtr.Name(), []byte(s1+"-"+s2), 0600)
						} else {
							os.WriteFile(filePtr.Name(), []byte(string(file_data)+"\n"+s1+"-"+s2), 0600)

						}
						break
					case 2:
						if string(file_data) == "" {
							os.WriteFile(filePtr.Name(), []byte(s2+"-"+s1), 0600)
						} else {
							os.WriteFile(filePtr.Name(), []byte(string(file_data)+"\n"+s2+"-"+s1), 0600)

						}
						break
					}
				}
			}
		}
	}
	defer filePtr.Close()
}

func numErrFunc(s1 string, s2 string) bool {
	var b bool
	re1, _ := regexp.MatchString(`\d`, s1)
	re2, _ := regexp.MatchString(`\d`, s2)
	if re1 == true || re2 == true {
		fmt.Println("Error.Enter string without numbers")
		if re1 == true {
			re1 = false
		}
		if re2 == true {
			re2 = false
		}
		b = true
	} else {
		b = false
	}
	return b
}

func RuEng(s1 string) int {
	var k int

	for _, r := range s1 {
		if unicode.Is(unicode.Cyrillic, r) {

			k = 1
			break
		} else {

			k = 2
			break
		}

	}
	return k
}

func Print() {
	filePtr, err := os.ReadFile("Translate book.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(filePtr) + "\n")

	}
}
func Find() {
	var b bool
	var letter string
	filePtr, err := os.Open("Translate book.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		b = true
		for b {
			fmt.Println("Enter the word you want to translate")
			fmt.Scanf("%s", &letter)
			b = numErrFunc(letter, "")
			if !b {

				scanner := bufio.NewScanner(filePtr)
				for scanner.Scan() {

					re, _ := regexp.MatchString(letter, string(scanner.Text()))
					if re {
						k := RuEng(letter)
						if k == 1 {
							words := strings.Split(string(scanner.Text()), "-")
							i := 0
							for _, word := range words {

								if string(word) != letter && i == 0 {
									fmt.Println(string(scanner.Text()))
								}
								if string(word) == letter && i == 0 {

									out := strings.Replace(string(scanner.Text()), string(word)+"-", "", 1)
									fmt.Println(string(out))
								}
								i = 1
							}

						}
						if k == 2 {
							words := strings.Split(string(scanner.Text()), "-")
							i := 0
							for _, word := range words {

								if string(word) != letter && i == 1 {
									fmt.Println(string(scanner.Text()))
								}
								if string(word) == letter && i == 1 {

									out := strings.Replace(string(scanner.Text()), "-"+string(word), "", 1)
									fmt.Println(string(out))
								}
								i = 1
							}
						}
					}

				}

			}
		}
	}
	defer filePtr.Close()
}
func Delete() {
	filePtr, err := os.Open("Translate book.txt")
	var str string
	var b bool
	b = true
	if err != nil {
		log.Fatal(err)
	} else {
		for b {
			fmt.Println("Enter the word you want to delete")
			fmt.Scanf("%s", &str)
			b = numErrFunc(str, "")
			if !b {
				scanner := bufio.NewScanner(filePtr)
				var lines []string
				n := 0
				for scanner.Scan() {
					text := scanner.Text()
					re, _ := regexp.MatchString(str, text)
					n++
					if re {

						for scanner.Scan() {

							text = scanner.Text()
							break

						}

					}

					lines = append(lines, text)
				}
				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}

				err = os.WriteFile(filePtr.Name(), []byte(strings.Join(lines, "\n")), 0644)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
}
func Sort() {
	filePtr, err := os.Open("Translate book.txt")
	if err != nil {
		log.Fatal(err)
	} else {

		scanner := bufio.NewScanner(filePtr)
		var arr []string

		for scanner.Scan() {
			text := scanner.Text()
			arr = append(arr, text)

		}
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr); j++ {
				if arr[i] < arr[j] {
					str := arr[i]
					arr[i] = arr[j]
					arr[j] = str
				}
			}
		}
		err = os.WriteFile(filePtr.Name(), []byte(strings.Join(arr, "\n")), 0644)
		if err != nil {
			log.Fatalln(err)
		}

	}
}

func menu() {

	var i int
	for {
		fmt.Println("Choose function:\n1. Add new word\n2. Print all\n3. Find word\n4. Delete word\n5. Exit")
		fmt.Scanf("%d", &i)

		if i < 1 || i > 6 {
			fmt.Println("Error. Enter the correct function number from one to five")
		}

		switch i {

		case 1:
			Add()
			Sort()
		case 2:
			Print()
		case 3:
			Find()
		case 4:
			Delete()
		case 5:
			return

		}
	}
}
func main() {
	menu()
}
