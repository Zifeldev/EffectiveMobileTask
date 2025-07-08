# People API

RESTful API для создания и получения информации о людях с автоматическим обогащением через внешние сервисы:
- [genderize.io](https://genderize.io)
- [agify.io](https://agify.io)
- [nationalize.io](https://nationalize.io)

---

## Функциональность

-  Создание пользователя с обогащением (gender, age, country)
-  Получение списка с фильтрацией и пагинацией
-  Swagger-документация
-  Построено с использованием Go, Gin, PostgreSQL, Goose, Squirrel

---

##  Запуск

### 1. Настрой `.env`

```env
DB_URL=postgres://postgres:123@localhost:5432/people_db?sslmode=disable
PORT=8080
LOG_LEVEL=debug
(Это мой env)
```
### 2. Запуск миграции
goose -dir ./internal/migrations postgres "$DB_URL" up
за место DB_URL вводим свой

Также можно сделать по другому

### 3. Запуск сервера
go run cmd/main.go 
или 
cd cmd && go run main.go

### Примеры запросов
POST /people
Создание человека (обогащается внешними API):
curl -X POST http://localhost:8080/people \
  -H "Content-Type: application/json" \
  -d '{"name":"Azat","surname":"Buranbayev","patronymic":"Erzhanovich"}'

### GET /people
Фильтрация и пагинация:
curl "http://localhost:8080/people?gender=male&min_age=20&page=1&limit=5"

### Swagger UI
Документация доступна по адресу:
http://localhost:8080/swagger/index.html
(Если вы конечно не поменяли порт)

### Структура проекта

| Путь                         | Описание                                               |
|-----------------------------|--------------------------------------------------------|
| `cmd/main.go`               | Точка входа, запуск сервера                            |
| `internal/handler/`         | HTTP-хендлеры (Gin)                                    |
| `internal/model/`           | Структуры, включая `Person`, `CreatePersonRequest` и т.п. |
| `internal/repository/`      | Работа с базой данных (Squirrel + pgx)                 |
| `internal/service/`         | Логика обогащения (внешние API)                        |
| `internal/config/`          | Загрузка конфигурации из `.env`                        |
| `internal/db/`              | Подключение к БД                                       |
| `internal/migrations/`      | Goose миграции SQL                                     |
| `go.mod`, `go.sum`          | Зависимости проекта                                    |






