package log

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type Level int

// WARNING: some functional logic may assume incremental ordering of levels
const (
	ERROR   Level = iota // 0 - Always output errors (stop execution)
	WARNING              // 1 - Always output (keep executing)
	INFO                 // 2 - General processing information
	TRACE                // 3 - Also, output functional info. (signature, parameter); include UTC timestamps
	DEBUG                // 4 - Also, output internal logic and data (timestamps included)
)

var LevelNames = map[Level]string{
	DEBUG:   "debug",
	ERROR:   "error",
	INFO:    "info",
	TRACE:   "trace",
	WARNING: "warning",
}

var DEFAULT_LEVEL = TRACE

type MyLog struct {
	logLevel      Level
	indentEnabled bool
	indentSpaces  uint
	indentCounter uint
}

func NewLogger() MyLog {
	return MyLog{
		logLevel:      DEFAULT_LEVEL,
		indentEnabled: false,
		indentSpaces:  2,
		indentCounter: 0,
	}
}

func (log *MyLog) SetLevel(level Level) {
	log.logLevel = level
}

func (log *MyLog) SetIndentSpaces(spaces uint) {
	// Put some sensible limit on spaces
	if spaces > 8 {
		log.indentSpaces = spaces
	}
}

func (log MyLog) Trace(tag string, value interface{}) {
	// Skip 2 on call stack (i.e., this method and the method we are calling)
	log.dumpInterface(TRACE, tag, value, 2)
}

func (log MyLog) Debug(tag string, value interface{}) {
	// Skip 2 on call stack (i.e., this method and the method we are calling)
	log.dumpInterface(DEBUG, tag, value, 2)
}

func (log MyLog) Info(tag string, value interface{}) {
	// Skip 2 on call stack (i.e., this method and the method we are calling)
	log.dumpInterface(INFO, tag, value, 2)
}

func (log MyLog) Warning(tag string, value interface{}) {
	// Skip 2 on call stack (i.e., this method and the method we are calling)
	log.dumpInterface(WARNING, tag, value, 2)
}

func (log MyLog) Error(tag string, value interface{}) {
	// Skip 2 on call stack (i.e., this method and the method we are calling)
	log.dumpInterface(ERROR, tag, value, 2)
}

func (log MyLog) Enter() {
	// Skip 2 on call stack (i.e., this method and the method we are calling)
	log.dumpInterface(INFO, "ENTER", nil, 2)
}
func (log MyLog) Exit() {
	// Skip 2 on call stack (i.e., this method and the method we are calling)
	log.dumpInterface(INFO, "EXIT", nil, 2)
}

func (log MyLog) dumpInterface(lvl Level, tag string, value interface{}, skip int) {

	if log.logLevel == DEBUG {
		fmt.Printf("dumpInterface(): %s\n", DumpStruct("logger", log))
	}

	// retrieve all the info we might need
	pc, fn, line, ok := runtime.Caller(skip)

	// TODO: create a logging package that can indent based upon stack size
	// Note: the "Callers()" method will not append() so allocate a large array
	// var mystack []uintptr = make([]uintptr, 10)
	// stacksize := runtime.Callers(0, mystack)
	//fmt.Printf("stacksize=%v\n", stacksize)

	// TODO: Provide means to order component output;
	// for example, to add Timestamp component first (on each line) before Level
	if ok {
		// Setup "string builder" and initialize with log-level prefix
		sb := bytes.NewBufferString(fmt.Sprintf("[%s] ", LevelNames[lvl]))

		// Append UTC timestamp if TRACE (or DEBUG) enabled
		if lvl == TRACE || lvl == DEBUG {
			// UTC time shows fractions of a second
			// TODO: add setting to show milli or micro seconds supported by "time" package
			tmp := time.Now().UTC().String()
			// create a (left) slice of the timestamp omitting the " +0000 UTC" portion
			//ts = fmt.Sprintf("[%s] ", tmp[:strings.Index(tmp, "+")-1])
			sb.WriteString(fmt.Sprintf("[%s] ", tmp[:strings.Index(tmp, "+")-1]))
		}

		// Append basic filename, line number, function name
		basicFile := fn[strings.LastIndex(fn, "/")+1:]
		function := runtime.FuncForPC(pc)
		// TODO: add logger flag to show full module paths (not just module.function)
		basicModFnName := function.Name()[strings.LastIndex(function.Name(), "/")+1:]

		sb.WriteString(fmt.Sprintf("%s(%d) %s()", basicFile, line, basicModFnName))

		// Append (optional) tag
		if tag != "" {
			sb.WriteString(fmt.Sprintf(": %s", tag))
		}
		// Append (optional) value
		if value != nil {
			sb.WriteString(fmt.Sprintf(": %+v", value))
		}
		// TODO: use a general output writer (set to stdout, stderr, or filestream)
		fmt.Println(sb.String())
	} else {
		os.Stderr.WriteString("Error: Unable to retrieve call stack. Exiting...")
		os.Exit(-2)
	}
}

func DumpStruct(structName string, field interface{}) error {

	formattedStruct, err := FormatStruct(structName, field)
	if err != nil {
		return err
	}
	fmt.Print(formattedStruct)
	return nil
}

func DumpArgs() {
	args := os.Args
	for i, a := range args {
		fmt.Printf("os.Arg[%d]: `%v`\n", i, a)
	}
}

func DumpSeparator(sep byte, repeat int) error {
	if repeat <= 80 {
		sb := bytes.NewBufferString("")
		for i := 0; i < repeat; i++ {
			sb.WriteByte(sep)
		}
		fmt.Println(sb.String())
		return nil
	} else {
		return errors.New("invalid repeat length (>80)")
	}
}
