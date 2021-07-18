package infrastrcture

type Config struct {
	DB struct {
		Production struct {
			Host     string
			UserName string
			Password string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {
	c := new(Config)

	c.DB.Production.Host = "localhost"
	c.DB.Production.UserName = "root"
	c.DB.Production.Password = "pass"
	c.DB.Production.DBName = "weapon_roulette"

	c.Routing.Port = "8080"

	return c
}
