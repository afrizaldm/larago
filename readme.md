# Laravel-Like Go Framework

Framework ini merupakan hasil pengembangan saya sendiri, dirancang untuk memudahkan transisi bagi developer yang terbiasa dengan Laravel namun ingin memanfaatkan performa unggul dari bahasa pemrograman Go. Berdasarkan pengalaman dan pemahaman mendalam saya dalam kedua ekosistem ini, framework ini dibangun dengan prinsip kesederhanaan namun tetap mampu memenuhi kebutuhan aplikasi web berskala besar.

Framework web berbasis Go ini terinspirasi dari Laravel, menyediakan fitur-fitur esensial seperti routing, middleware, ORM, autentikasi, dan migrasi database. Tujuannya adalah mempermudah pengembangan aplikasi web melalui sintaks sederhana yang tetap optimal dalam performa, menjadikannya ideal bagi developer Laravel yang ingin beralih ke ekosistem Go.


> **⚠️ Catatan:**
> Proyek ini masih bersifat eksperimental dan dalam tahap pengembangan aktif. Beberapa fitur mungkin berubah atau belum lengkap. Gunakan dengan risiko Anda sendiri.



## Fitur
- Routing yang mudah
- Middleware yang dapat dikustomisasi
- ORM terintegrasi untuk manajemen database
- Sistem autentikasi bawaan
- Dukungan migrasi database
- Arsitektur MVC (Model-View-Controller)

## Instalasi

Untuk menginstal dan mengatur framework Go ini, ikuti langkah-langkah berikut:

### Prasyarat
- Go versi 1.18 atau lebih tinggi harus sudah terinstal. Anda dapat mengunduh dan menginstalnya dari [Situs Resmi Go](https://golang.org/dl/).

### Langkah-Langkah Instalasi

1. **Kloning repository**:
   ```bash
   git clone https://github.com/afrizaldm/larago.git
   cd framework-anda
    ```

# Perintah CLI

CLI ini menyediakan berbagai perintah untuk membantu dalam membangun, menjalankan, dan mengelola aplikasi Go. Berikut adalah daftar perintah yang tersedia:

1. **Perintah Build**:
   - **Instruksi**:
     - Membersihkan folder `build` jika ada.
     - Membuat folder `build`.
     - Menyalin folder `public` ke dalam `build/public`.
     - Menyalin folder `resources/views` ke dalam `build/views`.
     - Membuat executable dengan perintah `go build -o build/main.exe ./cmd/main.go`.
   - **Penggunaan**:
     ```bash
     > goat build
     ```
   
   **Aplikasi hasil build akan tersimpan di dalam folder `build`.**

2. **Perintah Serve**:
   - **Instruksi**:
     - Menjalankan aplikasi dalam mode debug menggunakan perintah `go run ./cmd/main.go --debug`.
   - **Penggunaan**:
     ```bash
     > goat serve
     ```

3. **Perintah Preview**:
   - **Instruksi**:
     - Menjalankan server untuk preview aplikasi.
   - **Penggunaan**:
     ```bash
     > goat preview
     ```

4. **Perintah Clean**:
   - **Instruksi**:
     - Menghapus folder `build` jika ada.
   - **Penggunaan**:
     ```bash
     > goat build:clean
     ```

Perintah-perintah ini dirancang untuk memperlancar proses pengembangan dan meningkatkan produktivitas saat mengelola aplikasi Go Anda.

---

# Laravel-Like Go Framework

This framework is my own development, designed to ease the transition for developers familiar with Laravel who want to take advantage of Go’s superior performance. With my experience and deep understanding of both ecosystems, this framework is built to remain simple while effectively meeting the demands of large-scale web applications.

This Go-based web framework is inspired by Laravel, providing essential features such as routing, middleware, ORM, authentication, and database migrations. It aims to simplify web application development with a straightforward syntax and optimal performance, making it an ideal choice for Laravel developers looking to transition to Go.

> **⚠️ Note:**
> This project is still experimental and under active development. Some features may change or be incomplete. Use it at your own risk.


## Features
- Easy routing
- Customizable middleware
- Integrated ORM for database management
- Built-in authentication system
- Database migration support
- MVC (Model-View-Controller) architecture

## Installation

To install and set up this Go framework, follow these steps:

### Prerequisites
- Go 1.18 or higher must be installed. You can download and install it from [Go Official Website](https://golang.org/dl/).

### Step-by-Step Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/afrizaldm/larago.git
   cd your-framework
    ```

# CLI Commands

This CLI provides various commands to help in building, running, and managing the Go application. Below are the available commands:

1. **Build Command**:
   - **Instructions**:
     - Cleans the `build` folder if it exists.
     - Creates a `build` folder.
     - Copies the `public` folder into `build/public`.
     - Copies the `resources/views` folder into `build/views`.
     - Compiles the executable using the command `go build -o build/main.exe ./cmd/main.go`.
   - **Usage**:
     ```bash
     > goat build
     ```
   
   **The built application will be saved in the `build` folder.**

2. **Serve Command**:
   - **Instructions**:
     - Runs the application in debug mode using the command `go run ./cmd/main.go --debug`.
   - **Usage**:
     ```bash
     > goat serve
     ```

3. **Preview Command**:
   - **Instructions**:
     - Starts the server for previewing the application.
   - **Usage**:
     ```bash
     > goat preview
     ```

4. **Clean Command**:
   - **Instructions**:
     - Deletes the `build` folder if it exists.
   - **Usage**:
     ```bash
     > goat build:clean
     ```

These commands are designed to streamline the development process and improve productivity when managing your Go application.

---

<!-- 4. **Migrate Command**:
   - **Perintah**:
     - Menjalankan migrasi basis data.
   - **Penggunaan**: `artisan migrate`

1. **Rollback Command**:
   - **Perintah**:
     - Melakukan rollback migrasi basis data.
   - **Penggunaan**: `artisan rollback`

2. **Seed Command**:
   - **Perintah**:
     - Melakukan seeding basis data.
   - **Penggunaan**: `artisan seed`

3. **Test Command**:
   - **Perintah**:
     - Menjalankan tes unit untuk aplikasi dengan perintah `go test ./...`.
   - **Penggunaan**: `artisan test` -->

<!-- Pastikan untuk menyesuaikan perintah-perintah ini dengan kebutuhan aplikasi Anda. Juga, pastikan untuk menyesuaikan perintah migrasi, rollback, dan seeding dengan sistem manajemen basis data yang Anda gunakan. -->
