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
	uts, uas, quiz float64
}

type studentScore struct {
	id, studentId, courseId, sks int
}

type students [NMAX]student
type courses [NMAX]course
type studentCourses [NMAX]studentScore

func main() {
	var answer int
	var students students
	var courses courses
	var studentCourses studentCourses
	var nStudent, nCourses, nStudentCourses int

	header()
	fmt.Scan(&answer)

	if answer == 1 {
		menuMahasiswa(&students, &nStudent)
	} else if answer == 2 {
		// Bikin menuMatkul(&courses, &nCourses)
		// Terus isinya samain aja kayak yang mahasiswa
		matkul(&courses, &nCourses)
	} else if answer == 3 {
		showScore(&studentCourses, &nStudentCourses)
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
	fmt.Print("Pilih menu: "); 
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
	fmt.Print("Pilih Menu: "); fmt.Scan(&answer)

	if answer == 1 {
		showMahasiswa(students, n)
	} else if answer == 2 {
		inputMahasiswa(students, n)
	} else if answer == 9 {
		clear()
		main()
	}
}

func showMahasiswa(students *students, n *int) {
	if *n == 0 {
		fmt.Println("Data kosong.")
		menuMahasiswa(students, n)
		return
	}

	for i := 0; i < *n; i++ {
		data := students[i]
		fmt.Println(data.id, data.nim, data.name)	
	}

	menuMahasiswa(students, n)
}

func inputMahasiswa(students *students, n *int) {
	var active bool = true
	var i int = 0

	if *n > 0 {
		i = *n
	}

	for active {
		var s student
		fmt.Print("NIM: "); fmt.Scan(&s.nim)
		fmt.Print("NAMA: "); fmt.Scan(&s.name)
		s.id = i+1
		students[i] = s
		i++
		*n = i

		fmt.Print("Apakah Anda ingin menambah lagi data mahasiswa? (true/false): "); fmt.Scan(&active)
	}
	menuMahasiswa(students, n)
}

func matkul(courses *courses, n *int) {
	for i := 0; i < *n; i++ {
		data := courses[i]
		fmt.Println(data.id, data.name, data.quiz, data.uts, data.uas)
	}
}

func showScore(studentCourses *studentCourses, n *int) {
	for i := 0; i < *n; i++ {
		data := studentCourses[i]
		fmt.Println(data.id, data.studentId, data.courseId, data.sks)
	}
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
  cmd.Stdout = os.Stdout
  cmd.Run()
}

