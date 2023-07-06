package inp

import (
	"bufio"
	"encoding/base64"
	"io"
)

func ReadLine(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	var inp string
	if scanner.Scan() {
		inp = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return inp
}

func ReadLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

func ReadBase64(r io.Reader) []byte {
	decoder := base64.NewDecoder(base64.StdEncoding, r)
	inp, err := io.ReadAll(decoder)
	if err != nil {
		panic(err)
	}
	return inp
}
