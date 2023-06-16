package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"strings"
	"strconv"
	"reflect"
	_ "unsafe" // use for Sizeof
)

type Person struct {
	form_id    string
	name       string
	email      string
	arr_date   string
	arr_time   string
	flight_num string
}

type Ride struct {
	date          string
	time          string
	peoplePerRide []Person
}

var ride_per_day = 2
var people_per_ride = 46

func main() {
	// store information of all people from CSV
	people := readCSV()
	fmt.Println(people)
	rides := calc(people, "9月20日")
}

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
	for /*i := 0; i < 53; i++*/ {
		flightInfo, err := r.Read()
		if err == io.EOF { //if end of file
			break
		}
		if err != nil { // if there is eror
			log.Fatal(err)
		}
		//generate fake dateset names and emails
		//a := [53]string{"name", "李华", "王强", "张伟", "刘洋", "陈杰", "杨秀", "黄磊", "赵婷", "孙涛", "周军", "吴杰", "郑浩", "朱峰", "马娟", "胡平", "林刚", "何勇", "高翔", "罗明", "谢霞", "宋飞", "韩鹏", "钟艳", "许健", "徐鑫", "程洁", "汪霞", "石鹏", "戴华", "魏军", "柳洋", "范峰", "彭平", "鲍霞", "安强", "代刚", "乐婷", "苏鹏", "农磊", "齐勇", "房婷", "费磊", "纪健", "滕峰", "段霞", "景婷", "章明", "巴勇", "侯霞", "宁飞", "柴健", "贝峰"}
		//b := [53]string{"email", "lihua01@uw.edu", "wangq02@uw.edu", "zhangw03@uw.edu", "liuy04@uw.edu", "chenj05@uw.edu", "yangx06@uw.edu", "huangl07@uw.edu", "zhaot08@uw.edu", "sunt09@uw.edu", "zhouj10@uw.edu", "wuj11@uw.edu", "zhengh12@uw.edu", "zhuf13@uw.edu", "maj14@uw.edu", "hup15@uw.edu", "ling16@uw.edu", "hey17@uw.edu", "gaor18@uw.edu", "luom19@uw.edu", "xiex20@uw.edu", "songf21@uw.edu", "hanp22@uw.edu", "zhongy23@uw.edu", "xuj24@uw.edu", "xup25@uw.edu", "chengj26@uw.edu", "wangx27@uw.edu", "ship28@uw.edu", "daih29@uw.edu", "weij30@uw.edu", "liuy31@uw.edu", "fanf32@uw.edu", "pengp33@uw.edu", "baox34@uw.edu", "ana35@uw.edu", "daig36@uw.edu", "ley37@uw.edu", "sup38@uw.edu", "nongl39@uw.edu", "qiy40@uw.edu", "fangt41@uw.edu", "feil42@uw.edu", "jij43@uw.edu", "tengf44@uw.edu", "duanx45@uw.edu", "jingt46@uw.edu", "zhangm47@uw.edu", "bay48@uw.edu", "houx49@uw.edu", "ningf50@uw.edu", "chaij51@uw.edu", "beif52@uw.edu"}
		//fmt.Printf("%s, %s, %s, %s, %s, %s, %s, %s, %s\n", flightInfo[0], a[i], b[i], flightInfo[3], flightInfo[4], flightInfo[5], flightInfo[6], flightInfo[7], flightInfo[8])
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

//calculate the 30minute time frame that has the most people landed on one day
//return the ride informationi that include the date, the time frame, and the
//people on that ride
func calc(people []Person, date string) []Ride {
	allRides := []Ride{}
	groupByArrTime := make(map[string][]Person)
	for _ , v range people {
		if v.arr_date == date {
			arrT = v.arr_time
			value, ok := groupByArrTime[arrT]
			if ok == true {
				groupByArrTime[arrT] = append(groupByArrTime[arrT], v)
			} else {
				pplLanded := []Person{v}
			}
		}
	}
	keys := reflect.ValueOf(groupByArrTime).MapKeys()
    fmt.Println(keys) 
	sort.Slice(keys, func(i, j int) bool {return convertToDatetime(date, keys[i]).Before(convertToDatetime(date, keys[j]))})
	fmt.Println("sorted keys:", keys)
	for i := 1; i < len(keys)
	retrun allRides
}

func convertToDatetime(date string, time string) {
	date := strings.Split(date, "月")
	day := strings.Trim(date[1], "日")
	monthInt, err := strconv.Atoi(date[0])
	fmt.Println(monthInt, err, reflect.TypeOf(monthInt))
	dayInt, err := strconv.Atoi(day)
	fmt.Println(dayInt, err, reflect.TypeOf(dayInt))

	time := strings.Split(time, ":")
	timeInt = []int{}
	for i := 0; i < 2; i++ {
		res, err := strconv.Atoi(time[i])
		timeInt = append(timeInt, res)
	}
	return time.Date(2023, monthInt, dayInt, timeInt[0], timeInt[1], 0, 100, time.Local)
}

func subtractTime(time1, time2 time.Time) float64 {
	diff := time2.Sub(time1).Minutes()
	return diff
}