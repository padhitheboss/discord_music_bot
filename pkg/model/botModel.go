package model

import (
	"fmt"
)

type BotCommand struct {
	Id       int
	Name     string
	Command  string
	Query    string
	User     string
	UserType string
}
type PlayerStatus struct {
	Playing bool
	RunChan chan bool
}

var Q PlayQueue
var P PlayerStatus
var PlayFromQueueRunning = false

func (b *BotCommand) Do() string {
	var s Song
	var response string
	fmt.Println(b.Query)
	if b.Command == "!play" || b.Command == "!p" {
		P.Playing = true
		P.RunChan = make(chan bool)
		s.CreateRequest(b.Query, b.User)
		Q.Lock.Lock()
		Q.AddAtBack(s)
		Q.Lock.Unlock()
		response = s.Title + "Added to Queue Duration=" + s.Duration
	} else if b.Command == "!stop" {
		P.Playing = false
		fmt.Println("Closing Channel")
		close(P.RunChan)
	} else if b.Command == "!skip" {
		P.Playing = false
		fmt.Println("Closing Channel")
		close(P.RunChan)

		if len(Q.Playlist) == 0 {
			response = "no song in queue"
		}
		for PlayFromQueueRunning {
		}
		P.Playing = true
		P.RunChan = make(chan bool)
	} else if b.Command == "!resume" {
		if P.Playing {
			response = "already playing"
		} else {
			P.Playing = true
			P.RunChan = make(chan bool)
		}
	} else if b.Command == "!queue" || b.Command == "!q" {
		for i := 0; i < len(Q.Playlist); i++ {
			response += "\n" + Q.Playlist[i].Title
		}
		if len(Q.Playlist) == 0 {
			response = "playlist is empty"
		}
	} else if b.Command == "!clear" || b.Command == "!c" {
		Q.Lock.Lock()
		Q.ClearQueue()
		Q.Lock.Unlock()
		response = "cleared playlist queue"
	} else if b.Command == "!playnext" || b.Command == "!pn" {
		s.CreateRequest(b.Query, b.User)
		Q.Lock.Lock()
		Q.AddAtFront(s)
		Q.Lock.Unlock()
		response = s.Title + "Added to Queue Duration=" + s.Duration
	}
	fmt.Println(Q.Playlist)
	return response
}
