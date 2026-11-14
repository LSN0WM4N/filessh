package format

import "time"

// UnixTime converts a Unix timestamp (seconds since the epoch) to a
// date and time string in the local time zone using the format
// "DD/MM/YYYY HH:mm".
//
// Example: 1700000000 -> "14/11/2023 23:13" (depending on the time zone).
func UnixTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("02/01/2006 15:04")
}
