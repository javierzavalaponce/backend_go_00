package core_models

type Config struct {
	ApiPort string
	ApiHost string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbSslMode  string

	Db2Host     string
	Db2Port     string
	Db2User     string
	Db2Password string
	Db2Name     string
	Db2SslMode  string

	LogLevel       string
	LogOutput      string
	SessionSecret  string
	MigrationsPath string
	SeedersPath    string
}
