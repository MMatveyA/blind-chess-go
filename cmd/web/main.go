// Сайт-форум для шахматистов
package main

import (
    "flag"
    "log"
    "net/http"

    handler "github.com/MMatveyA/blind-chess-go/internal/transport"
)

type Config struct {
    Addr string
    Port string
    StaticDir string
}

func main() {
    // Конфигурация веб сервера
    cfg := new(Config)
    flag.StringVar(&cfg.Addr, "addr", "localhost", "HTTP network address")
    flag.StringVar(&cfg.Port, "port", ":4000", "HTTP network port")
    flag.StringVar(&cfg.StaticDir, "static-dir", "./web/static", "Path to static assets")
    flag.Parse()

    mux := http.NewServeMux()
    mux.HandleFunc("/", handler.home)
    mux.HandleFunc("/post", handler.showPost)
    mux.HandleFunc("/post/create", handler.createPost)

    log.Println("Запуск сервера на http://127.0.0.1:4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
