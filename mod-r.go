package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	env := os.Args[1] + "/env"
	file, _ := os.Open(env)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	kvs := map[string]string{}
	for scanner.Scan() {
		text := scanner.Text()
		i := strings.Index(text, "\t")
		if i > 0 {
			kvs[text[0:i]] = text[i+1:]
		}
	}
	println("Recv:", kvs["mymsg"])
}
