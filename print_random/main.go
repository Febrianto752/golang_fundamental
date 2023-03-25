package main

import (
	"fmt"
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
	for i := 1; i <= 4; i++ {
		go printDataCoba(i)
	}
	for i := 1; i <= 4; i++ {
		go printDataBisa(i)
	}

	time.Sleep(1 * time.Second)
}

func (dataCoba DataCoba) GetCoba() [3]string {
	return [3]string{"coba1", "coba2", "coba3"}
}

func (dataBisa DataBisa) GetBisa() [3]string {
	return [3]string{"bisa1", "bisa2", "bisa3"}
}

func printDataCoba(i int) {
	var Data DataCoba
	fmt.Println(Data.GetCoba(), i)
}
func printDataBisa(i int) {
	var Data DataBisa
	fmt.Println(Data.GetBisa(), i)
}
