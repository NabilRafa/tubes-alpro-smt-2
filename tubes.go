package main
import "fmt"

func menu () {
	fmt.Println ("---MENU UTAMA---")
	fmt.Println ("Selamat Datang, kamu mau cari tau stok bahan makanan mu & tanggal expired nya? ayok ikuti instruksi ku!")
	fmt.Println ()
	fmt.Println (""1" Untuk masukkan keterangan bahan makanan")
	fmt.Println (""2" Untuk keluar")
	fmt.Println ()
	fmt.Print ("Opsi : ")
	
}
	
	
func menuInputData (jmlData int) {
	fmt.Printf ("Karena kamu mau masukin %d data makanan, sekarang kamu tulis nama, jumlah stok, dan expired date makanannya nya ya!\n", jmlData)
}

func inputData (A *tabBahan int) {
	
}

func opsiSatu (opsi int) {
	var i, jmlData int
	
	if opsi == 1 {
		fmt.Println ("Oke, sekarang berapa banyak stok makanan yang mau kamu hitung?")
		fmt.Scan (&jmlData)
		menuInputData (jmlData)
		
		for i = 1; i < jmlData; i++ {
		
			fmt.Scan (&dataMakanan[i].nama, &dataMakanan[i].jumlahStok, &dataMakanan[i].expiredDate) 
			inputData  (dataMakanan[i].nama, dataMakanan[i].jumlahStok, dataMakanan[i].expiredDate)
			
		}
	
	} else if opsi == 2{
		break
	
	} else {
		fmt.Println ("Opsi tidak valid")
		fmt.Println ("Silahkan ulangi ya!")
		break
		
	}
}


const NMAX int = 100


type makanan struct {
	nama string
	banyakStok int
	expiredDate int
}

func main () {
	var dataMakanan makanan
	var pilihOpsi int
	var namaMakanan string
	var i int
	
	menu ()
	fmt.Scan (&pilihOpsi)
	pilihOpsi (pilihOpsi)
	fmt.Println ()
}