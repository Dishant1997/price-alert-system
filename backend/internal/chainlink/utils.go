package chainlink

import "log"

func LogError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}