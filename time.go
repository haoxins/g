package g

import (
	"fmt"
	"time"
)

// TimeConverter converts between time.Time and string.
type TimeConverter struct{}

// FromString converts string to time.Time.
func (tc *TimeConverter) FromString(value string, layout string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse time: %s", value))
	}
	return t
}

// ToString converts time.Time to string.
func (tc *TimeConverter) ToString(t time.Time, layout string) string {
	return t.Format(layout)
}

// GetTime gets the time.Time from the time string.
func GetTime(timeStr string, layout string) time.Time {
	var tc TimeConverter
	return tc.FromString(timeStr, layout)
}

// GetCurrentTime gets the current time.
func GetCurrentTime(layout string) string {
	var tc = &TimeConverter{}
	var now = time.Now()
	return tc.ToString(now, layout)
}

// SetCurrentTime sets the current time to the target.
func SetCurrentTime(target *string, layout string) {
	var tc = &TimeConverter{}
	var now = time.Now()
	*target = tc.ToString(now, layout)
}
