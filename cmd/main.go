package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"redovalnica/ocene"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "CLI aplikacija za upravljanje redovalnice",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Value: 1,
				Usage: "najmanjše število ocen potrebnih za pozitivno oceno",
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Value: 0,
				Usage: "najmanjša možna ocena",
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Value: 10,
				Usage: "največja možna ocena",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			student1 := ocene.Student{Ime: "Ana", Priimek: "Novak"}
			student2 := ocene.Student{Ime: "Boris", Priimek: "Kralj"}
			student3 := ocene.Student{Ime: "Janez", Priimek: "Novak"}

			ocene.DodajOceno(&student1, 10)
			ocene.DodajOceno(&student1, 9)
			ocene.DodajOceno(&student1, 8)

			ocene.DodajOceno(&student2, 6)
			ocene.DodajOceno(&student2, 7)
			ocene.DodajOceno(&student2, 5)
			ocene.DodajOceno(&student2, 8)

			ocene.DodajOceno(&student3, 4)
			ocene.DodajOceno(&student3, 5)
			ocene.DodajOceno(&student3, 3)
			ocene.DodajOceno(&student3, 5)

			minOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")

			fmt.Printf("Nastavitve: minOcen=%d, minOcena=%d, maxOcena=%d\n\n", minOcen, minOcena, maxOcena)

			fmt.Println("Student 1: ", student1.Ime, student1.Priimek)
			ocene.IzpisVsehOcen(student1)
			ocene.IzpisiKoncniUspeh(student1)
			fmt.Println()

			fmt.Println("Student 2: ", student2.Ime, student2.Priimek)
			ocene.IzpisVsehOcen(student2)
			ocene.IzpisiKoncniUspeh(student2)
			fmt.Println()

			fmt.Println("Student 3: ", student3.Ime, student3.Priimek)
			ocene.IzpisVsehOcen(student3)
			ocene.IzpisiKoncniUspeh(student3)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
