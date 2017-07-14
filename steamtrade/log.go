package steamtrade


import (
	"fmt"
	"time"
	"encoding/json"
)


const (
	LOG_LEVEL_DEBUG = 0 + iota
	LOG_LEVEL_INFO
	LOG_LEVEL_WARNING
	LOG_LEVEL_ERROR
	LOG_LEVEL_FATAL
)


type LogFormat struct {
	Level uint8 `json:"level"`
	LevelStr string `json:"level_str"`
	Message string `json:"message"`
	Request string `json:"request"`
	Time time.Time `json:"time"`
}


func Log(message string, request string, level uint8) {
	if level < Config.MinLogLevel {
		return
	}
	
	log_levels := []string{ "Debug", "Info", "Warning", "Error", "Fatal" }

	if int(level) > len(log_levels) - 1 {
		Log("level > LOG_LEVEL_FATAL", "", LOG_LEVEL_FATAL)
		level = LOG_LEVEL_FATAL;
	}

	log := LogFormat{
		Level: level,
		LevelStr: log_levels[level],
		Message: message,
		Request: request,
		Time: time.Now(),
	}
	
	text, err := json.Marshal(log)
	if err != nil {
		fmt.Prinln(err.Error())
	}
	fmt.Prinln(text)
}
