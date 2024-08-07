package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const subsDirectory = "./subs"
const fragmentDurationSec = 10


func generateSubtitle(startTime int, endTime int, text string, pad *bool) string {
	timeStr := printTimeRange(startTime, endTime)
	padStr := func() string { if *pad { return "\n" } else {  return  "" }} ()
	subtitle := padStr + fmt.Sprintf("WEBVTT\n\n%s\n%s", timeStr, text)
	return subtitle
}


func printTimeRange(startTime int, endTime int) string {
	hours := startTime / 3600
	minutes := (startTime % 3600) / 60
	seconds := startTime % 60
	formattedStartTime := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	hours = endTime / 3600
	minutes = (endTime % 3600) / 60
	seconds = endTime % 60
	formattedEndTime := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)

	return fmt.Sprintf("%s.000 --> %s.000", formattedStartTime, formattedEndTime)
}


func main() {
	
	flag.Usage = func() {
		w := flag.CommandLine.Output() 
	
		fmt.Fprintf(w, "Usage of %s: \n", os.Args[0])
	
		flag.PrintDefaults()
	}
	    
	mode := flag.String("mode", "file", `Specify a mode of writing subtitles:
		 'file' writes each sub to a separate file, 
		 'stdout' writes each to stdout, 
		 'udp' writes each to udp connection, see udpport setting
		 'pipe' writes to mkfifo, see fifoname setting. 
		 'listen'  runs UDP listener to debug UDP output
		 Default is 'file'`)
	
	userDefinedText := flag.String("text", "", `User defined text for subtitles, default is number of current subtitle.
		You can use {s} to substitute current time and {n} for current subtitle number`)

	udpPort := flag.Int("udpport", 9723, "Specify a port for udp, default is 9723")

	textDur := flag.Int("textdur", 10, "Specify a text duration in subtitle, default is 10")

	offset := flag.Int("offset", 0, "Specify an offset for timecodes, default is 0")

	delay := flag.Int("delay", 10, "Specify a delay in seconds between each subtitle, default is 10")

	//@todo intervals := flag.Int("intervals", 1, "Specify quantity of intervals in given subtitle")

	pad := flag.Bool("pad", false, "Whether to add \\n newline  before new subtitle")
	
	incrementTimeCodes := flag.Bool("inc", true, "If true, timecodes in subtitles will incrementing to textduration, default is true")

	fifoName := flag.String("fifoname", "/tmp/subtitle_pipe", `Specify a name for mkfifo named pipe, 
		default is 'subtitle_pipe' in the current folder`)

	flag.Parse()

	c := 0

	port := strconv.Itoa(*udpPort)
	udpAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:"+port)

	if _, err := os.Stat(subsDirectory); os.IsNotExist(err) {
		os.Mkdir(subsDirectory, 0755)
	}
	from := *offset
	to := *textDur+*offset

	for {		

		if (*userDefinedText != "") {
			currentTime := time.Now()
			formattedTime := currentTime.Format("2006-01-02 15:04:05")
			output := strings.Replace(*userDefinedText, "{s}", formattedTime, -1)
			output  = strings.Replace(output, "{n}", strconv.Itoa(c), -1)
		}
		subs := generateSubtitle(from, to, fmt.Sprintf("Subtitle piece number %d", c), pad)

		c=c+1
		if (*incrementTimeCodes) {
			from = from+*textDur+*offset
			to  = from+*textDur
		}

		fileName := fmt.Sprintf("%s/subtitle_%d.vtt", subsDirectory, c)
		if (*mode == "file") {
			err := os.WriteFile(fileName, []byte(subs), 0755)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				continue
			}
		}
		if (*mode == "stdout") {
			fmt.Println(subs)
		}

		if (*mode == "pipe") {

			f, err := os.OpenFile(*fifoName, os.O_WRONLY, os.ModeNamedPipe)
			if os.IsNotExist(err) {
				fmt.Printf("Named pipe need to create")
				if err := syscall.Mkfifo(*fifoName, 0600); err != nil {
					panic(err)
				}
			}
			if err != nil {
				panic(err)
			}
			_, err = f.WriteString(subs+"\n")
			if err != nil {
				panic(err)
			}

			fmt.Println("Written to named pipe")
			_ = f.Close()			
		}

		if (*mode == "udp") {		
				// Send text over UDP
			conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
				IP:   net.ParseIP("127.0.0.1"),
				Port: *udpPort,
			})
			if err != nil {
				fmt.Println("Error opening UDP Port:", err)
				return
			}
			defer conn.Close()
			
			_, err = conn.Write([]byte(subs))
			if err != nil {
				fmt.Println("Error sending string to UDP port:", err)
				return
			}
		
		}

		if (*mode == "listen") {			
			conn, err := net.ListenUDP("udp", udpAddr)
			if err != nil {
				fmt.Println("Error listen UDP:", err)
				return
			}

			defer conn.Close()
			

			fmt.Println("UDP server listening on", conn.LocalAddr().String())
			for {

				// Receive and handle incoming UDP packets
				buf := make([]byte, 1024)
				for {
					n, addr, err := conn.ReadFromUDP(buf)
					if err != nil {
						fmt.Println("Error:", err)
						continue
					}

					fmt.Printf("Received %d bytes from %s\n", n, addr.String())
					fmt.Println("Text:", string(buf[:n]))
				}
			}
		}

		time.Sleep(time.Duration(*delay) * time.Second)		
	}
}