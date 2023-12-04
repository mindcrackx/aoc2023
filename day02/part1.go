package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func part1() error {

	const (
		maxRed   = 12
		maxGreen = 13
		maxBlue  = 14
	)

	file, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	var sum int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		splt := strings.Split(line, ":")
		gameID, _ := strconv.Atoi(strings.Split(splt[0], " ")[1])

		possible := true

		rounds := strings.Split(splt[1], ";")
		for _, r := range rounds {
			vals := strings.Split(r, ",")
			for _, v := range vals {
				vv := strings.Split(strings.TrimSpace(v), " ")
				x, _ := strconv.Atoi(vv[0])
				switch vv[1] {
				case "red":
					if x > maxRed {
						possible = false
					}
				case "blue":
					if x > maxBlue {
						possible = false
					}
				case "green":
					if x > maxGreen {
						possible = false
					}
				}
			}
		}

		if possible {
			sum += gameID
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	slog.Info("part1", "result", sum)

	return nil
}
