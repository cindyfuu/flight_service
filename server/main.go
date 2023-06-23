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
	time          string
	peoplePerRide []Person
}

// para struct to store information about parameters passed in from client side
type para struct {
	ride_per_day      int
	time_frame_in_min int
	//people_per_ride int
}

func main() {
	// store information of all people from CSV
	people := readCSV()
	fmt.Println(people)
	rides := calc(people, "9月20日")
	info := para{
		ride_per_day:      2,
		time_frame_in_min: 30,
	}
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
func calc(people []Person, date string) []Ride {
	allRides := []Ride{}
	// create a map which key is the arr_time and value is the slice of person
	groupByArrTime := make(map[string][]Person)
	for _, v := range people {
		if v.arr_date == date {
			arrT := v.arr_time
			value, ok := groupByArrTime[arrT]
			if ok == true {
				groupByArrTime[arrT] = append(groupByArrTime[arrT], v)
			} else {
				pplLanded := []Person{v}
			}
		}
	}
	// value of keys are strings
	keys := reflect.ValueOf(groupByArrTime).MapKeys()
	fmt.Println(keys)
	keysInTime := []int{}
	for i := 0; i < len(keys); i++ {
		keysInTime[i] = convertToDatetime(date, keys[i])
	}
	sort.Slice(keysInTime, func(i, j int) bool { return keysInTime[i].Before(keysInTime[j]) })
	fmt.Println("sorted keys:", keysInTime)
	pplperT := []int{}
	for i := 0; i < len(keys); i++ {
		num := 0
		start := keysInTime[i]
		end := keysInTime[i].Add(timeFrameInMin * time.Minute)
		for j := i; j < len(keys); j++ {
			if keysInTime[j].Before(end) {
				num += len(groupByArrTime[j])
			}
		}
		pplPerT[i] = num
	}
	maxV := 0
	maxIndex := 0
	// need to consider this scenario: [1 3 4 7 7 5]
	for index, value := range pplPerT {
		if value > maxV {
			maxV = value
			maxIndex = index
		}
	}
	return allRides
}

// convertToDatetime converts the date and time strings to a time.Time object
func convertToDatetime(date string, time string) time.Time {
	dateSlice := strings.Split(date, "月")
	day := strings.Trim(dateSlice[1], "日")
	monthInt, err := strconv.Atoi(dateSlice[0])
	fmt.Println(monthInt, err, reflect.TypeOf(monthInt))
	dayInt, err := strconv.Atoi(day)
	fmt.Println(dayInt, err, reflect.TypeOf(dayInt))

	timeSlice := strings.Split(time, ":")
	timeInt := []int{}
	for i := 0; i < 2; i++ {
		res, err := strconv.Atoi(timeSlice[i])
		timeInt = append(timeInt, res)
	}
	return time.Date(2023, time.Month(monthInt), dayInt, timeInt[0], timeInt[1], 0, 100, time.Local)
}
