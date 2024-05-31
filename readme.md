# CLI untuk Aplikasi Go

CLI ini menyediakan berbagai perintah yang memudahkan dalam pengembangan dan pengelolaan aplikasi Go. Berikut adalah daftar perintah yang tersedia:

## Perintah yang Tersedia

1. **Build Command**: 
   - **Perintah**:
     - Membersihkan folder build jika ada.
     - Membuat folder build.
     - Menyalin folder `public` ke dalam folder `build/public`.
     - Menyalin folder `resources/views` ke dalam folder `build/views`.
     - Membuat executable dengan perintah `go build -o build/main.exe ./cmd/main.go`.
   - **Penggunaan**: `go run main.go build`

2. **Serve Command**:
   - **Perintah**:
     - Menjalankan aplikasi dalam mode debug dengan perintah `go run ./cmd/main.go --debug`.
   - **Penggunaan**: `go run main.go serve`

3. **Clean Command**:
   - **Perintah**:
     - Menghapus folder `build` jika ada.
   - **Penggunaan**: `go run main.go clean`

<!-- 4. **Migrate Command**:
   - **Perintah**:
     - Menjalankan migrasi basis data.
   - **Penggunaan**: `artisan migrate`

5. **Rollback Command**:
   - **Perintah**:
     - Melakukan rollback migrasi basis data.
   - **Penggunaan**: `artisan rollback`

6. **Seed Command**:
   - **Perintah**:
     - Melakukan seeding basis data.
   - **Penggunaan**: `artisan seed`

7. **Test Command**:
   - **Perintah**:
     - Menjalankan tes unit untuk aplikasi dengan perintah `go test ./...`.
   - **Penggunaan**: `artisan test` -->

Pastikan untuk menyesuaikan perintah-perintah ini dengan kebutuhan aplikasi Anda. Juga, pastikan untuk menyesuaikan perintah migrasi, rollback, dan seeding dengan sistem manajemen basis data yang Anda gunakan.
