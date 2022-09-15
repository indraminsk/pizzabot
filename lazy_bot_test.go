package main

import "testing"

func TestConvertDeltaToPath(t *testing.T) {
	tests := []struct {
		delta            int
		direct, opposite string
		expected         string
	}{
		{3, MoveEast, MoveWest, "EEE"},
		{-3, MoveEast, MoveWest, "WWW"},
		{3, MoveNorth, MoveSouth, "NNN"},
		{-3, MoveNorth, MoveSouth, "SSS"},
	}

	for _, test := range tests {
		output := convertDeltaToPath(test.delta, test.direct, test.opposite)

		if output != test.expected {
			t.Fatalf("expected: %v, output: %v", test.expected, output)
		}
	}
}

func TestBuildLazyPath(t *testing.T) {
	input := DistrictSchemeType{
		lanes:  5,
		houses: 5,
		receivers: ReceiversLHFormatSchemeType{
			1: HousesListType{3},
			4: HousesListType{4},
		},
		laneList: LanesListType{1, 4},
	}

	expected := "ENNNDEEEND"

	output := buildLazyPath(input)

	if output != expected {
		t.Fatalf("expected: %v, output: %v", expected, output)
	}
}
