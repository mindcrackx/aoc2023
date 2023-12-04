package main

import (
	"bufio"
	"log/slog"
	"os"
	"strings"
)

type Card struct {
	Number  int
	Winning []string
	Nums    []string
}

func (c Card) Calculate() []int {
	var counter int
	ret := []int{}
	for _, num := range c.Nums {
		for _, win := range c.Winning {
			if num == win {
				counter += 1
				ret = append(ret, c.Number+counter)
				break
			}
		}
	}

	return ret
}

func part2() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cards := make([]Card, 0, 1024)

	cardNum := 0
	for scanner.Scan() {
		cardNum += 1
		line := scanner.Text()
		splt := strings.Split(strings.Split(line, ":")[1], "|")

		winning := strings.Fields(splt[0])
		numbers := strings.Fields(splt[1])

		cards = append(cards, Card{
			Number:  cardNum,
			Winning: winning,
			Nums:    numbers,
		})

	}
	if err := scanner.Err(); err != nil {
		return err
	}

	startNums := make([]int, 0, len(cards))
	for _, c := range cards {
		startNums = append(startNums, c.Number)
	}

	result := calculate(startNums, cards)

	slog.Info("part2", "result", result)

	return nil
}

func calculate(nums []int, cards []Card) int {
	count := len(nums)
	for _, n := range nums {
		matches := cards[n-1].Calculate()
		if len(matches) > 0 {
			count += calculate(matches, cards)
		}
	}
	return count
}
