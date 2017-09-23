// The MIT License (MIT)

// Copyright (c) 2016 Jerry Bai

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"syscall"

	"github.com/berkaroad/saashard"
	"github.com/berkaroad/saashard/config"
	"github.com/berkaroad/saashard/server"
	"github.com/berkaroad/saashard/utils/simplelog"
)

var (
	configFile = flag.String("config", "/opt/saashard/conf/ss.yaml", "saashard config file")
	logLevel   = flag.String("log-level", "", "log level [debug|info|warn|error], default error")
	version    = flag.Bool("v", false, "the version of saashard")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	memprofile = flag.String("memprofile", "", "write memory profile to file")
)

const (
	sqlLogName = "sql.log"
	sysLogName = "sys.log"
	maxLogSize = 100 * 1024 * 1024
)

const banner string = `
                                     █████                              █████
                                    ░░███                              ░░███ 
  █████   ██████    ██████    █████  ░███████    ██████   ████████   ███████ 
 ███░░   ░░░░░███  ░░░░░███  ███░░   ░███░░███  ░░░░░███ ░░███░░███ ███░░███ 
░░█████   ███████   ███████ ░░█████  ░███ ░███   ███████  ░███ ░░░ ░███ ░███ 
 ░░░░███ ███░░███  ███░░███  ░░░░███ ░███ ░███  ███░░███  ░███     ░███ ░███ 
 ██████ ░░████████░░████████ ██████  ████ █████░░████████ █████    ░░████████
░░░░░░   ░░░░░░░░  ░░░░░░░░ ░░░░░░  ░░░░ ░░░░░  ░░░░░░░░ ░░░░░      ░░░░░░░░ 

`

func main() {
	fmt.Print(banner)
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	fmt.Printf("Git commit:%s\n", saashard.Version)
	fmt.Printf("Build time:%s\n", saashard.Compile)
	if *version {
		return
	}
	if len(*cpuprofile) != 0 {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	if len(*memprofile) != 0 {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
	if len(*configFile) == 0 {
		fmt.Println("must use a config file")
		return
	}

	cfg, err := config.ParseConfigFile(*configFile)
	if err != nil {
		fmt.Printf("parse config file error:%v\n", err.Error())
		return
	}

	var svr *server.Server
	svr, err = server.NewServer(cfg)
	if err != nil {
		simplelog.Error("%s %s %s", "main", "main", err.Error())
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGPIPE,
	)

	go func() {
		for {
			sig := <-sc
			if sig == syscall.SIGINT || sig == syscall.SIGTERM || sig == syscall.SIGQUIT {
				simplelog.Info("%s %s %s signal=%s", "main", "main", "Got signal", sig)
				svr.Close()
			} else if sig == syscall.SIGPIPE {
				simplelog.Info("%s %s %s", "main", "main", "Ignore broken pipe signal")
				signal.Ignore(sig)
			}
		}
	}()

	svr.Run()
}
