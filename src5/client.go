package main

import (
  "bufio"
  "fmt"
  "net"
  "os"
  "strings"
  "time"
)

// прием данных из сокета и вывод на печать
func readSock(conn net.Conn) {

  if conn == nil {
    panic("Connection is nil")
  }
  buf := make([]byte, 256)
  eof_count := 0
  for {
    // чистим буфер
    for i := 0; i < 256; i++ {
      buf[i] = 0
    }

    readed_len, err := conn.Read(buf)
    if err != nil {
      if err.Error() == "EOF" {
        eof_count++
        time.Sleep(time.Second * 2)
        fmt.Println("EOF")
        if eof_count > 7 {

          fmt.Println("Timeout connection")
          break
        }
        continue
      }
      if strings.Index(err.Error(), "use of closed network connection") > 0 {

        fmt.Println("connection not exist or closed")
        continue
      }
      panic(err.Error())
    }
    if readed_len > 0 {
      fmt.Println(string(buf))
    }

  }
}

// ввод данных с консоли и вывод их в канал
func readConsole(ch chan string) {
  for {
    line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    if len(line) > 250 {
      fmt.Println("Error: message is very lagre")
      continue
    }
    fmt.Print(">")
    out := line[:len(line)-1] // убираем символ возврата каретки
    ch <- out                 // отправляем данные в канал
  }
}

func main() {
  ch := make(chan string)

  defer close(ch) // закрываем канал при выходе из программы

  conn, _ := net.Dial("tcp", "127.0.0.1:8081")
  if conn == nil {
    panic("Connection is nil")

  }

  go readConsole(ch)
  go readSock(conn)

  for {
    val, ok := <-ch
    if ok { // если есть данные, то их пишем в сокет
      // val_len := len(val)
      out := []byte(val)
      _, err := conn.Write(out)
      if err != nil {
        fmt.Println("Write error:", err.Error())
        break
      }

    } else {
      // данных в канале нет, задержка на 2 секунды
      time.Sleep(time.Second * 2)
    }

  }
  fmt.Println("Finished...")

  conn.Close()
}
