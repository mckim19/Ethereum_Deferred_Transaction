package main

import (
   "log"
   "net"
   "time"
)

func main() {
   conn, err := net.Dial("tcp", "166.104.144.107:8000")
   if nil != err {
      log.Println(err)
   }
	  // sending msg to server
	  var s string
	  s = "hello"
	  conn.Write([]byte(s))
	  time.Sleep(time.Duration(1)*time.Second)
	
	  // receiving msg from server
	  data := make([]byte, 4096)
	  n, err := conn.Read(data)
	  if err != nil {
		log.Println(err)
		return
	  }


	  log.Println("Server send : " + string(data[:n]))

}