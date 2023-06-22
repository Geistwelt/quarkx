/**
* Author: Xiangyu Wu
* Date: 2023-06-20
* From: hyperledger/fabric/common/flogging/fabenc/formatter.go + hyperledger/fabric/common/flogging/fabenc/encoder.go
 */

package enc

import (
	"fmt"
	"io"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	zaplogfmt "github.com/sykesm/zap-logfmt"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

var formatRegexp = regexp.MustCompile(`%{(color|sequence|level|message|module|shortfunc|time)(?::(.*?))?}`)

type FormatEncoder struct {
	zapcore.Encoder
	formatters []Formatter
	pool       buffer.Pool
}

type Formatter interface {
	Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field)
}

func NewFormatEncoder(formatters ...Formatter) *FormatEncoder {
	return &FormatEncoder{
		Encoder: zaplogfmt.NewEncoder(zapcore.EncoderConfig{
			MessageKey:     "",
			LevelKey:       "",
			TimeKey:        "",
			NameKey:        "",
			CallerKey:      "",
			StacktraceKey:  "",
			LineEnding:     "\n",
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeTime: func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
				pae.AppendString(t.Format("2006-01-02T15:04:05.999Z07:00"))
			},
		}),
		formatters: formatters,
		pool:       buffer.NewPool(),
	}
}

func (f *FormatEncoder) Clone() zapcore.Encoder {
	return &FormatEncoder{
		Encoder:    f.Encoder.Clone(),
		formatters: f.formatters,
		pool:       f.pool,
	}
}

func (f *FormatEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	line := f.pool.Get()
	for _, formatter := range f.formatters {
		formatter.Format(line, entry, fields)
	}

	encodedFields, err := f.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	if line.Len() > 0 && encodedFields.Len() != 1 {
		line.AppendString(" ")
	}
	line.AppendString(encodedFields.String())
	encodedFields.Free()

	return line, err
}

func ParseFormat(spec string) ([]Formatter, error) {
	cursor := 0
	formatters := []Formatter{}

	matches := formatRegexp.FindAllStringSubmatchIndex(spec, -1)
	for _, m := range matches {
		fullStart, fullEnd := m[0], m[1]
		kindStart, kindEnd := m[2], m[3]
		formatStart, formatEnd := m[4], m[5]

		if fullStart > cursor {
			formatters = append(formatters, StringFormat{Value: spec[cursor:fullStart]})
		}

		var format string
		if formatStart > -1 {
			format = spec[formatStart:formatEnd]
		}

		formatter, err := NewFormatter(spec[kindStart:kindEnd], format)
		if err != nil {
			return nil, err
		}

		formatters = append(formatters, formatter)
		cursor = fullEnd
	}

	if cursor != len(spec) {
		formatters = append(formatters, StringFormat{Value: spec[cursor:]})
	}

	return formatters, nil
}

func NewFormatter(kind, format string) (Formatter, error) {
	switch kind {
	case "color":
		return newColorFormatter(format)
	case "sequence":
		return newSequenceFormatter(format), nil
	case "level":
		return newLevelFormatter(format), nil
	case "message":
		return newMessageFormatter(format), nil
	case "module":
		return newModuleFormatter(format), nil
	case "shortfunc":
		return newShortFuncFormatter(format), nil
	case "time":
		return newTimeFormatter(format), nil
	default:
		return nil, fmt.Errorf("unknown logger kind: %s", kind)
	}
}

type StringFormat struct {
	Value string
}

func (sf StringFormat) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	fmt.Fprintf(w, "%s", sf.Value)
}

type ColorFormatter struct {
	Bold  bool
	Reset bool
}

func newColorFormatter(f string) (ColorFormatter, error) {
	switch f {
	case "bold":
		return ColorFormatter{Bold: true}, nil
	case "reset":
		return ColorFormatter{Reset: true}, nil
	case "":
		return ColorFormatter{}, nil
	default:
		return ColorFormatter{}, fmt.Errorf("invalid color option: %s", f)
	}
}

