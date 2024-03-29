package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error()string{
	return fmt.Sprintf("At %v, %s",e.When,e.What)
}

func run() error {
	return &MyError{
		time.Now(),"it Didn't work",
	}
}
func main(){
	if err := run();err!=nil{
		fmt.Println(err)
	}
}
