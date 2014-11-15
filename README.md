container
==

data structures in golang.

## Queue

Like the other data structures in the standard library container package, the base queue, implementation empty interface for the type. Below is an example of how to setup a queue with a custom type.

```go
package main

import (
  "fmt"
  "github.com/acmacalister/container/queue"
  )

type myStruct struct {
  name string
}

type myStructQueue struct {
  q queue.Queue
}

func (sq *myStructQueue) Push(s *myStruct) {
  sq.q.Push(s)
}

func (sq *myStructQueue) Pop() (myStruct, error) {
  val, err := sq.q.Pop()
  if err != nil {
    return "", err
  }
  return val.(myStruct), nil
}

func (sq *myStructQueue) Size() int {
  return sq.q.Size()
}

func (sq *myStructQueue) Empty() bool {
  return sq.q.Empty()
}

func main() {
  q := myStructQueue{}
  q.Push(myStruct{name: "Live"})
  q.Push(myStruct{name: "Die"})
  q.Push(myStruct{name: "Repeat"})
  val, err := q.Pop()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(val)
}
```

# Stack

```go
Cool stack example here.
```