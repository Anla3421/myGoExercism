package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	output, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}
	return output
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	output, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}

	if time.Now().Unix() > output.Unix() {
		return true
	}
	return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	output, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}
	if output.Hour() >= 12 && output.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	output, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}
	// `You have an appointment on Thursday, July 25, 2019, at 13:45.`
	text := `You have an appointment on %s, %s %v, %v, at %v:%v.`
	msg := fmt.Sprintf(text, output.Weekday(), output.Month(), output.Day(), output.Year(), output.Hour(), output.Minute())
	return msg
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	output := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	return output
}
