package main

import (
	"Bus/db"
	"Bus/pkg/setting"
	"Bus/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	session, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	engine := router.InitRouter(session)
	addr := fmt.Sprintf("%s:%d", setting.Host, setting.HTTPPort)
	log.Println("Running server in: ", addr)
	log.Fatal(http.ListenAndServe(addr, engine))
}
