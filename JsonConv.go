package main

import (
	"fmt"
	"os"
	"os/exec"
	//"os/exec"
)

type ethiface struct {
	ethName string
	ethType string
	vlanID  int
	enabled bool
	vlanTag bool
}

func main() {
	var iface [5]ethiface
	var i int

	iface = [5]ethiface{
		{"eth0", "ianaift:ethernetCsmacd", 0, false, false},
		{"eth1", "ianaift:ethernetCsmacd", 125, true, true},
		{"lo1", "ianaift:softwareLoopback", 0, true, false},
	}

	f, err := os.Create("ietf-interface.xml")
	_, err = f.WriteString("<interfaces xmlns=\"urn:ietf:params:xml:ns:yang:ietf-interfaces\" xmlns:ianaift=\"urn:ietf:params:xml:ns:yang:iana-if-type\" xmlns:vlan=\"http://example.com/vlan\">\n")
	/* output each array element's value */
	for i = 0; i < 3; i++ {
		fmt.Printf("a[%d][%d] = %s\n", i, i+1, iface[i].ethName)
		_, err = f.WriteString("  <interface>\n")
		_, err = f.WriteString("    <name>")
		_, err = f.WriteString(iface[i].ethName)
		_, err = f.WriteString("</name>\n")
		_, err = f.WriteString("    <type>")
		_, err = f.WriteString(iface[i].ethType)
		_, err = f.WriteString("</type>\n")
		if iface[i].enabled == true {
			_, err = f.WriteString("    <enabled>true</enabled>\n")
		} else {
			_, err = f.WriteString("    <enabled>false</enabled>\n")
		}
		if iface[i].vlanTag == true {
			_, err = f.WriteString("    <vlan:vlan-tagging>true</vlan:vlan-tagging>\n")
			_, err = f.WriteString("  </interface>\n")
			_, err = f.WriteString("  <interface>\n")
			_, err = f.WriteString("    <name>")
			_, err = f.WriteString(iface[i].ethName)
			_, err = f.WriteString(".")
			_, err = f.WriteString(fmt.Sprintf("%d", iface[i].vlanID))
			_, err = f.WriteString("</name>\n")
			_, err = f.WriteString("    <type>ianaift:l2vlan</type>\n")
			_, err = f.WriteString("    <enabled>true</enabled>\n")
			_, err = f.WriteString("    <vlan:base-interface>")
			_, err = f.WriteString(iface[i].ethName)
			_, err = f.WriteString("</vlan:base-interface>\n")
			_, err = f.WriteString("    <vlan:vlan-id>")
			_, err = f.WriteString(fmt.Sprintf("%d", iface[i].vlanID))
			_, err = f.WriteString("</vlan:vlan-id>\n")
		}
		_, err = f.WriteString("  </interface>\n")
	}
	_, err = f.WriteString("</interfaces>")
	fmt.Printf("err %T\n", err)

	out, err := exec.Command("mv", "ietf-interface.xml", "libyang/build").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)

	_, _ = exec.Command("/usr/bin/expect", "-f", "./script.exp").Output()
	//output := string(out[:])
	//fmt.Println(output)
}
