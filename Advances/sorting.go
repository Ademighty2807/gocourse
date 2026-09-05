package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type personSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

func (s *personSorter) Len() int {
	return len(s.people)
}

func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (by By) Sort(people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

func main() {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Anna", 35},
	}
	fmt.Println("Unsorted by age:", people)
	// Do assending sort by age and name
	ByAge := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	ByName := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}

	ageDesc := func(p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}

	nameDesc := func(p1, p2 *Person) bool {
		return p1.Name > p2.Name
	}

	lenByName := func(p1, p2 *Person) bool {
		return len(p1.Name) < len(p2.Name)
	}

	By(ByAge).Sort(people)
	fmt.Println("After by age sorting:", people)
	By(ByName).Sort(people)
	fmt.Println("After by name sorting:", people)
	By(ageDesc).Sort(people)
	fmt.Println("After by age descending sorting:", people)
	By(nameDesc).Sort(people)
	fmt.Println("After by name descending sorting:", people)
	By(lenByName).Sort(people)
	fmt.Println("After by name length sorting:", people)

	// ========== SORT.SLICE ==========
	stringSlice := []string{"banana", "apple", "cherry", "date"}
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})
	fmt.Println("After sort.Slice:", stringSlice)
}

// ==============================

// type ByAge []Person
// type ByName []Person

// func (a ByAge) Len() int {
// 	return len(a)
// }
// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }
// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }

// func (a ByName) Len() int {
// 	return len(a)
// }
// func (a ByName) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }
// func (a ByName) Less(i, j int) bool {
// 	return a[i].Name < a[j].Name
// }
