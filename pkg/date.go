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
