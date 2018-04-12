package domain

type Config struct {
	Port        int    `json:"port"`
	JwtSecret   string `json:"jwt_secret"`
	JwtLifetime int    `json:"jwt_lifetime_sec"`
}
