package main

import "fmt"

type datamovie struct {
	title        string
	rating       float64
	thnRilis     int
	nVote        int
}

const NMAX int = 303

var listFilm [NMAX]datamovie

func main() {
	fmt.Println("Masukkan daftar film:")
	inputfilm()
	var opsi int
	menu()
	fmt.Print(">> Opsi: ")
	fmt.Scan(&opsi)
	for opsi != 5 {
		menu()
		fmt.Print(">> Opsi: ")
		fmt.Scan(&opsi)
		switch opsi {
		case 1:
			showList()
		case 2:
			fmt.Println("Algoritma apa yang akan kamu gunakan?")
			fmt.Println("1. Insertion sort")
			fmt.Println("2. Quick sort")
			pilihAlgoritmaOpsi2()
		case 3:
			fmt.Println("Algoritma apa yang akan kamu gunakan?")
			fmt.Println("1. Insertion sort")
			fmt.Println("2. Quick sort")
			pilihAlgoritmaOpsi3()
		case 4:
			opsiSearch()
		case 5:
			fmt.Println("Terima kasih!")
			fmt.Println("<<============================>>")
		default:
			fmt.Println("Tolong masukkan opsi yang valid!")
			fmt.Println("")
		}
	}
}

func inputfilm() {
	var i int = 0
	for i < NMAX {
		fmt.Scan(&listFilm[i].title, &listFilm[i].rating, &listFilm[i].nVote, &listFilm[i].thnRilis)
		i += 1
	}
}

func menu() {
	fmt.Println("<<==========[[MENU]]==========>>")
	fmt.Println("")
	fmt.Println("1. Tampilkan daftar film")
	fmt.Println("2. Urutan film berdasarkan rating")
	fmt.Println("3. Urutan film berdasarkan jumlah vote dari peminat")
	fmt.Println("4. Cari film berdasarkan tahun rilis")
	fmt.Println("5. Exit")
	fmt.Println("")
	fmt.Println("<<============================>>")
}

func showList() {
	var i int = 0
	for i < NMAX {
		fmt.Println((i + 1), listFilm[i].title, ", ", listFilm[i].rating, ", ", listFilm[i].thnRilis, ", ", listFilm[i].nVote)
		i += 1
	}
}

func insertionSortOpsi2() {
	for i := 1; i < NMAX; i++ {
		currentFilm := listFilm[i]
		j := i - 1
		for j >= 0 && listFilm[j].rating < currentFilm.rating {
			listFilm[j+1] = listFilm[j]
			j--
		}
		listFilm[j+1] = currentFilm
	}

	fmt.Println("Film berhasil diurutkan berdasarkan rating (descending).")
}

func insertionSortOpsi3() {
	for i := 1; i < NMAX; i++ {
		currentFilm := listFilm[i]
		j := i - 1
		for j >= 0 && listFilm[j].nVote < currentFilm.nVote {
			listFilm[j+1] = listFilm[j]
			j--
		}
		listFilm[j+1] = currentFilm
	}

	fmt.Println("Film berhasil diurutkan berdasarkan jumlah vote (descending).")
}

func quickSortOpsi2() {
	stack := make([][2]int, 0)
	stack = append(stack, [2]int{0, NMAX - 1})

	for len(stack) > 0 {
		// Pop pasangan low dan high dari stack
		highLow := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		low, high := highLow[0], highLow[1]

		if low < high {
			// Partisi array
			pivot := listFilm[high].rating
			i := low - 1

			for j := low; j < high; j++ {
				if listFilm[j].rating > pivot { // Urut descending
					i++
					listFilm[i], listFilm[j] = listFilm[j], listFilm[i]
				}
			}

			// Tempatkan pivot di posisi yang benar
			listFilm[i+1], listFilm[high] = listFilm[high], listFilm[i+1]
			p := i + 1

			// Push bagian kiri dan kanan ke stack
			stack = append(stack, [2]int{low, p - 1})
			stack = append(stack, [2]int{p + 1, high})
		}
	}

	fmt.Println("Film berhasil diurutkan berdasarkan rating (descending).")
}

