package main

import (
	"fmt"
	"time"
)

func main() {
	timerChan := make(chan time.Time)
  
	go func(){
	  time.Sleep(5 * time.Second)
	  timerChan <- time.Now()
	}()
  
	fmt.Println("Time from synchronous call: ", time.Now())
	fmt.Println("Time from channel: ", <-timerChan)
  }