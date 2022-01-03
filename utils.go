package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/imroc/req"
)

var host = "http://127.0.0.1:3030/api/v1"

func helperFunc() {
	fmt.Println("- help 'to get help commands'")
	fmt.Println("- host http://127.0.0.1:3030 'to set the gitrowsPackApi host, default is localhost'")
	fmt.Println("- show databases 'to get the list of databases'")
	fmt.Println("- show collections 'to get list of collections from a selected database (use dbname)'")
	fmt.Println("- use dbname.collectionname 'to select the database and the collection that will be used'")
	fmt.Println("- create dbname.collectionname 'to create a new collection from its name'")
	fmt.Println("- drop dbname.collectionname 'to drop/delete the whole collection'")
	fmt.Println("- get dbname.collectionname 'to get a specific element from the collection'")
	fmt.Println("- getall dbname.collectionname 'to get all elements from the collection'")
	fmt.Println("- insert dbname.collectionname 'to insert a new element in the collection'")
	fmt.Println("- update dbname.collectionname 'to update an existing element from the collection'")
	fmt.Println("- delete dbname.collectionname 'to delete an existing element from the collection'")
	fmt.Println("- bump dbname.collectionname 'to replace the whole collection'")
	fmt.Println("\nEx: use shop.users")
}

func getInputCommand() string {
	fmt.Print("\n> ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
		return ""
	}
	return input
}

func errorCheck(cond bool) bool {
	if cond {
		log.Fatal("[x] gitrowspack api server not ready !")
		return false
	}
	return true
}

func ping() bool {
	var p pingResponse

	r, err := req.Get(host + "/ping")

	errorCheck(err != nil)

	resp, err := r.ToString()

	errorCheck(err != nil)
	// We parse stuffs
	json.Unmarshal([]byte(resp), &p)

	errorCheck(r.Response().StatusCode != 200)

	fmt.Println(p.Version)
	fmt.Println("H::" + p.ServerIP)
	fmt.Println(p.Message)

	return true
}
