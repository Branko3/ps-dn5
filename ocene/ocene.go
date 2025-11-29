package ocene

import "fmt"

type Student struct {
	Ime     string
	Priimek string
	ocene   []int
}

func DodajOceno(student *Student, ocena int) {
	student.ocene = append(student.ocene, ocena)
}

func IzpisVsehOcen(student Student) {
	for _, ocena := range student.ocene {
		fmt.Println(ocena)
	}
}

func IzpisiKoncniUspeh(student Student) {
	if len(student.ocene) == 0 {
		fmt.Println("Ni ocen.")
		return
	}
	fmt.Printf("Koncni uspeh: %.2f\n", povprecje(student.ocene))
}

func povprecje(ocene []int) float64 {
	suma := 0.0
	for _, ocena := range ocene {
		suma += float64(ocena)
	}
	return suma / float64(len(ocene))
}
