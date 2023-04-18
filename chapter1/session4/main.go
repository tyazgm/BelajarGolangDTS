package main

import (
	"fmt"
	"os"
	"strconv"
)

type Siswa struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	dataSiswa := make([]Siswa, 0)

	GetData(&dataSiswa)

	arg := os.Args
	nomorAbsen, err := strconv.Atoi(arg[1])
	if err != nil {
		panic(err)
	}

	PrintData(dataSiswa[nomorAbsen-1])
}

func GetData(dataSiswa *[]Siswa) {

	*dataSiswa = append(*dataSiswa,
		Siswa{
			Nama:      "Andi",
			Alamat:    "Jalan Kenanga No.3",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Untuk tugas akhir",
		}, Siswa{
			Nama:      "Tyaz",
			Alamat:    "Kenten Indah",
			Pekerjaan: "Tidak Bekerja",
			Alasan:    "Menambah portofolio",
		}, Siswa{
			Nama:      "Budi",
			Alamat:    "Jalan yang membawaku bertemu denganmu",
			Pekerjaan: "Part-Time Traveler",
			Alasan:    "Ikut teman",
		})
}

func PrintData(dataSiswa Siswa) {
	fmt.Printf(" Nama	: %s\n Alamat	: %s\n Pekerjaan: %s\n Alasan	: %s",
		dataSiswa.Nama, dataSiswa.Alamat, dataSiswa.Pekerjaan, dataSiswa.Alasan)
}
