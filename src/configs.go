package src

import(
	"os"
	"encoding/json"
)

type config struct{
	Address string 
	Port int
	HtmlFile string

}

var Config config

func LogConfig(configFile string) error {
	data,err:=os.ReadFile(configFile)

	if err != nil {
		//错误处理
	}
	return json.Unmarshal(data, &Config)
}