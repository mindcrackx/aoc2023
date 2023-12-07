package days

import (
	"cmp"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/mindcrackx/aoc2023/utils"
)

type Map struct {
	Name   string
	Values []MapValue
}

func (m *Map) Next(x int) int {
	next := x
	for _, v := range m.Values {
		if (x >= v.SourceRangeStart) && (x < (v.SourceRangeStart + v.RangeLength)) {
			next = x + v.DestinationRangeStart - v.SourceRangeStart
		}
	}
	return next
}

type MapValue struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

type Seed struct {
	Seed        int
	Soil        int
	Fertilizer  int
	Water       int
	Light       int
	Temperature int
	Humidity    int
	Location    int
}

func Five_1(input io.Reader) (string, error) {
	dataBytes, err := io.ReadAll(input)
	if err != nil {
		return "", err
	}

	data := strings.Split(string(dataBytes), "\n\n")
	if len(data) != 8 {
		return "", fmt.Errorf("expected 8 chunks but got %d", len(data))
	}

	seedsStr := strings.Fields(strings.Split(data[0], ":")[1])
	seeds := make([]int, 0, len(seedsStr))
	for _, val := range seedsStr {
		x := utils.MustAtoi(val)
		seeds = append(seeds, x)
	}

	maps := make([]Map, 0, 7)

	mapNames := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	for i, name := range mapNames {
		m := Map{
			Name:   name,
			Values: make([]MapValue, 0, 64),
		}
		for _, line := range strings.Split(data[i+1], "\n")[1:] {
			fields := strings.Fields(line)
			if len(fields) != 3 {
				continue
			}
			m.Values = append(m.Values, MapValue{
				DestinationRangeStart: utils.MustAtoi(fields[0]),
				SourceRangeStart:      utils.MustAtoi(fields[1]),
				RangeLength:           utils.MustAtoi(fields[2]),
			})
		}
		maps = append(maps, m)
	}

	locations := make([]int, 0, len(seeds))
	for i := range seeds {
		next := seeds[i]
		for _, m := range maps {
			next = m.Next(next)
		}
		locations = append(locations, next)
	}

	result := locations[0]
	for _, loc := range locations[1:] {
		if loc < result {
			result = loc
		}
	}

	return strconv.Itoa(result), nil
}

func (m *Map) NextRanges(inputRanges []Range) []Range {
	var answerRanges []Range
	var resultRanges []Range
	resultRanges = append(resultRanges, inputRanges...)

	for _, v := range m.Values {
		var newRanges []Range
		for _, ir := range resultRanges {
			// [start                                end)
			//           [src_start    src_end]
			// [BEFORE  ][INTERSECTION        ][AFTER   )
			before := Range{
				Start: ir.Start,
				End:   min(ir.End, v.SourceRangeStart),
			}
			inter := Range{
				Start: max(ir.Start, v.SourceRangeStart),
				End:   min(ir.End, v.SourceRangeStart+v.RangeLength),
			}
			after := Range{
				Start: max(ir.Start, v.SourceRangeStart+v.RangeLength),
				End:   ir.End,
			}

			if before.End > before.Start {
				newRanges = append(newRanges, before)
			}
			if inter.End > inter.Start {
				answerRanges = append(answerRanges, Range{
					Start: inter.Start - v.SourceRangeStart + v.DestinationRangeStart,
					End:   inter.End - v.SourceRangeStart + v.DestinationRangeStart,
				})
			}
			if after.End > after.Start {
				newRanges = append(newRanges, after)
			}
		}

		resultRanges = newRanges
	}

	return append(answerRanges, resultRanges...)
}

type Range struct {
	Start int
	End   int
}

func Five_2(input io.Reader) (string, error) {
	dataBytes, err := io.ReadAll(input)
	if err != nil {
		return "", err
	}

	data := strings.Split(string(dataBytes), "\n\n")
	if len(data) != 8 {
		return "", fmt.Errorf("expected 8 chunks but got %d", len(data))
	}

	seedsStr := strings.Fields(strings.Split(data[0], ":")[1])
	seedRanges := make([]Range, 0, len(seedsStr)/2)
	for i := 0; i < len(seedsStr)-1; i += 2 {
		r := Range{Start: utils.MustAtoi(seedsStr[i])}
		r.End = r.Start + utils.MustAtoi(seedsStr[i+1])
		seedRanges = append(seedRanges, r)
	}

	maps := make([]Map, 0, 7)

	mapNames := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	for i, name := range mapNames {
		m := Map{
			Name:   name,
			Values: make([]MapValue, 0, 64),
		}
		for _, line := range strings.Split(data[i+1], "\n")[1:] {
			fields := strings.Fields(line)
			if len(fields) != 3 {
				continue
			}
			m.Values = append(m.Values, MapValue{
				DestinationRangeStart: utils.MustAtoi(fields[0]),
				SourceRangeStart:      utils.MustAtoi(fields[1]),
				RangeLength:           utils.MustAtoi(fields[2]),
			})
		}
		maps = append(maps, m)
	}

	var results []int
	for _, r := range seedRanges {
		newRanges := []Range{r}
		for _, m := range maps {
			newRanges = m.NextRanges(newRanges)
		}

		minLocationRange := slices.MinFunc(newRanges, func(a, b Range) int {
			return cmp.Compare(a.Start, b.Start)
		})
		results = append(results, minLocationRange.Start)
	}

	result := slices.Min(results)

	return strconv.Itoa(result), nil

}
