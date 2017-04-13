package main

import (
	"flag"
	"io/ioutil"
	"log"
	"regexp"
	"github.com/tomasen/fcgi_client"
	"fmt"
	"os"
	"bufio"
	"strings"
	"bytes"
)

var (
	statusLineRegexp = regexp.MustCompile(`(?m)^(.*):\s+(.*)$`)
	FCGI_ADDR = ""
	FCGI_PROT = ""
)

func GetPHPStatus() (body []byte, err error) {
	fcgi, err := fcgiclient.Dial(FCGI_PROT, FCGI_ADDR)
	if err != nil {
		return
	}
	defer fcgi.Close()

	env := make(map[string]string)
	env["SCRIPT_NAME"] = "/status"
	env["SCRIPT_FILENAME"] = "/status"
	env["QUERY_STRING"] = ""
	env["REQUEST_METHOD"] = "GET"

	resp, err := fcgi.Get(env)

	body, err = ioutil.ReadAll(resp.Body)

	return body, err
}

func PrintPHPStatus(phpAddr string) {
	if phpAddr == "" {
		log.Fatal("The php-fpm-addr flag is required.")
	} else if strings.Contains(phpAddr, "tcp://") {
		FCGI_PROT = "tcp"
		FCGI_ADDR = strings.Replace(phpAddr, "tcp://", "", 1)
	} else if strings.Contains(phpAddr, "unix:") {
		FCGI_PROT = "unix"
		FCGI_ADDR = strings.Replace(phpAddr, "unix:", "", 1)
	} else {
		log.Fatal("The php-fpm-addr must like tcp://127.0.0.0.1:9001 or unix:/path/to/unix/socket")
	}

	var w bytes.Buffer
	body, err := GetPHPStatus()

	if err != nil {
		NewMetricsFromMatches([][]string{}).WriteTo(&w, FCGI_ADDR)
	} else {
		matches := statusLineRegexp.FindAllStringSubmatch(string(body), -1)
		NewMetricsFromMatches(matches).WriteTo(&w, FCGI_ADDR)
	}
	fmt.Println(w.String())

}

func main() {
	filePath := flag.String("php-fpm-list", "", "PHP-FPM list file path")
	flag.Parse()

	if *filePath == "" {
		log.Fatal("The php-fpm-list flag is required.")
	}

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		PrintPHPStatus(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

