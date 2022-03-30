package pkg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GetReadableDateStringFromDate(date time.Time) string {
	return fmt.Sprintf("%v %v, %v", date.Day(), date.Month(), date.Year())
}

func GetLocalDateFromYearMonthDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func GetUtcDateFromYearMonthDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// "2022-12-31 -> time.Time"
func GetUtcDateFromYearMonthDayFormattedString(yyyyMmDddd string) (time.Time, error) {
	splitted := strings.Split(yyyyMmDddd, "-") // should return ['2022', '12', 31]
	if len(splitted) != 3 {
		return time.Now(), errors.New("invalid date string cannot be parsed into a valid date")
	}

	yyyy, err := strconv.ParseInt(splitted[0], 10, 32)
	mm, err := strconv.ParseInt(splitted[1], 10, 32)
	dd, err := strconv.ParseInt(splitted[2], 10, 32)

	if mm < 1 || mm > 12 || dd < 1 || dd > 31 || yyyy < 2000 || yyyy > 2999 {
		err = errors.New("invalid date found")
	}

	if err != nil {
		return time.Now(), errors.New("invalid date string cannot be parsed into a valid date")
	}

	var month time.Month
	switch mm {
	case 1:
		month = time.January
	case 2:
		month = time.February
	case 3:
		month = time.March
	case 4:
		month = time.April
	case 5:
		month = time.May
	case 6:
		month = time.June
	case 7:
		month = time.July
	case 8:
		month = time.August
	case 9:
		month = time.September
	case 10:
		month = time.October
	case 11:
		month = time.November
	case 12:
		month = time.December
	}

	date := GetUtcDateFromYearMonthDay(int(yyyy), month, int(dd))
	return date, nil
}
