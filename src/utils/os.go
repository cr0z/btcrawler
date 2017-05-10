/**
 * Created by 93201 on 2017/5/10.
 */
package utils

import (
	"os"
	"os/signal"
	"syscall"
	"io/ioutil"
	"strconv"
)

func HoldSignal(f func(os.Signal)) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	s := <-c
	signal.Stop(c)
	defer close(c)
	f(s)

}

func SavePID(path string) error {
	return ioutil.WriteFile(path, []byte(strconv.Itoa(os.Getpid())), os.ModePerm)
}
