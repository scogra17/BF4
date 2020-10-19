package main

import (
	"exponential_backoff/exponentialbackoff"
	"exponential_backoff/someobject"
)


func main() {
	object := someobject.ICanFail{
		SomeString: "Bonjour",
		SomeInt:    0,
	}

	stoppingCriteria := exponentialbackoff.StoppingCriteria{
		int(0),
		int(10),
	}

	exponentialbackoff.ExponentialBackoff(&object, []int{1000, 2000}, &stoppingCriteria, 1000)
}