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
	id   int
	name string
}

type studentScore struct {
	id, studentId, courseId, sks int
	uts, uas, quiz               float64
}

type students [NMAX]student
type courses [NMAX]course
type studentScores [NMAX]studentScore

var answer int
var studentsData students
var coursesData courses
var studentScoresData studentScores
var nStudent, nCourses, nstudentScores int

func main() {
	header()
	fmt.Scan(&answer)

	if answer == 1 {
		clear()
		menuMahasiswa(&studentsData, &nStudent)
	} else if answer == 2 {
		menuMatkul(&coursesData, &nCourses)
	} else if answer == 3 {
		menuNilaiMahasiswa(&studentScoresData, &nstudentScores)
	} else if answer == 4 {
		// TODO: Transkrip nilai
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
		fmt.Print("Pilih id untuk mengedit data mahasiswa: ")
		fmt.Scan(&id)
		editMahasiswa(id, students, n)
		showMahasiswa(students, n)
		menuMahasiswa(students, n)
	} else if answer == 4 {
		var id int
		fmt.Print("Pilih id untuk menghapus data mahasiswa: ")
		fmt.Scan(&id)
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
		fmt.Print("ID: ")
		fmt.Scan(&s.id)
		fmt.Print("NIM: ")
		fmt.Scan(&s.nim)
		fmt.Print("NAMA: ")
		fmt.Scan(&s.name)
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
	fmt.Printf("NIM (%s): ", students[idx].nim)
	fmt.Scan(&s.nim)
	fmt.Printf("NAMA (%s): ", students[idx].name)
	fmt.Scan(&s.name)

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
	for i <= *n-2 {
		students[i] = students[i+1]
		i++
	}
	*n--
}

func searchMahasiswaById(id int, students students, n int) int {
	var start int = 0
	var end int = n - 1

	for start <= end {
		var mid int = start + (end-start)/2
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
	} else if answer == 3 {
		var id int
		fmt.Print("Pilih id untuk mengedit data matakuliah: ")
		fmt.Scan(&id)
		editMatkul(id, courses, n)
		showMatkul(courses, n)
		menuMatkul(courses, n)
	} else if answer == 4 {
		var id int
		fmt.Print("Pilih id untuk menghapus data matakuliah: ")
		fmt.Scan(&id)
		deleteMatkul(id, courses, n)
		showMatkul(courses, n)
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
		fmt.Print("ID: ")
		fmt.Scan(&c.id)
		fmt.Print("NAMA: ")
		fmt.Scan(&c.name)
		courses[i] = c
		i++
		*n = i

		fmt.Print("Apakah Anda ingin menambah lagi data mahasiswa? (true/false): ")
		fmt.Scan(&active)
	}
}

func editMatkul(id int, courses *courses, n *int) {
	var idx int = searchMatkulById(id, *courses, *n)

	if idx == -1 {
		fmt.Printf("Data matakuliah dengan id %d tidak ditemukan.\n", id)
		return
	}

	var c course
	fmt.Printf("NAMA (%s): ", courses[idx].name)
	fmt.Scan(&c.name)

	if c.name != "" {
		courses[idx].name = c.name
	}
}

func deleteMatkul(id int, courses *courses, n *int) {
	var idx, i int
	idx = searchMatkulById(id, *courses, *n)

	if idx == -1 {
		fmt.Printf("Data matakuliah dengan id %d tidak ditemukan.\n", id)
		return
	}

	i = idx
	for i <= *n-2 {
		courses[i] = courses[i+1]
		i++
	}
	*n--
}

func searchMatkulById(id int, courses courses, n int) int {
	var start int = 0
	var end int = n - 1

	for start <= end {
		var mid int = start + (end-start)/2
		if courses[mid].id == id {
			return mid
		} else if courses[mid].id < id {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func menuNilaiMahasiswa(studentScores *studentScores, n *int) {
	fmt.Println("-----------------------------------------")
	fmt.Println("1. Tampilkan Data Nilai Mahasiswa")
	fmt.Println("2. Tambah Data Nilai Mahasiswa")
	fmt.Println("3. Edit Data Nilai Mahasiswa")
	fmt.Println("4. Hapus Data Nilai Mahasiswa")
	fmt.Println("9. Kembali ke Menu Utama")
	fmt.Println("-----------------------------------------")

	var answer int
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&answer)

	if answer == 1 {
		showNilaiMahasiswa(studentScores, n)
		menuNilaiMahasiswa(studentScores, n)
	} else if answer == 2 {
		// inputNilaiMahasiswa(studentScores, n)
		menuNilaiMahasiswa(studentScores, n)
	} else if answer == 3 {
		var id int
		fmt.Print("Pilih id untuk mengedit data mahasiswa: ")
		fmt.Scan(&id)
		// editMahasiswa(id, studentScores, n)
		showNilaiMahasiswa(studentScores, n)
		menuNilaiMahasiswa(studentScores, n)
	} else if answer == 4 {
		var id int
		fmt.Print("Pilih id untuk menghapus data mahasiswa: ")
		fmt.Scan(&id)
		// deleteMahasiswa(id, studentScores, n)
		showNilaiMahasiswa(studentScores, n)
		menuNilaiMahasiswa(studentScores, n)
	} else if answer == 9 {
		clear()
		main()
	}
}

func showNilaiMahasiswa(studentScores *studentScores, n *int) {
	if *n == 0 {
		fmt.Println("Data kosong.")
		return
	}

	for i := 0; i < *n; i++ {
		data := studentScores[i]
		fmt.Println(data.id, data.studentId, data.courseId, data.sks, data.quiz, data.uts, data.uas)
	}
}

func inputNilaiMahasiswa(studentScores *studentScores, students students, courses courses, nStudent, nCourse int, n *int) {
	var active bool = true
	var i int = 0

	if *n > 0 {
		i = *n
	}

	for active {
		var studentIdx int = -1
		var searchStudent bool = true

		for searchStudent {
			var idx int
			fmt.Print("Masukkan ID Mahasiswa: ")
			fmt.Scan(&idx)
			studentIdx = searchMahasiswaById(idx, students, nStudent)

			if studentIdx == -1 {
				fmt.Printf("Data mahasiswa dengan ID %d tidak ditemukan.\n", idx)
				fmt.Print("Apakah Anda ingin memasukkan kembali ID mahasiswa? (true/false): ")
				fmt.Scan(&searchStudent)
				return
			}
		}

		if studentIdx != -1 {
			fmt.Printf("Mahasiswa yang dipilih dengan nama %s\n", students[studentIdx].name)
		}

		var courseIdx int = -1
		var searchCourse bool = true

		for searchCourse {
			var idx int
			fmt.Print("Masukkan ID Mata Kuliah: ")
			fmt.Scan(&idx)
			// TODO: pakai func ini kalo dari Jihan udah selesai
			// courseIdx = searchMahasiswaById(idx, students, nStudent)

			if courseIdx == -1 {
				fmt.Printf("Data mata kuliah dengan ID %d tidak ditemukan.\n", idx)
				fmt.Print("Apakah Anda ingin memasukkan kembali ID mata kuliah? (true/false): ")
				fmt.Scan(&searchCourse)
				return
			}
		}

		if courseIdx != -1 {
			fmt.Printf("Nama mata kuliah yang dipilih adalah %s\n", courses[courseIdx].name)
		}

		var ss studentScore
		fmt.Print("ID: ")
		fmt.Scan(&ss.id)
		fmt.Print("SKS: ")
		fmt.Scan(&ss.sks)
		fmt.Print("QUIZ: ")
		fmt.Scan(&ss.quiz)
		fmt.Print("UTS: ")
		fmt.Scan(&ss.uts)
		fmt.Print("UAS: ")
		fmt.Scan(&ss.uas)

		ss.courseId = courses[courseIdx].id
		ss.studentId = students[studentIdx].id

		studentScores[i] = ss
		i++
		*n = i

		fmt.Print("Apakah Anda ingin menambah lagi data nilai mahasiswa? (true/false): ")
		fmt.Scan(&active)
	}
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
