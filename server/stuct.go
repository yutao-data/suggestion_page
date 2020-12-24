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

type Passwd struct {
	Passwd string `json:"passwd"`
}
type SuggestionList struct {
	SuggestionList []Suggestion `json:"suggestion_list,omitempty"`
}

type Suggestion struct {
	Id int64 `json:"id,omitempty"`
	Type bool `json:"type,omitempty"`
	First bool `json:"first,omitempty"`
	Time int64 `json:"time,omitempty"`
	Content string `json:"content,omitempty"`
	Passwd string `json:"passwd"`
	Show bool `json:"show"`
}

var InsertStmt *sql.Stmt
var QueryByIdStmt *sql.Stmt
var QueryAllFirstStmt *sql.Stmt
var SetSuggestionShowStmt *sql.Stmt

func InitAllStmt() {
	var err error
	if DB == nil {
		log.Fatal("DB is nil")
	}
	InsertStmt, err = DB.Prepare(
		`INSERT INTO suggestion(id, type, first, time, content, show) 
		VALUES (?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal("Init Stmt failed ", err)
	}
	QueryByIdStmt, err = DB.Prepare(
		`SELECT * FROM suggestion WHERE id=?`)
	if err != nil {
		log.Fatal("Init QueryByIdStmt failed ", err)
	}
	QueryAllFirstStmt, err = DB.Prepare(
		`SELECT * FROM suggestion WHERE first=true AND show=true`)
	if err != nil {
		log.Fatal("Init QueryAllFirstStmt failed ", err)
	}
	SetSuggestionShowStmt, err = DB.Prepare(
		`UPDATE suggestion SET show=false WHERE id=? AND first=true`)
	if err != nil {
		log.Fatal("Init SetSuggestionStmt failed ", err)
	}
}
