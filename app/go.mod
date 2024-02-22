module reservator

go 1.20

require (
	github.com/gofiber/fiber/v2 v2.52.1
	gorm.io/driver/mysql v1.5.0
	gorm.io/gorm v1.25.1
)

require github.com/davecgh/go-spew v1.1.1 // indirect

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/steinfletcher/apitest v1.5.15
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
)

replace reservator/api/model => ./api/model
