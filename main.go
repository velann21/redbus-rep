package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Bus struct {
	Source string `json:"source"`
	Destination string `json:"destination"`
}

func (bus *Bus) GetBuses(rw http.ResponseWriter, req *http.Request){
	busDb := &BusDatabase{
		datas: map[string]string{},
	}
	busDb.UploadBus()

	body := req.Body
	err := json.NewDecoder(body).Decode(bus)
	if err != nil{
		return
	}
	busName := busDb.GetBuses(bus.Source, bus.Destination)
	err = json.NewEncoder(rw).Encode(busName)
	if err != nil {
		return
	}
}

type BusDatabase struct {
	datas map[string]string
}

func (db *BusDatabase) UploadBus(){
	db.datas["erode-chennai"] = "kpn"
	db.datas["chennai-erode"] = "A1"
}
func (db *BusDatabase) GetBuses(source, dest string)string{
	busSourceDest := fmt.Sprintf("%s-%s", source, dest)
	bus := db.datas[busSourceDest]
	return bus
}

func handleRequests() {
	bus := &Bus{}
	http.HandleFunc("/buses", bus.GetBuses)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
