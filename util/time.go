package util

import (
	"time"
)

type bucket int

var (
	Today     bucket = 0
	Yesterday bucket = 1
	ThisMonth bucket = 2
	LastMonth bucket = 3
	ThisYear  bucket = 4
	LastYear  bucket = 5
)

func (b bucket) Start() (t time.Time) {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()

	switch b {
	case Today:
		t = time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	case Yesterday:
		t = time.Date(year, month, day-1, 0, 0, 0, 0, time.Local)
	case ThisMonth:
		t = time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	case LastMonth:
		t = time.Date(year, month-1, 1, 0, 0, 0, 0, time.Local)
	case ThisYear:
		t = time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	case LastYear:
		t = time.Date(year-1, 1, 1, 0, 0, 0, 0, time.Local)
	default:
		t = time.Time{}
	}
	return
}

func (b bucket) End() (t time.Time) {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()

	switch b {
	case Today:
		t = time.Date(year, month, day+1, 0, 0, 0, 0, time.Local)
	case Yesterday:
		t = time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	case ThisMonth:
		t = time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local)
	case LastMonth:
		t = time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	case ThisYear:
		t = time.Date(year+1, 1, 1, 0, 0, 0, 0, time.Local)
	case LastYear:
		t = time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	default:
		t = time.Time{}
	}
	return
}
