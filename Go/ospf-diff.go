package main

import (
	"github.com/scottdware/go-junos"
	"encoding/xml"	
	"log"
	"fmt"
)

//Define some structures

type OSPFConfigOutput struct {
	OSPFConfig []OSPFNeighborConfig `xml:"protocols>ospf>area>interface"`
}
type OSPFNeighborConfig struct {
	InterfaceName string `xml:"name"`
}
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

func main() {
	// Establish our session first.
    jnpr, err := junos.NewSession("an1.wav-edi.fluency.net.uk", "nick", "bn6VDSyn")
    if err != nil {
        log.Fatal(err)
    }
    defer jnpr.Close()

    var ospfconfig OSPFConfigOutput
    config, err := jnpr.GetConfig("protocols>ospf>area", "xml")
	if err != nil {
		fmt.Println(err)
	}
	
	var ospf OSPFOutput
    contents, err := jnpr.RunCommand("show ospf neighbor", "xml")
    if err != nil {
    // handle error
    fmt.Println(err)
    }

	err = xml.Unmarshal([]byte(config), &ospfconfig)
	if err != nil {
		fmt.Println(err)
	}


	fmt.Printf("The below interfaces are configured for ospf\n")
	for _, table1 := range ospfconfig.OSPFConfig {
		if table1.InterfaceName != "lo0.0" {
			fmt.Printf("   Interface Name: %s\n", table1.InterfaceName)
		}
		
	}
	err = xml.Unmarshal([]byte(contents), &ospf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("The below interfaces have an active ospf neighbor\n")
	for _, table := range ospf.OSPFEntries {
		fmt.Printf("   Interface Name: %s", table.InterfaceName)
	
		fmt.Println()
	}
}