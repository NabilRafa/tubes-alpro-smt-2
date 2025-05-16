package main
import "fmt"

func menu () {
	fmt.Println ("---MENU UTAMA---")
	fmt.Println ("Selamat Datang, kamu mau cari tau stok bahan makanan mu & tanggal expired nya? ayok ikuti instruksi ku!")
	fmt.Println ()
	fmt.Println ("1 Untuk masukkan keterangan bahan makanan")
	fmt.Println ("2 Untuk keluar")
	fmt.Println ()
	fmt.Print ("Opsi : ")
	
}
	
func bacaOpsiSatu (A tabMakanan) {
	var i int
	
	fmt.Scan (
}

func menuInputData (jmlData int) {
	fmt.Printf ("Karena kamu mau masukin %d data makanan, sekarang kamu tulis nama, jumlah stok, dan expired date makanannya nya ya!\n", jmlData)
}

//func inputData (namaMakanan, jmlStok, expDate int) {
	
//}


const NMAX int = 100


type bahan struct {
	nama string
	banyakStok int
	expiredDate_Tanggal int
	expiredDate_Bulan int
	expiredDate_Tahun int
	
}

type tabMakanan [NMAX] bahan

func main () {
	var dataMakanan tabMakanan
	var pilihOpsi int
	var namaMakanan string
	var i, jmlData int
	
	menu ()
	fmt.Scan (&pilihOpsi)
	fmt.Println ()
	
	if pilihOpsi == 1 {
		opsiSatu 
		
		menuInputData (jmlData)
		
		for i = 1; i < jmlData; i++ {
		
			fmt.Scan (&dataMakanan[i].nama, &dataMakanan[i].banyakStok, &dataMakanan[i].expiredDate_Tanggal, &dataMakanan[i].expiredDate_Bulan, &dataMakanan[i].expiredDate_Tahun) 
			inputData (dataMakanan.nama[i], dataMakanan.banyakStok[i], dataMakanan.expiredDate[i])
			
		}
	
	} else if pilihOpsi == 2{
	
	} else {
		fmt.Println ("Opsi tidak valid")
		
	}
}