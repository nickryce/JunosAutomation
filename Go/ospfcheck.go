package main

import (
	"encoding/xml"
	"github.com/scottdware/go-junos"
	"log"
	"fmt"
	"flag"

)

//Define some structures

type OSPFOutput struct {
	OSPFEntries []OSPFNeighborEntry `xml:"ospf-neighbor"`
}
type OSPFNeighborEntry struct {
	NeighborAddress 	string `xml:"neighbor-address"`
	InterfaceName 		string `xml:"interface-name"`
	OSPFState   		string `xml:"ospf-neighbor-state"`
	NeighborID   		string `xml:"neighbor-id"`
	NeighborPriority	string `xml:"neighbor-priority"`
	ActivityTimer		string `xml:"activity-timer"`
}


var hostname = flag.String("hostname", "", "Hostname of the device you wish to connect to")

var user = flag.String("user", "", "Username to use when connecting to the device")

var pass = flag.String("pass", "", "Password associated with chosen username")

func main() {

	//Parse the flags set above
	flag.Parse()

   // Establish our session first.
    jnpr, err := junos.NewSession(*hostname, *user, *pass)
    if err != nil {
        log.Fatal(err)
    }
    defer jnpr.Close()

    var ospf OSPFOutput
    contents, err := jnpr.RunCommand("show ospf neighbor", "xml")
    if err != nil {
    // handle error
    fmt.Println(err)
    }

    err = xml.Unmarshal([]byte(contents), &ospf)
	if err != nil {
		fmt.Println(err)
	}

	for _, table := range ospf.OSPFEntries {
		fmt.Printf("---OSPF Neighbors Seen---\n")
		fmt.Printf("   Interface Name: %s\n", table.InterfaceName)
		fmt.Printf("   Neighbor Address: %s\n", table.NeighborAddress)
		fmt.Printf("   Neighor State: %s\n", table.OSPFState)
	
		fmt.Println()
	}
}