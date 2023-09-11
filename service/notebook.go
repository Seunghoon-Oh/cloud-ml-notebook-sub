package service

import (
	"encoding/json"
	"fmt"

	"github.com/Seunghoon-Oh/cloud-ml-notebook-subscriber/network"
	circuit "github.com/rubyist/circuitbreaker"
)

var notebookCb *circuit.Breaker
var notebookClient *circuit.HTTPClient

func SetupNotebookCircuitBreaker() {
	notebookClient, notebookCb = network.GetHttpClient()
}

func CreateNotebook() {
	if notebookCb.Ready() {
		resp, err := notebookClient.Post("http://cloud-ml-notebook-manager.cloud-ml-notebook:8082/notebook", "", nil)
		if err != nil {
			fmt.Println(err)
			notebookCb.Fail()
			return
		}
		notebookCb.Success()
		defer resp.Body.Close()
		rsData := network.ResponseData{}
		json.NewDecoder(resp.Body).Decode(&rsData)
		fmt.Println(rsData.Data)
		return
	}
}
