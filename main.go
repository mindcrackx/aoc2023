package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/mindcrackx/aoc2023/days"
	"github.com/mindcrackx/aoc2023/utils"
)

const AocYear = 2023

func setup() map[int]map[int]func(io.Reader) (string, error) {
	lookup := make(map[int]map[int]func(io.Reader) (string, error))

	for day := 1; day <= 25; day++ {
		lookup[day] = make(map[int]func(io.Reader) (string, error))
	}

	lookup[1][1] = days.One_1
	lookup[1][2] = days.One_2

	return lookup
}

func main() {
	if err := run(); err != nil {
		slog.Error("during run", "err", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 3 {
		return fmt.Errorf("expected 2 args (day and part), but got %d", len(os.Args)-1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return fmt.Errorf("converting day %q to int: %w", os.Args[1], err)
	}
	if day <= 0 || day > 25 {
		return fmt.Errorf("expected day between 1 and 25, but got %d", day)
	}

	part, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return fmt.Errorf("converting part %q to int: %w", os.Args[2], err)
	}
	if part <= 0 || part > 2 {
		return fmt.Errorf("expected part to be 1 or 2, but got %d", part)
	}

	lookup := setup()
	slog.Info("running...", "day", day, "part", part)

	file, err := os.Open(fmt.Sprintf("./inputs/%02d.txt", day))
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("opening input file for day %d: %w", day, err)
		}

		// no cached file, download
		slog.Info("downloading input", "day", day)

		sessionCookie := os.Getenv("AOC_SESSION_COOKIE")
		if sessionCookie == "" {
			return errors.New("no session cookie in env var 'AOC_SESSION_COOKIE' found")
		}

		err = utils.DownloadInput(sessionCookie, AocYear, day, "")
		if err != nil {
			return fmt.Errorf("downloading input file: %w", err)
		}

		file, err = os.Open(fmt.Sprintf("./inputs/%02d.txt", day))
		if err != nil {
			return err
		}
	}
	defer file.Close()

	start := time.Now()

	result, err := lookup[day][part](file)
	if err != nil {
		return fmt.Errorf("day=%d part=%d: %w", day, part, err)
	}

	end := time.Now()

	slog.Info("timing", "day", day, "part", part, "duration", end.Sub(start))

	fmt.Println("Result:")
	fmt.Println(result)

	return nil
}
