package days

import (
	"bufio"
	"io"
	"regexp"
	"slices"
	"strconv"
)

func One_1(input io.Reader) (string, error) {
	var result int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := []rune(scanner.Text())

		var left, right rune

		// left side
		for i := 0; i < len(line); i++ {
			if slices.Contains([]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}, line[i]) {
				left = line[i]
				break
			}
		}

		// right side
		for i := len(line) - 1; i >= 0; i-- {
			if slices.Contains([]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}, line[i]) {
				right = line[i]
				break
			}
		}

		lineResult, err := strconv.Atoi(string(left) + string(right))
		if err != nil {
			return "", err
		}
		result += lineResult
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(result), nil
}

func One_2(input io.Reader) (string, error) {
	var result int

	wordMap := map[string]rune{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	re := regexp.MustCompile(`^(zero|one|two|three|four|five|six|seven|eight|nine)`)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		var left, right rune

		// left side
		for i := 0; i < len(line); i++ {
			if slices.Contains([]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}, rune(line[i])) {
				left = rune(line[i])
				break
			}

			match := re.FindString(line[i:])
			if match != "" {
				left = wordMap[match]
				break
			}
		}

		// right side
		for i := len(line) - 1; i >= 0; i-- {
			if slices.Contains([]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}, rune(line[i])) {
				right = rune(line[i])
				break
			}

			match := re.FindString(line[i:])
			if match != "" {
				right = wordMap[match]
				break
			}
		}

		lineResult, err := strconv.Atoi(string(left) + string(right))
		if err != nil {
			return "", err
		}
		result += lineResult
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(result), nil
}
