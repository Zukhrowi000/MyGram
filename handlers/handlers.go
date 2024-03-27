package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct untuk menangani permintaan unggah foto
type UploadHandler struct{}

// Metode untuk menangani permintaan POST untuk mengunggah foto
func (h UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Simulasi pengolahan unggahan foto
	// Di sini Anda akan menambahkan logika untuk menyimpan foto ke penyimpanan, mengambil metadata, dll.
	// Untuk contoh ini, kita hanya mencetak pesan ke konsol.
	fmt.Println("Photo uploaded successfully")

	// Kirim respons OK
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Photo uploaded successfully"))
}

// Struct untuk menangani permintaan untuk menambahkan filter pada foto
type FilterHandler struct{}

// Metode untuk menangani permintaan POST untuk menambahkan filter pada foto
func (h FilterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Simulasi pengolahan penambahan filter pada foto
	// Di sini Anda akan menambahkan logika untuk menambahkan filter ke foto, dll.
	// Untuk contoh ini, kita hanya mencetak pesan ke konsol.
	fmt.Println("Filter added to photo")

	// Kirim respons OK
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Filter added to photo"))
}

// Handler untuk rute utama
func mainHandler(w http.ResponseWriter, r *http.Request) {
	// Menangani permintaan ke rute utama, misalnya, menampilkan halaman utama aplikasi.
	// Di sini Anda dapat menambahkan logika untuk menampilkan halaman HTML atau respons JSON.
	// Untuk contoh ini, kita hanya memberikan respons sederhana.
	response := map[string]string{"message": "Welcome to MyGram!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Membuat instance dari setiap handler
	uploadHandler := UploadHandler{}
	filterHandler := FilterHandler{}

	// Menghubungkan handler dengan rute yang sesuai
	http.Handle("/upload", uploadHandler)
	http.Handle("/filter", filterHandler)
	http.HandleFunc("/", mainHandler)

	// Memulai server HTTP di port 8080
	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
