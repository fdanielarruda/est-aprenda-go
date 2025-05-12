/*
Partindo do código abaixo, ordene os []user por idade e sobrenome.
https://play.golang.org/p/BVRZTdlUZ_
Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.
*/

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type orderByAge []user

func (u orderByAge) Len() int           { return len(u) }
func (u orderByAge) Less(i, j int) bool { return u[i].Age < u[j].Age }
func (u orderByAge) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

type orderByLast []user

func (u orderByLast) Len() int           { return len(u) }
func (u orderByLast) Less(i, j int) bool { return u[i].Last < u[j].Last }
func (u orderByLast) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	fmt.Println("")

	fmt.Println("Ordenado por idade")
	sort.Sort(orderByAge(users))
	showUsers(users)

	fmt.Println("")

	fmt.Println("Ordenado por sobrenome")
	sort.Sort(orderByLast(users))
	showUsers(users)
}

func showUsers(users []user) {
	for _, u := range users {
		fmt.Printf("%v %v (%v)\n", u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)

		for _, saying := range u.Sayings {
			fmt.Printf("\t%v\n", saying)
		}
	}
}
