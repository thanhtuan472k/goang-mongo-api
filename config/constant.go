package config

type Database struct {
	URI  string
	Name string
}

type Jwt struct {
	SecretKey string
}

// ENV .env struct
type ENV struct {
	// App port
	AppPort string

	// Database
	Database Database

	// Jwt
	Jwt Jwt
}
