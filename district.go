package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type LanesListType []int
type HousesListType []int

type ReceiversLHFormatSchemeType map[int]HousesListType

type DistrictSchemeType struct {
	lanes, houses int
	receivers     ReceiversLHFormatSchemeType
	laneList      LanesListType
}

// buildLaneList sorts slice with lanes' numbers and then stores sorted data
func buildLaneList(markupData *DistrictSchemeType) {
	var (
		lanesList LanesListType
	)

	for lane := range markupData.receivers {
		lanesList = append(lanesList, lane)
	}

	sort.Slice(lanesList, func(i, j int) bool {
		return lanesList[i] < lanesList[j]
	})

	markupData.laneList = lanesList
}

// sortHousesPerLane sorts houses' numbers per lane
func sortHousesPerLane(markupData *DistrictSchemeType) {
	var (
		isNeedRevers bool
	)

	for _, lane := range markupData.laneList {
		houses := markupData.receivers[lane]

		if isNeedRevers {
			sort.Slice(houses, func(i, j int) bool {
				return houses[i] > houses[j]
			})
		} else {
			sort.Slice(houses, func(i, j int) bool {
				return houses[i] < houses[j]
			})
		}

		isNeedRevers = !isNeedRevers
		markupData.receivers[lane] = houses
	}
}

func buildDistrictScheme(districtLegend string) (DistrictSchemeType, error) {
	var (
		err error

		legendItems, sizeDistrictMatrix []string
		districtScheme                  DistrictSchemeType
	)

	if len(strings.TrimSpace(districtLegend)) == 0 {
		return districtScheme, errors.New("error [build district scheme]: empty legend")
	}

	legendItems = strings.Split(districtLegend, " ")
	sizeDistrictMatrix = strings.Split(legendItems[0], "x")

	if len(sizeDistrictMatrix) != 2 {
		return districtScheme, errors.New("error [build district scheme]: wrong extension")
	}

	districtScheme.lanes, err = strconv.Atoi(sizeDistrictMatrix[0])
	if err != nil {
		return districtScheme, errors.New("error [build district scheme]: wrong extension")
	}

	districtScheme.houses, err = strconv.Atoi(sizeDistrictMatrix[1])
	if err != nil {
		return districtScheme, errors.New("error [build district scheme]: wrong extension")
	}

	districtLegend = strings.Replace(districtLegend, legendItems[0]+" ", "", 1)
	legendItems = strings.Split(districtLegend, "(")

	districtScheme.receivers = make(ReceiversLHFormatSchemeType)

	for _, point := range legendItems[1:] {
		var (
			pointItems  []string
			lane, house int
		)

		pointItems = strings.Split(point, "")

		lane, err = strconv.Atoi(pointItems[0])
		if err != nil {
			return districtScheme, err
		}

		house, err = strconv.Atoi(pointItems[3])
		if err != nil {
			return districtScheme, err
		}

		if (lane >= districtScheme.lanes) || (house >= districtScheme.houses) {
			return districtScheme, errors.New(fmt.Sprintf("error [build district scheme]: lane out of field range in point (%d, %d)", lane, house))
		}

		districtScheme.receivers[lane] = append(districtScheme.receivers[lane], house)
	}

	if len(districtScheme.receivers) == 0 {
		return districtScheme, errors.New("error [build district scheme]: empty list of receivers")
	}

	buildLaneList(&districtScheme)
	sortHousesPerLane(&districtScheme)

	return districtScheme, nil
}
