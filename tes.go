package main
import "fmt"

const NMAX int = 2

//Menu Utama
func menuUtama () {
	fmt.Println ("---MENU UTAMA---")
	fmt.Println ("Selamat Datang, kamu mau cari tau stok bahan makanan mu & tanggal expired nya? ayok ikuti instruksi ku!")
	fmt.Println ()
	fmt.Println ("1 Untuk masukkan keterangan bahan makanan")
	fmt.Println ("2 Untuk keluar")
	fmt.Println ()
	fmt.Print ("Opsi : ")
	
}

//Input nama, banyak stok, & expired date bahan makanan	
func bacaOpsiSatu (A *tabMakanan, n int) {
	var i int
	
		for i = 1; i < n; i++ {
			fmt.Scan (&A[i].nama, &A[i].banyakStok, &A[i].expiredDate_Tanggal, &A[i].expiredDate_Bulan, &A[i].expiredDate_Tahun)
		}
}

//Mencetak list bahan makanan serta atributnya yang sudah di input oleh user
func cetakArray (A tabMakanan, n int) {
	var i int
	
		for i = 1; i < n; i++ {
			fmt.Printf("%20s %15s %6s-%6s-%6s", "Nama Bahan", "Banyak Stok", "Exp Date", "Exp Bulan", "Exp Tahun")
			fmt.Printf("%20d %15d %6d-%6d-%6d", A[i].nama, &A[i].banyakStok, &A[i].expiredDate_Tanggal, &A[i].expiredDate_Bulan, &A[i].expiredDate_Tahun)
			
		}
}

//Tipe data bahan makanan
type bahan struct {
	nama string
	banyakStok int
	expiredDate_Tanggal int
	expiredDate_Bulan int
	expiredDate_Tahun int
	
}

//Penamaan array
type tabMakanan [NMAX] bahan

func main () {
	var dataMakanan tabMakanan
	var pilihOpsi int
	
	menuUtama ()
	fmt.Scan (&pilihOpsi)
	fmt.Println ()
	
	if pilihOpsi == 1 {
		bacaOpsiSatu (&dataMakanan, pilihOpsi)
		cetakArray (&dataMakanan, &pilihOpsi)
		
	}
}1package main
import "fmt"

const NMAX int = 2

//Menu Utama
func menuUtama () {
	fmt.Println ("---MENU UTAMA---")
	fmt.Println ("Selamat Datang, kamu mau cari tau stok bahan makanan mu & tanggal expired nya? ayok ikuti instruksi ku!")
	fmt.Println ()
	fmt.Println ("1 Untuk masukkan keterangan bahan makanan")
	fmt.Println ("2 Untuk keluar")
	fmt.Println ()
	fmt.Print ("Opsi : ")
	
}

//Input nama, banyak stok, & expired date bahan makanan	
func bacaOpsiSatu (A *tabMakanan, n int) {
	var i int
	
		for i = 1; i < n; i++ {
			fmt.Scan (&A[i].nama, &A[i].banyakStok, &A[i].expiredDate_Tanggal, &A[i].expiredDate_Bulan, &A[i].expiredDate_Tahun)
		}
}

//Mencetak list bahan makanan serta atributnya yang sudah di input oleh user
func cetakArray (A tabMakanan, n int) {
	var i int
	
		for i = 1; i < n; i++ {
			fmt.Printf("%20s %15s %6s-%6s-%6s", "Nama Bahan", "Banyak Stok", "Exp Date", "Exp Bulan", "Exp Tahun")
			fmt.Printf("%20d %15d %6d-%6d-%6d", A[i].nama, &A[i].banyakStok, &A[i].expiredDate_Tanggal, &A[i].expiredDate_Bulan, &A[i].expiredDate_Tahun)
			
		}
}

//Tipe data bahan makanan
type bahan struct {
	nama string
	banyakStok int
	expiredDate_Tanggal int
	expiredDate_Bulan int
	expiredDate_Tahun int
	
}

//Penamaan array
type tabMakanan [NMAX] bahan

func main () {
	var dataMakanan tabMakanan
	var pilihOpsi int
	
	menuUtama ()
	fmt.Scan (&pilihOpsi)
	fmt.Println ()
	
	if pilihOpsi == 1 {
		bacaOpsiSatu (&dataMakanan, pilihOpsi)
		cetakArray (dataMakanan, pilihOpsi)
		
	}
}