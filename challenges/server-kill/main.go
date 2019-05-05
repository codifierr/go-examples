package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {

	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		errors.Wrap(err, "can't open pidfile")
	}

	if err := os.Remove(pidFile); err != nil {

		log.Printf("warning: can't remove pid file - %s\n", err)

	}

	strPid := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(strPid)

	if err != nil {
		errors.Wrap(err, "bad process id")
	}

	//simulate kill
	fmt.Printf("Killing server with pid : %d\n", pid)

	return nil

}

func main() {
	if err := killServer("server.pid"); err != nil {

		//fmt.FPrintf(os.Stderr, "error : %s\n", err)
		os.Exit(1)

	}
}
