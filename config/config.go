package config

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

type Config struct {
	Environment string // debug, test, release

	ServerHost string
	ServerPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string

	DefaultOffset int
	DefaultLimit  int
	PostgresMaxConnections int32

	SecretKey	string
}

func Load() Config {
	cfg := Config{}

	cfg.ServerHost = "localhost"
	cfg.ServerPort = ":8001"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "isanov"
	cfg.PostgresDatabase = "salesdb"
	cfg.PostgresPassword = "i1saanoff"
	cfg.PostgresPort = "5432"
	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10
	cfg.PostgresMaxConnections = 20
	cfg.SecretKey = "1234"

	return cfg
}
