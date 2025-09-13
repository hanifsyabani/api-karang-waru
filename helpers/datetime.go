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