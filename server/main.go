package main

import (
	"encoding/csv"
	_ "fmt"
	"io"
	"log"
	"os"
	_ "time"
	_ "unsafe" // use for Sizeof
)

type person struct {
	form_id    string
	name       string
	email      string
	arr_date   string
	arr_time   string
	flight_num string
}

type ride struct {
	date   string
	time   string
	people []person
}

func main() {
	// Open the file
	csvfile, err := os.Open("../flight_info.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	// Iterate through the records
	for i := 0; i < 53; i++ {
		flightInfo, err := r.Read()
		if err == io.EOF { //if end of file
			break
		}
		if err != nil { // if there is eror
			log.Fatal(err)
		}
		//a := [53]string{"", "李华", "王强", "张伟", "刘洋", "陈杰", "杨秀", "黄磊", "赵婷", "孙涛", "周军", "吴杰", "郑浩", "朱峰", "马娟", "胡平", "林刚", "何勇", "高翔", "罗明", "谢霞", "宋飞", "韩鹏", "钟艳", "许健", "徐鑫", "程洁", "汪霞", "石鹏", "戴华", "魏军", "柳洋", "范峰", "彭平", "鲍霞", "安强", "代刚", "乐婷", "苏鹏", "农磊", "齐勇", "房婷", "费磊", "纪健", "滕峰", "段霞", "景婷", "章明", "巴勇", "侯霞", "宁飞", "柴健", "贝峰"}
		//b := [53]string{"", "lihua01@uw.edu", "wangq02@uw.edu", "zhangw03@uw.edu", "liuy04@uw.edu", "chenj05@uw.edu", "yangx06@uw.edu", "huangl07@uw.edu", "zhaot08@uw.edu", "sunt09@uw.edu", "zhouj10@uw.edu", "wuj11@uw.edu", "zhengh12@uw.edu", "zhuf13@uw.edu", "maj14@uw.edu", "hup15@uw.edu", "ling16@uw.edu", "hey17@uw.edu", "gaor18@uw.edu", "luom19@uw.edu", "xiex20@uw.edu", "songf21@uw.edu", "hanp22@uw.edu", "zhongy23@uw.edu", "xuj24@uw.edu", "xup25@uw.edu", "chengj26@uw.edu", "wangx27@uw.edu", "ship28@uw.edu", "daih29@uw.edu", "weij30@uw.edu", "liuy31@uw.edu", "fanf32@uw.edu", "pengp33@uw.edu", "baox34@uw.edu", "ana35@uw.edu", "daig36@uw.edu", "ley37@uw.edu", "sup38@uw.edu", "nongl39@uw.edu", "qiy40@uw.edu", "fangt41@uw.edu", "feil42@uw.edu", "jij43@uw.edu", "tengf44@uw.edu", "duanx45@uw.edu", "jingt46@uw.edu", "zhangm47@uw.edu", "bay48@uw.edu", "houx49@uw.edu", "ningf50@uw.edu", "chaij51@uw.edu", "beif52@uw.edu"}
		//fmt.Printf("%s, %s, %s, %s, %s, %s, %s, %s, %s\n", flightInfo[0], a[i], b[i], flightInfo[3], flightInfo[4], flightInfo[5], flightInfo[6], flightInfo[7], flightInfo[8])
		personInfo := person{
			form_id:    flightInfo[0],
			name:       flightInfo[1],
			email:      flightInfo[2],
			arr_date:   flightInfo[6],
			arr_time:   flightInfo[7],
			flight_num: flightInfo[8],
		}
	}
}
