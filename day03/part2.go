package main

import (
	"io"
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

	type Point struct {
		x int
		y int
	}
	type Val struct {
		start int
		end   int
		value int
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")

	// symbols
	symbols := make([]Point, 0, 1024)
	for y, line := range lines {
		for x, val := range line {
			if strings.Contains("*", string(val)) {
				symbols = append(symbols, Point{x, y})
			}
		}
	}

	// values
	values := make([][]Val, len(lines))
	for i := 0; i < len(lines); i++ {
		values[i] = make([]Val, 0, 64)
	}
	for y, line := range lines {
		start := -1
		value := ""
		for x, v := range line {
			if strings.Contains("0123456789", string(v)) {
				value += string(v)
				if start == -1 {
					start = x
				}
			} else {
				if start != -1 {
					v, err := strconv.Atoi(value)
					if err != nil {
						return err
					}
					values[y] = append(values[y], Val{start: start, end: x - 1, value: v})
				}

				start = -1
				value = ""
			}
			if start != -1 && x == len(line)-1 {
				v, err := strconv.Atoi(value)
				if err != nil {
					return err
				}
				values[y] = append(values[y], Val{start: start, end: x - 1, value: v})
			}
		}
	}

	// find gears
	gears := make([][]int, len(symbols))
	for i := 0; i < len(symbols); i++ {
		gears[i] = make([]int, 0, 64)
	}

	for y, row := range values {
		for _, v := range row {
			for i, s := range symbols {
				if (y == s.y || max(0, y-1) == s.y || min(len(values)-1, y+1) == s.y) &&
					((s.x >= max(0, v.start-1)) && (s.x <= min(len(lines[0]), v.end+1))) {
					gears[i] = append(gears[i], v.value)
				}
			}
		}
	}

	var sum int
	for _, g := range gears {
		if len(g) == 2 {
			sum += g[0] * g[1]
		}
	}

	slog.Info("part2", "result", sum)

	return nil

}
