package main

import (
	"github.com/wpferg/services/httpHandlers"
	"github.com/wpferg/services/storage"
	"github.com/wpferg/services/structs"
	"log"
	"net/http"
	"strconv"
)

const PORT = 8080

func createMessage(message string, sender string) structs.Message {
return structs.Message{
Sender:  sender,
Message: message,
}
}

func main() {
log.Println("Creating dummy messages")

storage.Add(createMessage("Testing", "1234"))
storage.Add(createMessage("Testing Again",""))
storage.Add(createMessage("Testing A Third Time", "9012"))

log.Println("Attempting to start HTTP Server.")
http.HandleFunc("/", httpHandlers.HandleRequest)

var err = http.ListenAndServe(":"+strconv.Itoa(PORT), nil)

if err != nil {
log.Panicln("Server failed starting. Error: %s", err)
}
}