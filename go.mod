module api-testing

go 1.23

toolchain go1.23.10

require (
	github.com/DATA-DOG/go-sqlmock v1.5.2
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.30.0
)

require (
	github.com/go-chi/chi/v5 v5.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
)

require (
	custom-api-server v0.0.0
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)

replace custom-api-server => ../../DAY2/Keploy-Session-2/backend
