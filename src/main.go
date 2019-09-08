package main

import (
	"flag"
	"fmt"
	"github.com/nhalstead/ftp"
	"os"
	"strings"
	"time"
)

func main() {
	var (
		host = flag.String("host", "", "FTP Host, host:port")
		username = flag.String("username", "anonymous", "Username to connect to as, Default is `anonymous`")
		password = flag.String("password", "anonymous", "Password to connect to with, Default is `anonymous`")
		timeout = flag.Duration("timeout", 10*time.Second, "Timeout, Default is 10 Seconds")
	)
	flag.Parse()

	if *host == "" {
		// Closed as it has NO HOST to connect to.
		os.Exit(2)
	}

	// https://github.com/jlaffaye/ftp/blob/master/ftp.go#L195
	c, err := ftp.Dial(*host, ftp.DialWithTimeout(*timeout))

	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	err = c.Login(*username, *password)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	// Do that command to get the status of the storage.
	// SITE QUOTA
	e, err := c.Quotas()
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	// Respond with the Storage Quota, If it can be found.
	if val, ok := e["UploadedBytes"]; ok {
		s := strings.Split(val, "/") // Get just the total storage (in bytes)
		if s[1] != "" {
			fmt.Print(s[1]) // Select Part 2, Total Storage
		} else {
			fmt.Print(s[0]) // Select First Part as a backup
		}
	} else {
		fmt.Print("0")
	}

	if err := c.Quit(); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

}
