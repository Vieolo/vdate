package vdate

import "time"

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
