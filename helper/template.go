package helper

import "time"

func Url(path string) string{
	return path
}

func DateFormat(date time.Time, layout string) string {
	return date.Format(layout)
}