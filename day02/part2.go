package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func part2() error {
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

		var maxRed, maxBlue, maxGreen int

		rounds := strings.Split(splt[1], ";")
		for _, r := range rounds {
			vals := strings.Split(r, ",")
			for _, v := range vals {
				vv := strings.Split(strings.TrimSpace(v), " ")
				x, _ := strconv.Atoi(vv[0])
				switch vv[1] {
				case "red":
					if x > maxRed {
						maxRed = x
					}
				case "blue":
					if x > maxBlue {
						maxBlue = x
					}
				case "green":
					if x > maxGreen {
						maxGreen = x
					}
				}
			}
		}

		sum += maxRed * maxBlue * maxGreen
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	slog.Info("part2", "result", sum)

	return nil

}
