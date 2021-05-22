package kuttilog_test

import (
	"fmt"
	"testing"

	"github.com/kuttiproject/kuttilog"
)

func TestSetLogLevel(t *testing.T) {
	defaultloglevel := kuttilog.Loglevel()
	maxloglevel := kuttilog.MaxLevel()
	t.Logf("\nDefault log level is: %v\nMax log level is: %v", defaultloglevel, maxloglevel)

	t.Logf("Trying to change to negative log level")
	kuttilog.Setloglevel(-1)
	if kuttilog.Loglevel() != defaultloglevel {
		t.Error("SetLogLevel allowed setting negative level")
		t.FailNow()
	}

	t.Logf("Trying to change to greater than max log level")
	kuttilog.Setloglevel(maxloglevel + 1)
	if kuttilog.Loglevel() != defaultloglevel {
		t.Error("SetLogLevel allowed setting level greater than max.")
		t.FailNow()
	}

}

func TestV(t *testing.T) {
	t.Logf("Testing V()")

	testvtable := []struct {
		setlevel int
		level    int
		result   bool
	}{
		{setlevel: 2, level: 0, result: true},
		{setlevel: 2, level: 1, result: true},
		{setlevel: 2, level: 2, result: true},
		{setlevel: 2, level: 3, result: false},
		{setlevel: 2, level: 4, result: false},
		{setlevel: 2, level: 5, result: false},
		{setlevel: 2, level: -1, result: false},
		{setlevel: 1, level: 0, result: true},
		{setlevel: 1, level: 1, result: true},
		{setlevel: 1, level: 2, result: false},
		{setlevel: 1, level: 3, result: false},
		{setlevel: 1, level: 4, result: false},
		{setlevel: 1, level: 5, result: false},
		{setlevel: 1, level: -1, result: false},
	}

	for _, testrow := range testvtable {
		kuttilog.Setloglevel(testrow.setlevel)
		receivedresult := kuttilog.V(testrow.level)
		if receivedresult != testrow.result {
			t.Errorf(
				"Set Level: %v, level value: %v, expected result: %v, received result: %v",
				testrow.setlevel,
				testrow.level,
				testrow.result,
				receivedresult,
			)
			t.Fail()
		}
	}

}

// Mock logger for testing
var testloggerlevelprefixes = []string{
	"",
	"",
	"",
	"[V]",
	"[D]",
}

type testlogger struct {
	logstring string
}

func (d *testlogger) MaxLevel() int {
	return kuttilog.Debug
}

func (d *testlogger) LevelPrefix(level int) string {
	return testloggerlevelprefixes[level]
}

func (d *testlogger) Print(v ...interface{}) {
	d.logstring = fmt.Sprint(v...)
}

func (d *testlogger) Printf(format string, v ...interface{}) {
	d.logstring = fmt.Sprintf(format, v...)
}

func (d *testlogger) Println(v ...interface{}) {
	d.logstring = fmt.Sprintln(v...)
}

// End Mock logger

