package main
import "fmt"

const NMAX int = 100

//Tipe data bahan makanan
type bahan struct {
	nama string
	banyakStok int
	expiredDate_Tanggal int
}

//Penamaan array
type tabMakanan [NMAX] bahan


//Menu Utama
func menuUtama () {
	fmt.Println ("")	
	fmt.Println ("=== Aplikasi Manajemen Stok Bahan Makanan ===")
	fmt.Println ("1. Masukkan keterangan bahan makanan")
	fmt.Println ("2. Tambahkan bahan makanan")
	fmt.Println ("3. Ubah data bahan makanan")
	fmt.Println ("4. Hapus data bahan makanan")
	fmt.Println ("5. Cari nama makanan")
	fmt.Println ("6. Gunakan bahan makanan")
	fmt.Println ("7. Cek tanggal kedaluwarsa")
	fmt.Println ("8. Mengurutkan bahan makanan berdasarkan Tanggal kedaluwarsa/jumlah stok makanan")
	fmt.Println ("99. Keluar Program")
	fmt.Print ("Pilih Opsi kamu : ")
	
}


//Input nama, banyak stok, & expired date bahan makanan	
func bacaOpsiSatu (A *tabMakanan, n int) {
	var i int
	
	fmt.Println ()
	fmt.Scan (n)
	
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


//Input nama, banyak stok, & expired date bahan makanan	
func bacaOpsiDua (A *tabMakanan, n *int) {
	
	fmt.Println ()
	
		fmt.Println ()
		fmt.Printf ("Bahan ke %d : \n", *n+1)
		
		fmt.Print ("Nama Bahan Tambahan : ")
		fmt.Scan (&A[*n+1].nama)
			
		fmt.Print ("Banyak Stok : ")
		fmt.Scan (&A[*n+1].banyakStok)
		
		fmt.Print ("Tanggal Kedaluwarsa di bulan Juni : ")
		fmt.Scan (&A[*n+1].expiredDate_Tanggal)
		
			
		
		fmt.Println ()
		
	*n = *n + 1
	fmt.Println ()
		
}

//Mengubah data makanan
func bacaOpsiTiga (A *tabMakanan, n int) {
	var i int
	var namaBahan string
	
	fmt.Scan (&namaBahan)
	
	for i = 1; i < n; i++ {
		if A[i].nama == namaBahan {
			fmt.Print ("Nama Bahan Baru : ")
			fmt.Scan (&A[i].nama)
				
			fmt.Print ("Banyak Stok : ")
			fmt.Scan (&A[i].banyakStok)
				
			fmt.Print ("Tanggal Kedaluwarsa di bulan Juni : ")
			fmt.Scan (&A[i].expiredDate_Tanggal)
			fmt.Println ()
		}		
	}
}


//Menghapus data makanan
func bacaOpsiEmpat (A *tabMakanan, n int) {
	var i int
	var namaBahan string
	
	fmt.Println ("Ketik nama bahan makanan yang mau di hapus : ")
	fmt.Scan (&namaBahan)
	
		if A[i].nama == namaBahan {
			fmt.Print ("Nama Bahan Baru : ")
			fmt.Scan (&A[i].nama)
				
			fmt.Print ("Banyak Stok : ")
			fmt.Scan (&A[i].banyakStok)
			
			fmt.Print ("Tanggal Kedaluwarsa di bulan Juni : ")
			fmt.Scan (&A[i].expiredDate_Tanggal)
			fmt.Println ()
			
		}	
}


//Mencari nama bahan makanan, misal kamu mau cari bahan telor, nanti telor bakal ke cetak beserta total stok berapa & tanggal Kedaluwarsa kapan
func bacaOpsiLima(A *tabMakanan, n int) {
	var i int
	var namaCari string
	fmt.Scan(n) 
	fmt.Print("Masukkan nama bahan yang ingin dicari: ")
	fmt.Scan(&namaCari)

	for i = 1; i <= n; i++ {
		if A[i].nama == namaCari {
			fmt.Println("-----------------------------------------")
			fmt.Printf("Nama: %s, Stok: %d, Expired Date: %d\n", A[i].nama, A[i].banyakStok, A[i].expiredDate_Tanggal)
			fmt.Println("-----------------------------------------")
			fmt.Println ()
			return
		}
	}
	fmt.Println("Bahan tidak ditemukan.")
}




//Menggunakan bahan makanan sebanyak "x" dengan cara mengetik nama bahan makanan. Misal mau pake telor sebanyak x, nanti datanya bakal ke cetak dan telor nya berkurang.
func bacaOpsiEnam(A *tabMakanan, n int) {
	var i, jumlah int
	var namaBahan string
	
	fmt.Scan(n) 
	fmt.Print("Masukkan nama bahan yang ingin digunakan: ")
	fmt.Scan(&namaBahan)

	for i = 1; i <= n; i++ {
		if A[i].nama == namaBahan {
			fmt.Printf("Stok tersedia: %d\n", A[i].banyakStok)
			fmt.Print("Masukkan jumlah yang ingin digunakan: ")
			fmt.Scan(&jumlah)

			if jumlah > A[i].banyakStok {
				fmt.Println("Stok tidak cukup.")
			} else {
				A[i].banyakStok -= jumlah
				fmt.Printf("Stok %s sekarang: %d\n", A[i].nama, A[i].banyakStok)
			}
			return
		}
	}
	fmt.Println("Bahan tidak ditemukan.")
}



//Mengecek apakah bahan makanan sudah mau kedaluwarsa atau belum
func bacaOpsiTujuh (A *tabMakanan, n int) {
	var i int
	var namaBahan string
	
	fmt.Println ("Ketik nama bahan makanan yang mau di cek!")
	fmt.Scan (&namaBahan)
	
	for i = 1; i < n; i++ {
    	if A[i].nama == namaBahan {
        	if (A[i].expiredDate_Tanggal) <= 3 {
        		fmt.Printf ("Mendekati Tanggal Kedaluwarsa!")
        		
        	} else if (A[i].expiredDate_Tanggal) > 3 && (A[i].expiredDate_Tanggal <= 10) {
        		fmt.Printf ("Kedaluwarsa masih lumayan lama")
        		
        	} else if (A[i].expiredDate_Tanggal) > 10 {
        		fmt.Printf ("Kedaluwarsa masih lama")
        		
        	}
		
	    }
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


//Running Program
func main () {
	var dataMakanan tabMakanan
	var pilihOpsi int
	var nData int
	
	for {
	menuUtama ()
	fmt.Scan (&pilihOpsi)
	fmt.Println ()	
	
		switch pilihOpsi {
			
			case 1 :
				fmt.Print ("Berapa banyak bahan makanan yang mau kamu masukkan? : ")
				fmt.Scan (&nData)
				bacaOpsiSatu (&dataMakanan, nData)
				cetakArray (&dataMakanan, nData)
					
			case 2 :
				bacaOpsiDua (&dataMakanan, &nData)
				cetakArray (&dataMakanan, nData)
			
			case 3 :
				fmt.Println ("Ketik nama bahan makanan yang mau di ubah : ")
				fmt.Scan (&dataMakanan)
				bacaOpsiTiga (&dataMakanan, nData)
				cetakArray (&dataMakanan, nData)
				
			case 4 :
				fmt.Scan (&nData)
				bacaOpsiEmpat (&dataMakanan, nData)
				cetakArray (&dataMakanan, nData)
				
			case 5 :
				bacaOpsiLima (&dataMakanan, nData)
				cetakArray (&dataMakanan, nData)
					
			case 6 :
				bacaOpsiEnam (&dataMakanan, nData)
				cetakArray (&dataMakanan, nData)
					
			case 7 :
				fmt.Scan (&nData)
				bacaOpsiTujuh (&dataMakanan, nData)
			
			case 99 :
				fmt.Println ("Terimakasih telah menggunakan program :D")
				return
				
			default :
				fmt.Println ("Opsi tidak valid, coba lagi")
		}	
	}
}
