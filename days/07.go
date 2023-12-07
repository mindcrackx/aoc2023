package days

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/mindcrackx/aoc2023/utils"
)

type HandType string

const (
	HandTypeInvalid HandType = ""
	FiveOfAKind     HandType = "FiveOfAKind"
	FourOfAKind     HandType = "FourOfAKind"
	FullHouse       HandType = "FullHouse"
	ThreeOfAKind    HandType = "ThreeOfAKind"
	TwoPair         HandType = "TwoPair"
	OnePair         HandType = "OnePair"
	HighCard        HandType = "HighCard"
)

var HandTypeStrength = []HandType{FiveOfAKind, FourOfAKind, FullHouse, ThreeOfAKind, TwoPair, OnePair, HighCard}

type CamelCard string

const (
	CardInvalid CamelCard = ""
	CardA       CamelCard = "A"
	CardK       CamelCard = "K"
	CardQ       CamelCard = "Q"
	CardJ       CamelCard = "J"
	CardT       CamelCard = "T"
	Card9       CamelCard = "9"
	Card8       CamelCard = "8"
	Card7       CamelCard = "7"
	Card6       CamelCard = "6"
	Card5       CamelCard = "5"
	Card4       CamelCard = "4"
	Card3       CamelCard = "3"
	Card2       CamelCard = "2"
)

var CamelCardStrength = []CamelCard{CardA, CardK, CardQ, CardJ, CardT, Card9, Card8, Card7, Card6, Card5, Card4, Card3, Card2}
var CamelCardStrengthWithJoker = []CamelCard{CardA, CardK, CardQ, CardT, Card9, Card8, Card7, Card6, Card5, Card4, Card3, Card2, CardJ}

type Hand struct {
	Hand string
	Bet  int
}

func (h Hand) Type() HandType {
	type CardCount struct {
		Type  CamelCard
		Count int
	}
	counts := make([]CardCount, 0, len(CamelCardStrength))

	for _, c := range CamelCardStrength {
		cnt := strings.Count(h.Hand, string(c))
		counts = append(counts, CardCount{
			Type:  c,
			Count: cnt,
		})
	}

	// cmp(a, b) should return a negative number when a < b, a positive number when a > b and zero when a == b.
	slices.SortFunc(counts, func(a, b CardCount) int {
		if a.Count < b.Count {
			return 1
		}
		if a.Count > b.Count {
			return -1
		}
		return 0
	})

	switch counts[0].Count {
	case 1:
		return HighCard

	case 2:
		if counts[1].Count == 2 {
			return TwoPair
		}
		return OnePair

	case 3:
		if counts[1].Count == 2 {
			return FullHouse
		}
		return ThreeOfAKind

	case 4:
		return FourOfAKind

	case 5:
		return FiveOfAKind
	}

	if counts[0].Count == 1 {
		return HighCard
	}

	return HandTypeInvalid
}

func (h Hand) TypeWithJoker() HandType {
	type CardCount struct {
		Type  CamelCard
		Count int
	}
	counts := make([]CardCount, 0, len(CamelCardStrength))

	for _, c := range CamelCardStrength {
		cnt := strings.Count(h.Hand, string(c))
		counts = append(counts, CardCount{
			Type:  c,
			Count: cnt,
		})
	}

	// cmp(a, b) should return a negative number when a < b, a positive number when a > b and zero when a == b.
	slices.SortFunc(counts, func(a, b CardCount) int {
		if a.Count < b.Count {
			return 1
		}
		if a.Count > b.Count {
			return -1
		}
		return 0
	})

	jokerCount := 0
	for i := range counts {
		if counts[i].Type == CardJ {
			jokerCount = counts[i].Count
			counts[i].Count = 0
			break
		}
	}

	// cmp(a, b) should return a negative number when a < b, a positive number when a > b and zero when a == b.
	slices.SortFunc(counts, func(a, b CardCount) int {
		if a.Count < b.Count {
			return 1
		}
		if a.Count > b.Count {
			return -1
		}
		return 0
	})

	counts[0].Count += jokerCount

	switch counts[0].Count {
	case 1:
		return HighCard

	case 2:
		if counts[1].Count == 2 {
			return TwoPair
		}
		return OnePair

	case 3:
		if counts[1].Count == 2 {
			return FullHouse
		}
		return ThreeOfAKind

	case 4:
		return FourOfAKind

	case 5:
		return FiveOfAKind
	}

	if counts[0].Count == 1 {
		return HighCard
	}

	return HandTypeInvalid
}

func Seven_1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	handMap := make(map[HandType][]Hand)
	for _, name := range []HandType{FiveOfAKind, FourOfAKind, FullHouse, ThreeOfAKind, TwoPair, OnePair, HighCard} {
		handMap[name] = make([]Hand, 0)
	}

	for scanner.Scan() {
		line := scanner.Text()
		splt := strings.Split(line, " ")
		h := Hand{
			Hand: splt[0],
			Bet:  utils.MustAtoi(splt[1]),
		}

		handMap[h.Type()] = append(handMap[h.Type()], h)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	for key := range handMap {
		// cmp(a, b) should return a negative number when a < b, a positive number when a > b and zero when a == b.
		slices.SortFunc(handMap[key], func(a, b Hand) int {
			for i := 0; i < 5; i++ {
				if slices.Index(CamelCardStrength, CamelCard(a.Hand[i])) < slices.Index(CamelCardStrength, CamelCard(b.Hand[i])) {
					return -1
				}
				if slices.Index(CamelCardStrength, CamelCard(a.Hand[i])) > slices.Index(CamelCardStrength, CamelCard(b.Hand[i])) {
					return 1
				}
			}
			return 0
		})
	}

	rankings := make([]Hand, 0, 1024)
	for _, ht := range HandTypeStrength {
		rankings = append(rankings, handMap[ht]...)
	}

	var result int
	rankingCount := len(rankings)
	for i, h := range rankings {
		result += h.Bet * (rankingCount - i)
	}

	return strconv.Itoa(result), nil
}

func Seven_2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	handMap := make(map[HandType][]Hand)
	for _, name := range []HandType{FiveOfAKind, FourOfAKind, FullHouse, ThreeOfAKind, TwoPair, OnePair, HighCard} {
		handMap[name] = make([]Hand, 0)
	}

	for scanner.Scan() {
		line := scanner.Text()
		splt := strings.Split(line, " ")
		h := Hand{
			Hand: splt[0],
			Bet:  utils.MustAtoi(splt[1]),
		}

		handMap[h.TypeWithJoker()] = append(handMap[h.TypeWithJoker()], h)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	for key := range handMap {
		// cmp(a, b) should return a negative number when a < b, a positive number when a > b and zero when a == b.
		slices.SortFunc(handMap[key], func(a, b Hand) int {
			for i := 0; i < 5; i++ {
				if slices.Index(CamelCardStrengthWithJoker, CamelCard(a.Hand[i])) < slices.Index(CamelCardStrengthWithJoker, CamelCard(b.Hand[i])) {
					return -1
				}
				if slices.Index(CamelCardStrengthWithJoker, CamelCard(a.Hand[i])) > slices.Index(CamelCardStrengthWithJoker, CamelCard(b.Hand[i])) {
					return 1
				}
			}
			return 0
		})
	}

	rankings := make([]Hand, 0, 1024)
	for _, ht := range HandTypeStrength {
		rankings = append(rankings, handMap[ht]...)
	}

	var result int
	rankingCount := len(rankings)
	for i, h := range rankings {
		result += h.Bet * (rankingCount - i)
	}

	return strconv.Itoa(result), nil

}
