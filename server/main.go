package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Person struct to store information about a person
type Person struct {
	form_id    string
	name       string
	email      string
	arr_date   string
	arr_time   string
	flight_num string
}

// Ride struct to store information about a ride
type Ride struct {
	date          string
	start         string
	end           string
	peoplePerRide []Person
}

// request struct to store information about parameters passed in from client side
type request struct {
	ride_per_day      int
	time_frame_in_min int
	//people_per_ride int
	people []Person
	date   string
}

type Response struct {
	topNumberRides []Ride
}

// Pair struct will hold a key-value pair
type RidePair struct {
	Key   Ride
	Value int
}

// PairList is a slice of Pairs that implements sort.Interface to sort by Value.
type RidePairList []RidePair

func main() {
	// store information of all people from CSV
	people := readCSV()
	requestInfo := request{
		ride_per_day:      2,
		time_frame_in_min: 30,
		people:            people,
		date:              "9月20日",
	}
	allRides := calc(&requestInfo)
	final := Response{
		topNumberRides: getTopRides(allRides, requestInfo.ride_per_day),
	}
	fmt.Println(final)
}

func (p RidePairList) Len() int           { return len(p) }
func (p RidePairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p RidePairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func getTopRides(allRides []Ride, num int) []Ride {
	p := make(RidePairList, len(allRides))

	for i, ride := range allRides {
		p[i] = RidePair{ride, len(ride.peoplePerRide)}
	}

	sort.Sort(sort.Reverse(p))

	// Limit to top 2 rides
	topRides := make([]Ride, num)
	for i := 0; i < num; i++ {
		topRides[i] = p[i].Key
	}

	return topRides
}

// readCSV reads flight_info_clean.csv into a slice of person struct
func readCSV() []Person {
	people := []Person{}
	// Open the file
	csvfile, err := os.Open("../flight_info_clean.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)
	// Iterate through the records
	for {
		flightInfo, err := r.Read()
		if err == io.EOF { //if end of file
			break
		}
		if err != nil { // if there is an error
			log.Fatal(err)
		}
		if flightInfo[0] != "form_id" && flightInfo[8] != "flight_num" {
			personInfo := Person{
				form_id:    flightInfo[0],
				name:       flightInfo[1],
				email:      flightInfo[2],
				arr_date:   flightInfo[6],
				arr_time:   flightInfo[7],
				flight_num: flightInfo[8],
			}
			people = append(people, personInfo)
		}
	}
	return people
}

// calc calculates the 30-minute time frame that has the most people landed on one day
// It returns the ride information that includes the date, the time frame, and the people on that ride
func calc(info *request) []Ride {
	allRides := []Ride{}
	// create a map which key is the arr_time and value is the slice of person
	groupByArrTime := make(map[string][]Person)
	for _, v := range info.people {
		if v.arr_date == info.date {
			arrT := v.arr_time
			_, ok := groupByArrTime[arrT]
			if ok == true {
				groupByArrTime[arrT] = append(groupByArrTime[arrT], v)
			} else {
				groupByArrTime[arrT] = []Person{v}
			}
		}
	}
	// value of keys are strings
	keys := make([]string, len(groupByArrTime))
	for k := range groupByArrTime {
		keys = append(keys, k)
	}
	fmt.Println("keys:", keys)
	keysInTime := []time.Time{}
	for i := 0; i < len(keys); i++ {
		keysInTime[i] = convertToDatetime(info.date, keys[i])
	}
	sort.Slice(keysInTime, func(i, j int) bool { return keysInTime[i].Before(keysInTime[j]) })
	fmt.Println("sorted keys:", keysInTime)
	//try to calculate number of people landed for each 30mintue time frame
	groupByArrTimeInTime := make(map[time.Time][]Person)
	for key, value := range groupByArrTime {
		newKey := convertToDatetime(info.date, key)
		groupByArrTimeInTime[newKey] = value
	}
	allRides = calculateTimeInterval(0, 1, keysInTime, groupByArrTimeInTime, allRides, info.time_frame_in_min)
	return allRides
}

func calculateTimeInterval(start int, end int, keys []time.Time, groupByArrTime map[time.Time][]Person, allRides []Ride, timeFrame int) []Ride {
	if start == len(keys) || end == len(keys) {
		return allRides
	}
	ifIn30Min := keys[start].Add(time.Duration(timeFrame) * time.Minute).After(keys[end])
	if keys[start].Equal(keys[end]) {
		allRides = calculateTimeInterval(start, end+1, keys, groupByArrTime, allRides, timeFrame)
	} else if ifIn30Min {
		temp := groupByArrTime[keys[start]]
		for i := start + 1; i <= end; i++ {
			temp = append(temp, groupByArrTime[keys[i]]...)
		}
		allRides = calculateTimeInterval(start, end+1, keys, groupByArrTime, allRides, timeFrame)
	} else {
		allRides = calculateTimeInterval(start+1, end, keys, groupByArrTime, allRides, timeFrame)
	}
	return allRides
}

// convertToDatetime converts the date and time strings to a time.Time object
func convertToDatetime(date string, exactTime string) time.Time {
	dateSlice := strings.Split(date, "月")
	day := strings.Trim(dateSlice[1], "日")
	monthInt, err := strconv.Atoi(dateSlice[0])
	fmt.Println(monthInt, err, reflect.TypeOf(monthInt))
	dayInt, err := strconv.Atoi(day)
	fmt.Println(dayInt, err, reflect.TypeOf(dayInt))

	timeSlice := strings.Split(exactTime, ":")
	timeInt := []int{}
	for i := 0; i < 2; i++ {
		res, _ := strconv.Atoi(timeSlice[i])
		timeInt = append(timeInt, res)
	}
	return time.Date(2023, time.Month(monthInt), dayInt, timeInt[0], timeInt[1], 0, 100, time.Local)
}
