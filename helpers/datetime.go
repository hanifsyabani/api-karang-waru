package helpers

import (
	"api-karang-waru/config"
	"log"
	"time"
)


var (
	appTimezone *time.Location
)

func init() {
	timezoneStr := config.GetEnv("APP_TIMEZONE", "Asia/Jakarta")
	var err error
	appTimezone,err = time.LoadLocation(timezoneStr)
	if err != nil {
		log.Fatalf("Failed to load timezone: %v", err)
	}
}

func FormatTimeHuman(t time.Time) string {
	return  t.In(appTimezone).Format("2006-01-02 15:04:05")
}

func ParseTimeHuman(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil
	}
	t, err := time.ParseInLocation("2006-01-02", dateStr, appTimezone)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}