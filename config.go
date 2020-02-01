package npk

type DatabaseConfig struct {
    Host string
    Port string
    SSLMode string
    Username string
    Password string
    Catalog string
}

type Configuration struct {
    Database DatabaseConfig
    SessionKey string
    Listen string
}