func TestPrint(t *testing.T) {
	t.Log("Testing Print()")

	testprinttable := []struct {
		setlevel   int
		level      int
		inputtext  string
		outputtext string
	}{
		{setlevel: 2, level: 0, inputtext: "Hello", outputtext: "Hello"},
		{setlevel: 2, level: 1, inputtext: "Hello", outputtext: "Hello"},
		{setlevel: 2, level: 2, inputtext: "Hello", outputtext: "Hello"},
		{setlevel: 2, level: 3, inputtext: "Hello", outputtext: ""},
		{setlevel: 2, level: 4, inputtext: "Hello", outputtext: ""},
		{setlevel: 2, level: 5, inputtext: "Hello", outputtext: ""},
		{setlevel: 2, level: -1, inputtext: "Hello", outputtext: ""},
		{setlevel: 0, level: 0, inputtext: "Hello", outputtext: "Hello"},
		{setlevel: 0, level: 1, inputtext: "Hello", outputtext: ""},
		{setlevel: 0, level: 2, inputtext: "Hello", outputtext: ""},
		{setlevel: 0, level: 3, inputtext: "Hello", outputtext: ""},
		{setlevel: 0, level: 4, inputtext: "Hello", outputtext: ""},
		{setlevel: 0, level: 5, inputtext: "Hello", outputtext: ""},
		{setlevel: 0, level: -1, inputtext: "Hello", outputtext: ""},
		{setlevel: 4, level: 0, inputtext: "Hello", outputtext: "Hello"},
		{setlevel: 4, level: 1, inputtext: "Hello", outputtext: "Hello"},
		{setlevel: 4, level: 2, inputtext: "Hello", outputtext: "Hello"},
		{setlevel: 4, level: 3, inputtext: "Hello", outputtext: "[V]Hello"},
		{setlevel: 4, level: 4, inputtext: "Hello", outputtext: "[D]Hello"},
		{setlevel: 4, level: 5, inputtext: "Hello", outputtext: ""},
		{setlevel: 4, level: -1, inputtext: "Hello", outputtext: ""},
	}

	tl := &testlogger{}
	kuttilog.SetLogger(tl)
	for _, testrow := range testprinttable {
		tl.logstring = ""
		kuttilog.Setloglevel(testrow.setlevel)
		kuttilog.Print(testrow.level, testrow.inputtext)
		if tl.logstring != testrow.outputtext {
			t.Errorf(
				"Set Level: %v, Level: %v, Input: %v, Expected output: %v, Received output: %v",
				testrow.setlevel,
				testrow.level,
				testrow.inputtext,
				testrow.outputtext,
				tl.logstring,
			)
			t.Fail()
		}
	}
	kuttilog.ResetLogger()
}

func TestPrintf(t *testing.T) {
	t.Log("Testing Printf()")

	testprintftable := []struct {
		setlevel   int
		level      int
		inputtext  string
		parameter  int
		outputtext string
	}{
		{setlevel: 2, level: 0, inputtext: "Hello %v\n", parameter: 42, outputtext: "Hello 42\n"},
		{setlevel: 2, level: 1, inputtext: "Hello %v\n", parameter: 42, outputtext: "Hello 42\n"},
		{setlevel: 2, level: 2, inputtext: "Hello %v\n", parameter: 42, outputtext: "Hello 42\n"},
		{setlevel: 2, level: 3, inputtext: "Hello %v\n", parameter: 42, outputtext: ""},
		{setlevel: 2, level: 4, inputtext: "Hello %v\n", parameter: 42, outputtext: ""},
		{setlevel: 2, level: 5, inputtext: "Hello %v\n", parameter: 42, outputtext: ""},
		{setlevel: 2, level: -1, inputtext: "Hello %v\n", parameter: 42, outputtext: ""},
		{setlevel: 0, level: 0, inputtext: "Hello %v\n", parameter: 42, outputtext: "Hello 42\n"},
		{setlevel: 0, level: 1, inputtext: "Hello %v\n", parameter: 42, outputtext: ""},
		{setlevel: 0, level: 2, inputtext: "Hello %v\n", parameter: 42, outputtext: ""},
		{setlevel: 0, level: 3, inputtext: "Hello %v\n", parameter: 42, outputtext: ""},
		{setlevel: 0, level: 4, inputtext: "Hello %v\n", parameter: 42, outputtext: ""},
		{setlevel: 0, level: 5, inputtext: "Hello %v\n", parameter: 42, outputtext: ""},
		{setlevel: 0, level: -1, inputtext: "Hello %v\n", parameter: 42, outputtext: ""},
		{setlevel: 4, level: 0, inputtext: "Hello %v\n", parameter: 42, outputtext: "Hello 42\n"},
		{setlevel: 4, level: 1, inputtext: "Hello %v\n", parameter: 42, outputtext: "Hello 42\n"},
		{setlevel: 4, level: 2, inputtext: "Hello %v\n", parameter: 42, outputtext: "Hello 42\n"},
		{setlevel: 4, level: 3, inputtext: "Hello %v\n", parameter: 42, outputtext: "[V]Hello 42\n"},
		{setlevel: 4, level: 4, inputtext: "Hello %v\n", parameter: 42, outputtext: "[D]Hello 42\n"},
		{setlevel: 4, level: 5, inputtext: "Hello %v\n", parameter: 42, outputtext: ""},
		{setlevel: 4, level: -1, inputtext: "Hello %v\n", parameter: 42, outputtext: ""},
	}

	tl := &testlogger{}
	kuttilog.SetLogger(tl)
	for _, testrow := range testprintftable {
		tl.logstring = ""
		kuttilog.Setloglevel(testrow.setlevel)
		kuttilog.Printf(testrow.level, testrow.inputtext, testrow.parameter)
		if tl.logstring != testrow.outputtext {
			t.Errorf(
				"\n  Set Level: %v\n  Level: %v\n  Input: %v\n  Expected output: %#v\n  Received output: %#vEND",
				testrow.setlevel,
				testrow.level,
				testrow.inputtext,
				testrow.outputtext,
				tl.logstring,
			)
			t.Fail()
		}
	}
	kuttilog.ResetLogger()
}

