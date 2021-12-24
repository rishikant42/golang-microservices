package main

import (
	"fmt"
	"sync"
)

type AtomicInt struct {
	value int
	lock  sync.Mutex
}

func (i *AtomicInt) Increase() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value++
}

func (i *AtomicInt) Decrease() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value--
}

func (i *AtomicInt) Value() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.value
}

var (
	counter     = 0
	atomicValue = AtomicInt{}
	lock        sync.Mutex
)

func main() {
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		wg2.Add(1)
		go updateCounter(&wg1)
		go updateAtomicValue(&wg2)
	}
	wg1.Wait()
	wg2.Wait()
	fmt.Println("counter:", counter)
	fmt.Println("Atomic value:", atomicValue.Value())
}

func updateAtomicValue(wg *sync.WaitGroup) {
	atomicValue.Increase()
	wg.Done()
}

func updateCounter(wg *sync.WaitGroup) {
	lock.Lock()
	defer lock.Unlock()
	counter++
	wg.Done()
}
