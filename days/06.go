package days

import (
	"io"
	"strconv"
	"strings"

	"github.com/mindcrackx/aoc2023/utils"
)

func Six_1(input io.Reader) (string, error) {
	data, err := io.ReadAll(input)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")
	var times, distances = make([]int, 0, 4), make([]int, 0, 4)

	for _, x := range strings.Fields(strings.Split(lines[0], ":")[1]) {
		times = append(times, utils.MustAtoi(x))
	}
	for _, x := range strings.Fields(strings.Split(lines[1], ":")[1]) {
		distances = append(distances, utils.MustAtoi(x))
	}

	results := make([]int, 0, len(times))
	for i, t := range times {
		further := 0
		for x := 0; x <= t; x++ {
			res := (t - x) * x
			// slog.Info("", "t", t, "d", distances[i], "x", x, "dist", (t-x)*x)
			if res > distances[i] {
				further += 1
			}
		}
		results = append(results, further)
	}

	result := results[0]
	for _, r := range results[1:] {
		result *= r
	}

	return strconv.Itoa(result), nil
}

func Six_2(input io.Reader) (string, error) {
	data, err := io.ReadAll(input)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")
	var time, distance int

	a := strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", "")
	time = utils.MustAtoi(a)
	b := strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", "")
	distance = utils.MustAtoi(b)

	result := 0
	for x := 0; x <= time; x++ {
		res := (time - x) * x
		// slog.Info("", "t", t, "d", distances[i], "x", x, "dist", (t-x)*x)
		if res > distance {
			result += 1
		}
	}

	return strconv.Itoa(result), nil

}
