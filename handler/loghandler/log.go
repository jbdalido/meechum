package logHandler

import (
	"github.com/jbdalido/meechum"
	"log"
)

type LogHandler struct {
	Prefix string
}

func NewLogHandler() *LogHandler {
	return &LogHandler{}
}

func (l *LogHandler) Fire(r *meechum.Result) error {
	log.Printf("[Handler][log] STD:%s LEVEL:%s ECODE:%d", r.StdOut, r.Level, r.Code)
	return nil
}

func (l *LogHandler) Levels() []meechum.ErrorCode {
	return []meechum.ErrorCode{
		meechum.WARNING,
		meechum.FATAL,
	}
}

func (l *LogHandler) String() string {
	return "log"
}
