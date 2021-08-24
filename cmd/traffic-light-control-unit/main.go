package main

import (
	"bufio"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
	"github.com/Tinkerforge/go-api-bindings/rgb_led_v2_bricklet"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const (
	host     = ""
	port     = "55554"
	protocol = "tcp"

	tfAddr = "localhost:4223"
	tfUid  = "VRg"
)

func main() {
	ipcon := ipconnection.New()
	defer ipcon.Close()
	rgbledv2Bricklet, _ := rgb_led_v2_bricklet.New(tfUid, &ipcon) // Create device object.

	err := ipcon.Connect(tfAddr) // Connect to brickd.
	if err != nil {
		log.Fatal("Can't connect to LED Bricklet", err)
	}
	defer ipcon.Disconnect()

	// Flash LED on startup
	time.Sleep(3 * time.Second)
	rgbledv2Bricklet.SetRGBValue(255, 255, 255)
	time.Sleep(3 * time.Second)
	rgbledv2Bricklet.SetRGBValue(0, 0, 0)

	log.Println("Starting " + protocol + " server on " + host + ":" + port)
	listener, err := net.Listen(protocol, host+":"+port)
	if err != nil {
		log.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Println("Error connecting:", err.Error())
			return
		}
		log.Println("Client " + c.RemoteAddr().String() + " connected.")
		go handleConnection(c, rgbledv2Bricklet)
	}
}

func handleConnection(conn net.Conn, bricklet rgb_led_v2_bricklet.RGBLEDV2Bricklet) {
	buffer, err := bufio.NewReader(conn).ReadByte()
	//buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	//buffer, err := ioutil.ReadAll(bufio.NewReader(conn))

	if err != nil {
		log.Println("Client left")
		conn.Close()
		return
	}

	msg := strings.TrimSpace(string(buffer))
	log.Println("Message: " + msg)

	if msg == "4" {
		log.Println("Traffic light: red")
		bricklet.SetRGBValue(255, 0, 0)
		conn.Write([]byte("\nred\n"))
	} else if msg == "1" {
		log.Println("Traffic light: green")
		bricklet.SetRGBValue(0, 255, 0)
		conn.Write([]byte("\ngreen\n"))
	} else if msg == "2" {
		log.Println("Traffic light: yellow")
		bricklet.SetRGBValue(255, 255, 0)
		conn.Write([]byte("\nyellow\n"))
	} else if msg == "0" {
		log.Println("Traffic light: off")
		bricklet.SetRGBValue(0, 0, 0)
		conn.Write([]byte("\noff\n"))
	}

	handleConnection(conn, bricklet)
}
