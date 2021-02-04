package serviceconfig

type Server struct {
	Name    string `json:"name"`
	Version string `json:"-"`
	Addr    string `json:"addr" default:":7070"`
}

type XMLConfig struct {
	Url          string `json:"url"`
	Xmlns        string `json:"xmlns"`
	XmlVersion   string `json:"version"`
	XmlDirection string `json:"direction"`
	XmlMsgType   string `json:"msg_type"`
}

type Config struct {
	Server    *Server    `json:"server"`
	XmlConfig *XMLConfig `json:"xmlConfig"`
}

func NewConfig(name, version string) *Config {
	return &Config{
		Server: &Server{
			Name:    name,
			Version: version,
		},
	}
}
