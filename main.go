package main

import (
	"exponential_backoff/exponentialbackoff"
	"exponential_backoff/sloths"
	"fmt"
)


func main() {
	flashSlothmore := sloths.FlashSlothmore{
		MaxMillisecondsToProcessTransaction: 4000,
		MillisecondsToFallAsleep:            3000,
	}

	stoppingCriteria := exponentialbackoff.StoppingCriteria{
		MaxRetriesAllowed: 10,
		RetriesCompleted:  0,
	}

	durationsForBackoffsMilliseconds := []int{100, 200, 400, 800}
	jitterInMilliseconds := 100

	err, resp := exponentialbackoff.ExponentialBackoff(
		&flashSlothmore,
		durationsForBackoffsMilliseconds,
		&stoppingCriteria,
		jitterInMilliseconds)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}

}