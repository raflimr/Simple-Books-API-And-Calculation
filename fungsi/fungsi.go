package fungsi

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func HitungSegitiga(alas, tinggi int, hitung string, hasil chan int) {
	luas := (alas * tinggi) / 2
	keliling := 3 * alas
	if hitung == "luas" {
		hasil <- luas
	} else {
		hasil <- keliling
	}
}

func HitungPersegi(sisi int, hitung string, hasil chan int) {
	luas := sisi * sisi
	keliling := 4 * sisi
	if hitung == "luas" {
		hasil <- luas
	} else {
		hasil <- keliling
	}
}
func HitungPersegiPanjang(panjang, lebar int, hitung string, hasil chan int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	if hitung == "luas" {
		hasil <- luas
	} else {
		hasil <- keliling
	}
}

func HitungLingkaran(jariJari float32, hitung string, hasil chan float32) {
	luas := 3.14 * jariJari * jariJari
	keliling := 2 * 3.14 * jariJari
	fmt.Println(luas)
	fmt.Println(keliling)
	if hitung == "luas" {
		hasil <- luas
	} else {
		hasil <- keliling
	}
}

func HitungJajarGenjang(sisi, alas, tinggi int, hitung string, hasil chan int) {
	luas := alas * tinggi
	keliling := (2 * alas) + (2 * sisi)
	if hitung == "luas" {
		hasil <- luas
	} else {
		hasil <- keliling
	}
}

func BangunDatar(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	hitung := queryValues.Get("hitung")
	channel := make(chan int)
	namabangundatar := ps.ByName("bangundatar")
	switch namabangundatar {
	case "segitiga-sama-sisi":
		alas := queryValues.Get("alas")
		tinggi := queryValues.Get("tinggi")
		alasInt, _ := strconv.Atoi(alas)
		tinggiInt, _ := strconv.Atoi(tinggi)
		go HitungSegitiga(alasInt, tinggiInt, hitung, channel)
		hasil := <-channel
		fmt.Fprint(w, hasil)
	case "persegi":
		sisi := queryValues.Get("sisi")
		sisiInt, _ := strconv.Atoi(sisi)
		go HitungPersegi(sisiInt, hitung, channel)
		hasil := <-channel
		fmt.Fprintln(w, hasil)
	case "persegi-panjang":
		panjang := queryValues.Get("panjang")
		lebar := queryValues.Get("lebar")
		panjangInt, _ := strconv.Atoi(panjang)
		lebarInt, _ := strconv.Atoi(lebar)
		go HitungPersegiPanjang(panjangInt, lebarInt, hitung, channel)
		hasil := <-channel
		fmt.Fprintln(w, hasil)
	case "lingkaran":
		jariJari := queryValues.Get("jariJari")
		jariJariInt, _ := strconv.Atoi(jariJari)
		jariJariFloat := float32(jariJariInt)
		fmt.Println(jariJariInt)
		channelFlt := make(chan float32)
		go HitungLingkaran(jariJariFloat, hitung, channelFlt)
		hasil := <-channelFlt
		fmt.Fprintln(w, hasil)
	case "jajar-genjang":
		sisi := queryValues.Get("sisi")
		alas := queryValues.Get("alas")
		tinggi := queryValues.Get("tinggi")
		sisiInt, _ := strconv.Atoi(sisi)
		alasInt, _ := strconv.Atoi(alas)
		tinggiInt, _ := strconv.Atoi(tinggi)
		go HitungJajarGenjang(sisiInt, alasInt, tinggiInt, hitung, channel)
		hasil := <-channel
		fmt.Fprintln(w, hasil)
	default:
		fmt.Fprintln(w, "Nama Bangun Datar Tidak Ada")
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

// Filter
func GetTotalPage(halaman string) string {
	halamanInt, _ := strconv.Atoi(halaman)
	if halamanInt <= 100 {
		return "tipis"
	} else if halamanInt >= 101 && halamanInt <= 200 {
		return "sedang"
	} else {
		return "tebal"
	}
}

func GetReleaseYear(year int) (string, error) {
	if year <= 1980 && year >= 2021 {
		return "", errors.New("Minimal buku tahun 1980 dan maksimum buku tahun 2021 yang bisa diinput")
	} else {
		return strconv.Itoa(year), nil
	}
}

func GetPrice(harga int) string {
	hargaStr := strconv.Itoa(harga)
	var result string
	counter := 0

	for i := len(hargaStr) - 1; i >= 0; i-- {
		result = string(hargaStr[i]) + result
		counter += 1

		if counter == 3 && i != 0 {
			result = "." + result
			counter = 0
		}
	}
	return "Rp. " + result + ",-"
}
