package main

import (
	"database/sql"
	"encoding/json"
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
	mux.HandleFunc("/api/get_all_suggestion_list/", apiGetAllSuggestionListFunc)

	// static file router here
	mux.Handle("/", http.FileServer(http.Dir(rootPath)))

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

func apiGetAllSuggestionListFunc(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		shortCutMethodNotSupport(w, req)
		return
	}
	// query database
	rows, err := QueryAllFirstStmt.Query()
	if err != nil {
		shortCutHandleError(w, req, err)
		return
	}
	var suggestionList SuggestionList
	var suggestion Suggestion
	for rows.Next() {
		err = rows.Scan(&suggestion.Id, &suggestion.Type, &suggestion.First, &suggestion.Time, &suggestion.Content)
		if err != nil {
			shortCutHandleError(w, req, err)
			return
		}
		suggestionList.SuggestionList = append(suggestionList.SuggestionList, suggestion)
	}
	log.Println("query all first suggestion")
	json.NewEncoder(w).Encode(&suggestionList)
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
		false, // type
		false, // frist
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
		// todo rows.Scan(&suggestion)
		err = rows.Scan(&suggestion.Id, &suggestion.Type, &suggestion.First, &suggestion.Time, &suggestion.Content)
		if err != nil {
			shortCutHandleError(w, req, err)
			return
		}
		suggestion_list.SuggestionList = append(suggestion_list.SuggestionList, suggestion)
	}
	suggestion_list.SuggestionList = append(suggestion_list.SuggestionList, request)
	log.Println("Query suggestion list by id ", request.Id)
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
	// generate suggestion information in object
	// because it needs to be return
	suggestion.Id = genRandomId()
	suggestion.Type = true
	suggestion.First = true
	suggestion.Time = time.Now().Unix()
	// insert database
	_, err = InsertStmt.Exec(
		suggestion.Id,
		suggestion.Type,
		suggestion.First,
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
		id = rand.Int63() % 10000
		err := QueryByIdStmt.QueryRow(id).Scan(&suggestion)
		if err == sql.ErrNoRows {
			break
		}
	}
	return id
}