func TestPrintln(t *testing.T) {
	t.Log("Testing Println()")

	// Println will ALWAYS add a space between the label prefix and the
	// value printed
	testprintlntable := []struct {
		setlevel   int
		level      int
		inputtext  string
		outputtext string
	}{
		{setlevel: 2, level: 0, inputtext: "Hello", outputtext: " Hello\n"},
		{setlevel: 2, level: 1, inputtext: "Hello", outputtext: " Hello\n"},
		{setlevel: 2, level: 2, inputtext: "Hello", outputtext: " Hello\n"},
		{setlevel: 2, level: 3, inputtext: "Hello", outputtext: ""},
		{setlevel: 2, level: 4, inputtext: "Hello", outputtext: ""},
		{setlevel: 2, level: 5, inputtext: "Hello", outputtext: ""},
		{setlevel: 2, level: -1, inputtext: "Hello", outputtext: ""},
		{setlevel: 0, level: 0, inputtext: "Hello", outputtext: " Hello\n"},
		{setlevel: 0, level: 1, inputtext: "Hello", outputtext: ""},
		{setlevel: 0, level: 2, inputtext: "Hello", outputtext: ""},
		{setlevel: 0, level: 3, inputtext: "Hello", outputtext: ""},
		{setlevel: 0, level: 4, inputtext: "Hello", outputtext: ""},
		{setlevel: 0, level: 5, inputtext: "Hello", outputtext: ""},
		{setlevel: 0, level: -1, inputtext: "Hello", outputtext: ""},
		{setlevel: 4, level: 0, inputtext: "Hello", outputtext: " Hello\n"},
		{setlevel: 4, level: 1, inputtext: "Hello", outputtext: " Hello\n"},
		{setlevel: 4, level: 2, inputtext: "Hello", outputtext: " Hello\n"},
		{setlevel: 4, level: 3, inputtext: "Hello", outputtext: "[V] Hello\n"},
		{setlevel: 4, level: 4, inputtext: "Hello", outputtext: "[D] Hello\n"},
		{setlevel: 4, level: 5, inputtext: "Hello", outputtext: ""},
		{setlevel: 4, level: -1, inputtext: "Hello", outputtext: ""},
	}

	tl := &testlogger{}
	kuttilog.SetLogger(tl)
	for _, testrow := range testprintlntable {
		tl.logstring = ""
		kuttilog.Setloglevel(testrow.setlevel)
		kuttilog.Println(testrow.level, testrow.inputtext)
		if tl.logstring != testrow.outputtext {
			t.Errorf(
				"\n  Set Level: %v\n  Level: %v\n  Input: %v\n  Expected output: %#v\n  Received output: %#vEND",
				testrow.setlevel,
				testrow.level,
				testrow.inputtext,
				testrow.outputtext,
				tl.logstring,
			)
			t.Fail()
		}
	}
	kuttilog.ResetLogger()
}
