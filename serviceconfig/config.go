package serviceconfig

type Server struct {
	Name    string `json:"name"`
	Version string `json:"-"`
	Addr    string `json:"addr" default:":7070"`
}

type Config struct {
	Server *Server `json:"server"`
}

func NewConfig(name, version string) *Config {
	return &Config{
		Server: &Server{
			Name:    name,
			Version: version,
		},
	}
}
