package npk

type DatabaseConfig struct {
    Host string
    Port string
    Username string
    Password string
    Catalog string
}

type Configuration struct {
    Database DatabaseConfig
}
