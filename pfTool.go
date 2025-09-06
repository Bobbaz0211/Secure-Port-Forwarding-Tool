package main
import (
     "os"
     "fmt"
     "encoding/json"
     "reflect"

)

const defCommand string = "ssh"

var baseArgs = []string{
    "-N",
    "-o",
    "ExitOnForwardFaliure=yes",
    "-o",
    "StreamLocalBindUnlink=yes",
    "-o",
    "ServerAliveInterval=5",
    "-o",
    "ServerAliveCountMax=1",
    "-L",


}

type configStruct struct {
	Host string `json:"host"`
	Type string `json:"Type"`
	Port int `json:"Port"`
}

func main(){
	data, err := os.ReadFile("./pfToolConfig.json")
	if err != nil {
		fmt.Printf("Error Reading Config file", err)
		os.Exit(1)
	}
	
	var cfg configStruct

	err =json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Printf("Error Unmarshlling Json", err)
		os.Exit(1)
	}

	result := make(map[string][]string)

	mapValue := reflect.ValueOf(cfg)
	mapType := reflect.TypeOf(cfg)

	fields := []string{}
	for i := 0; i< mapValue.NumField(); i++{
		fieldValue := mapValue.Field(i).Interface()
		fields = append(fields, fmt.Sprintf("%v", fieldValue))
	}

	result[mapType.Name()] = fields
	
	fmt.Println(result)

	
}
