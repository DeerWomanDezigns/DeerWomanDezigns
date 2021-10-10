package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-yaml/yaml"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func orderPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "orderPage!")
	fmt.Println("Endpoint Hit: orderPage")
}

func handleRequests(port int) {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/order", orderPage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

type Config struct {
	Port int `yaml:"port"`
}

func readConfig() Config {
	dat, readErr := ioutil.ReadFile("webConfig.yaml")
	check(readErr)
	var config Config
	yamlErr := yaml.Unmarshal(dat, &config)
	check(yamlErr)
	return config
}

func main() {
	config := readConfig()

	handleRequests(config.Port)
}
