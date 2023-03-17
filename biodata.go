package main

import (
	"fmt"
	"os"
	"strconv"
)

type Friend struct {
	name    string
	address string
	work    string
	reason  string
}

var Friends map[int]Friend = map[int]Friend{
	1: {
		name:    "Agung Syahputra",
		address: "Jl Percetakan Negara II 6, Dki Jakarta",
		work:    "Android Developer",
		reason:  "Ingin belajar tentang bahasa pemrograman golang",
	},
	2: {
		name:    "Chandra Prasetyo",
		address: "Jl Karangmenjangan 22, Jawa Timur",
		work:    "Backend Developer",
		reason:  "Ingin memperdalam bidan backend developer",
	},
	3: {
		name:    "Dimas Eka Putra",
		address: "Jl Tebah Psr Mayestik Bl B/96, Dki Jakarta",
		work:    "Frontend Developer",
		reason:  "Ingin belajar backend menggunakan golang",
	},
	4: {
		name:    "Novida Ria Andini",
		address: "Jl Trunojoyo 3, Dki Jakarta",
		work:    "grapik desainer",
		reason:  "Ingin mempelajari bahasa pemrograman golang",
	},
	5: {
		name:    "Nurul Afifah",
		address: "Jl FM Noto 13, Jawa Tengah",
		work:    "UI/UX Desainer",
		reason:  "Ingin mempelajari bahasa pemrograman golang",
	},
}

func main() {
	args := os.Args
	if len(args) == 2 {

		id, err := strconv.Atoi(args[1])

		if err == nil {
			if id < len(Friends)-1 {
				friendById := Friends[id]
				fmt.Println("Nama :", friendById.GetName())
				fmt.Println("Alamant :", friendById.GetAddress())
				fmt.Println("Pekerjaan :", friendById.GetWork())
				fmt.Println("Alasan memilih kelas Golang :", friendById.GetReason())
			} else {
				fmt.Printf("Absen no.%d tidak ada", id)
			}
		}

	}

}

func (friend *Friend) GetName() string {
	return friend.name
}
func (friend *Friend) GetAddress() string {
	return friend.address
}
func (friend *Friend) GetWork() string {
	return friend.work
}
func (friend *Friend) GetReason() string {
	return friend.reason
}
