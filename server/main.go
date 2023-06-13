package main

import (
	_ "fmt"
	_ "time"
	_ "unsafe" // use for Sizeof
)

type person struct {
	husky_id   string
	form_id    int
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

func main() {}
