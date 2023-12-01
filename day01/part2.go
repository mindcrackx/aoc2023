package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part2() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	calibrations := make([]int, 0, 1024)

	wordMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		re := regexp.MustCompile(`^(one|two|three|four|five|six|seven|eight|nine)`)

		var leftI int
		var left, right string
		for i := 0; i < len(line); i++ {
			if strings.Contains("0123456789", string(line[i])) {
				leftI = i
				left = string(line[i])
				break
			}

			matches := re.FindStringSubmatch(line[i:])
			if matches != nil {
				leftI = i
				left = wordMap[matches[1]]
				break
			}
		}

		for i := len(line) - 1; i >= leftI; i-- {
			if strings.Contains("0123456789", string(line[i])) {
				right = string(line[i])
				break
			}

			matches := re.FindStringSubmatch(line[i:])
			if matches != nil {
				right = wordMap[matches[1]]
				break
			}
		}

		slog.Info("part2", "left", left, "right", right)
		num, err := strconv.Atoi(left + right)
		if err != nil {
			return err
		}

		calibrations = append(calibrations, num)

	}
	if err := scanner.Err(); err != nil {
		return err
	}

	var sum int
	for _, c := range calibrations {
		sum += c
	}

	slog.Info("part2", "val", sum)

	return nil

}
