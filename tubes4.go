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
func bacaOpsiSatu (A *tabMakanan, n *int) {
	var i, temp int
	
	
	for i = 0; i < *n; i++ {
		fmt.Println ()
		fmt.Printf ("Bahan ke %d : \n", i+1)
		
		fmt.Print ("Nama Bahan : ")
		fmt.Scan (&A[i].nama)
			
		fmt.Print ("Banyak Stok : ")
		fmt.Scan (&A[i].banyakStok)
			
		fmt.Print ("Tanggal Kedaluwarsa di bulan Juni : ")
		fmt.Scan (&temp)
			
			if temp < 0 || temp >= 32 {
				fmt.Println ("Tanggal tidak valid, coba lagi")
				fmt.Println ("")
				i--
				*n--
				return
			
			} else {
				A[i].expiredDate_Tanggal = temp 
				
				
			}
	}
}	


//Input nama, banyak stok, & expired date bahan makanan	
func bacaOpsiDua (A *tabMakanan, n *int) {
	var temp int
	
	fmt.Println ()
	fmt.Printf ("Bahan ke %d : \n", *n+1)
		
	fmt.Print ("Nama Bahan Tambahan : ")
	fmt.Scan (&A[*n].nama)
			
	fmt.Print ("Banyak Stok : ")
	fmt.Scan (&A[*n].banyakStok)
			
	fmt.Print ("Tanggal Kedaluwarsa di bulan Juni : ")
	fmt.Scan (&temp)
		
		if temp >= 32 {
			fmt.Println ("Tanggal tidak valid")
			return
				
		}
				
		A[*n].expiredDate_Tanggal = temp 
		fmt.Println ()	
	*n++
	fmt.Println ()	
}

//Mengubah data makanan
func bacaOpsiTiga (A *tabMakanan, n int) {
	var i, temp int
	var namaBahan string
	
	fmt.Scan (&namaBahan)
	
	for i = 0; i < n; i++ {
		if A[i].nama == namaBahan {
			fmt.Println ("Nama Bahan Baru : ")
			fmt.Scan (&A[i].nama)
				
			fmt.Println ("Banyak Stok : ")
			fmt.Scan (&A[i].banyakStok)
					
			fmt.Println ("Tanggal Kedaluwarsa di bulan Juni : ")
			fmt.Scan (&temp)
			
				if temp >= 32 {
					fmt.Println ("Tanggal tidak valid")
					return
				
				}
				
				A[i].expiredDate_Tanggal = temp 
				fmt.Println ()
				return
			
		}		
	}
	fmt.Println ("Bahan tidak ditemukan")
}


//Menghapus data makanan
func bacaOpsiEmpat (A *tabMakanan, n *int) {
	var i, j int
	var namaBahan string
	
	fmt.Println ("Ketik nama bahan makanan yang mau di hapus : ")
	fmt.Scan (&namaBahan)
	
	for i = 0; i < *n; i++ {
		if A[i].nama == namaBahan {
			for j = i; j < *n-1; j++ {
				A[j] = A[j+1]
				i--
			}
		
		*n--
		
		fmt.Printf ("Bahan makanan %s berhasil dihapus\n", namaBahan)
		fmt.Println()
		return
		}
	}
	fmt.Println ("Bahan tidak ditemukan")
}


//Mencari nama bahan makanan, misal kamu mau cari bahan telor, nanti telor bakal ke cetak beserta total stok berapa & tanggal Kedaluwarsa kapan
func bacaOpsiLima (A *tabMakanan, n int) {
	var i int
	var namaCari string
	
	fmt.Print("Masukkan nama bahan yang ingin dicari: ")
	fmt.Scan(&namaCari)

	for i = 0; i < n; i++ {
		if A[i].nama == namaCari {
			fmt.Printf("Nama: %s\nStok: %d\nExpired Date: %d\n", A[i].nama, A[i].banyakStok, A[i].expiredDate_Tanggal)
			return
		}
	}
	fmt.Println("Bahan tidak ditemukan.")
}


//Menggunakan bahan makanan sebanyak "x" dengan cara mengetik nama bahan makanan. Misal mau pake telor sebanyak x, nanti datanya bakal ke cetak dan telor nya berkurang.
func bacaOpsiEnam (A *tabMakanan, n *int) {
	var i, jumlah int
	var namaBahan string 
	
	fmt.Println ("Masukkan Bahan yang ingin digunakan :")
	fmt.Scan(&namaBahan) 
	
	for i = 0; i < *n; i++ {
		if A[i].nama == namaBahan {
			fmt.Println()
			fmt.Printf("Stok tersedia: %d\n", A[i].banyakStok)
			fmt.Print("Masukkan jumlah yang ingin digunakan: ")
			fmt.Scan(&jumlah)

			if jumlah > A[i].banyakStok {
				fmt.Println("Stok tidak cukup.")
				
			} else {
				A[i].banyakStok -= jumlah
				fmt.Println()
				fmt.Printf("Stok %s sekarang: %d\n", A[i].nama, A[i].banyakStok)
				
				if A[i].banyakStok < 1 {
					hapusNol (A, n, namaBahan)
				}
			}
			return
		}
	}
	fmt.Println("Bahan tidak ditemukan.")
}


