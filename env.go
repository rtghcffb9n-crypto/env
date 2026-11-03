package env

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	red    = "\x1b[31m"
	yellow = "\x1b[33m"
	green  = "\x1b[32m"
	reset  = "\x1b[0m"
)

func Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		expression := strings.TrimSpace(scanner.Text())

		if expression == "" || strings.HasPrefix(expression, "#") {
			continue
		}

		key, value, success := parse(expression)
		if !success {
			log.Fatal(red + "CHECK AND CORRECT '.env' FILE!" + reset)
		}

		os.Setenv(key, value)
		log.Printf("YOU HAVE LOADED '%s' VARIABLE", key)
	}

	return scanner.Err()
}

func parse(expression string) (key, value string, success bool) {
	pos := strings.IndexByte(expression, '=')
	if pos == -1 {
		return "", "", false
	}

	key = strings.TrimSpace(expression[:pos])
	value = strings.TrimSpace(expression[pos+1:])

	return key, value, true
}
