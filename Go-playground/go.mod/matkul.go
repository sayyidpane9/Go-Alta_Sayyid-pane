package main

import (
	"fmt"
	"sort"
)

type MataKuliah struct {
	Judul string
	SKS   int
}

type Mahasiswa struct {
	Nama        string
	MataKuliah  []MataKuliah
}

func (m *Mahasiswa) TambahMataKuliah(judul string, sks int) {
	m.MataKuliah = append(m.MataKuliah, MataKuliah{Judul: judul, SKS: sks})
}

func (m *Mahasiswa) DapatkanMataKuliah() []MataKuliah {
	// membuat salinan array MataKuliah agar tidak mengubah array asli saat diurutkan
	mataKuliah := make([]MataKuliah, len(m.MataKuliah))
	copy(mataKuliah, m.MataKuliah)

	// mengurutkan array MataKuliah berdasarkan judul mata kuliah
	sort.Slice(mataKuliah, func(i, j int) bool {
		return mataKuliah[i].Judul < mataKuliah[j].Judul
	})

	return mataKuliah
}

func main() {
	// membuat objek mahasiswa
	m := Mahasiswa{Nama: "Paijo Pratama"}

	// menambahkan mata kuliah ke objek mahasiswa
	m.TambahMataKuliah("Kalkulus", 4)
	m.TambahMataKuliah("Fisika", 3)
	m.TambahMataKuliah("Pemrograman Berorientasi Objek", 3)
	m.TambahMataKuliah("Kewarganegaraan", 3)

	// mendapatkan mata kuliah dari objek mahasiswa dan mencetak judul mata kuliah
	mataKuliah := m.DapatkanMataKuliah()
	for _, mk := range mataKuliah {
		fmt.Println(mk.Judul)
	}
	fmt.Println(" ")
}
