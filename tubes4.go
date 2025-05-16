package main
import "fmt"

const NMAX int = 100

//Menu Utama
func menuUtama () {
	fmt.Println ("---MENU UTAMA---")
	fmt.Println ("Selamat Datang, kamu mau cari tau stok bahan makanan mu & tanggal expired nya? ayok ikuti instruksi ku!")
	fmt.Println ()
	fmt.Println ("1. Masukkan keterangan bahan makanan")
	fmt.Println ("2. Keluar Program")
	fmt.Println ()
	fmt.Print ("Pilih Opsi : ")
	
}

//Input nama, banyak stok, & expired date bahan makanan	
func bacaOpsiSatu (A *tabMakanan, n int) {
	var i, NMAX int
	
	fmt.Println ()
	fmt.Scan (n)
	fmt.Println("Sekarang masukkan nama, banyak stok, & tanggal expired bahan makanan kamu ya!") 
	
	if n < NMAX {
	    NMAX = n
	}
		
		for i = 1; i <= n; i++ {
			fmt.Println ()
			fmt.Printf ("Bahan ke %d : \n", i)
			
			fmt.Print ("Nama Bahan : ")
			fmt.Scan (&A[i].nama)
			
			fmt.Print ("Banyak Stok : ")
			fmt.Scan (&A[i].banyakStok)
			
			fmt.Print ("Tanggal Kedaluwarsa di bulan Juni : ")
			fmt.Scan (&A[i].expiredDate_Tanggal)
			fmt.Println ()
			
		}
}

//Mencetak list bahan makanan serta atributnya yang sudah di input oleh user
func cetakArray (A *tabMakanan, n int) {
	var i int
	
	fmt.Printf("%-20s %-15s %-15s\n", "Nama Bahan", "Banyak Stok", "Exp Date")
	
		for i = 1; i <= n; i++ {
			fmt.Printf("%-20s %-15d %-15d\n", A[i].nama, A[i].banyakStok, A[i].expiredDate_Tanggal)
			
		}
}

//Tipe data bahan makanan
type bahan struct {
	nama string
	banyakStok int
	expiredDate_Tanggal int
	
}

//Penamaan array
type tabMakanan [NMAX] bahan

//Running Program
func main () {
	var dataMakanan tabMakanan
	var pilihOpsi, nData int
	
	menuUtama ()
	fmt.Scan (&pilihOpsi)
	fmt.Println ()
	fmt.Print ("Berapa banyak bahan makanan yang mau kamu masukkan? : ")
	
	if pilihOpsi == 1 {
	fmt.Scan (&nData)
		bacaOpsiSatu (&dataMakanan, nData)
		cetakArray (&dataMakanan, nData)
		
	}
}