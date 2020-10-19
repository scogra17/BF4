package exponentialbackoff

import (
	"fmt"
	"math/rand"
	"time"
)

type Failable interface {
	CanFail() (error, string)
}

type StoppingCriteria struct {
	RetriesCompleted    int // has zero value of 0
	MaxRetriesCompleted int
}

func ExponentialBackoff(object Failable, durations []int, stoppingCriteria *StoppingCriteria, jitterInMilliseconds int) (error, string) {
	// Try the operation once
	err, resp := object.CanFail()

	if len(durations) != 0 && (err != nil && stoppingCriteria.RetriesCompleted < stoppingCriteria.MaxRetriesCompleted) {
		fmt.Println("**Call to object.DoSomething failed**")
		fmt.Println("stoppingCriteria.RetriesCompleted: ", stoppingCriteria.RetriesCompleted)
		fmt.Println("durations: ", durations)
		// Retry operations
		sleepDurationInt := durations[0]
		randomJitter := rand.Intn(jitterInMilliseconds) - jitterInMilliseconds / 2
		fmt.Println("randomJitter: ", randomJitter)
		sleepDurationJitterIncluded := sleepDurationInt + randomJitter
		fmt.Println("sleepDurationJitterIncluded: ", sleepDurationJitterIncluded)
		sleepDurationJitterIncludedDuration := time.Duration(sleepDurationJitterIncluded)
		time.Sleep(sleepDurationJitterIncludedDuration*time.Millisecond)

		stoppingCriteria.RetriesCompleted++

		ExponentialBackoff(object, durations[1:], stoppingCriteria, jitterInMilliseconds)
	}

	fmt.Print("*****RETURN!*******")
	fmt.Print("err", err)
	fmt.Print("resp", resp)
	return err, resp
}

