package main

import "fmt"

type node struct {
  coreSpeed int // Hz
  numCores  int
  linkSpeed int // bits per second
  numLinks  int
}

type message struct {
  timeSent int
  timeRecv int
}

type task struct {
  cycles int // clock cycles to execute
  inputs []message
  outputs []message
}

func execute(results chan<- int, traces []task, hw node) {
  currentTime := 0

  for index, _ := range traces {
    currentTime += traces[index].cycles
  }

  results <- currentTime
}

func main() {
  fmt.Println("Hello, World!")

  results := make(chan int)

  node1 := node{}
  var trace1 []task

  node2 := node{}
  trace2 := []task{task{cycles: 2}}

  go execute(results, trace1, node1)

  go execute(results, trace2, node2)

  time1 := <-results
  time2 := <-results

  totalTime := max(time1, time2)

  fmt.Println("Total execution time: ", totalTime)
}
