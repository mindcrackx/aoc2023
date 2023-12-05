package days_test

import (
	"strings"
	"testing"

	"github.com/mindcrackx/aoc2023/days"
)

func TestNineteen_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		expect string
	}{}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := days.Nineteen_1(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("got unexpected error: %s", err.Error())
			}

			if result != tt.expect {
				t.Fatalf("expected %q but got %q instead", tt.expect, result)
			}
		})
	}
}

func TestNineteen_2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		expect string
	}{}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := days.Nineteen_2(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("got unexpected error: %s", err.Error())
			}

			if result != tt.expect {
				t.Fatalf("expected %q but got %q instead", tt.expect, result)
			}
		})
	}
}
