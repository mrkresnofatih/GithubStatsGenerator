package executors

import (
	"log"
	"time"
)

func CleanUp() {
	time.Sleep(2 * time.Second)
	log.Println("Hooray! All Done!")
}
