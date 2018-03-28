package date

import (
	"strings"
	"time"
)

const (
	yyyy = "2006"
	yy   = "06"
	MMMM = "January"
	MMM  = "Jan"
	MM   = "01"
	dddd = "Monday"
	ddd  = "Mon"
	dd   = "02"

	HHT = "03"
	HH  = "15"
	mm  = "04"
	SS  = "05"
	ss  = "05"
	tt  = "PM"
	Z   = "MST"
	ZZZ = "MST"
)

func DateFormat(tlong int64, duration time.Duration, format string) string {
	var (
		sec, mills int64
	)

	r := int64(time.Second / duration)

	if r == 0 {
		sec = tlong
		mills = 0
	} else {
		sec = tlong / r
		mills = (tlong - sec*r) * int64(duration)
	}

	layout := convertFormat(format)
	t := time.Unix(sec, mills)
	return t.Format(layout)
}

func DateFormat2(t time.Time, format string) string {
	return t.Format(convertFormat(format))
}

func ParseDate(dateTimeStr string, format string) time.Time {
	t, _ := time.Parse(convertFormat(format), dateTimeStr)
	return t
}

func convertFormat(format string) string {
	var goFormate = format
	if strings.Contains(goFormate, "YYYY") {
		goFormate = strings.Replace(goFormate, "YYYY", yyyy, -1)
	} else if strings.Contains(goFormate, "yyyy") {
		goFormate = strings.Replace(goFormate, "yyyy", yyyy, -1)
	} else if strings.Contains(goFormate, "YY") {
		goFormate = strings.Replace(goFormate, "YY", yy, -1)
	} else if strings.Contains(goFormate, "yy") {
		goFormate = strings.Replace(goFormate, "yy", yy, -1)
	}

	//month
	if strings.Contains(goFormate, "MMMM") {
		goFormate = strings.Replace(goFormate, "MMMM", MMMM, -1)
	} else if strings.Contains(goFormate, "MMM") {
		goFormate = strings.Replace(goFormate, "MMM", MMM, -1)
	} else if strings.Contains(goFormate, "MM") {
		goFormate = strings.Replace(goFormate, "MM", MM, -1)
	}

	if strings.Contains(goFormate, "mm") { //minute
		goFormate = strings.Replace(goFormate, "mm", mm, -1)
	}

	//day
	if strings.Contains(goFormate, "dddd") {
		goFormate = strings.Replace(goFormate, "dddd", dddd, -1)
	} else if strings.Contains(goFormate, "ddd") {
		goFormate = strings.Replace(goFormate, "ddd", ddd, -1)
	} else if strings.Contains(goFormate, "dd") {
		goFormate = strings.Replace(goFormate, "dd", dd, -1)
	}

	if strings.Contains(goFormate, "tt") {
		if strings.Contains(goFormate, "HH") {
			goFormate = strings.Replace(goFormate, "HH", HHT, -1)
		} else if strings.Contains(goFormate, "hh") {
			goFormate = strings.Replace(goFormate, "hh", HHT, -1)
		}
		goFormate = strings.Replace(goFormate, "tt", tt, -1)
	} else {
		if strings.Contains(goFormate, "HH") {
			goFormate = strings.Replace(goFormate, "HH", HH, -1)
		} else if strings.Contains(goFormate, "hh") {
			goFormate = strings.Replace(goFormate, "hh", HH, -1)
		}
		goFormate = strings.Replace(goFormate, "tt", "", -1)
	}

	//second
	if strings.Contains(goFormate, "SS") {
		goFormate = strings.Replace(goFormate, "SS", SS, -1)
	} else if strings.Contains(goFormate, "ss") {
		goFormate = strings.Replace(goFormate, "ss", SS, -1)
	}

	if strings.Contains(goFormate, "ZZZ") {
		goFormate = strings.Replace(goFormate, "ZZZ", ZZZ, -1)
	} else if strings.Contains(goFormate, "zzz") {
		goFormate = strings.Replace(goFormate, "zzz", ZZZ, -1)
	} else if strings.Contains(goFormate, "Z") {
		goFormate = strings.Replace(goFormate, "Z", Z, -1)
	} else if strings.Contains(goFormate, "z") {
		goFormate = strings.Replace(goFormate, "z", Z, -1)
	}

	if strings.Contains(goFormate, "tt") {
		goFormate = strings.Replace(goFormate, "tt", tt, -1)
	}
	return goFormate
}
