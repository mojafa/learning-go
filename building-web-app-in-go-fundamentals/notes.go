package main

import (
	"html/template"
	"log"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}

// not accesing the disk to get tempplate we put it in the app config
// Repository patter
type Repository struct {
	App *AppConfig
}

// import cycles is an issie fix by using models

/*
Routing in go we use pat and chi
*/

// Middleware
//  allows to process requests as they come in and make decisons on ehat to do wityh them

//  like striplashes on  url

//  golang/nosurf - gen and prevents CSRF

//  sessions also use middlewares

//  gorilla and alexedwrads/scs - packages for sessions

//  variable shadowing - when you have a variable in a nested scope with the same name as a variable in the outer scope
//  webservers are not state aware by default
