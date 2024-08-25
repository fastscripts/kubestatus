package config

import (
	"encoding/json"
	"log"

	"math/bits"
	"os"
	"reflect"
	"strconv"
)

type AppConfig struct {
	Port           int    `json:"Port" env:"PORT"`
	MetricsPort    int    `json:"MetricsPort" env:"METRICS_PORT"`
	Devmode        bool   `json:"Devmode" env:"DEVMODE"`
	KubeAccessType string `json:"KubeAccessType" env:"KUBE_ACCESS_TYPE"`
	KubeConfigPath string `json:"KubeConfigPath" env:"KUBE_CONFIG_PATH"`
}

func (app *AppConfig) LoadJSONConfiguration(file string) {

	configFileHandle, err := os.Open(file)
	if err != nil {
		log.Print("configuration file", file, "not found, using defaults ", err)
	} else {
		jsonParser := json.NewDecoder(configFileHandle)
		jsonParser.Decode(app)
	}
	defer configFileHandle.Close()

}

// LoadENVConfiguration try to fill out configuration struct with values of ENV variables
func (app *AppConfig) LoadENVConfiguration() {
	ptrType := reflect.TypeOf(app)
	ref := ptrType.Elem()
	// config has to be a pointer to struct
	if ptrType.Kind() == reflect.Ptr && ref.Kind() == reflect.Struct {
		for i := 0; i < ref.NumField(); i++ {
			element := ref.Field(i)

			//value := reflect.ValueOf(config).Elem().FieldByName(element.Name)
			//fmt.Println("Type: ", element.Type, "; Name: ", element.Name, "; Tags: ", element.Tag, "; Value: ", value)

			value := ""

			// check if we have an environment variable with tag Name
			tagName := element.Tag.Get("env")
			if len(tagName) > 0 {
				value = os.Getenv(tagName)
			} else {
				// if no tag is defined, check if we have an environment variable with element name
				value = os.Getenv(element.Name)
			}
			if len(value) > 0 {

				field := reflect.ValueOf(app).Elem().FieldByName(element.Name)
				if field.IsValid() && field.CanSet() {
					if field.Kind() == reflect.Int {
						machineType := bits.UintSize
						if machineType == 64 {
							convertedValue, err := strconv.ParseInt(value, 10, 64)
							if err == nil {
								if !field.OverflowInt(convertedValue) {
									field.SetInt(convertedValue)
								}
							}
						} else {
							convertedValue, err := strconv.ParseInt(value, 10, 32)
							if err == nil {
								if !field.OverflowInt(convertedValue) {
									field.SetInt(convertedValue)
								}
							}
						}
					} else if field.Kind() == reflect.Int64 {
						convertedValue, err := strconv.ParseInt(value, 10, 64)
						if err == nil {
							if !field.OverflowInt(convertedValue) {
								field.SetInt(convertedValue)
							}
						}
					} else if field.Kind() == reflect.Int32 {
						convertedValue, err := strconv.ParseInt(value, 10, 32)
						if err == nil {
							if !field.OverflowInt(convertedValue) {
								field.SetInt(convertedValue)
							}
						}
					} else if field.Kind() == reflect.Int16 {
						convertedValue, err := strconv.ParseInt(value, 10, 16)
						if err == nil {
							if !field.OverflowInt(convertedValue) {
								field.SetInt(convertedValue)
							}
						}
					} else if field.Kind() == reflect.Uint {
						machineType := bits.UintSize
						if machineType == 64 {
							convertedValue, err := strconv.ParseUint(value, 10, 64)
							if err == nil {
								if !field.OverflowUint(convertedValue) {
									field.SetUint(convertedValue)
								}
							}
						} else {
							convertedValue, err := strconv.ParseUint(value, 10, 32)
							if err == nil {
								if !field.OverflowUint(convertedValue) {
									field.SetUint(convertedValue)
								}
							}
						}
					} else if field.Kind() == reflect.Uint64 {
						convertedValue, err := strconv.ParseUint(value, 10, 64)
						if err == nil {
							if !field.OverflowUint(convertedValue) {
								field.SetUint(convertedValue)
							}
						}
					} else if field.Kind() == reflect.Uint32 {
						convertedValue, err := strconv.ParseUint(value, 10, 32)
						if err == nil {
							if !field.OverflowUint(convertedValue) {
								field.SetUint(convertedValue)
							}
						}
					} else if field.Kind() == reflect.Uint16 {
						convertedValue, err := strconv.ParseUint(value, 10, 16)
						if err == nil {
							if !field.OverflowUint(convertedValue) {
								field.SetUint(convertedValue)
							}
						}
					} else if field.Kind() == reflect.Float32 {
						convertedValue, err := strconv.ParseFloat(value, 32)
						if err == nil {
							if !field.OverflowFloat(convertedValue) {
								field.SetFloat(convertedValue)
							}
						}
					} else if field.Kind() == reflect.Float64 {
						convertedValue, err := strconv.ParseFloat(value, 64)
						if err == nil {
							if !field.OverflowFloat(convertedValue) {
								field.SetFloat(convertedValue)
							}
						}
					} else if field.Kind() == reflect.Bool {
						convertedValue, err := strconv.ParseBool(value)
						if err == nil {
							field.SetBool(convertedValue)
						}
					} else if field.Kind() == reflect.String {
						field.SetString(value)
					}
				}
			}
		}
	}
}
