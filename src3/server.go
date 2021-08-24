package main

import (
  "bufio"
  "fmt"
  "net"
)

// функция запускается как горутина
func process(conn net.Conn) {
  // определим, что перед выходом из функции, мы закроем соединение
  defer conn.Close()
  for {
    // Будем прослушивать все сообщения разделенные \n
    message, _ := bufio.NewReader(conn).ReadString('\n')
    // Распечатываем полученое сообщение
    fmt.Print("Message Received:", string(message))
    // Отправить новую строку обратно клиенту
    _, err := conn.Write([]byte(message + "\n"))
    // анализируем на ошибку
    if err != nil {
      fmt.Print(err.Error())
      break // выходим из цикла
    }

  }
}

func main() {

  fmt.Println("Start server...")

  // Устанавливаем прослушивание порта
  ln, _ := net.Listen("tcp", ":8081")

  // Запускаем цикл обработки соединений
  for {
    // Принимаем входящее соединение
    conn, _ := ln.Accept()
    // запускаем функцию process(conn)   как горутину
    go process(conn)
  }
}
