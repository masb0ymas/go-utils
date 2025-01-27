package pkg

import (
	"log"
	"time"
)

var countryTz = map[string]string{
	"ID": "Asia/Jakarta",
	"MY": "Asia/Kuala_Lumpur",
	"SG": "Asia/Singapore",
}

func TimeIn(name string) time.Time {
	loc, err := time.LoadLocation(countryTz[name])

	if err != nil {
		log.Fatal(err)
	}

	return time.Now().In(loc)
}

func IsValidDate(date *string) bool {
	return date != nil && *date != "" && *date != "0001-01-01T00:00:00Z"
}
