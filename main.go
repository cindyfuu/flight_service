package main

import (
	_ "fmt"
	_ "time"
	_ "unsafe" // use for Sizeof
)

type person struct {
	husky_id   string
	Form_id    int
	name       string
	Email      string
	Arr_date   string
	Arr_time   string
	flight_num string
}

type ride struct {
	date   string
	time   string
	people []person
}
