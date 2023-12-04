package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		slog.Error("during run:", "err", err)
		os.Exit(1)
	}
}

func run() error {
	var err error
	part := 0
	if len(os.Args) > 1 {
		part, err = strconv.Atoi(os.Args[1])
		if err != nil {
			return fmt.Errorf("could not parse part %q as int: %w", os.Args[1], err)
		}
	}

	switch part {
	case 0:
		if err := part1(); err != nil {
			return fmt.Errorf("part 1: %w", err)
		}
		if err := part2(); err != nil {
			return fmt.Errorf("part 2: %w", err)
		}
	case 1:
		return part1()
	case 2:
		return part2()
	default:
		return fmt.Errorf("invalid part number %d, must be 0 1 2", part)
	}

	return nil
}
