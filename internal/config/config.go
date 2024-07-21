package config

// Config holds the application configuration.
type Config struct {
    MongoDBURI  string
    MongoDBName string
    Port        string
}

func LoadFromEnv() *Config {
    return &Config{
        MongoDBURI:  os.Getenv("MONGODB_URI"),
        MongoDBName: os.Getenv("MONGODB_NAME"),
        Port:        os.Getenv("PORT"),
    }
}