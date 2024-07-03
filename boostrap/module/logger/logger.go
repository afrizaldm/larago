package logger

import (
	"fmt"
	"os"
	"simple-api/config"
	"time"
)

var logRoot string = "logs"

func EnsureLogsFolder(folderName string) error {
	// Cek apakah folder sudah ada
	_, err := os.Stat(folderName)
	if os.IsNotExist(err) {
		// Jika folder tidak ada, buat folder baru
		err := os.Mkdir(folderName, 0755) // 0755 adalah mode permission untuk folder baru
		if err != nil {
			return err
		}
		fmt.Printf("Folder '%s' successfully created.\n", folderName)
	} else if err != nil {
		return err
	}

	return nil
}

func Write(description string) error {

	appConfig := config.NewAppConfig().Load()

	if !appConfig.APP_ACTIVE_LOGGING {
		return nil
	}

	EnsureLogsFolder(logRoot)

	// Dapatkan tanggal hari ini
	currentTime := time.Now()

	// Nama sub folder dengan kombinasi tahun_bulan
	logTime := currentTime.Format("2006-01-02 15:04:05.000")

	// Format untuk waktu: "2006-01-02 15:04:05.000"
	folder_year_month := currentTime.Format("2006-01")
	EnsureLogsFolder(fmt.Sprintf("logs/%s", folder_year_month))

	// Format baris log: "tanggal jam detik dan mili detik type description"
	logEntry := fmt.Sprintf("%s: %s\n", logTime, description)

	// Nama file dengan format DD.txt (misal: 01.txt)
	name := currentTime.Format("02") // Format tanggal sebagai "01" hingga "31"
	fileName := fmt.Sprintf("logs/%s/%s.txt", folder_year_month, name)

	// Coba buka atau buat file baru
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Menulis log ke dalam file
	if _, err := file.WriteString(logEntry); err != nil {
		return err
	}

	return nil
}
