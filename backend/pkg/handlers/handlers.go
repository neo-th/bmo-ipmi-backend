package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tarm/serial"
)

var serialPort *serial.Port

// Initialize the serial connection to Arduino
func InitSerial() {
	// Get the serial device from the environment variable
	serialDevice := os.Getenv("SERIAL_DEVICE")
	if serialDevice == "" {
		serialDevice = "/dev/ttyACM0" // Default to /dev/ttyACM0 if not set
	}

	// Configure the serial connection
	config := &serial.Config{Name: serialDevice, Baud: 9600, ReadTimeout: time.Second * 2}
	var err error
	serialPort, err = serial.OpenPort(config)
	if err != nil {
		log.Fatalf("Failed to open serial port: %v", err)
	}

	log.Printf("Serial port %s opened successfully", serialDevice)
}

// Send command to Arduino and get response
func sendCommand(command string) (string, error) {
	_, err := serialPort.Write([]byte(command + "\n"))
	if err != nil {
		return "", err
	}

	buffer := make([]byte, 128)
	n, err := serialPort.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:n]), nil
}

// Handler for power on request
func PowerOnHandler(w http.ResponseWriter, r *http.Request) {
	response, err := sendCommand("POWER ON")
	if err != nil {
		http.Error(w, "Failed to power on node", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, response)
}

// Handler for power off request
func PowerOffHandler(w http.ResponseWriter, r *http.Request) {
	response, err := sendCommand("POWER OFF")
	if err != nil {
		http.Error(w, "Failed to power off node", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, response)
}

// Handler for reset request
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	response, err := sendCommand("RESET")
	if err != nil {
		http.Error(w, "Failed to reset node", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, response)
}
