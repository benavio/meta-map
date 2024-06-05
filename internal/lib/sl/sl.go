package sl

import (
	"log/slog"
	"time"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func InfLatitude(Latitude float64) slog.Attr {
	return slog.Attr{
		Key:   "Latitude",
		Value: slog.Float64Value(Latitude),
	}
}

func InfLongitude(Longitude float64) slog.Attr {
	return slog.Attr{
		Key:   "Longitude",
		Value: slog.Float64Value(Longitude),
	}
}

func InfDate(Date time.Time) slog.Attr {
	return slog.Attr{
		Key:   "Date",
		Value: slog.TimeValue(Date),
	}
}
