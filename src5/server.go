package main

import (
  "fmt"
  "net"
  "strings"
)

// функция запускается как горутина
func process(conns map[int]net.Conn, n int) {
  var clientNo int
  buf := make([]byte, 256)
  // получаем доступ к текущему соединению
  conn := conns[n]
  // определим, что перед выходом из функции, мы закроем соединение
  fmt.Println("Accept cnn:", n)
  defer conn.Close()

  for {
    readed_len, err := conns[n].Read(buf)
    if err != nil {
      if err.Error() == "EOF" {
        fmt.Println("Close ", n)
        delete(conns, n)
        break
      }
      fmt.Println(err)
    }

    // Распечатываем полученое сообщение
    // fmt.Println("Received Message:", read_len, buf)
    var message = ""
      message = string(buf[:readed_len])

    // парсинг полученного сообщения
    _, err = fmt.Sscanf(message, "%d", &clientNo) // определи номер клиента
    if err != nil {
      // обработка ошибки формата
      conn.Write([]byte("error format message\n"))
      continue
    }
    pos := strings.Index(message, " ") // нашли позицию разделителя

    if pos > 0 {
      out_message := message[pos+1:] // отчистили сообщение от номера клиента
      // Распечатываем полученое сообщение

      // if buf[0] == 0 {
      conn = conns[clientNo]
      if conn == nil {
        conns[n].Write([]byte("client is close"))
        continue
      }

      // }
      out_buf := []byte(fmt.Sprintf("%d->>%s\n", clientNo, out_message))

       // Отправить новую строку обратно клиенту
      _, err2 := conn.Write(out_buf)

      // анализируем на ошибку
      if err2 != nil {
        fmt.Println("Error:", err2.Error())

        break // выходим из цикла
      }
    }

  }
}

func main() {

  fmt.Println("Start server...")
  // создаем пул соединений
  conns := make(map[int]net.Conn, 1024)
  i := 0

  // Устанавливаем прослушивание порта
  ln, _ := net.Listen("tcp", ":8081")
  // Запускаем цикл обработки соединений
  for {
    // Принимаем входящее соединение
    conn, _ := ln.Accept()
    // сохраняем соединение в пул
    conns[i] = conn
    // запускаем функцию process(conn)   как горутину
    go process(conns, i)
    i++
  }
}
