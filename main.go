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
	grade string
}

type studentGroup struct {
	name string
	uts, uas, quiz, totalScore, totalSks float64
}

type students [NMAX]student
type courses [NMAX]course
type studentScores [NMAX]studentScore
type studentSummary [NMAX]studentGroup

var studentsData students
var coursesData courses
var studentScoresData studentScores
var studentSummaryData studentSummary
var nStudent, nCourses, nstudentScores, nStudentSummary int

func main() {
	var answer int
	header()
	fmt.Scan(&answer)

	if answer == 1 {
		clear()
		menuMahasiswa(&studentsData, &nStudent)
	} else if answer == 2 {
		clear()
		menuMatkul(&coursesData, &nCourses)
	} else if answer == 3 {
		clear()
		menuNilaiMahasiswa(&studentScoresData, studentsData, coursesData, studentSummaryData, nStudent, nCourses, nStudentSummary, &nstudentScores)
	} else if answer == 4 {
		clear()
		transcript(studentsData, coursesData, studentScoresData, nStudent, nstudentScores)
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
	fmt.Println("\n-----------------------------------------")
	fmt.Print("\tAplikasi IGracias Console")
	fmt.Println("\n-----------------------------------------")
	fmt.Println("By: Muhammad Farid")
	fmt.Println("By: Jihan Alifah Maritza")
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
		clear()
		showMahasiswa(*students, *n)
		menuMahasiswa(students, n)
	} else if answer == 2 {
		clear()
		inputMahasiswa(students, n)
		menuMahasiswa(students, n)
	} else if answer == 3 {
		clear()
		var id int
		showMahasiswa(*students, *n)
		fmt.Print("Pilih id untuk mengedit data mahasiswa: ")
		fmt.Scan(&id)
		editMahasiswa(id, students, n)
		showMahasiswa(*students, *n)
		menuMahasiswa(students, n)
	} else if answer == 4 {
		clear()
		showMahasiswa(*students, *n)
		var id int
		fmt.Print("Pilih id untuk menghapus data mahasiswa: ")
		fmt.Scan(&id)
		deleteMahasiswa(id, students, n, studentScoresData, nstudentScores)
		showMahasiswa(*students, *n)
		menuMahasiswa(students, n)
	} else if answer == 9 {
		clear()
		main()
	}
}

