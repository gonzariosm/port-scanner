package main

import (
	"github.com/gonzariosm/port-scanner/port"
)

func main() {
	port.GetOpenPorts("www.freecodecamp.com", port.PortRange{Start: 75, End: 85})
	port.GetOpenPorts("104.26.10.78", port.PortRange{Start: 8079, End: 8090})
	port.GetOpenPorts("104.26.10.78", port.PortRange{Start: 440, End: 450})
	port.GetOpenPorts("137.74.187.104", port.PortRange{Start: 440, End: 450})
	port.GetOpenPorts("scanme.nmap.org", port.PortRange{Start: 20, End: 80})
}
