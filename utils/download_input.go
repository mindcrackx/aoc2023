package utils

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
)

func DownloadInput(sessionCookie string, year, day int, targetFolder string) error {
	if targetFolder == "" {
		targetFolder = "./inputs"
	}

	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "github.com/mindcrackx/aoc2023")
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = sessionCookie
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got unexpected status code %d", resp.StatusCode)
	}

	filePath := fmt.Sprintf("%s/%02d.txt", targetFolder, day)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("file %q already exists: %w", filePath, err)
		}
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
