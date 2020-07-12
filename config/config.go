package config

// Configurations exported
type Configurations struct {
	Server      ServerConfigurations
	Database    DatabaseConfigurations
	ExamplePath string
	ExampleVar  string
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
}