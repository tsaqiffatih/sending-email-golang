package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Read sensitive data from .env
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	to := os.Getenv("TO_EMAIL")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	hotelName := "Bandung Merdeka Hotel"
	bookingID := "TB9Y25219"
	checkIn := "21 Februari 2023"
	checkOut := "22 Februari 2023"
	alamat := "Jl. Cirebon, Jawa Barat, Indonesia"
	totalHarga := "Rp 72,000"
	roomType := "1 double bedroom"
	guestCount := "2 Tamu"
	stayDetails := "1 Kamar, 1 Malam"

	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<style>
			body { font-family: Arial, sans-serif; background-color: #f8f8f8; margin: 0; padding: 0; }
			.container { max-width: 600px; margin: 20px auto; background: #fff; padding: 20px; border-radius: 8px; border: 1px solid #ddd; }
			.header { text-align: center; font-size: 24px; font-weight: bold; color: #ff2e2e; margin-bottom: 20px; }
			.section { margin-bottom: 20px; }
			.section-title { font-weight: bold; margin-bottom: 10px; }
			.section-content { padding: 10px; background: #f5f5f5; border-radius: 6px; }
			.footer { font-size: 14px; color: #555; text-align: center; margin-top: 20px; }
			.highlight { color: #333; font-weight: bold; }
			.button { background-color: #ff2e2e; color: white; text-decoration: none; padding: 10px 15px; border-radius: 4px; display: inline-block; margin-top: 10px; }
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">Booking Hotels</div>
			<div class="section">
				<p>Dear <span class="highlight">Bapak/Ibu,</span></p>
				<p>Kami senang dapat menyambut Anda kembali. Berikut ini adalah detail reservasi Anda di <span class="highlight">%s</span>.</p>
			</div>
			<div class="section">
				<div class="section-title">Detail Booking:</div>
				<div class="section-content">
					<ul>
						<li><strong>Booking ID:</strong> %s</li>
						<li><strong>Check-in:</strong> %s</li>
						<li><strong>Check-out:</strong> %s</li>
						<li><strong>Alamat:</strong> %s</li>
						<li><strong>Jumlah Tamu:</strong> %s</li>
						<li><strong>Rincian Menginap:</strong> %s</li>
						<li><strong>Tipe Kamar:</strong> %s</li>
					</ul>
				</div>
			</div>
			<div class="section">
				<div class="section-title">Total Pembayaran:</div>
				<div class="section-content">
					<p><strong>%s</strong></p>
					<a href="#" class="button">Bayar Sekarang</a>
				</div>
			</div>
			<div class="section">
				<div class="section-title">Peraturan Hotel:</div>
				<div class="section-content">
					<p>Mohon untuk mengikuti aturan hotel selama masa menginap.</p>
					<ul>
						<li>Check-in mulai pukul 12:00 siang.</li>
						<li>Mohon tunjukkan ID saat check-in.</li>
						<li>Untuk tamu tambahan, harap menghubungi resepsionis.</li>
					</ul>
				</div>
			</div>
			<div class="footer">
				<p>Email ini dikirim secara otomatis. Jangan membalas email ini.</p>
			</div>
		</div>
	</body>
	</html>
	`, hotelName, bookingID, checkIn, checkOut, alamat, guestCount, stayDetails, totalHarga, roomType)

	// Create new email message
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Konfirmasi Pembayaran Berhasil")
	m.SetBody("text/html", body)

	// Connect to SMTP server
	port := 587 // default port
	fmt.Sscan(smtpPort, &port)
	d := gomail.NewDialer(smtpHost, port, from, password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Gagal mengirim email:", err)
		return
	}

	fmt.Println("Email berhasil dikirim!")
}
