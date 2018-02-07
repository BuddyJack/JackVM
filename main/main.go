package main

import (
	"JackVM/cmd"
	"fmt"
)

func main() {
	command := cmd.ParseCmd()
	if command.VersionFlag {
		fmt.Println("java version \"0.0.0_00\" by go builded")
	} else if command.HelpFlag || (command.Class == "" && command.JarOption == "") {
		cmd.PrintUsage()
	} else {
		startJvm(command)
	}
}

func startJvm(cmd *cmd.CmdInfo) {
	if cmd.Class == "" {
		fmt.Printf("jar:%s args:%v \n", cmd.JarOption, cmd.Args)
	} else {
		fmt.Printf("classpath:%s class:%s args:%v \n", cmd.CpOption, cmd.Class, cmd.Args)
	}
}
