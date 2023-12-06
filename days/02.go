package days

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/mindcrackx/aoc2023/utils"
)

func Two_1(input io.Reader) (string, error) {
	const (
		maxRed   = 12
		maxGreen = 13
		maxBlue  = 14
	)

	var sum int

	scanner := bufio.NewScanner(input)

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
				x := utils.MustAtoi(vv[0])
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
		return "", err
	}

	return strconv.Itoa(sum), nil
}

func Two_2(input io.Reader) (string, error) {
	var sum int

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		splt := strings.Split(line, ":")

		var maxRed, maxBlue, maxGreen int

		rounds := strings.Split(splt[1], ";")
		for _, r := range rounds {
			vals := strings.Split(r, ",")
			for _, v := range vals {
				vv := strings.Split(strings.TrimSpace(v), " ")
				x := utils.MustAtoi(vv[0])
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
		return "", err
	}

	return strconv.Itoa(sum), nil
}
