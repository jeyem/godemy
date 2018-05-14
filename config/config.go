package config

var (
	DefaultConfig = Config{
		Debug:     true,
		MongoDB:   "videolearn",
		MongoPort: 27017,
		MongoHost: "127.0.0.1",
		Port:      2000,
	}
)

type Config struct {
	MongoDB   string `json:"mongo_db" valid:"required"`
	MongoHost string `json:"mongo_host" required:"required"`
	MongoPort int    `json:"mongo_port" required:"required"`
	Port      int    `json:"port"`
	Debug     bool   `json:"debug"`
}
