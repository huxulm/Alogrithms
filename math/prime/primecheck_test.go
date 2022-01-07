package prime

import "testing"

func TestNaiveApproach(t *testing.T) {
	tests := []struct {
		title  string
		input  int64
		expect bool
	}{
		{"smallest prime", 2, true},
		{"random prime", 3, true},
		{"neither prime nor composite", 1, false},
		{"random non-prime", 10, false},
		{"another random prime", 23, true},
	}

	for _, c := range tests {
		t.Run(c.title, func(t *testing.T) {
			if output := NaiveApproach(int(c.input)); c.expect != output {
				t.Errorf("Expect %v but got %v", c.expect, output)
			}
		})
	}
}

func TestPairsApproach(t *testing.T) {
	tests := []struct {
		title  string
		input  int64
		expect bool
	}{
		{"smallest prime", 2, true},
		{"random prime", 3, true},
		{"neither prime nor composite", 1, false},
		{"random non-prime", 10, false},
		{"another random prime", 23, true},
	}

	for _, c := range tests {
		t.Run(c.title, func(t *testing.T) {
			if output := PairApproach(int(c.input)); c.expect != output {
				t.Errorf("Expect %v but got %v", c.expect, output)
			}
		})
	}
}
