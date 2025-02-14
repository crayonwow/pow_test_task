package tcp

type Config struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	Timeout int64  `json:"time_out"`
}
