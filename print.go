package main

import (
	"fmt"
	"io"
	"strings"

	palgate "github.com/zcahana/palgate-sdk"
)

var columns = []struct {
	displayName   string
	width         int
	valueProvider func(record *palgate.LogRecord) string
}{
	{
		displayName: "DATE",
		width:       10,
		valueProvider: func(record *palgate.LogRecord) string {
			return record.Date()
		},
	},
	{
		displayName: "TIME",
		width:       8,
		valueProvider: func(record *palgate.LogRecord) string {
			return record.Time()
		},
	},
	{
		displayName: "OPERATION TYPE",
		width:       16,
		valueProvider: func(record *palgate.LogRecord) string {
			return record.Type.String()
		},
	},
	{
		displayName: "STATUS",
		width:       10,
		valueProvider: func(record *palgate.LogRecord) string {
			return record.OperationStatus.String()
		},
	},
	{
		displayName: "SERIAL",
		width:       12,
		valueProvider: func(record *palgate.LogRecord) string {
			return record.SerialNumber
		},
	},
	{
		displayName: "USER ID",
		width:       14,
		valueProvider: func(record *palgate.LogRecord) string {
			return record.UserID
		},
	},
	{
		displayName: "NAME",
		width:       20,
		valueProvider: func(record *palgate.LogRecord) string {
			return record.Name()
		},
	},
}

func Print(records []palgate.LogRecord, writer io.Writer) error {
	format := buildFormatString()
	headerArgs := buildHeaderArgs()

	_, err := fmt.Fprintf(writer, format, headerArgs...)
	if err != nil {
		return fmt.Errorf("error writing header column: %v", err)
	}

	for _, record := range records {
		recordArgs := buildRecordArgs(&record)
		_, err := fmt.Fprintf(writer, format, recordArgs...)
		if err != nil {
			return fmt.Errorf("error writing record column: %v", err)
		}
	}

	return nil
}

func buildFormatString() string {
	var builder strings.Builder

	for _, column := range columns {
		builder.WriteString(fmt.Sprintf("%%-%ds\t", column.width))
	}
	builder.WriteRune('\n')

	return builder.String()
}

func buildHeaderArgs() []interface{} {
	args := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		args = append(args, column.displayName)
	}

	return args
}

func buildRecordArgs(record *palgate.LogRecord) []interface{} {
	args := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		args = append(args, column.valueProvider(record))
	}

	return args
}
