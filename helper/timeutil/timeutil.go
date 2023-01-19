package timeutil

import (
	"log"
	"time"
)

func BeginOfDay(t time.Time) time.Time {
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return t
}

func BeginOfNextDay(t time.Time) time.Time {
	log.Println(t)
	t = BeginOfDay(t).Add(24 * time.Hour)
	log.Println(t)
	return t
}

func SameDay(t1, t2 time.Time) bool {
	d1, m1, y1 := t1.Date()
	d2, m2, y2 := t2.Date()
	return d1 == d2 && m1 == m2 && y1 == y2
}

func SameWeek(t1, t2 time.Time) bool {
	w1, y1 := t1.ISOWeek()
	w2, y2 := t2.ISOWeek()
	return w1 == w2 && y1 == y2
}

func SameMonth(t1, t2 time.Time) bool {
	_, m1, y1 := t1.Date()
	_, m2, y2 := t2.Date()
	return m1 == m2 && y1 == y2
}
