package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
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
	start         time.Time
	end           time.Time
	peoplePerRide [][]Person
	count         int
}

// Slightly moderated ride struct that the start and end field are in string which is more readable than time.Time
type RideFinal struct {
	date          string
	start         string
	end           string
	peoplePerRide [][]Person
	count         int
}

// request struct to store information about parameters passed in from client side
type Request struct {
	ride_per_day      int
	time_frame_in_min int
	people            []Person
	date              string
}

// Response struct to store all information returning to client
type Response struct {
	topNumberRides []RideFinal
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
	requestInfo := Request{
		ride_per_day:      2,
		time_frame_in_min: 30,
		people:            people,
		date:              "09/20",
	}
	allRides := calc(&requestInfo)
	final := Response{
		topNumberRides: getTopRides(allRides, requestInfo.ride_per_day),
	}
	fmt.Println(final.topNumberRides[0])
	fmt.Println()
	fmt.Println(final.topNumberRides[1])
}

// readCSV reads flight_info_clean.csv into a slice of person struct
func readCSV() []Person {
	people := []Person{}
	// Open the file
	csvfile, err := os.Open("../flight_info_new.csv")
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

// calc function calculates all possible 30 min time frame ride
// Para: info - request from client
// Return: a slice of ride struct that include all possible ride
func calc(info *Request) []Ride {
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
	// keys is a slice of string where value are the keys of groupByArrTime
	keys := []string{}
	for k := range groupByArrTime {
		keys = append(keys, k)
	}
	// keysInTime is a slice of time.Time where value are the time.Time type of keys slice
	keysInTime := []time.Time{}
	for i := 0; i < len(keys); i++ {
		keysInTime = append(keysInTime, convertToDatetime(info.date, keys[i]))
	}
	//sort keysInTime in ascending order
	sort.Slice(keysInTime, func(i, j int) bool { return keysInTime[i].Before(keysInTime[j]) })
	//try to calculate number of people landed for each 30mintue time frame
	//groupByArrTimeInTime is the same map as groupByArrTime except the keys are in time.Time
	groupByArrTimeInTime := make(map[time.Time][]Person)
	for key, value := range groupByArrTime {
		newKey := convertToDatetime(info.date, key)
		groupByArrTimeInTime[newKey] = value
	}
	allRides := calcTimeInter(keysInTime, info, groupByArrTimeInTime)
	return allRides
}

// calculate all possible 30 min rides
// para: keys - arrival time in time.Time type
//
//	info - request info from client
//	groupByArrTime - record all arrrival time in time.Time and its respectively people landing at that time
//
// return: all possible rides in the time frame
func calcTimeInter(keys []time.Time, info *Request, groupByArrTime map[time.Time][]Person) []Ride {
	allRides := []Ride{}
	start := 0
	end := 0
	temp := [][]Person{}
	count := 0
	for !(start == end && start == len(keys)-1) {
		ifIn30Min := keys[start].Add(time.Duration(info.time_frame_in_min) * time.Minute).After(keys[end])
		if keys[start].Equal(keys[end]) {
			end++
		} else if ifIn30Min {
			if len(temp) == 0 {
				temp = append(temp, groupByArrTime[keys[start]])
				count = len(groupByArrTime[keys[start]])
			}
			temp = append(temp, groupByArrTime[keys[end]])
			count = count + len(groupByArrTime[keys[end]])
			fmt.Println(keys[start])
			fmt.Println(keys[end])
			ride := Ride{
				date:          info.date,
				start:         keys[start],
				end:           keys[end],
				peoplePerRide: temp,
				count:         count,
			}
			allRides = append(allRides, ride)
			end++
		} else {
			num := len(temp[0])
			temp = temp[1:]
			count = count - num
			start++
		}
	}
	return allRides
}

// convertToDatetime converts the date and time strings to a time.Time object
func convertToDatetime(date string, exactTime string) time.Time {
	dateSlice := strings.Split(date, "/")
	monthInt, _ := strconv.Atoi(dateSlice[0])
	dayInt, _ := strconv.Atoi(dateSlice[1])

	timeSlice := strings.Split(exactTime, ":")
	timeInt := []int{}
	for i := 0; i < 2; i++ {
		res, _ := strconv.Atoi(timeSlice[i])
		timeInt = append(timeInt, res)
	}
	return time.Date(2023, time.Month(monthInt), dayInt, timeInt[0], timeInt[1], 0, 100, time.Local)
}

// RidePairList implements sort.Interface for []Ride based on the count field
func (rpl RidePairList) Len() int           { return len(rpl) }
func (rpl RidePairList) Less(i, j int) bool { return rpl[i].Value > rpl[j].Value } // reverse order
func (rpl RidePairList) Swap(i, j int)      { rpl[i], rpl[j] = rpl[j], rpl[i] }

// Binary search to find the latest ride (before current ride) that doesn't conflict with the current ride.
// arr[i] should be sorted in increasing order of start time
func latestNonConflict(arr []Ride, i int) int {
	lo := 0
	hi := i - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid].end.Before(arr[i].start) {
			if arr[mid+1].end.Before(arr[i].start) {
				lo = mid + 1
			} else {
				return mid
			}
		} else {
			hi = mid - 1
		}
	}
	return -1
}

// getTopRides returns the top k rides that have the most people and that do not overlap in time
func getTopRides(allRides []Ride, num int) []RideFinal {
	// Sort rides according to start time
	sort.Slice(allRides, func(i, j int) bool { return allRides[i].start.Before(allRides[j].start) })

	// Create an array to store solutions of subproblems. table[i] stores the maximum people count ending at arr[i]
	n := len(allRides)
	table := make([]int, n)
	table[0] = allRides[0].count // First value in table should be count of first ride

	// Fill table[] using recursive property
	for i := 1; i < n; i++ {
		// Find count including current ride
		incl := allRides[i].count
		l := latestNonConflict(allRides, i)
		if l != -1 {
			incl += table[l]
		}

		// Store maximum of including and excluding
		table[i] = max(incl, table[i-1])
	}

	// The last entry in table[] stores the maximum count
	maxCount := table[n-1]

	// Initialize result
	res := make([]RideFinal, 0, num)

	// Traverse through table[] to find out which rides are included in result
	for i := n - 1; i >= 0; i-- {
		// If this ride is included
		if (i == 0 && maxCount > 0) || maxCount != table[i-1] {
			ride := RideFinal{
				date:          allRides[i].date,
				start:         allRides[i].start.Format("01/02/2006 15:04"),
				end:           allRides[i].end.Format("01/02/2006 15:04"),
				peoplePerRide: allRides[i].peoplePerRide,
				count:         allRides[i].count,
			}
			// This ride is included in result
			res = append(res, ride)
			// Since this ride is included its count should be subtracted
			maxCount -= allRides[i].count
			if len(res) == num {
				break
			}
		}
	}
	return res
}

// max returns the maximum of two integers
func max(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}
