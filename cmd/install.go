package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Docker installation and environment initialization",
	Long:  "Docker installation and environment initialization",
	Run: func(cmd *cobra.Command, args []string) {
		shellCmds:=make([]*exec.Cmd, 0, 9)
		shellCmds = append(shellCmds, exec.Command("/bin/sh","-c","sudo apt-get update"))
		shellCmds = append(shellCmds, exec.Command("/bin/sh","-c","sudo apt-get -y install apt-transport-https ca-certificates curl gnupg lsb-release"))
		shellCmds = append(shellCmds, exec.Command("/bin/sh","-c","curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -"))
		shellCmds = append(shellCmds, exec.Command("/bin/sh","-c","sudo add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\""))
		shellCmds = append(shellCmds, exec.Command("/bin/sh","-c","sudo apt-get update"))
		shellCmds = append(shellCmds, exec.Command("/bin/sh","-c","sudo apt install docker-ce docker-ce-cli containerd.io"))
		shellCmds = append(shellCmds, exec.Command("/bin/sh","-c","sudo systemctl status docker"))
		shellCmds = append(shellCmds, exec.Command("/bin/sh","-c","q"))
		shellCmds = append(shellCmds, exec.Command("/bin/sh","-c","sudo usermod -aG docker $USER"))
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
