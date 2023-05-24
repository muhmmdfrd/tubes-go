package main

import "fmt"

const NMAX int = 1000

type student struct {
	id int
	nim, name string
	courseId int
}

type course struct {
	id int
	name string
	uts, uas, quiz float64
}

type students [NMAX]student
type courses [NMAX]course

func main() {
	var answer int
	var students students
	var courses courses
	var nStudent, nCourses int

	header(&answer)

	switch (answer) {
	case 1:
			mahasiswa(&students, &nStudent)
		break;
	case 2:
			matkul(&courses, &nCourses)
		break;
	case 3:
		break;
	case 4:
		break;
	default:
		break
	}
}

func header(answer *int) {
	fmt.Println("\n_________________________________________")
	fmt.Print("\tAplikasi IGracias Console")
	fmt.Println("\n-----------------------------------------")
	fmt.Println("1. Data Mahasiswa")
	fmt.Println("2. Data Mata Kuliah")
	fmt.Println("3. Data Nilai Mahasiswa")
	fmt.Println("4. Transkrip Nilai")

	fmt.Scan(&answer)
}


func mahasiswa(students *students, n *int) {
	for i := 0; i < *n; i++ {
		fmt.Println(students[i].name)
	}
}

func matkul(courses *courses, n *int) {
	for i := 0; i < *n; i++ {
		fmt.Println(courses[i].name)
	}
}