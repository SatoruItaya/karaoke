package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	flag.Parse()

	orgList := readFile(flag.Arg(0))
	var totalList []string
	var aisareList []string

	m := make(map[string]int)

	for _, element := range orgList {
		if _, ok := m[element]; !ok {
			m[element] = 1
			totalList = append(totalList, element)
		} else if m[element] == 1 {
			m[element] += 1
		} else if m[element] == 2 {
			aisareList = append(aisareList, element)
		}
	}

	fmt.Printf("トータル: %v曲\n", len(totalList))
	fmt.Printf("愛され大賞: %v曲\n", len(aisareList))

	writeFile(totalList, "total")
	writeFile(aisareList, "aisare")
}

func readFile(fileName string) []string {

	b, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(b)

	list := make([]string, 0)

	for s.Scan() {
		line := s.Text()

		if line == "" {
			continue
		}

		list = append(list, line)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return list
}

func writeFile(list []string, fileName string) {

	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for i, n := range list {
		if _, err := writer.WriteString(strconv.Itoa(i+1) + " " + n + "\n"); err != nil {
			log.Fatal(err)
		}
	}

	writer.Flush()
}
