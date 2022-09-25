package main

import "fmt"

type Student struct {
	No_Absen  int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	var student = []Student{

		{No_Absen: 1, Nama: "Sena,", Alamat: " Jl. Tempe,", Pekerjaan: " Karyawan Swasta,", Alasan: " Tuntutan pekerjaan"},
		{No_Absen: 2, Nama: "Nasha,", Alamat: " Jl. Tahu,", Pekerjaan: " Pegawai Negeri Sipil,", Alasan: " Tertarik dengan bahasa Golang"},
		{No_Absen: 3, Nama: "Alfia,", Alamat: " Jl. Kangkung,", Pekerjaan: " Atlet Basket,", Alasan: " Ingin menambah pengetahuan"},
		{No_Absen: 4, Nama: "Sanjaya,", Alamat: " Jl. Bayam,", Pekerjaan: " Pelukis,", Alasan: " Ingin menggambar logo Google"},
		{No_Absen: 5, Nama: "Wita,", Alamat: " Jl. Terong,", Pekerjaan: " Penari,", Alasan: " Karena ingin mengisi CV untuk pekerjaan yang baru"},
		{No_Absen: 6, Nama: "Fachry,", Alamat: " Jl. Wortel,", Pekerjaan: " Binaraga,", Alasan: " Untuk melatih kakuatan tangan mengetik"},
		{No_Absen: 7, Nama: "Dwi,", Alamat: " Jl. Jagung,", Pekerjaan: " Wirausaha,", Alasan: " Ingin membuat web perusahaan"},
		{No_Absen: 8, Nama: "Syafiq,", Alamat: " Jl. Brokoli,", Pekerjaan: " Koki,", Alasan: " Salah mendaftar kelas di Digitalent"},
		{No_Absen: 9, Nama: "Lutpi,", Alamat: " Jl. Kol,", Pekerjaan: " Dokter,", Alasan: " Ingin membuat aplikasi kesehatan"},
		{No_Absen: 10, Nama: "Ojan,", Alamat: " Jl. Kentang,", Pekerjaan: " Arsitek,", Alasan: " Karena bahasa Golang lebih mudah dibanding bahasa lain"},
	}

	var absen int

	/*
	   Pengulangan apabila No. Absen yang dimasukkan salah menggunakan outerLoop

	   Contoh :
	   Masukkan Nomor Absen : 13
	   Data Siswa Tidak Tersedia. Silahkan masukkan kembali No. Absen.

	*/

outerLoop:
	for i := 0; i < 100; i++ {
		fmt.Println("Masukkan nomor absen : ")
		fmt.Scanln(&absen)

		if i <= 10 {
			for i, v := range student {
				if absen == i+1 {
					fmt.Printf("%+v\n", v)
					break outerLoop
				}
			}
		}
		fmt.Println("Data Siswa Tidak Tersedia. Silahkan masukkan kembali No. Absen.")
		fmt.Println()
	}
}
