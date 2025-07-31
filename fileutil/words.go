package fileutil

import (
	"bufio"
	"github.com/tronget/go-utils/strutil"
	"os"
	"strings"
)

func WordsFrequency(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	hm := make(map[string]int)

	for scanner.Scan() {
		current := scanner.Text()
		word := strutil.TrimSpec(current)
		word = strings.ToLower(word)
		if word != "" {
			hm[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return hm, nil
}

func ListWords(filename string) (words []string, err error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func CountWords(filename string) (count int, err error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}
