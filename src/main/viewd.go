<<<<<<< HEAD
package main

//
// see directions in pbc.go
//

import "time"
import "6.5840/viewservice"
import "os"
import "fmt"

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: viewd port\n")
		os.Exit(1)
	}

	viewservice.StartServer(os.Args[1])

	for {
		time.Sleep(100 * time.Second)
	}
}
=======
package main

//
// see directions in pbc.go
//

import "time"
import "6.5840/viewservice"
import "os"
import "fmt"

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: viewd port\n")
		os.Exit(1)
	}

	viewservice.StartServer(os.Args[1])

	for {
		time.Sleep(100 * time.Second)
	}
}
>>>>>>> 7a5641aae21a48923e23a4a1d556bcb3bf4607a5
