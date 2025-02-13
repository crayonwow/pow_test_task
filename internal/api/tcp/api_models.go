package tcp

// reserved 0-10000
const (
	MessageTypeEmpty MessageTypeV1 = iota
	MessageTypeError
)

const (
	MessageTypeGetChallengeRequest MessageTypeV1 = 10001 + iota
	MessageTypeGetChallengeResponse
	MessageTypeGetWisdomRequest
	MessageTypeGetWisdomResponse
)

type APIChallengeResponseV1 struct {
	Data       []byte `json:"data"`
	Difficulty int64  `json:"difficulty"`
}

type APIWisdomResponseV1 struct {
	Text string `json:"text"`
}
