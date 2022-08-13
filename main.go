package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	device "github.com/d2r2/go-hd44780"
	"github.com/d2r2/go-i2c"
)

const ( // Define constants
	speedtest_file = "/tmp/speedtest.txt" // "Down: 12.34 Mbit/s"
	weather_file   = "/tmp/weather.txt"   // "Temp: 12.34"
	status_file    = "/tmp/status.txt"    // "1" or "0"
	Address        = 0x27                     // I2C address of LCD
)

func checkError(err error) { // Used for giving error messages
	if err != nil {
		log.Fatal(err)
	}
}

func file_read(file string, default_text string) string { // Will read the file, if it is empty it will return the default text given
	text_byte, err := ioutil.ReadFile(file)
	checkError(err)
	text := string(text_byte)
	stats, err := os.Stat(file)
	checkError(err)

	if stats.Size() == 0 { // If the file is empty, it will return the default text
		return default_text
	} else { // If the file is not empty, it will return the text from the file
		return text
	}

}

func printmodificationdate(file string) int64 { // Prints the last modification date of a file
	fileinfo, err := os.Stat(file)
	checkError(err)
	return fileinfo.ModTime().Unix()
}

func checkStatus() bool { // Checks if the status file is 1 or 0
	status_current_byte, err := ioutil.ReadFile(status_file)
	checkError(err)
	status_current := string(status_current_byte)

	return strings.Contains(status_current, "1")

}

func main() {

	var last_time_speed, last_time_weather int64 = 0, 0

	// initialize LCD
	i2c, err := i2c.NewI2C(Address, 1)
	checkError(err)
	lcd, err := device.NewLcd(i2c, device.LCD_20x4)

	checkError(err)
	lcd.Clear()

	err = lcd.BacklightOn()
	checkError(err)
	fmt.Println("INIT DONE")
	defer i2c.Close()

	for {
		for time.Now().Hour() < 8 { // Sleep till 8 o'clock
			time.Sleep(time.Second * 5)
			lcd.BacklightOff()
			continue
		}
		if !checkStatus() { // If the status is 0, the program will sleep
			time.Sleep(time.Second * 1)
			lcd.BacklightOff()
			for {
				if checkStatus() {
					break
				}
			}
		}
		lcd.BacklightOn()

		err = lcd.ShowMessage(time.Now().Format("Mon 02/01/2006"), device.SHOW_LINE_1)
		checkError(err)
		err = lcd.ShowMessage(time.Now().Format("15:04:05"), device.SHOW_LINE_2)
		checkError(err)

		if printmodificationdate(speedtest_file) != last_time_speed {
			last_time_speed = printmodificationdate(speedtest_file)

			speedtest_current := file_read(speedtest_file, "Down: 0.00 Mbit/s")
			lcd.ShowMessage("                    ", device.SHOW_LINE_3)

			err = lcd.ShowMessage(string(speedtest_current), device.SHOW_LINE_3)
			checkError(err)

		}
		if printmodificationdate(weather_file) != last_time_weather {
			last_time_weather = printmodificationdate(weather_file)

			weather_current := file_read(weather_file, "0C NULL")
			lcd.ShowMessage("                    ", device.SHOW_LINE_4)

			err = lcd.ShowMessage(string(weather_current), device.SHOW_LINE_4)
			checkError(err)
		}

		time.Sleep(time.Millisecond * 200)
	}
}
