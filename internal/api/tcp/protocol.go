package tcp

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
		Type       MessageTypeV1
		DataLenght uint32 // Lenght of data
		Data       data
	}

	WisdomMessageV1 string
)

func (m *MessageV1) clear() {
	m.Data.Headers = map[string]string{}
	m.Data.Payload = nil
	m.DataLenght = 0
	m.Type = 0
}

func (m *MessageV1) Decode(reader io.Reader) error {
	rawHeaders := make([]byte, messageV1MinLenBytes)
	readBytes := 0
	rb, err := reader.Read(rawHeaders)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
	}
	readBytes += rb
	m.Type = MessageTypeV1(binary.BigEndian.Uint32(rawHeaders[0:4]))
	m.DataLenght = binary.BigEndian.Uint32(rawHeaders[5:])

	if m.DataLenght > dataMaxLenBytes {
		return errInvalidSize
	} else if m.DataLenght == 0 {
		return nil
	}
	rawBody := make([]byte, m.DataLenght)
	for {
		_, err := reader.Read(rawBody)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			break
		}
	}
	d := data{}
	err = json.Unmarshal(rawBody, &d)
	if err != nil {
		return fmt.Errorf("Unmarshal: %w", err)
	}

	m.Data = d
	return nil
}

func (m *MessageV1) Encode() ([]byte, error) {
	b, err := json.Marshal(m.Data)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}
	l := len(b)
	m.DataLenght = uint32(l)
	result := make([]byte, messageV1MinLenBytes+l)
	result = binary.BigEndian.AppendUint32(result, uint32(m.Type))
	result = binary.BigEndian.AppendUint32(result, m.DataLenght)
	result = append(result, b...)
	return result, nil
}
