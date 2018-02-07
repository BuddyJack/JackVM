package cmd

import (
	"flag"
	"fmt"
)

/**
	定义命令行结构体

 */
type CmdInfo struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	Class       string
	JarOption   string
	Jar         string
	Args        []string
}

func ParseCmd() *CmdInfo {
	cmd := &CmdInfo{}
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print version,then exit")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.JarOption, "jar", "", "jar")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		if cmd.JarOption != "" {
			cmd.Jar = cmd.JarOption
			cmd.Args = args[0:]
		} else {
			cmd.Class = args[0]
			cmd.Args = args[1:]
		}
	}
	return cmd
}

func PrintUsage() {
	fmt.Println("用法: java [-options] class [args...]")
	fmt.Println("           (执行类)")
	fmt.Println("或者 java [-options] -jar jarfile [args...]")
	fmt.Println("           (执行 jar 文件)")
	fmt.Println("其中选项包含:")
	fmt.Println("       -cp <目录和 zip/jar 文件的类搜索路径>")
	fmt.Println("       -classpath <目录和 zip/jar 文件的类搜索路径> 用 : 分隔的目录, JAR 档案和 ZIP 档案列表, 用于搜索类文件。")
	fmt.Println("       -version      输出产品版本并退出")
	fmt.Println("有关详细信息, 请参阅 https://github.com/BuddyJack/JackVM")
}
