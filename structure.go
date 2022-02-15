// EXAMPLE OF A STRUCTURE IN GO LENGUAGE BY DANI95RICO

package main

import "fmt"

type contactInfo struct {
	email string
	num   int
}
type persona struct {
	nombre   string
	apellido string
	//contact  contactInfo
	contactInfo
}

func main() {
	dani := persona{
		nombre:   "Dani",
		apellido: "Rico",
		contactInfo: contactInfo{
			email: "dani@gmail.com",
			num:   689494998,
		},
	}

	dani.print()
	dani.updateName("Daniel")
	dani.print()
}

func (p persona) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *persona) updateName(newNombre string) {
	(*pointerToPerson).nombre = newNombre
}
