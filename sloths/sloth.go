package sloths

import (
	"errors"
	"math/rand"
	"time"
)

// FlashSlothmore manages the maximum milliseconds it takes for Flash
// Slothmore, Department of Mammal Vehicles employee, to process a transaction
// and the number of milliseconds it takes for him to fall asleep
// FYI: https://zootopia.fandom.com/wiki/Flash_Slothmore
type FlashSlothmore struct {
	MaxMillisecondsToProcessTransaction int
	MillisecondsToFallAsleep int
}

// CanFail determines if Flash Slothmore is able to process a transaction
// before falling asleep
func (f *FlashSlothmore) CanFail() (error, string){
	// Use a seed to ensure rand.Intn provides a newly seeded int on each run
	seed := rand.NewSource(time.Now().UnixNano())
	randWithSeed := rand.New(seed)
	transactionDelayMilliseconds := randWithSeed.Intn(f.MaxMillisecondsToProcessTransaction)

	time.Sleep(time.Duration(transactionDelayMilliseconds) * time.Millisecond)

	if transactionDelayMilliseconds > f.MillisecondsToFallAsleep {
		return errors.New("Flash Slothmore fell asleep before he could" +
			" process the transaction :("), ""
	}
	return nil, "Success! Flash Slothmore processed the transaction before" +
		" falling asleep!"
}