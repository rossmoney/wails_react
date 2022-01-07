package sys

import (
	"log"
	"math"
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

	// Current User
	/*fmt.Println("Hi " + user.Name + " (id: " + user.Uid + ")")
	  fmt.Println("Username: " + user.Username)
	  fmt.Println("Home Dir: " + user.HomeDir)*/

}

func (s *Sys) All() *Sys {
	user := _GetCurrentUser()
	cpuUsage := _GetCPUUsage()
	return &Sys{
		CPU:  *cpuUsage,
		User: user,
	}
}
