package main

import (
   "io"
   "log"
   "net"
)

func main() {
   l, err := net.Listen("tcp", ":8000")
   if nil != err {
      log.Println(err);
   }
   defer l.Close()

   for {
      conn, err := l.Accept()
      if nil != err {
         log.Println(err);
         continue
      }
      defer conn.Close()
      go ConnHandler(conn)
   }
}

func ConnHandler(conn net.Conn) {
   recvBuf := make([]byte, 4096)
   for {
      n, err := conn.Read(recvBuf)
      if nil != err {
         if io.EOF == err {
            log.Println(err);
            return
         }
         log.Println(err);
         return
      }
      if 0 < n {
         data := recvBuf[:n]
         log.Println(string(data))
         _, err = conn.Write(data[:n])
         if err != nil {
            log.Println(err)
            return
         }
      }
   }
}