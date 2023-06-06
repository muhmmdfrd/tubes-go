package main

import (
	"fmt"
	"os"
	"os/exec"
)

const NMAX int = 2000

type student struct {
	id        int
	nim, name string
}

type course struct {
	id             int
	name           string
}

type studentScore struct {
	id, studentId, courseId, sks int
	uts, uas, quiz float64
}

type students [NMAX]student
type courses [NMAX]course
type studentCourses [NMAX]studentScore

var answer int
var studentsData students
var coursesData courses
var studentCoursesData studentCourses
var nStudent, nCourses, nStudentCourses int

func main() {
	header()
	fmt.Scan(&answer)

	if answer == 1 {
		clear()
		menuMahasiswa(&studentsData, &nStudent)
	} else if answer == 2 {
		menuMatkul(&coursesData, &nCourses)
	} else if answer == 3 {
		showScore(&studentCoursesData, &nStudentCourses)
	}
}

func header() {
	fmt.Println("\n-----------------------------------------")
	fmt.Print("\tAplikasi IGracias Console")
	fmt.Println("\n-----------------------------------------")
	fmt.Println("By: Muhammad Farid")
	fmt.Println("By: Jihan Alifah Maritza")
	fmt.Println("-----------------------------------------")
	fmt.Println("1. Data Mahasiswa")
	fmt.Println("2. Data Mata Kuliah")
	fmt.Println("3. Data Nilai Mahasiswa")
	fmt.Println("4. Transkrip Nilai")
	fmt.Println("-----------------------------------------")
	fmt.Print("Pilih menu: ")
}

func menuMahasiswa(students *students, n *int) {
	fmt.Println("-----------------------------------------")
	fmt.Println("1. Tampilkan Data Mahasiswa")
	fmt.Println("2. Tambah Data Mahasiswa")
	fmt.Println("3. Edit Data Mahasiswa")
	fmt.Println("4. Hapus Data Mahasiswa")
	fmt.Println("9. Kembali ke Menu Utama")
	fmt.Println("-----------------------------------------")

	var answer int
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&answer)

	if answer == 1 {
		showMahasiswa(students, n)
		menuMahasiswa(students, n)
	} else if answer == 2 {
		inputMahasiswa(students, n)
		menuMahasiswa(students, n)
	} else if answer == 3 {
		var id int
		fmt.Print("Pilih id untuk mengedit data mahasiswa: "); fmt.Scan(&id)
		editMahasiswa(id, students, n)
		showMahasiswa(students, n)
		menuMahasiswa(students, n)
	} else if answer == 4 {
		var id int
		fmt.Print("Pilih id untuk menghapus data mahasiswa: "); fmt.Scan(&id)
		deleteMahasiswa(id, students, n)
		showMahasiswa(students, n)
		menuMahasiswa(students, n)
	} else if answer == 9 {
		clear()
		main()
	}
}

func showMahasiswa(students *students, n *int) {
	if *n == 0 {
		fmt.Println("Data kosong.")
		return
	}

	for i := 0; i < *n; i++ {
		data := students[i]
		fmt.Println(data.id, data.nim, data.name)
	}
}

func inputMahasiswa(students *students, n *int) {
	var active bool = true
	var i int = 0

	if *n > 0 {
		i = *n
	}

	for active {
		var s student
		fmt.Print("ID: "); fmt.Scan(&s.id)
		fmt.Print("NIM: "); fmt.Scan(&s.nim)
		fmt.Print("NAMA: "); fmt.Scan(&s.name)
		students[i] = s
		i++
		*n = i

		fmt.Print("Apakah Anda ingin menambah lagi data mahasiswa? (true/false): ")
		fmt.Scan(&active)
	}
}

func editMahasiswa(id int, students *students, n *int) {
	var idx int = searchMahasiswaById(id, *students, *n)

	if idx == -1 {
		fmt.Printf("Data mahasiswa dengan id %d tidak ditemukan.\n", id)
		return
	}

	var s student
	fmt.Printf("NIM (%s): ", students[idx].nim); fmt.Scan(&s.nim)
	fmt.Printf("NAMA (%s): ", students[idx].name); fmt.Scan(&s.name)

	if s.nim != "" {
		students[idx].nim = s.nim
	}

	if s.name != "" {
		students[idx].name = s.name
	}
}

func deleteMahasiswa(id int, students *students, n *int) {
	var idx, i int
	idx = searchMahasiswaById(id, *students, *n)

	if idx == -1 {
		fmt.Printf("Data mahasiswa dengan id %d tidak ditemukan.\n", id)
		return
	}

	i = idx
	for i <= *n - 2 {
		students[i] = students[i+1]
		i++
	}
	*n--
}

func searchMahasiswaById(id int, students students, n int) int {
	var start int = 0
	var end int = n-1

	for start <= end {
		var mid int = start + (end-start) / 2
		if students[mid].id == id {
			return mid
		} else if students[mid].id < id {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
  return -1
}

func menuMatkul(courses *courses, n *int) {
	fmt.Println("-----------------------------------------")
	fmt.Println("1. Tampilkan Data Mata Kuliah")
	fmt.Println("2. Tambah Data Mata Kuliah")
	fmt.Println("3. Edit Data Mata Kuliah")
	fmt.Println("4. Hapus Data Mata Kuliah")
	fmt.Println("9. Kembali ke Menu Utama")
	fmt.Println("-----------------------------------------")

	var answer int
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&answer)

	if answer == 1 {
		showMatkul(courses, n)
		menuMatkul(courses, n)
	} else if answer == 2 {
		inputMatkul(courses, n)
		menuMatkul(courses, n)
	} else if answer == 9 {
		clear()
		main()
	}
}

func showMatkul(courses *courses, n *int) {
	if *n == 0 {
		clear()
		fmt.Println("Data kosong.")
		return
	}

	for i := 0; i < *n; i++ {
		data := courses[i]
		fmt.Println(data.id, data.name)
	}
}

func inputMatkul(courses *courses, n *int) {
	var active bool = true
	var i int = 0

	if *n > 0 {
		i = *n
	}

	for active {
		var c course
		fmt.Print("ID: "); fmt.Scan(&c.id)
		fmt.Print("NAMA: "); fmt.Scan(&c.name)
		courses[i] = c
		i++
		*n = i

		fmt.Print("Apakah Anda ingin menambah lagi data mahasiswa? (true/false): ")
		fmt.Scan(&active)
	}
}

func showScore(studentCourses *studentCourses, n *int) {
	for i := 0; i < *n; i++ {
		data := studentCourses[i]
		fmt.Println(data.id, data.studentId, data.courseId, data.sks, data.quiz, data.uts, data.uas)
	}
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
