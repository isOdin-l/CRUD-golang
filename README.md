# REST API сервер

[![JWT](https://img.shields.io/badge/JWT-black?style=flat-square&logo=JSON%20web%20tokens)](https://www.jwt.io/introduction)
[![Go](https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat-square&logo=docker&logoColor=white)](https://www.docker.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=flat-square&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Swagger](https://img.shields.io/badge/Swagger-Documentation-85EA2D?style=flat-square&logo=swagger)](https://swagger.io/)

Целью данного проекта было научиться разработке сервера на Go, следуя дизайну REST API.

- Сервер написан на [go-chi/chi](https://github.com/go-chi/chi)
- Используется подход Чистой Архитектуры в построении структуры приложения
- Структура репозитория по стандарту [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- Работа с БД PostgreSQL - [jackc/pgx](https://github.com/jackc/pgx)
- Сборка и запуск всего приложения - Docker с использованием multi-stage building
- Парсинг переменных окружения в структуру - [caarlos0/env](https://github.com/caarlos0/env) 
- Аутентификация осуществляется с помощью JWT - [golang-jwt/jwt](https://github.com/golang-jwt/jwt)
- Реализован собственный парсер ```tools/bindchi``` URL параметров, тела запроса и контекста в структуру
- Валидация структур по структурным тегам - [go-playground/validator](https://github.com/go-playground/validator)
- Добавлен Swagger для наглядности - [swaggo/swag](https://github.com/swaggo/swag)

## Окно Swagger

<img src="assets/swagger.png" alt="Swagger UI" style="width: 90%;">

Доступ к swagger после запуска docker-контейнера: http://localhost:SERVER_PORT/swagger/index.html

## Запуск сервера #
* Клонирование репозитория
```bash
git clone https://github.com/isOdin-l/REST-CRUD-golang
cd REST-CRUD-golang
```

* Сборка и запуск в docker контейнере
```bash
docker-compose up --build
```

## Special thanks
Огромная благодарность [@TobbyMax](https://github.com/TobbyMax) за ценные советы, review и помощь в реализации проекта.