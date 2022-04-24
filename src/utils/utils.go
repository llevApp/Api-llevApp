package utils

import (
	"bytes"
	"fmt"

	"github.com/sirupsen/logrus"
)

//LogFormat LogFormat
type LogFormat struct{}

func (logFormat *LogFormat) Format(entry *logrus.Entry) ([]byte, error) {
	var buffer *bytes.Buffer

	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}

	buffer.WriteString("[" + entry.Level.String()[0:4] + "] ")
	buffer.WriteString(entry.Time.Format("2006/01/02 15:04:05.000 "))

	for key, value := range entry.Data {
		buffer.WriteByte('[')
		buffer.WriteString(key)
		buffer.WriteByte(':')
		fmt.Fprint(buffer, value)
		buffer.WriteString("] ")
	}

	buffer.WriteString(entry.Message)
	buffer.WriteByte('\n')

	return buffer.Bytes(), nil
}
