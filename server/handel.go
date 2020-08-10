package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var buffSize = 4096

func StartServer() {
	// init random seed
	rand.Seed(time.Now().Unix())

	mux := http.NewServeMux()

	// api router here
	mux.HandleFunc("/api/add_suggestion/", apiSuggestionAddHandleFunc)
	mux.HandleFunc("/api/reply_suggestion/", apiReplySuggestionByIdFunc)
	mux.HandleFunc("/api/get_suggestion_list_by_id/", apiGetSuggestionListByIdFunc)

	// router here
	mux.HandleFunc("/", indexHandelFunc)

	var s = &http.Server {
		Addr: ":8039",
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
}

func shortCutMethodNotSupport(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(403)
	json.NewEncoder(w).Encode(&MessageMethodNotSupport)
}
func shortCutHandleError(w http.ResponseWriter, req *http.Request, err error) {
	log.Println("Error: ", err.Error())
	w.WriteHeader(500)
	message := Message {
		Title: "Error",
		Content: err.Error(),
	}
	json.NewEncoder(w).Encode(&message)
}

func apiReplySuggestionByIdFunc(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		shortCutMethodNotSupport(w, req)
	}
	var request Suggestion
	var err error
	json.NewDecoder(req.Body).Decode(&request)
	// insert into database
	_, err = InsertStmt.Exec(
		request.Id,
		2,
		time.Now().Unix(),
		request.Content)
	if err != nil {
		shortCutHandleError(w, req, err)
		return
	}
	log.Println("Reply a suggestion: ", request.Content)
	json.NewEncoder(w).Encode(MessageSuccess)
}

func apiGetSuggestionListByIdFunc(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		shortCutMethodNotSupport(w, req)
	}
	var request Suggestion
	json.NewDecoder(req.Body).Decode(&request)
	// query database
	rows, err := QueryByIdStmt.Query(request.Id)
	if err != nil {
		shortCutHandleError(w, req, err)
		return
	}
	// new suggestion list
	var suggestion_list SuggestionList
	for rows.Next() {
		var suggestion Suggestion
		err = rows.Scan(&suggestion.Id, &suggestion.Type, &suggestion.Time, &suggestion.Content)
		if err != nil {
			shortCutHandleError(w, req, err)
			return
		}
		suggestion_list.SuggestionList = append(suggestion_list.SuggestionList, suggestion)
	}
	json.NewEncoder(w).Encode(&suggestion_list)
}

func apiSuggestionAddHandleFunc(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		shortCutMethodNotSupport(w, req)
		return
	}
	// decode json
	// only need "type" and "content"
	var suggestion Suggestion
	var err error
	jsonDecoder := json.NewDecoder(req.Body)
	err = jsonDecoder.Decode(&suggestion)
	if err != nil {
		log.Println("Error at decode json ", err)
		return
	}
	// generate suggestion information
	suggestion.Id = genRandomId()
	suggestion.Type = 1
	suggestion.Time = time.Now().Unix()

	// insert database
	_, err = InsertStmt.Exec(
		suggestion.Id,
		suggestion.Type,
		suggestion.Time,
		suggestion.Content)
	if err != nil {
		shortCutHandleError(w, req, err)
		return
	}
	log.Println("New suggestion: ", suggestion.Content)
	json.NewEncoder(w).Encode(&suggestion)
}

func genRandomId() int64 {
	var suggestion Suggestion
	var id int64
	for {
		id = rand.Int63()
		err := QueryByIdStmt.QueryRow(id).Scan(&suggestion)
		if err == sql.ErrNoRows {
			break
		}
	}
	return id
}

func indexHandelFunc(w http.ResponseWriter, req *http.Request) {
	// because / match everything
	// so we need to check url here
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	fmt.Fprintf(w, "Welcome to home page")
}
