package main

import (
  "bufio"
  "fmt"
  "net"
  "strconv"
)

// функция запускается как горутина
func process(conns map[int]net.Conn, n int) {
  // получаем доступ к текущему соединению
  conn := conns[n]
  // определим, что перед выходом из функции, мы закроем соединение
  defer conn.Close()
  for {
    // Будем прослушивать все сообщения разделенные \n
    message, _ := bufio.NewReader(conn).ReadString('\n')
    // Распечатываем полученое сообщение
    fmt.Print("Message Received:", string(message))
    // Отправить новую строку обратно клиенту
    _, err := conn.Write([]byte(strconv.Itoa(n) + "->> " + message + "\n"))
    // анализируем на ошибку
    if err != nil {
      fmt.Print(err.Error())
      break // выходим из цикла
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
