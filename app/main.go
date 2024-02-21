package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Регистрация маршрутов
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	r.HandleFunc("/user", CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	// Запуск сервера
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}

// Обработчики
func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Получение списка пользователей")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Получение пользователя")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Создание пользователя")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Обновление пользователя")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Удаление пользователя")
}
