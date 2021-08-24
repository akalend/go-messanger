package main

import (
  "bufio"
  "fmt"
  "net"
  "strings" // требуется только ниже для обработки примера
)

func main() {

  fmt.Println("Start server...")

  // Устанавливаем прослушивание порта
  ln, _ := net.Listen("tcp", ":8081")

  // Принимаем входящее соединение
  conn, _ := ln.Accept()

  // Запускаем цикл обработки входящих данных
  for {
    // Будем прослушивать все сообщения разделенные \n
    message, _ := bufio.NewReader(conn).ReadString('\n')
    // Распечатываем полученое сообщение
    fmt.Print("Message Received:", string(message))
    // Отправить новую строку обратно клиенту
    conn.Write([]byte(message + "\n"))
  }
}
