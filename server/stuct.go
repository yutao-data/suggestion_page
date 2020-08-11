package main

import (
	"database/sql"
	"log"
)

type Message struct {
	Title string `json:"title"`
	Content string `json:"Content"`
}
var MessageMethodNotSupport = Message {
	Title: "Method not suppoert",
	Content: "Your HTTP method is not support by the server",
}
var MessageSuccess = Message {
	Title: "Success",
	Content: "Action success",
}

type SuggestionList struct {
	SuggestionList []Suggestion `json:"suggestion_list"`
}

type Suggestion struct {
	Id int64 `json:"id,omitempty"`
	Type int64 `json:"type,omitempty"`
	Time int64 `json:"time,omitempty"`
	Content string `json:"content,omitempty"`
}

var InsertStmt *sql.Stmt
var QueryByIdStmt *sql.Stmt

func InitAllStmt() {
	var err error
	if DB == nil {
		log.Fatal("DB is nil")
	}
	InsertStmt, err = DB.Prepare(
		`INSERT INTO suggestion(id, type, time, content) 
		VALUES (?, ?, ?, ?)`)
	if err != nil {
		log.Fatal("Init Stmt failed ", err)
	}
	QueryByIdStmt, err = DB.Prepare(
		`SELECT * FROM suggestion WHERE id=?`)
	if err != nil {
		log.Fatal("Init QueryByIdStmt failed ", err)
	}
}
