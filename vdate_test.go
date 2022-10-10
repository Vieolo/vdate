package vdate_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vieolo/vdate"
)

func TestFormatDate(t *testing.T) {
	assert.Equal(t, "2022-02-14", vdate.FormatDate(time.Date(2022, 2, 14, 0, 0, 0, 0, time.UTC), vdate.DF_YYYYMMDD))
	assert.Equal(t, "14/02/2022", vdate.FormatDate(time.Date(2022, 2, 14, 0, 0, 0, 0, time.UTC), vdate.DF_DDMMYYYY))
	assert.Equal(t, "02/14/2022", vdate.FormatDate(time.Date(2022, 2, 14, 0, 0, 0, 0, time.UTC), vdate.DF_MMDDYYYY))
}
