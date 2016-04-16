package main

import (

	"net"
	"os"
	"fmt"
)


func main() {

  fmt.Print(len(os.Args))
  if len(os.Args) !=2 {
	fmt.Fprintf(os.Stderr, "Usage: %s ip-address\n", os.Args[0])
	os.Exit(1)
  }
  name := os.Args[1]

  addr := net.ParseIP(name)

  if addr == nil {
	  fmt.Println("Invalid address")
  } else {
	  fmt.Println("The address is " , addr.String())
  }

  mask := addr.DefaultMask()
  network := addr.Mask(mask)
  ones, bits := mask.Size()

  fmt.Println("Default mask lenth  : ", bits)
  fmt.Println("Leadming ones count : ", ones)
  fmt.Println("hex: ", mask.String())
  fmt.Println("network: ", network.String())
  os.Exit(0)

}