func (cf ColorFormatter) LevelColor(level zapcore.Level) Color {
	switch level {
	case zapcore.DebugLevel:
		return ColorCyan
	case zapcore.InfoLevel:
		return ColorBlue
	case zapcore.WarnLevel:
		return ColorYellow
	case zapcore.ErrorLevel:
		return ColorRed
	case zapcore.PanicLevel, zapcore.DPanicLevel, zapcore.FatalLevel:
		return ColorMagenta
	default:
		return ColorNone
	}
}

func (cf ColorFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	switch {
	case cf.Reset:
		fmt.Fprint(w, ResetColor())
	case cf.Bold:
		fmt.Fprint(w, cf.LevelColor(entry.Level).Bold())
	default:
		fmt.Fprint(w, cf.LevelColor(entry.Level).Normal())
	}
}

type LevelFormatter struct {
	FormatVerb string
}

func newLevelFormatter(format string) LevelFormatter {
	return LevelFormatter{FormatVerb: "%" + stringOrDefault(format, "s")}
}

func (lf LevelFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	fmt.Fprintf(w, lf.FormatVerb, entry.Level.CapitalString())
}

type MessageFormatter struct {
	FormatVerb string
}

func newMessageFormatter(fortmat string) MessageFormatter {
	return MessageFormatter{FormatVerb: "%" + stringOrDefault(fortmat, "s")}
}

func (mf MessageFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	fmt.Fprintf(w, mf.FormatVerb, strings.TrimRight(entry.Message, "\n"))
}

type ModuleFormatter struct {
	FormatVerb string
}

func newModuleFormatter(format string) ModuleFormatter {
	return ModuleFormatter{FormatVerb: "%" + stringOrDefault(format, "s")}
}

func (mf ModuleFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	fmt.Fprintf(w, mf.FormatVerb, entry.LoggerName)
}

var sequence uint64

func SetSequence(s uint64) {
	atomic.StoreUint64(&sequence, s)
}

type SequenceFormatter struct {
	FormatVerb string
}

func newSequenceFormatter(format string) SequenceFormatter {
	return SequenceFormatter{FormatVerb: "%" + stringOrDefault(format, "d")}
}

func (sf SequenceFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	fmt.Fprintf(w, sf.FormatVerb, atomic.AddUint64(&sequence, 1))
}

type ShortFuncFormatter struct {
	FormatVerb string
}

func newShortFuncFormatter(format string) ShortFuncFormatter {
	return ShortFuncFormatter{FormatVerb: "%" + stringOrDefault(format, "s")}
}

func (sff ShortFuncFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	f := runtime.FuncForPC(entry.Caller.PC)
	if f == nil {
		fmt.Fprintf(w, sff.FormatVerb, "(unknown)")
		return
	}
	fname := f.Name()
	funcIdx := strings.LastIndex(fname, ".")
	fmt.Fprintf(w, sff.FormatVerb, fname[funcIdx+1:])
}

type TimeFormatter struct {
	Layout string
}

func newTimeFormatter(layout string) TimeFormatter {
	return TimeFormatter{Layout: stringOrDefault(layout, "2006-01-02T15:04:05.999Z07:00")}
}

func (tf TimeFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	fmt.Fprint(w, entry.Time.Format(tf.Layout))
}

type MultiFormatter struct {
	mutex      sync.RWMutex
	formatters []Formatter
}

func NewMultiFormatter(formatters ...Formatter) *MultiFormatter {
	return &MultiFormatter{
		formatters: formatters,
	}
}

func (mf *MultiFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	mf.mutex.RLock()
	for _, formatter := range mf.formatters {
		formatter.Format(w, entry, fields)
	}
}

func (mf *MultiFormatter) SetFormatters(formatters []Formatter) {
	mf.mutex.Lock()
	mf.formatters = formatters
	mf.mutex.Unlock()
}

func stringOrDefault(str, d string) string {
	if str != "" {
		return str
	}
	return d
}