//Menghapus data makanan bila stok habis
func hapusNol (A *tabMakanan, n *int, namaBahan string) {
	var i, j int
		
	for i = 0; i < *n; i++ {
		if A[i].nama == namaBahan {
			for j = i; j < *n-1; j++ {
				A[j] = A[j+1]
				i--
			}
		
		*n--
		
		fmt.Printf ("Bahan makanan %s dihapus karena stok habis\n", namaBahan)
		fmt.Println()
		return
		}
	}
	fmt.Println ("Bahan tidak ditemukan")
}


//Mengecek apakah bahan makanan sudah mau kedaluwarsa atau belum
func bacaOpsiTujuh (A *tabMakanan, n int) {
	var i, tglNow int
	var selisihHari int
	var namaBahan string
	
	
	fmt.Println ("Ketik nama bahan makanan yang mau di cek!")
	fmt.Scan (&namaBahan)
	fmt.Println()
	
	fmt.Println ("Ketik tanggal hari ini!")
	fmt.Scan (&tglNow)
	fmt.Println()
	
	if tglNow <= 31 {
		for i = 0; i < n; i++ {
			if A[i].nama == namaBahan {
				selisihHari = ((A[i].expiredDate_Tanggal) - tglNow)
				if tglNow < (A[i].expiredDate_Tanggal){
					if selisihHari <= 3 {
						fmt.Printf ("Sisa %d hari, Mendekati Tanggal Kedaluwarsa!", selisihHari)
						fmt.Println()
						
					} else if selisihHari > 3 && selisihHari <= 10 {
						fmt.Printf ("Sisa %d hari, Kedaluwarsa masih lumayan lama", selisihHari)
						fmt.Println()
						
					} else if selisihHari > 10 {
						fmt.Printf ("Sisa %d hari, Kedaluwarsa masih lama", selisihHari)
						fmt.Println()
					}
					
				} else if tglNow > (A[i].expiredDate_Tanggal) {
					fmt.Printf ("Sudah kedaluwarsa %d hari yang lalu!", selisihHari * (-1))
					
				}
			
			}
		}
		return
	}
	fmt.Println ("Tanggal hari ini tidak valid")
}


//Mengurutkan bahan makanan berdasarkan jumlah stok (Ascending)
func bacaOpsiDelapan(A *tabMakanan, n int) {
	var i, j, x int
	var temp bahan
	
	fmt.Println("Pilih jenis sorting Banyak Stok:")
	fmt.Println("1. Ascending (dari stok paling sedikit)")
	fmt.Println("2. Descending (dari stok paling banyak)")
				
	fmt.Print("Pilihan kamu: ")
	fmt.Scan(&x)
	
	//Ascending
	if x == 1 {
		for i = 0; i < n-1; i++ {
			for j = 0; j < n-i-1; j++ {
				if A[j].banyakStok > A[j+1].banyakStok {
					temp = A[j]
					A[j] = A[j+1]
					A[j+1] = temp
				}
			}
		}
		fmt.Println ()
		fmt.Println ("Mengurutkan bahan makanan berdasarkan stok (Ascending)")
	
	//Descending
	} else if x == 2 {
		for i = 0; i < n-1; i++ {
			for j = 0; j < n-i-1; j++ {
				if A[j].banyakStok < A[j+1].banyakStok {
					temp = A[j]
					A[j] = A[j+1]
					A[j+1] = temp
				}
		
			}
		}
		fmt.Println ()
		fmt.Println ("Mengurutkan bahan makanan berdasarkan stok (Descending)")
		
		
	} else {
		fmt.Println ("Input tidak valid")
		fmt.Println ()
		return
	}
}


//Mencetak list bahan makanan serta atributnya yang sudah di input oleh user
func cetakArray (A *tabMakanan, n int) {
	var i int
	
	fmt.Println ("")
	fmt.Println ("--------------------------------------------------------------")
	
	fmt.Printf("%-20s %-15s %-15s\n", "Nama Bahan", "Banyak Stok", "Exp Date")
	
		for i = 0; i < n; i++ {
			fmt.Printf("%-20s %-15d %-15d\n", A[i].nama, A[i].banyakStok, A[i].expiredDate_Tanggal)
			
		}
		fmt.Println ("--------------------------------------------------------------")
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
				bacaOpsiSatu (&dataMakanan, &nData)
				cetakArray (&dataMakanan, nData)
					
			case 2 :
				bacaOpsiDua (&dataMakanan, &nData)
				cetakArray (&dataMakanan, nData)
			
			case 3 :
				fmt.Println ("Ketik nama bahan makanan yang mau di ubah : ")
				bacaOpsiTiga (&dataMakanan, nData)
				cetakArray (&dataMakanan, nData)
				
			case 4 :
				bacaOpsiEmpat (&dataMakanan, &nData)
				cetakArray (&dataMakanan, nData)
				
			case 5 :
				bacaOpsiLima (&dataMakanan, nData)
				cetakArray (&dataMakanan, nData)
					
			case 6 :
				bacaOpsiEnam (&dataMakanan, &nData)
				cetakArray (&dataMakanan, nData)
					
			case 7 :
				bacaOpsiTujuh (&dataMakanan, nData)
				cetakArray (&dataMakanan, nData)
				
			case 8:
				bacaOpsiDelapan (&dataMakanan, nData)
				cetakArray(&dataMakanan, nData)
				
			case 99 :
				fmt.Println ("Terimakasih telah menggunakan program :D")
				return
				
			default :
				fmt.Println ("Opsi tidak valid, coba lagi")
		
		}	
	}
}
