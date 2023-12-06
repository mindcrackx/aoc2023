package days

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Four_1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var sum int

	cardNum := 0
	for scanner.Scan() {
		cardNum += 1
		line := scanner.Text()
		splt := strings.Split(strings.Split(line, ":")[1], "|")

		winning := strings.Fields(splt[0])
		numbers := strings.Fields(splt[1])

		var points int
		var found bool

		for _, num := range numbers {
			for _, win := range winning {
				if num == win {
					if !found {
						points += 1
						found = true
						break
					}
					points *= 2
					break
				}
			}
		}

		if found {
			sum += points
		}

	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(sum), nil
}

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

func Four_2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

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
		return "", err
	}

	startNums := make([]int, 0, len(cards))
	for _, c := range cards {
		startNums = append(startNums, c.Number)
	}

	result := calculate(startNums, cards)

	return strconv.Itoa(result), nil
}
