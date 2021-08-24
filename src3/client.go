package main

import (
  "bufio"
  "fmt"
  "net"
  "os"
)

func main() {

  // Подключаемся к сокету
  conn, _ := net.Dial("tcp", "127.0.0.1:8081")
  // отложенное закрытие соединения, которое срабатывает при выходе из функции
  defer conn.Close()
  for {
    // Чтение входных данных от stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    text, _ := reader.ReadString('\n')
    // Отправляем в socket
    _, err := fmt.Fprintf(conn, text+"\n")
    // анализируем на ошибку
    if err != nil {
      fmt.Print(err.Error())
      break // выходим из цикла
    }
    // Прослушиваем ответ
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: " + message)
  }
}
