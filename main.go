package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/skycoin/skycoin/src/aether/wifi"
)

func main() {
	if hasInternet() {
		log.Println("got internet. No need to connect on wifi")
		os.Exit(0)
	}
	ifs, err := network.NewWifiInterfaces()
	if err != nil {
		log.Fatal(err)
	}
	for i := range ifs {
		nets, err := ifs[i].Scan()
		if err != nil {
			log.Fatal(err)
		}
		for j := range nets {
			if nets[j].Address == "" {
				continue
			}

			//fmt.Printf("%#v\n", nets[j])
			if strings.ToLower(nets[j].EncryptionKeyStatus) != "on" {
				ifs[i].Connection.SSID = nets[j].ESSID
				ifs[i].Connection.DHCPEnabled = true
				ifs[i].Start()
				// TODO check for success and bail?
			}
		}
	}
}

func hasInternet() bool {
	resp, err := http.Head("http://www.google.com/")
	if err != nil {
		return false
	}
	resp.Body.Close()
	return true
}

/*
type WifiConnection struct {
    ConnectionName string
    InterfaceName  string
    //
    Mode             string
    SSID             string
    Channel          string
    Frequency        string
    SecurityProtocol string // [NONE, WEP, WPA]
    SecurityKey      string
    DHCPEnabled      bool
    Addresses        []Address
    Routes           []Route
    Nameservers      []net.IP
    DefaultGateway   net.IP
}
*/
