package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"

)


func main() {
  log.Println("Starting Server")
  log.Println("Starting mongo db session")

  r := mux.NewRouter()
  r.HandleFunc("/sockets", serveWs)
  r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
  http.Handle("/", r)

  log.Println("Listening on 8080")
  go h.run()
  err:=http.ListenAndServe("10.0.0.2:8080", nil)
  if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


}
