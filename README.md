- Разработка веб-приложения на Go, следуя дизайну REST API.
- Сервер написан с использование [go-chi/chi](https://github.com/go-chi/chi)
- Использовался подход Чистой Архитектуры в построении структуры приложения
- Структура репозитория по стандарту [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- Работа с БД PostgreSQL через [jackc/pgx](https://github.com/jackc/pgx)
- Сборка и запуск всего приложения с помощью Docker
- Конфигурация приложения и парсинг переменных окружения в структуру с помощью библиотеки [caarlos0/env](https://github.com/caarlos0/env) 
- Аутентификация осуществляется с помощью JWT
- Реализован собственный парсер ```tools/chiBinding``` URL параметров, тела запроса и контекста в структуру

# Запуск сервера #
1) Клонирование репозитория
```console
git clone https://github.com/isOdin-l/REST-CRUD-golang
```
2) Сборка и запуск в docker контейнере
```console
docker-compose up --build
```
