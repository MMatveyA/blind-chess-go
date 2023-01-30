package transport

import (
    "fmt"
    "net/http"
    "strconv"
)

// Домашняя страница
func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w,r)
        return
    }

    w.Write([]byte("Привет из Arch!"))
}

// Показ поста с ID
func showPost(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w,r)
        return
    }

    fmt.Fprintf(w, "Отображение статьи с ID %d", id)
}

// Создание нового поста
func createPost(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Неверный метод запроса", 405)
        return
    }

    w.Write([]byte("Форма для создание новой статьи"))
}
