package sys

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type CPUUsage struct {
	Average int `json:"avg"`
}

type Sys struct {
	CPU  CPUUsage
	User *user.User
}

func _GetCPUUsage() *CPUUsage {
	percents, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &CPUUsage{
		Average: int(math.Round(percents[0])),
	}
}

func _GetCurrentUser() *user.User {
	user, err := user.Current()
	if err != nil {
		log.Println(err)
		return nil
	}

	return user

}

func (s *Sys) All() *Sys {
	user := _GetCurrentUser()
	cpuUsage := _GetCPUUsage()
	return &Sys{
		CPU:  *cpuUsage,
		User: user,
	}
}

func (s *Sys) Exec(command string) {
	cmd := exec.Command(command)
	err := cmd.Run()

	if err != nil {
		log.Println(err)
	}
}

func (s *Sys) GetRemoteFile(fileName string, remoteUrl string) {
	fmt.Println("Downloading: " + remoteUrl)
	err := DownloadFile(fileName, remoteUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + remoteUrl)
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
