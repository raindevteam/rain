package main

import (
	"time"
	"sync"
	"fmt"
	//"runtime"
)

type Item struct {
	Id int
}

type Test struct {
	Name string
	Items map[int]*Item
	sync.Mutex
}

func (t *Test) MakeItem(id int) {
	//fmt.Println("here")
	t.Lock()

	time.Sleep(time.Second * 3)
	fmt.Println("then here")
	t.Items[id] = &Item{id}
	t.Unlock()
}

func (t *Test) IncItemId(id int) {
	t.Lock()
	fmt.Println("and here")
	t.Items[id].Id = 7
	t.Unlock()
}

func main() {
	test := &Test{"Test", make(map[int]*Item), sync.Mutex{}}
    funcs := []func(int){test.MakeItem, test.IncItemId}
    for _, fun := range funcs {
        go fun(1)
    }
	time.Sleep(time.Second * 4)
}
