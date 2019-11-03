// Package phoneBook provides ...
package phoneBook

import "fmt"

type Contact struct {
	Name         string
	Patronymic   string
	Surname      string
	Organization string
	Phones       []int
}

type Contacts []Contact

func (c Contacts) Print() {
	for i, v := range c {
		fmt.Printf("%3d:\t%s\t%s\t%s\t%s\t", i, v.Name, v.Surname, v.Patronymic, v.Organization)
		fmt.Printf("[")
		for i, p := range v.Phones {
			fmt.Printf("%d", p)
			if i != len(v.Phones)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("]")
		fmt.Printf("\n")
	}
}

type ByName Contacts

// Len is the number of elements in the collection.
func (b ByName) Len() int {
	return len(b)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (b ByName) Less(i int, j int) bool {
	return b[i].Name < b[j].Name
}

// Swap swaps the elements with indexes i and j.
func (b ByName) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

type ByOrganization []Contact

// Len is the number of elements in the collection.
func (b ByOrganization) Len() int {
	return len(b)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (b ByOrganization) Less(i int, j int) bool {
	return b[i].Organization < b[j].Organization
}

// Swap swaps the elements with indexes i and j.
func (b ByOrganization) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

type BySurname []Contact

// Len is the number of elements in the collection.
func (b BySurname) Len() int {
	return len(b)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (b BySurname) Less(i int, j int) bool {
	return b[i].Surname < b[j].Surname
}

// Swap swaps the elements with indexes i and j.
func (b BySurname) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

type ByOrganizationAndSurname []Contact

// Len is the number of elements in the collection.
func (b ByOrganizationAndSurname) Len() int {
	return len(b)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (b ByOrganizationAndSurname) Less(i int, j int) bool {
	if b[i].Organization == b[j].Organization {
		return BySurname(b).Less(i, j)
	}
	return ByOrganization(b).Less(i, j)
}

// Swap swaps the elements with indexes i and j.
func (b ByOrganizationAndSurname) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}