func showMahasiswa(students students, n int) {
	if n == 0 {
		fmt.Println("Data kosong.")
		return
	}

	for i := 0; i < n; i++ {
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

func deleteMahasiswa(id int, students *students, n *int, studentScores studentScores, nStudentScore int) {
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

func searchMahasiswaByNim(nim string, students students, n int) int {
	var low int = 0
	var high int = n - 1

	for low <= high {
		mid := (low + high) / 2

		if students[mid].nim == nim {
			return mid
		} else if students[mid].nim < nim {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func menuMatkul(courses *courses, n *int) {
	fmt.Println("\n-----------------------------------------")
	fmt.Print("\tAplikasi IGracias Console")
	fmt.Println("\n-----------------------------------------")
	fmt.Println("By: Muhammad Farid")
	fmt.Println("By: Jihan Alifah Maritza")
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
		clear()
		showMatkul(*courses, *n)
		menuMatkul(courses, n)
	} else if answer == 2 {
		clear()
		inputMatkul(courses, n)
		menuMatkul(courses, n)
	} else if answer == 3 {
		clear()
		showMatkul(*courses, *n)
		var id int
		fmt.Print("Pilih id untuk mengedit data matakuliah: ")
		fmt.Scan(&id)
		editMatkul(id, courses, n)
		showMatkul(*courses, *n)
		menuMatkul(courses, n)
	} else if answer == 4 {
		clear()
		showMatkul(*courses, *n)
		var id int
		fmt.Print("Pilih id untuk menghapus data matakuliah: ")
		fmt.Scan(&id)
		deleteMatkul(id, courses, n, studentScoresData, nstudentScores)
		showMatkul(*courses, *n)
		menuMatkul(courses, n)
	} else if answer == 9 {
		clear()
		main()
	}
}

func showMatkul(courses courses, n int) {
	if n == 0 {
		clear()
		fmt.Println("Data kosong.")
		return
	}

	for i := 0; i < n; i++ {
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

		fmt.Print("Apakah Anda ingin menambah lagi data matakuliah? (true/false): ")
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

func deleteMatkul(id int, courses *courses, n *int, studentScores studentScores, nStudentCourse int) {
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

func searchMatkulByStudentId(studentId int, studentstudentScores studentScores, n int) ([NMAX]int, int) {
	var ids [NMAX]int
	var counter int = 0

	for i := 0; i < n; i++ {
		if studentstudentScores[i].studentId == studentId && studentstudentScores[i].studentId != 0 {
			ids[counter] = studentstudentScores[i].courseId
			counter++
		}
	}

	return ids, counter
}

func menuNilaiMahasiswa(studentScores *studentScores, 
	students students, courses courses, studentSummary studentSummary,
	nStudent, nCourse, nStudentSummary int, 
	n *int) {

	fmt.Println("\n-----------------------------------------")
	fmt.Print("\tAplikasi IGracias Console")
	fmt.Println("\n-----------------------------------------")
	fmt.Println("By: Muhammad Farid")
	fmt.Println("By: Jihan Alifah Maritza")
	fmt.Println("-----------------------------------------")
	fmt.Println("1. Tampilkan Data Nilai Mahasiswa")
	fmt.Println("2. Tambah Data Nilai Mahasiswa")
	fmt.Println("3. Edit Data Nilai Mahasiswa")
	fmt.Println("4. Hapus Data Nilai Mahasiswa")
	fmt.Println("5. Tampilkan Data Mahasiswa Berdasarkan Mata Kuliah")
	fmt.Println("6. Tampilkan Data Mahasiswa secara Berurut")
	fmt.Println("9. Kembali ke Menu Utama")
	fmt.Println("-----------------------------------------")

	var answer int
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&answer)

	if answer == 1 {
		clear()
		showNilaiMahasiswa(*studentScores, students, courses, nStudent, nCourse, *n)
		menuNilaiMahasiswa(studentScores, students, courses, studentSummary, nStudent, nCourse, nStudentSummary, n)
	} else if answer == 2 {
		clear()
		inputNilaiMahasiswa(studentScores, students, courses, nStudent, nCourse, n)
		menuNilaiMahasiswa(studentScores, students, courses, studentSummary, nStudent, nCourse, nStudentSummary, n)
	} else if answer == 3 {
		clear()
		var id int
		fmt.Print("Pilih id untuk mengedit data mahasiswa: ")
		fmt.Scan(&id)
		editNilaiMahasiswa(id, studentScores, n)
		showNilaiMahasiswa(*studentScores, students, courses, nStudent, nCourse, *n)
		menuNilaiMahasiswa(studentScores, students, courses, studentSummary, nStudent, nCourse, nStudentSummary, n)
	} else if answer == 4 {
		clear()
		var id int
		fmt.Print("Pilih id untuk menghapus data mahasiswa: ")
		fmt.Scan(&id)
		deleteNilaiMahasiswa(id, studentScores, n)
		showNilaiMahasiswa(*studentScores, students, courses, nStudent, nCourse, *n)
		menuNilaiMahasiswa(studentScores, students, courses, studentSummary, nStudent, nCourse, nStudentSummary, n)
	} else if answer == 5 {
		clear()
		var id int
		fmt.Print("Pilih id mata kuliah: ")
		fmt.Scan(&id)
		searchStudentByMatkulId(id, *studentScores, students, courses, nCourse, nStudent, *n)
		menuNilaiMahasiswa(studentScores, students, courses, studentSummary, nStudent, nCourse, nStudentSummary, n)
	} else if answer == 6 {
		clear()
		groupStudents(&studentSummary, &nStudentSummary, *studentScores, *n, students, nStudent)
		showSortedNilaiMahasiswa(studentSummary, students, nStudentSummary)
		menuNilaiMahasiswa(studentScores, students, courses, studentSummary, nStudent, nCourse, nStudentSummary, n)
	} else if answer == 9 {
		clear()
		main()
	}
}

func groupStudents(
	studentSummary *studentSummary, nStudentSummary *int, 
	studentScores studentScores, nStudentScore int, 
	students students, nStudent int) {

	var mapping map[string]studentGroup = make(map[string]studentGroup)
	var counter int = 0

	for i := 0; i < nStudentScore; i++ {
		var d studentScore = studentScores[i]
		var studentName string = students[searchMahasiswaById(d.studentId, students, nStudent)].name

		if studentName != "" {
			if v, ok := mapping[studentName]; ok {
				v.quiz += d.quiz
				v.uts += d.uts
				v.uas += d.uas
				v.totalSks += float64(d.sks)
				v.totalScore += (d.quiz + d.uts + d.uas)
				mapping[studentName] = v
			} else if d.studentId != 0 {
				var total float64 = d.quiz + d.uts + d.uas
				mapping[studentName] = studentGroup{name: studentName, quiz: d.quiz, uts: d.uts, uas: d.uas, totalScore: total, totalSks: float64(d.sks) }
				counter++
			}
		}
	}

	for i := 0; i < counter; i++ {
		var name string = students[i].name
		studentSummary[i] = mapping[name]
	}
	
	*nStudentSummary = counter
}

func showSortedNilaiMahasiswa(studentSummary studentSummary, students students, n int) {
	var sort, field string

	fmt.Print("Urutkan berdasarkan (total/quiz/uts/uas/sks): ")
	fmt.Scan(&field)

	fmt.Print("Urutkan berdasarkan: (asc/desc): ")
	fmt.Scan(&sort)

	var isValid bool = sortNilaiMahasiswa(&studentSummary, n, sort, field)

	for i := 0; i < n && isValid; i++ {
		var data studentGroup = studentSummary[i]
		fmt.Println(data.name, data.totalSks, data.quiz, data.uts, data.uas, data.totalScore)
	}
}

func searchStudentByMatkulId(matkulId int, studentScores studentScores, studentsData students, courses courses, nCourse, nStudent, nStudentCourse int) {
	var idx = searchMatkulById(matkulId, courses, nCourse)
	var result students
	var counter int = 0

	if idx == -1 {
		fmt.Println("Mata kuliah dengan id", matkulId, "tidak ditemukan.")
		return
	}

	for i := 0; i < nStudentCourse; i++ {
		if studentScores[i].courseId == courses[idx].id {
			var idxStudent int = searchMahasiswaById(studentScores[i].studentId, studentsData, nStudent)

			if idxStudent != -1 {
				result[counter] = studentsData[idxStudent]
				counter++
			}
		}
	}

	fmt.Println("Mata Kuliah:", courses[idx].name)
	showMahasiswa(result, counter)
}

func editNilaiMahasiswa(id int, studentScores *studentScores, n *int) {
	var idx int = searchNilaiMahasiswaById(id, *studentScores, *n)

	if idx == -1 {
		fmt.Printf("Data nilai mahasiswa dengan id %d tidak ditemukan.\n", id)
		return
	}

	var s studentScore
	fmt.Printf("SKS (%d): ", studentScores[idx].sks)
	fmt.Scan(&s.sks)
	fmt.Printf("QUIZ (%f): ", studentScores[idx].quiz)
	fmt.Scan(&s.quiz)
	fmt.Printf("UTS (%f): ", studentScores[idx].uts)
	fmt.Scan(&s.uts)
	fmt.Printf("UAS (%f): ", studentScores[idx].uas)
	fmt.Scan(&s.uas)

	if s.sks != 0 {
		studentScores[idx].sks = s.sks
	}

	if s.quiz != 0 {
		studentScores[idx].quiz = s.quiz
	}

	if s.uts != 0 {
		studentScores[idx].uts = s.uts
	}

	if s.uas != 0 {
		studentScores[idx].uas = s.uas
	}

	var ss studentScore = studentScores[idx]
	studentScores[idx].grade = calculateGrade(ss.quiz, ss.uts, ss.uas)
}

func deleteNilaiMahasiswa(id int, studentsScoresData *studentScores, n *int) {
	var idx, i int
	idx = searchNilaiMahasiswaById(id, *studentsScoresData, *n)

	if idx == -1 {
		fmt.Printf("Data nilai mahasiswa dengan id %d tidak ditemukan.\n", id)
		return
	}

	i = idx
	for i <= *n-2 {
		studentsScoresData[i] = studentsScoresData[i+1]
		i++
	}
	*n--
}

func showNilaiMahasiswa(studentScores studentScores, students students, courses courses, nStudent, nCourse int, n int) {
	if n == 0 {
		fmt.Println("Data kosong.")
		return
	}

	for i := 0; i < n; i++ {
		var data studentScore = studentScores[i]
		var idxStudent int = searchMahasiswaById(data.studentId, students, nStudent)
		var idxCourse int = searchMatkulById(data.courseId, courses, nCourse)
		var student student = students[idxStudent]
		var course course = courses[idxCourse]

		fmt.Println(data.id, student.nim, student.name, course.name, data.sks, data.quiz, data.uts, data.uas, data.grade)
	}
}

func searchNilaiMahasiswaById(id int, studentScores studentScores, n int) int {
	var start int = 0
	var end int = n - 1

	for start <= end {
		var mid int = start + (end-start)/2
		if studentScores[mid].id == id {
			return mid
		} else if studentScores[mid].id < id {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
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
			showMahasiswa(students, nStudent)
			var idx int
			fmt.Print("Masukkan ID Mahasiswa: ")
			fmt.Scan(&idx)
			studentIdx = searchMahasiswaById(idx, students, nStudent)

			if studentIdx == -1 {
				fmt.Printf("Data mahasiswa dengan ID %d tidak ditemukan.\n", idx)
				fmt.Print("Apakah Anda ingin memasukkan kembali ID mahasiswa? (true/false): ")
				fmt.Scan(&searchStudent)
			} else {
				searchStudent = false
				active = false
			}
		}

		if studentIdx != -1 {
			fmt.Printf("Mahasiswa yang dipilih dengan nama %s\n", students[studentIdx].name)
		}

		var courseIdx int = -1
		var searchCourse bool = true

		for searchCourse {
			showMatkul(courses, nCourse)
			var idx int
			fmt.Print("Masukkan ID Mata Kuliah: ")
			fmt.Scan(&idx)

			courseIdx = searchMatkulById(idx, courses, nCourse)

			if courseIdx == -1 {
				fmt.Printf("Data mata kuliah dengan ID %d tidak ditemukan.\n", idx)
				fmt.Print("Apakah Anda ingin memasukkan kembali ID mata kuliah? (true/false): ")
				fmt.Scan(&searchCourse)
			} else {
				searchCourse = false
				active = false
			}
		}

		if courseIdx != -1 {
			fmt.Printf("Nama mata kuliah yang dipilih adalah %s\n", courses[courseIdx].name)
		}

		if courseIdx != -1 && studentIdx != -1 {
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

			ss.grade = calculateGrade(ss.quiz, ss.uts, ss.uas)
	
			ss.courseId = courses[courseIdx].id
			ss.studentId = students[studentIdx].id

			studentScores[i] = ss
			i++
			*n = i

			fmt.Print("Apakah Anda ingin menambah lagi data nilai mahasiswa? (true/false): ")
			fmt.Scan(&active)
		}
	}
}

func transcript(students students, courses courses, studentScoresData studentScores, nStudent, nstudentScores int) {
	fmt.Println("\n-----------------------------------------")
	fmt.Print("\tAplikasi IGracias Console")
	fmt.Println("\n-----------------------------------------")
	fmt.Println("By: Muhammad Farid")
	fmt.Println("By: Jihan Alifah Maritza")
	fmt.Println("\n-----------------------------------------")

	var active bool = true

	for active {
		var searchStudent bool = true
		var idx int = -1

		for {
			showMahasiswa(students, nStudent)
			var nim string
			fmt.Print("Cari nilai mahasiswa berdasarkan NIM: ")
			fmt.Scan(&nim)

			sortStudentByNim(&students, nStudent)

			idx = searchMahasiswaByNim(nim, students, nStudent)

			if idx == -1 {
				fmt.Printf("Data mahasiswa dengan NIM %s tidak ditemukan.\n", nim)
				fmt.Print("Apakah Anda ingin memasukkan kembali NIM mahasiswa? (true/false): ")
				fmt.Scan(&searchStudent)

				if !searchStudent {
					active = false
					break
				}
			} else {
				active = false
				break
			}
		}

		if idx != -1 {
			var student student = students[idx]
			var ids, length = searchMatkulByStudentId(student.id, studentScoresData, nstudentScores)
			var counter int = 0
			var mapping map[string]studentScore = make(map[string]studentScore)

			for i := 0; i < length; i++ {
				for j := 0; j < nstudentScores; j++ {
					var idx int = searchMatkulById(ids[i], courses, nCourses)
				
					if idx != -1 {
						if _, ok := mapping[courses[i].name]; !ok && 
						studentScoresData[j].courseId == ids[i] && studentScoresData[j].studentId == student.id {
							mapping[courses[idx].name] = studentScoresData[j]
							counter++
						}
					}
				}
			}

			for i := 0; i < counter; i++ {
				var idx int = searchMatkulById(ids[i], courses, nCourses)
				var courseName string = courses[idx].name
				var mapData studentScore = mapping[courseName]
				fmt.Println(student.nim, student.name, courseName, mapData.sks, mapData.quiz, mapData.uts, mapData.uas)
			}

			fmt.Print("Apakah Anda ingin melanjutkan menampilkan transkrip nilai mahasiswa? (true/false): ")
			fmt.Scan(&active)
		}
	}
	clear()
	main()
}

func calculateGrade(quiz, uts, uas float64) string {
	var total float64 = (quiz + uts + uas) / 3
	
	if 80 < total {
		return "A"
	} else if 70 < total && total <= 80 {
		return "AB"
	} else if 65 < total && total <= 70 {
		return "B"
	} else if 60 < total && total <= 65 {
		return "BC"
	} else if 50 < total && total <= 60 {
		return "C"
	} else if 40 < total && total <= 50 {
		return "D"
	} else if total <= 40 {
		return "E"
	}

	return "-"
}

func sortNilaiMahasiswa(studentSummary *studentSummary, nStudentScores int, sort, field string) bool {
	if field != "uas" && field != "total" && field != "uts" && field != "sks" && field != "quiz" {
		fmt.Println("Mohon masukkan nama field yang disediakan!")
		return false
	}

	if sort != "asc" && sort != "desc" {
		fmt.Println("Urutan tidak valid.")
		return false
	}

	for i := 0; i < nStudentScores-1; i++ {
		var min int = i
		
		for j := i+1; j < nStudentScores; j++ {
			var jData studentGroup = studentSummary[j]
			var minData studentGroup = studentSummary[min]
			var jInt, minInt float64 = 0, 0

			if field == "sks" {
				jInt = float64(jData.totalSks)
				minInt = float64(minData.totalSks)
			} else if field == "quiz" {
				jInt = jData.quiz
				minInt = minData.quiz
			} else if field == "uas" {
				jInt = jData.uas
				minInt = minData.uas
			} else if field == "uts" {
				jInt = jData.uts
				minInt = minData.uts
			} else if field == "total" {
				jInt = jData.totalScore
				minInt = minData.totalScore
			}

			if sort == "asc" {
				if jInt < minInt {
					min = j
				}
			} else if sort == "desc" {
				if jInt > minInt {
					min = j
				}
			}
		}
		studentSummary[i], studentSummary[min] = studentSummary[min], studentSummary[i]
	}
	return true
}

func sortStudentByNim(students *students, n int) {
	for i := 1; i < n; i++ {
		var temp student = students[i]
		var j int = i - 1

		for j >= 0 && students[j].nim > temp.nim {
			students[j+1] = students[j]
			j--
		}

		students[j+1] = temp
	}
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

