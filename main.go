package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	pInput := flag.String("input", "", "input json file")
	flag.Parse()

	if *pInput == "" {
		fmt.Println("Error. Input is empty")
		flag.PrintDefaults()
		os.Exit(-1)
	}

	inputBytes, err := ioutil.ReadFile(*pInput)
	if err != nil {
		fmt.Printf("Error reading input %s: %v\n", *pInput, err)
		os.Exit(-1)
	}

	var obj interface{}
	err = json.Unmarshal(inputBytes, &obj)
	if err != nil {
		fmt.Printf("Error unmarshalling input. Is it valid JSON? %s: %v\n", *pInput, err)
		os.Exit(-1)
	}

	if obj != nil {
		yamlBytes, err := yaml.Marshal(obj)
		if err != nil {
			fmt.Printf("Error marshaling into YAML from JSON file %s: %v\n", *pInput, err)
			os.Exit(-1)
		}

		fmt.Println(string(yamlBytes))
	}

}
