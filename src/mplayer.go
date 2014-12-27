package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"library"
	"mp"
)

var lib *library.MusicManager
var id int = 0
var ctrl, signal chan int

func handleLibCommands(tokens []string) {

	switch tokens[1] {
		case "list":
			for i := 1; i <= len(lib.Musics);i++ {
				e,_ := lib.Get(i)
				fmt.Println(i, ":", e.Name, e.Artist, e.Source, e.Type)
			}
		case "add":
			if len(tokens) == 7 {
				id ++
				lib.Add(&library.MusicEntry{strconv.Itoa(id),tokens[2],tokens[3],tokens[4],tokens[5],tokens[6]})
			}else {
				fmt.Println("USAGE: Lib add <name><artist><source><type>")
			}
		case "remove":
			if len(tokens) == 3 {
				musicIndex,_ := strconv.Atoi(tokens[2]) 
				lib.Remove(musicIndex)
			}else{
				fmt.Println("USAGE: Lib remove <name>")
			}
		default:
			fmt.Println("Unrecognized Lib command:", tokens[1])
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) !=2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "dose not exist.")
		return
	}
	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println(`
		Enter following commands to control the player:
		Lib list -- View the existing music Lib
		Lib add <name><artist><source><type> -- Add a music to the muisc Lib
		Lib remove <name> -- Remove the specified music from the Lib
		paly <name> -- Play the specified music
	`)
	lib = library.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for  {
		fmt.Print("Enter command->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "exit" {
			break
		}
		tokens := strings.Split(line, " ")

		if tokens[0] == "Lib" {
			handleLibCommands(tokens)
		}else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		}else {
			fmt.Println("Unrecognizend command:", tokens[0])
		}
	}
}