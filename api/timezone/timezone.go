package timezone

import "time"

func ConvertToAsiaSeoul(datetime *time.Time) *time.Time {
	converted := datetime.Add(-9 * time.Hour)
	return &converted
}
