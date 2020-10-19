package exponentialbackoff

import (
	"fmt"
	"math/rand"
	"time"
)

type Failable interface {
	CanFail() (error, string)
}

// StoppingCriteria manages the retry data for exponential backoff
type StoppingCriteria struct {
	RetriesCompleted  int // zero value of 0
	MaxRetriesAllowed int
}

func durationWithJitter(duration, jitter int) int {
	// Jitter should be a maximum of twice duration
	if jitter > duration*2 {
		jitter = duration*2
	}
	randomJitter := rand.Intn(jitter) - jitter / 2
	durationWithJitter := duration + randomJitter

	return durationWithJitter
}

func ExponentialBackoff(object Failable, durationsInMilliseconds []int, stoppingCriteria *StoppingCriteria, jitterInMilliseconds int) (error, string) {
	// First attempt
	err, resp := object.CanFail()

	// Retry if there are still unused durations remaining and stopping criteria
	// are not met
	if len(durationsInMilliseconds) != 0 && err != nil &&
		stoppingCriteria.RetriesCompleted < stoppingCriteria.MaxRetriesAllowed {
		stoppingCriteria.RetriesCompleted++

		backoffMilliseconds :=
			durationWithJitter(durationsInMilliseconds[0], jitterInMilliseconds)

		fmt.Printf("Will attempt retry #%d following backoff of %d ms\n",
			stoppingCriteria.RetriesCompleted, backoffMilliseconds)

		time.Sleep(time.Duration(backoffMilliseconds)*time.Millisecond)

		return ExponentialBackoff(
			object,
			durationsInMilliseconds[1:],
			stoppingCriteria,
			jitterInMilliseconds,
		)
	}
	return err, resp
}

