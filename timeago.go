// Package time provides a set of functions to return how much
// time has been passed between two dates.
package timeago

import (
	"time"
  "fmt"
  "math"
	"errors"
)

type DateAgoValues int

const (
    SecondsAgo DateAgoValues = iota
    MinutesAgo
    HoursAgo
    DaysAgo
    WeeksAgo
    MonthsAgo
    YearsAgo
)

// TimeAgoFromNowWithTime takes a specific end Time value
// and the current Time to return how much has been passed
// between them.
func TimeAgoFromNowWithTime(end time.Time) (string, error) {

	return TimeAgoWithTime(time.Now(), end)
}

// TimeAgoFromNowWithTime takes a specific layout as time
// format to parse the time string on end paramter to return
// how much time has been passed between the current time and
// the string representation of the time provided by user.
func TimeAgoFromNowWithString(layout, end string) (string, error) {

	t, e := time.Parse(layout, end)
	if e == nil {
		return TimeAgoWithTime(time.Now(), t)
	} else {
		err := errors.New("Invalid format")
		return "", err
	}
}

// TimeAgoWithTime takes a specific start/end Time values
// and calculate how much time has been passed between them.
func TimeAgoWithTime(start, end time.Time) (string, error) {
	duration := start.Sub(end)
	return stringForDuration(duration), nil
}

// TimeAgoWithString takes a specific layout as time
// format to parse the time string on start/end parameter to return
// how much time has been passed between them.
func TimeAgoWithString(layout, start, end string) (string, error) {

	timeStart, e := time.Parse(layout, start)
	if e != nil {
		err := errors.New("Invalid start time format")
		return "", err
	}

	timeEnd, e := time.Parse(layout, end)
	if e != nil {
		err := errors.New("Invalid end time format")
		return "", err
	}

	duration := timeStart.Sub(timeEnd)
	return stringForDuration(duration), nil
}

func stringForDuration(duration time.Duration) string {
	if duration.Hours() < 24 {
		if duration.Hours() >= 1 {
			return localizedStringFor(HoursAgo, int(round(duration.Hours())));
		} else if duration.Minutes() >= 1 {
			return localizedStringFor(MinutesAgo, int(round(duration.Minutes())));
		} else {
			return localizedStringFor(SecondsAgo, int(round(duration.Seconds())));
		}
	} else {
		if duration.Hours() >= 8760 {
			years := duration.Hours() / 8760
			return localizedStringFor(YearsAgo, int(years));
		} else if duration.Hours() >= 730 {
			months := duration.Hours() / 730
			return localizedStringFor(MonthsAgo, int(months));
		} else if duration.Hours() >= 168 {
			weeks := duration.Hours() / 168
			return localizedStringFor(WeeksAgo, int(weeks));
		} else {
			days := duration.Hours() / 24
			return localizedStringFor(DaysAgo, int(days));
		}
	}
}

func round(f float64) float64 {
    return math.Floor(f + .5)
}

func localizedStringFor(valueType DateAgoValues, value int) string {

    switch valueType {
        case YearsAgo:
            if value >= 2 {
                return fmt.Sprintf("%d tahun lalu", value);
            } else {
                return "Tahun lalu";
            }
        case MonthsAgo:
            if value >= 2 {
                return fmt.Sprintf("%d bulan lalu", value);
            } else {
                return "Bulan lalu";
            }
        case WeeksAgo:
            if value >= 2 {
                return fmt.Sprintf("%d minggu lalu", value);
            } else {
                return "Minggu lalu";
            }
        case DaysAgo:
            if value >= 2 {
                return fmt.Sprintf("%d hari lalu", value);
            } else {
                return "Kemarin";
            }
        case HoursAgo:
            if value >= 2 {
                return fmt.Sprintf("%d jam lalu", value);
            } else {
                return "1 jam lalu";
            }
        case MinutesAgo:
            if value >= 2 {
                return fmt.Sprintf("%d menit lalu", value);
            } else {
                return "1 menit lalu";
            }
        case SecondsAgo:
            if value >= 2 {
                return fmt.Sprintf("%d detik lalu", value);
            } else {
                return "Baru saja";
            }
    }
    return "";
}
