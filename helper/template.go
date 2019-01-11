package helper

import (
	"time"
	"html/template"
)

func Url(path string) string{
	return path
}

func DateFormat(date time.Time, layout string) string {
	return date.Format(layout)
}

func TmpHtml(h string)  template.HTML{
	return template.HTML(h)
}