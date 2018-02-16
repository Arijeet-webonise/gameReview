package app

// AppConfig will hold the global configuration parameters for the app to access.
// It will be set only during the application startup or test startup
type Config struct {
	Port string
}
