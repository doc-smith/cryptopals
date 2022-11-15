package main

import (
	"bufio"
	"fmt"
	"github.com/doc-smith/cryptopals/util/encoding/base64"
	"github.com/doc-smith/cryptopals/util/encoding/hex"
	"os"
)

func Solve(hexInput string) string {
	b, err := hex.DecodeHexString(hexInput)
	if err != nil {
		panic("input is not a valid hex string")
	}
	return base64.Encode(b)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fmt.Println(Solve(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
