package main

import (
	"reflect"
	"testing"
	"time"
)

func TestGetTopRidesOverlapped(t *testing.T) {
	layout := "01/02/2006 15:04"
	start1, _ := time.Parse(layout, "09/20/2023 08:00")
	end1, _ := time.Parse(layout, "09/20/2023 08:30")
	start2, _ := time.Parse(layout, "09/20/2023 08:30")
	end2, _ := time.Parse(layout, "09/20/2023 09:00")
	start3, _ := time.Parse(layout, "09/20/2023 09:30")
	end3, _ := time.Parse(layout, "09/20/2023 10:00")

	rides := []Ride{
		{
			date:          "09/20",
			start:         start1,
			end:           end1,
			peoplePerRide: nil,
			count:         10,
		},
		{
			date:          "09/20",
			start:         start2,
			end:           end2,
			peoplePerRide: nil,
			count:         9,
		},
		{
			date:          "09/20",
			start:         start3,
			end:           end3,
			peoplePerRide: nil,
			count:         8,
		},
	}

	expected := []RideFinal{
		{
			date:          "09/20",
			start:         "09/20/2023 09:30",
			end:           "09/20/2023 10:00",
			peoplePerRide: nil,
			count:         8,
		},
		{
			date:          "09/20",
			start:         "09/20/2023 08:00",
			end:           "09/20/2023 08:30",
			peoplePerRide: nil,
			count:         10,
		},
	}
	result := getTopRides(rides, 2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("getTopRides was incorrect, got: %v, want: %v.", result, expected)
	}
}
