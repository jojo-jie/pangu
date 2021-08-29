package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Uninstall old Docker versions",
	Long:  "Uninstall old Docker versions",
	Run: func(cmd *cobra.Command, args []string) {
		shellCmds:=make([]*exec.Cmd, 0, 4)
		shellCmds = append(shellCmds, exec.Command("/bin/sh","-c","sudo apt-get -y purge docker-ce docker-ce-cli containerd.io"))
		shellCmds = append(shellCmds, exec.Command("/bin/sh","-c","sudo rm -rf /var/lib/docker"))
		shellCmds = append(shellCmds, exec.Command("/bin/sh","-c","sudo rm -rf /var/lib/containerd"))
		shellCmds = append(shellCmds, exec.Command("/bin/sh","-c","apt-get remove docker docker-engine docker.io containerd runc"))
		for _, shellCmd := range shellCmds {
			fmt.Println("======******======")
			fmt.Println(shellCmd.Args)
			stdout, err := shellCmd.StdoutPipe()
			if err != nil {
				log.Fatalln(err)
			}
			shellCmd.Start()
			scan:=bufio.NewScanner(stdout)
			for scan.Scan() {
				fmt.Println(scan.Text())
			}
			shellCmd.Wait()
		}
	},
}




