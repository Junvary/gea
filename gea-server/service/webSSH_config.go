package service

import (
	"io/ioutil"
	"log"
	"time"
)

type messageType string

const (
	messageTypeCurrentUser = "currentUser"
	messageTypeAddr		 = "addr"
	messageTypeTerm      = "term"
	messageTypeLogin     = "login"
	messageTypePassword  = "password"
	messageTypePublickey = "publickey"
	messageTypeStdin     = "stdin"
	messageTypeStdout    = "stdout"
	messageTypeStderr    = "stderr"
	messageTypeResize    = "resize"
	messageTypeIgnore    = "ignore"
	messageTypeConsole   = "console"
	messageTypeAlert   	 = "alert"

	TermLinux		 = "linux"
	TermAnsi		 = "ansi"
	TermScoAnsi		 = "scoansi"
	TermXterm		 = "xterm"
	TermXterm256Color = "xterm-256color"
	TermVt100		 = "vt100"
	TermVt102		 = "vt102"
	TermVt220		 = "vt220"
	TermVt320		 = "vt320"
	TermWyse50		 = "wyse50"
	TermWyse60		 = "wyse60"
	TermDumb		 = "dumb"
)

type message struct {
	Type messageType `json:"type"`
	Data []byte      `json:"data"`
	Cols int         `json:"cols,omitempty"`
	Rows int         `json:"rows,omitempty"`
}

var (
	DefaultTerm = TermXterm
	DefaultConnTimeout = 15 * time.Second
	DefaultLogger = log.New(ioutil.Discard, "[gea-webSSH] ", log.Ltime|log.Ldate)
	DefaultBuffSize = uint32(8192)
)
