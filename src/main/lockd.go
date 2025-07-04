<<<<<<< HEAD
package main

// export GOPATH=~/6.5840
// go build lockd.go
// go build lockc.go
// ./lockd -p a b &
// ./lockd -b a b &
// ./lockc -l a b lx
// ./lockc -u a b lx
//
// on Athena, use /tmp/myname-a and /tmp/myname-b
// instead of a and b.

import "time"
import "6.5840/lockservice"
import "os"
import "fmt"

func main() {
	if len(os.Args) == 4 && os.Args[1] == "-p" {
		lockservice.StartServer(os.Args[2], os.Args[3], true)
	} else if len(os.Args) == 4 && os.Args[1] == "-b" {
		lockservice.StartServer(os.Args[2], os.Args[3], false)
	} else {
		fmt.Printf("Usage: lockd -p|-b primaryport backupport\n")
		os.Exit(1)
	}
	for {
		time.Sleep(100 * time.Second)
	}
}
=======
package main

// export GOPATH=~/6.5840
// go build lockd.go
// go build lockc.go
// ./lockd -p a b &
// ./lockd -b a b &
// ./lockc -l a b lx
// ./lockc -u a b lx
//
// on Athena, use /tmp/myname-a and /tmp/myname-b
// instead of a and b.

import "time"
import "6.5840/lockservice"
import "os"
import "fmt"

func main() {
	if len(os.Args) == 4 && os.Args[1] == "-p" {
		lockservice.StartServer(os.Args[2], os.Args[3], true)
	} else if len(os.Args) == 4 && os.Args[1] == "-b" {
		lockservice.StartServer(os.Args[2], os.Args[3], false)
	} else {
		fmt.Printf("Usage: lockd -p|-b primaryport backupport\n")
		os.Exit(1)
	}
	for {
		time.Sleep(100 * time.Second)
	}
}
>>>>>>> 7a5641aae21a48923e23a4a1d556bcb3bf4607a5
