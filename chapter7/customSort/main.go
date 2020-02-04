package main

import "sort"

import "fmt"

type Student struct {
	name  string
	id    int
	score int
}

type SortableStudents struct {
	students []*Student
	less     func(i, j int) bool
}

func (s SortableStudents) Len() int {
	return len(s.students)
}

func (s SortableStudents) Less(i, j int) bool {
	return s.less(i, j)
}

func (s SortableStudents) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i]
}

func printStudents(students []*Student) {
	for i, v := range students {
		fmt.Printf("%v. %v\n", i, *v)
	}
	fmt.Println()
}

func main() {
	students := []*Student{
		{"Tom", 1001, 90},
		{"Jack", 1002, 93},
		{"Jerry", 1003, 87},
		{"Alice", 1004, 96},
		{"Bob", 1005, 89},
	}
	sort.Sort(SortableStudents{students, func(i, j int) bool {
		return students[i].id < students[j].id
	}})
	printStudents(students)

	sort.Sort(SortableStudents{students, func(i, j int) bool {
		return students[i].name < students[j].name
	}})
	printStudents(students)

	sort.Sort(SortableStudents{students, func(i, j int) bool {
		return students[i].score > students[j].score
	}})
	printStudents(students)
}
