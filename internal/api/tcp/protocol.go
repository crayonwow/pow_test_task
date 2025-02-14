package tcp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
)

const (
	messageV1MinLenBytes = 16
	messageV1MaxLenBytes = messageV1MinLenBytes + dataMaxLenBytes
	dataMaxLenBytes      = 1000000 // 1mb
)

var (
	errEmptyMessage = errors.New("empty")
	errInvalidSize  = errors.New("invalid size")
)

const solutionHeader = "solution"

type (
	data struct {
		Headers map[string]string `json:"headers,omitempty"`
		Payload json.RawMessage   `json:"payload,omitempty"`
	}

	MessageTypeV1 uint32

	MessageV1 struct {
		Type MessageTypeV1
		Data data
	}

	WisdomMessageV1 string
)

func (m *MessageV1) clear() {
	m.Data.Headers = map[string]string{}
	m.Data.Payload = nil
	m.Type = 0
}

func (m *MessageV1) Decode(reader io.Reader) error {
	r := bufio.NewReader(reader)
	slog.Debug("before read")
	rawBytes, err := r.ReadBytes(delim)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return fmt.Errorf("read bytes: %w", err)
	}

	slog.Debug("read",
		slog.String("data", string(rawBytes)),
		slog.String("data_trimmed", string(rawBytes[:len(rawBytes)-2])),

		slog.Uint64("len", uint64(len(rawBytes))),
	)

	err = json.Unmarshal(bytes.TrimSuffix(rawBytes, []byte{delim}), m)
	if err != nil {
		return fmt.Errorf("Unmarshal: %w", err)
	}

	return nil
}

const delim byte = '|'

func (m *MessageV1) Encode() ([]byte, error) {
	result, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("marshal message: %w", err)
	}
	result = append(result, delim)

	return result, nil
}
