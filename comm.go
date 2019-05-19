package main

import (
	"bufio"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"go.bug.st/serial.v1"
)

// Configuration port
const BAUDRATE = 115200 // transmition baud rate
const PARITY = serial.NoParity
const DATABITS = 8
const STOPBITS = serial.OneStopBit

// Declaration vars
var port serial.Port
var serialPortArgument string

/*
*  FUNCTION:   main
*  PARAM:
*  RETURN:
*  DESCRIPTION: Main process.
 */
func main() {

	process()
}

/*
*  FUNCTION:   process
*  PARAM:
*  RETURN:
*  DESCRIPTION: Start process, read argument and read message for serial port.
 */
func process() {

	var err error

	flag.StringVar(&serialPortArgument, "serial", "", "a string")
	flag.Parse()

	if serialPortArgument == "" {
		fmt.Println("Error arguments param serial port ( -serial ) , pej: go run comm.go -serial=/dev/pts/1")
		return

	}
	fmt.Println("Serial port ready for connection: " + serialPortArgument)

	mode := &serial.Mode{
		BaudRate: BAUDRATE,
		Parity:   PARITY,
		DataBits: DATABITS,
		StopBits: STOPBITS,
	}

	port, err = serial.Open(serialPortArgument, mode)
	defer port.Close()
	if err != nil {
		println("ERROR SERIAL: " + err.Error())
		return
	}
	fmt.Println("Open serial sussessful")
	var wg sync.WaitGroup
	wg.Add(1)
	forever := make(chan bool)
	go func() {

		for {

			fmt.Println("Reading....: ")
			res := readSerial()

			if res == "SEND" {
				writeSerial("XXXXXXX" + "\r")

			} else if res == "ERROR" || res == "ABORT" {
				wg.Done()
				return
			}
		}

	}()
	close(forever)
	wg.Wait()

}

/*
*  FUNCTION:   	writeSerial
*  PARAM:		cad= String for write in serial port
*  RETURN:
*  DESCRIPTION: Method to write data to the serial port.
 */
func writeSerial(cad string) {

	println("Write serial: " + cad)
	_, err := port.Write([]byte(cad))

	if err != nil {

		println("Error when writing to the serial port: " + err.Error())
	}

}

/*
*  FUNCTION:   	readSerial
*  PARAM:
*  RETURN:		read string
*  DESCRIPTION: Method to read messages from the serial port
 */

func readSerial() string {
	reader := bufio.NewReader(port)
	reply, err := reader.ReadBytes('\x0a')
	if err != nil {
		println("Error reading data in serial port: " + err.Error())
	}
	var res string
	res = string(reply[:len(reply)])
	res = strings.TrimSpace(res)
	println("Response Data: " + res + "    Len Data: " + strconv.Itoa(len(res)))
	return res

}
