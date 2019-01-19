package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/textproto"
)

type Bot struct {
	server        string
	port          string
	nick          string
	user          string
	channel       string
	pass          string
	pread, pwrite chan string
	conn          net.Conn
}

func NewBot() *Bot {
	return &Bot{
		server:  "irc.freenode.net",
		port:    "6667",
		nick:    "stansbot1212",
		channel: "#somethingfortesting",
		pass:    "",
		conn:    nil,
		user:    "superbot1212",
	}
}

func (b *Bot) Connect() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", b.server+":"+b.port)
	if err != nil {
		log.Fatal("unable to connect to irc", err)
	}

	b.conn = conn
	log.Printf("connected to irc server")
	return b.conn, nil
}

func main() {
	ircbot := NewBot()
	conn, _ := ircbot.Connect()
	fmt.Fprintf(conn, "USER %s 8 * : %s\r\n", ircbot.nick, ircbot.nick)
	fmt.Fprintf(conn, "NICK %s\r\n", ircbot.nick)
	fmt.Fprintf(conn, "JOIN %s\r\n", ircbot.channel)
	defer conn.Close()

	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	for {
		line, err := tp.ReadLine()
		if err != nil {
			break
		}
		fmt.Printf("%s\n", line)
	}
}
