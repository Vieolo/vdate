package vdate

import "time"

// Formats the given `Time` instance
// `df` is the predefined format for the date declared as constants in this module. e.g. vdate.DF_YYYYMMDD
func FormatDate(t time.Time, df string) string {
	format := ""
	switch df {
	case DF_YYYYMMDD:
		format = "2006-01-02"
	case DF_DDMMYYYY:
		format = "02/01/2006"
	case DF_MMDDYYYY:
		format = "01/02/2006"
	}

	return t.Format(format)
}

// Checks whether the two given times are on the same day or not
func IsOnSameDay(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
