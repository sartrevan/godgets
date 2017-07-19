package main

import "fmt"
import "strconv"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (v IPAddr) String() string {
	return fmt.Sprintf("%s.%s.%s.%s", strconv.Itoa(int(v[0])), strconv.Itoa(int(v[1])), strconv.Itoa(int(v[2])), strconv.Itoa(int(v[3])))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, &ip)
	}
}
