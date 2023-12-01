package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func part1() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	calibrations := make([]int, 0, 1024)

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		var leftI, rightI int
		for i := 0; i < len(line); i++ {
			if strings.Contains("0123456789", string(line[i])) {
				leftI = i
				break
			}
		}

		for i := len(line) - 1; i >= leftI; i-- {
			if strings.Contains("0123456789", string(line[i])) {
				rightI = i
				break
			}
		}

		num, err := strconv.Atoi(string(line[leftI]) + string(line[rightI]))
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

	slog.Info("part1", "val", sum)

	return nil
}