func quickSortOpsi3() {
	stack := make([][2]int, 0)
	stack = append(stack, [2]int{0, NMAX - 1})

	for len(stack) > 0 {
		// Pop pasangan low dan high dari stack
		highLow := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		low, high := highLow[0], highLow[1]

		if low < high {
			// Partisi array
			pivot := listFilm[high].nVote
			i := low - 1

			for j := low; j < high; j++ {
				if listFilm[j].nVote > pivot { // Urut descending
					i++
					listFilm[i], listFilm[j] = listFilm[j], listFilm[i]
				}
			}

			// Tempatkan pivot di posisi yang benar
			listFilm[i+1], listFilm[high] = listFilm[high], listFilm[i+1]
			p := i + 1

			// Push bagian kiri dan kanan ke stack
			stack = append(stack, [2]int{low, p - 1})
			stack = append(stack, [2]int{p + 1, high})
		}
	}

	fmt.Println("Film berhasil diurutkan berdasarkan jumlah vote (descending).")
}

func opsiSearch() {
	var pilih int
	fmt.Println(">> Apakah anda ingin mencari dalam satu tahun saja, atau dalam jangkauan beberapa tahun?")
	fmt.Println("1. Dalam satu tahun tertentu.")
	fmt.Println("2. Dalam jangkauan satu tahun ke tahun lain")
	fmt.Print(">> ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		searchMovieInYr()
	} else if pilih == 2 {
		searchMovieInYrRange()
	} else {
		fmt.Println("Opsi invalid!")
		fmt.Println("")
		opsiSearch()
	}
}

func searchMovieInYr() {
	var year int
	fmt.Print(">> Masukkan tahun rilis yang ingin dicari: ")
	fmt.Scan(&year)

	fmt.Println("Film yang dirilis pada tahun", year, ":")
	found := false
	for i := 0; i < NMAX; i++ {
		if listFilm[i].thnRilis == year {
			fmt.Println((i + 1), listFilm[i].title, ", ", listFilm[i].rating, ", ", listFilm[i].thnRilis, ", ", listFilm[i].nVote)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada film yang ditemukan pada tahun", year, ".")
	}
}

func searchMovieInYrRange() {
	var yearA, yearB int
	fmt.Print(">> Pilih tahun awal: ")
	fmt.Scan(&yearA)
	fmt.Print(">> Pilih tahun akhir: ")
	fmt.Scan(&yearB)

	// Pastikan yearA <= yearB
	if yearA > yearB {
		yearA, yearB = yearB, yearA
	}

	fmt.Println("Film yang dirilis antara tahun", yearA, "dan", yearB, ":")
	found := false
	for i := 0; i < NMAX; i++ {
		if listFilm[i].thnRilis >= yearA && listFilm[i].thnRilis <= yearB {
			fmt.Println((i + 1), listFilm[i].title, ", ", listFilm[i].rating, ", ", listFilm[i].thnRilis, ", ", listFilm[i].nVote)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada film yang ditemukan antara tahun", yearA, "dan", yearB)
	}
}

func pilihAlgoritmaOpsi2() {
	var opsi int
	fmt.Print(">> ")
	fmt.Scan("&opsi")
	if opsi == 1 {
		insertionSortOpsi2()
	} else if opsi == 2 {
		quickSortOpsi2()
	} else {
		fmt.Println("Opsi invalid!")
		pilihAlgoritmaOpsi2()
	}
}

func pilihAlgoritmaOpsi3() {
	var opsi int
	fmt.Print(">> ")
	fmt.Scan("&opsi")
	if opsi == 1 {
		insertionSortOpsi3()
	} else if opsi == 2 {
		quickSortOpsi3()
	} else {
		fmt.Println("Opsi invalid!")
		pilihAlgoritmaOpsi3()
	}
}
