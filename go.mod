module github.com/KlareTeam/interview-challenges/go/pricematic

go 1.16

require (
	github.com/andybalholm/brotli v1.0.3 // indirect
	github.com/gofiber/fiber/v2 v2.18.0
	github.com/jackc/pgx/v4 v4.13.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/klauspost/compress v1.13.5 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/sys v0.0.0-20210831042530-f4d43177bf5e // indirect
	golang.org/x/text v0.3.7 // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.14
)

replace github.com/KlareTeam/interview-challenges/go/pricematic/products => ../products

replace github.com/KlareTeam/interview-challenges/go/pricematic/database => ../database
