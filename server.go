package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"flag"
	"fmt"
)

var Config_map_string map[string]string

func main() {

	development := flag.Bool("production", false, "run it in production mode")

	viper.SetConfigName("sign-server")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}



	if *development == true {
		Config_map_string = viper.GetStringMapString("development")
	} else {
		Config_map_string = viper.GetStringMapString("production")
	}

	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/upload", UploadHandler).Methods("PUT")
	r.HandleFunc("/file/{id}", GetFileHandler).Methods("GET")
	n.UseHandler(r)
	n.Run(fmt.Sprintf("0.0.0.0:%s",Config_map_string["port"]))
}
