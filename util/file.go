package util

import (
	"bufio"
	"os"
)

func ReadLines(path string) ([]string, error) {
	c, err := StreamLines(path)
	if err != nil {
		return nil, err
	}

	var lines []string
	for l := range c {
		lines = append(lines, l)
	}

	return lines, nil
}

func StreamLines(path string) (chan string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	c := make(chan string, 0)

	go func(*os.File, chan string) {

		defer close(c)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			txt := scanner.Text()
			c <- txt
		}
	}(file, c)

	return c, nil
}
