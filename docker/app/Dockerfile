# Используем официальный образ Go
FROM golang:1.21

# Установим рабочую директорию в контейнере
WORKDIR /app

# Копируем исходный код веб-приложения в контейнер
COPY ./app/ .

# Устанавливаем CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon@latest

# Указываем порт, который будет прослушивать приложение
EXPOSE 8080

# Запускаем приложение с помощью CompileDaemon для hot reloading
ENTRYPOINT CompileDaemon --build="go build -o /bin/main ." --command="./bin/main"