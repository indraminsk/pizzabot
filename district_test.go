package main

import (
	"reflect"
	"testing"
)

func TestBuildLaneList(t *testing.T) {
	input := DistrictSchemeType{
		lanes:  5,
		houses: 5,
		receivers: ReceiversLHFormatSchemeType{
			4: HousesListType{4},
			1: HousesListType{3},
		},
		laneList: nil,
	}

	expected := DistrictSchemeType{
		lanes:  5,
		houses: 5,
		receivers: ReceiversLHFormatSchemeType{
			4: HousesListType{4},
			1: HousesListType{3},
		},
		laneList: LanesListType{1, 4},
	}

	buildLaneList(&input)

	if !reflect.DeepEqual(input, expected) {
		t.Fatalf("expected: %v, output: %v", expected, input)
	}
}

func TestSortHousesPerLane(t *testing.T) {
	input := DistrictSchemeType{
		lanes:  5,
		houses: 5,
		receivers: ReceiversLHFormatSchemeType{
			4: HousesListType{4, 2},
			1: HousesListType{3, 1},
		},
		laneList: nil,
	}

	expected := DistrictSchemeType{
		lanes:  5,
		houses: 5,
		receivers: ReceiversLHFormatSchemeType{
			4: HousesListType{2, 4},
			1: HousesListType{1, 3},
		},
		laneList: nil,
	}

	sortHousesPerLane(&input)

	if !reflect.DeepEqual(input, expected) {
		t.Fatalf("expected: %v, output: %v", expected, input)
	}
}

func TestBuildDistrictScheme(t *testing.T) {
	input := "5x5 (1, 3) (4, 4)"
	expected := DistrictSchemeType{
		lanes:  5,
		houses: 5,
		receivers: ReceiversLHFormatSchemeType{
			1: HousesListType{3},
			4: HousesListType{4},
		},
		laneList: LanesListType{1, 4},
	}

	output, err := buildDistrictScheme(input)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("expected: %v, output: %v", expected, output)
	}
}

func TestBuildDistrictSchemeFailed(t *testing.T) {
	tests := []struct {
		input string
	}{
		{""},
		{" "},
		{"5x5 "},
		{"(1, 3) (4, 4)"},
		{"55 (1, 3) (4, 4)"},
		{"x5 (1, 3) (4, 4)"},
		{"5x (1, 3) (4, 4)"},
		{"5x5 (1,3) (4, 4)"},
		{"mxn (1, 3) (4, 4)"},
		{"5x5 (x, y) (4, 4)"},
		{"5x5 (1, 5) (4, 4)"},
	}

	for _, test := range tests {
		_, err := buildDistrictScheme(test.input)
		if err == nil {
			t.Fatal("expected error")
		}
	}
}
