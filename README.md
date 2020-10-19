# Bradfield Takehome - Exercise 4 - Exponential backoff

## About
This project implements an exponential backoff.

### Backstory 
[Flash Slothmore](https://zootopia.fandom.com/wiki/Flash_Slothmore) is an 
employee of the Department of Mammal Vehicles in Zootopia. He's the fastest
sloth on the payroll but he's, well, still a sloth. The DMV receives frequent
complaints that Flash works so slowly that he falls asleep before processing
some transactions. 

### Back to the Exercise
`sloths.FlashSlothmore` implements the `exponentialbackoff.Failable` interface
qualifying as an operation that has "a chance of failure".

The main function calls `ExponentialBackoff` with arbitarily selected arguments of 
the following types:
* Failable interface
* slice of ints representing backoff durations in milliseconds 
* StoppingCriteria
    * int representing the maximum number of retries allowed
    * int representing the maximum number of retries completed (see **Future Work**)
* int representing range of jitter values in milliseconds 

## System setup 
This project is written in Golang, version 1.15. 

To prepare your system to run this project: 

1. Verify you have Golang installed: `go --version`
2. If not, install: `brew install golang`
3. Verify your GOPATH environment variable is set: `echo $GOPATH`
4. If not set: `export GOPATH=<insert-path-to-directory-where-you-will-save-go-projects>`  
5. This project should be saved in the GOPATH

## Running the program
 
1. build the project `go build`
2. run main `go run main.go`

## Future work / shortcomings    
* Implement unit tests
* Fix code organization, e.g. RetriesCompleted should not be public  
