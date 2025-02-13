package protection

type Config struct {
	Limit             int64 `json:"limit"` // per seccond
	MaxDifficulty     int64 `json:"max_difficulty"`
	DefaultDifficulty int64 `json:"default_difficulty"`
}
