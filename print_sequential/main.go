package main

import (
	"fmt"
	"sync"
	"time"
)

type DataCoba struct {
}

type DataBisa struct {
}

// interface 1
type Coba interface {
	GetCoba() [3]string
}

// interface 2
type Bisa interface {
	GetBisa() [3]string
}

func main() {
	var mutex sync.Mutex
	var dataBisa DataBisa
	var dataCoba DataCoba

	for i := 1; i <= 4; i++ {

		go func(i int) {
			mutex.Lock()
			dataCoba.GetCoba(i)
			dataBisa.GetBisa(i)
			mutex.Unlock()
		}(i)

	}

	time.Sleep(1 * time.Second)
}

func (dataCoba DataCoba) GetCoba(i int) {
	fmt.Println([3]string{"coba1", "coba2", "coba3"}, i)
}

func (dataBisa DataBisa) GetBisa(i int) {
	fmt.Println([3]string{"bisa1", "bisa2", "bisa3"}, i)
}
