package main

import (
	"bufio"
	"log/slog"
	"os"
	"strings"
)

func part1() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int

	cardNum := 0
	for scanner.Scan() {
		cardNum += 1
		line := scanner.Text()
		splt := strings.Split(strings.Split(line, ":")[1], "|")

		winning := strings.Fields(splt[0])
		numbers := strings.Fields(splt[1])

		var points int
		var found bool

		for _, num := range numbers {
			for _, win := range winning {
				if num == win {
					if !found {
						points += 1
						found = true
						break
					}
					points *= 2
					break
				}
			}
		}

		if found {
			sum += points
		}

	}
	if err := scanner.Err(); err != nil {
		return err
	}

	slog.Info("part1", "result", sum)

	return nil
}
