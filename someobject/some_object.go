package someobject

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type ICanFail struct {
	SomeString string
	SomeInt int
}

func (s *ICanFail) CanFail() (error, string){
	start := time.Now()
	randomDuration := time.Duration(rand.Intn(4))*time.Second
	fmt.Println("Trying to DoSomething with randomDuration: ", randomDuration)
	time.Sleep(randomDuration)
	if time.Since(start) > 2 {
		return errors.New("timed out :("), ""
	}
	return nil, "Hello!"
}