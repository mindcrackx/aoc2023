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

	lookup[2][1] = days.Two_1
	lookup[2][2] = days.Two_2

	lookup[3][1] = days.Three_1
	lookup[3][2] = days.Three_2

	lookup[4][1] = days.Four_1
	lookup[4][2] = days.Four_2

	lookup[5][1] = days.Five_1
	lookup[5][2] = days.Five_2

	lookup[6][1] = days.Six_1
	lookup[6][2] = days.Six_2

	lookup[7][1] = days.Seven_1
	lookup[7][2] = days.Seven_2

	lookup[8][1] = days.Eight_1
	lookup[8][2] = days.Eight_2

	lookup[9][1] = days.Nine_1
	lookup[9][2] = days.Nine_2

	lookup[10][1] = days.Ten_1
	lookup[10][2] = days.Ten_2

	lookup[11][1] = days.Eleven_1
	lookup[11][2] = days.Eleven_2

	lookup[12][1] = days.Twelve_1
	lookup[12][2] = days.Twelve_2

	lookup[13][1] = days.Thirteen_1
	lookup[13][2] = days.Thirteen_2

	lookup[14][1] = days.Fourteen_1
	lookup[14][2] = days.Fourteen_2

	lookup[15][1] = days.Fifteen_1
	lookup[15][2] = days.Fifteen_2

	lookup[16][1] = days.Sixteen_1
	lookup[16][2] = days.Sixteen_2

	lookup[17][1] = days.Seventeen_1
	lookup[17][2] = days.Seventeen_2

	lookup[18][1] = days.Eighteen_1
	lookup[18][2] = days.Eighteen_2

	lookup[19][1] = days.Nineteen_1
	lookup[19][2] = days.Nineteen_2

	lookup[20][1] = days.Twenty_1
	lookup[20][2] = days.Twenty_2

	lookup[21][1] = days.Twentyone_1
	lookup[21][2] = days.Twentyone_2

	lookup[22][1] = days.Twentytwo_1
	lookup[22][2] = days.Twentytwo_2

	lookup[23][1] = days.Twentythree_1
	lookup[23][2] = days.Twentythree_2

	lookup[24][1] = days.Twentyfour_1
	lookup[24][2] = days.Twentyfour_2

	lookup[25][1] = days.Twentyfive_1
	lookup[25][2] = days.Twentyfive_2

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
