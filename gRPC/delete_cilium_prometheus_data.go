package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func deleteData() {
	resp, err := http.Post(`http://10.0.0.252:39000/api/v1/admin/tsdb/delete_series?match[]=hubble_port_distribution_total`, "", nil)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(len(string(body)))
	fmt.Println(string(body))

	resp, err = http.Post(`http://10.0.0.252:39000/api/v1/admin/tsdb/clean_tombstones`, "", nil)
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(len(string(body)))
	fmt.Println(string(body))
}

func main() {
	log.Println("begin")
	deleteData()
}
