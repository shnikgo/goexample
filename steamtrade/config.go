package steamtrade


import (
	"encoding/json"
	"io/ioutil"
)


var Config struct {
	AppId string
	AppKey string
	ServerPort string
	MinLogLevel uint8
	SentryMinLogLevel uint8
	SentryConnection string
}


func ReloadConfig(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		Log(err.Error(), "", LOG_LEVEL_FATAL)
		return err
	}
	err = json.Unmarshal(buf, &Config)
	if err != nil {
		Log(err.Error(), "", LOG_LEVEL_FATAL)
		return err
	}
	return nil
}
