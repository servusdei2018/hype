package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/servusdei2018/hype/pkg/config"
	"github.com/servusdei2018/hype/pkg/futil"
	"github.com/servusdei2018/hype/pkg/initialize"
	"github.com/servusdei2018/hype/pkg/post"
	"github.com/servusdei2018/hype/pkg/server"
)

const VERSION = `0.1.0`
const USAGE = `Hype is a tool for publishing plaintext blogs.

Usage:

	hype <command>
			
The commands are:

	init		initialize empty Hype repository
	post		create Hype post
	serve		serve Hype repository
	version		print Hype version`

var conf = new(config.Conf)
var initialized = false

func init() {
	if initialize.IsInit() {
		initialized = true
		if err := conf.Load(); err != nil {
			panic(err)
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s\n\n", USAGE)
		return
	}

	switch os.Args[1] {
	case "init":
		if err := initialize.Init(VERSION); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("initialized empty hype repository in %s\n", futil.WD)
	case "post":
		if !initialized {
			fmt.Println("fatal: not a hype repository (run hype init)")
			return
		}
		var err error
		var path string
		if err, path = post.NewPost(conf); err != nil {
			if err.Error() != "cancelled" {
				fmt.Println(err)
				return
			} else {
				fmt.Println("new post cancelled")
				return
			}
		}
		fmt.Printf("new post created in %s\n", path)
	case "serve":
		if !initialized {
			fmt.Println("fatal: not a hype repository (run hype init)")
			return
		}
		fmt.Printf("serving hype at 0.0.0.0:%s [Ctrl+C to Exit]\n", conf.Port)
		if err := server.Serve(conf); err != nil {
			fmt.Println(err)
			return
		}
	case "version":
		fmt.Printf("hype version %s %s/%s\n", VERSION, runtime.GOOS, runtime.GOARCH)
	default:
		fmt.Printf("%s\n\n", USAGE)
	}
}
