package main

import "fmt"

func main() {
	// input
	var studentScore int = 80
	var studentName string = "Lendra" // ISI NAMA MAHASISWA

	var studentGrade string
	if studentScore >= 80 && studentScore <= 100 {
		studentGrade = "A"
	} else if studentScore >= 65 && studentScore <= 79 {
		studentGrade = "B+"
	} else if studentScore >= 50 && studentScore <= 64 {
		studentGrade = "B"
	} else if studentScore >= 35 && studentScore <= 49 {
		studentGrade = "C"
	} else if studentScore >= 0 && studentScore <= 34 {
		studentGrade = "D"
	} else {
		studentGrade = "Nilai Tidak Valid"
	}

	fmt.Println("Nama Mahasiswa:", studentName)
	fmt.Println("Nilai:", studentGrade)

	switch studentGrade {
	case "A":
		println("Deskripsi Nilai: Sangat Baik")
		case "B+":
		println("Deskripsi Nilai: Baik")
		case "B":
		println("Deskripsi Nilai: Cukup")
		case "C":
		println("Deskripsi Nilai: Kurang")
		case "D":
		println("Deskripsi Nilai: Sangat Kurang")
		default:
		println("Deskripsi Nilai: Tidak Valid")
	}

}