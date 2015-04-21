package main

import (
	"encoding/xml"
	"github.com/scottdware/go-junos"
	"log"
	"fmt"
	"flag"

)

type vplsMACTable struct {
	L2aldEntries []L2aldMACEntry `xml:"l2ald-mac-entry"`
}

type L2aldMACEntry struct {
	L2RoutingInstance     string   `xml:"l2-mac-routing-instance"`
	L2BridgingDomain      string   `xml:"l2-mac-bridging-domain"`
	L2BridgeVlan          string   `xml:"l2-bridge-vlan"`
	L2MACAddrs            []string `xml:"l2-mac-address"`
	L2MACFlags            []string `xml:"l2-mac-flags"`
	L2MACLogicalInterface []string `xml:"l2-mac-logical-interface"`
}

var hostname = flag.String("hostname", "", "Hostname of the device you wish to connect to")

var user = flag.String("user", "", "Username to use when connecting to the device")

var pass = flag.String("pass", "", "Password associated with chosen username")

func main() {

	flag.Parse()

   // Establish our session first.
    jnpr, err := junos.NewSession(*hostname, *user, *pass)
    if err != nil {
        log.Fatal(err)
    }
    defer jnpr.Close()

	var vpls vplsMACTable
	contents, err := jnpr.RunCommand("show vpls mac-table", "xml")
    if err != nil {
    // handle error
    fmt.Println(err)
    }
	
	err = xml.Unmarshal([]byte(contents), &vpls)
	if err != nil {
		fmt.Println(err)
	}

	for _, table := range vpls.L2aldEntries {
		fmt.Printf("--- VPLS MAC Table ---\n")
		fmt.Printf("Routing Instance: %s\n", table.L2RoutingInstance)
		fmt.Printf("Bridge Domain: %s\n", table.L2BridgingDomain)
		fmt.Printf("Bridge VLAN: %s\n", table.L2BridgeVlan)
		for i := 0; i < len(table.L2MACAddrs); i++ {
			fmt.Printf("--> MAC: %s\n", table.L2MACAddrs[i])
		}
		fmt.Println()
	}
}

