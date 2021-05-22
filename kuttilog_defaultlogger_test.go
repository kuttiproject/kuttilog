package kuttilog

import (
	"bytes"
	"testing"
)

func TestDefaultlogger(t *testing.T) {
	t.Log("Testing Default Logger...")
	ResetLogger()

	Setloglevel(Info)

	var buf bytes.Buffer
	defaultLogger.log.SetOutput(&buf)

	Print(Info, "Hello")
	if buf.String() != " Hello\n" {
		t.Errorf("\nDefault Logger Print error:\n  Expected: %#v\n  Got: %#v",
			" Hello\n",
			buf.String())
	}

	buf.Reset()
	Println(Info, "Hello")
	if buf.String() != " Hello\n" {
		t.Errorf("\nDefault Logger Print error:\n  Expected: %#v\n  Got: %#v",
			" Hello\n",
			buf.String())
	}

	// Debug output should not happen at Info level
	buf.Reset()
	Printf(Debug, "Hello %v", 42)
	if buf.String() != "" {
		t.Errorf("\nDefault Logger Print error:\n  Expected: \"\"\n  Got: %#v",
			buf.String())
	}

	// Debug output should happen at Debug level
	Setloglevel(Debug)

	buf.Reset()
	Printf(Debug, "Hello %v", 42)
	if buf.String() != "[DEBUG]Hello 42\n" {
		t.Errorf("\nDefault Logger Print error:\n  Expected: %#v\n  Got: %#v",
			"[DEBUG]Hello 42\n",
			buf.String())
	}
}
