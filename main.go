package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	config, err := ParseConfig(os.Args)
	if err != nil {
		panic(err)
	}

	file, err := os.Open(config.Filename)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	lines, err := CompareLines(config, file)

	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

type Config struct {
	Filename   string
	Word       string
	IgnoreCase bool
}

func ParseConfig(args []string) (Config, error) {
	filename := args[1]
	word := args[2]
	case_, ok := os.LookupEnv("IGNORE_CASE")
	if !ok {
		case_ = "false"
	}

	ignore_case, err := strconv.ParseBool(case_)

	if err != nil {
		return Config{}, err
	}

	return Config{
		Filename:   filename,
		Word:       word,
		IgnoreCase: ignore_case,
	}, nil
}

func CompareLines(config Config, file *os.File) ([]string, error) {
	scanner := bufio.NewScanner(file)
	var lines []string

	for i := 1; scanner.Scan(); i++ {
		if !config.IgnoreCase {
			if strings.Contains(scanner.Text(), config.Word) {
				lines = append(lines, fmt.Sprintf("%d: %s", i, scanner.Text()))
			}
		} else {
			if strings.Contains(strings.ToLower(scanner.Text()), strings.ToLower(config.Word)) {
				lines = append(lines, fmt.Sprintf("%d: %s", i, scanner.Text()))
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
