package conf

type AppSettings struct {
	Server struct {
		Http struct {
			Addr    string `json:"addr"`
			Timeout int    `json:"timeout"`
		} `json:"http"`
		Grpc struct {
			Addr    string `json:"addr"`
			Timeout int    `json:"timeout"`
		} `json:"grpc"`
	} `json:"server"`
	Mongo struct {
		Server struct {
			ConnectionString string `json:"ConnectionString"`
			Database         string `json:"Database"`
		} `json:"server"`
	} `json:"mongo"`
	Redis struct {
		Core string `json:"core"`
	} `json:"redis"`
}
