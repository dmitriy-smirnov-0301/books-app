# Books App

Books App — это REST API для управления библиотекой книг. Оно позволяет выполнять CRUD-операции с книгами, используя **PostgreSQL** и **GORM**.

## Функциональность
- Получение списка всех книг (**GET /books**)
- Получение книги по ID (**GET /books/{id}**)
- Добавление книги (**POST /books**)
- Обновление книги (**PUT /books/{id}**)
- Удаление книги (**DELETE /books/{id}**)

### 1. Инициализация базы данных
Создайте базу данных и пользователя, выполнив SQL-скрипт:
```sh
psql -U postgres -f configs/init-db.sql
```

### 2. Установка зависимостей
```sh
go mod tidy
```

### 3. Запуск приложения
```sh
go run cmd/main.go
```

## 4. Запуск тестов
Запустите все тесты в проекте:
```sh
go test ./...
```

## Конфигурация
Приложение использует `config.json` для хранения параметров БД. Файл находится в `configs/` и имеет следующую структуру:
```json
{
  "database": {
    "host": "localhost",
    "user": "book_admin",
    "password": "book_password",
    "dbname": "books_db",
    "port": 5432,
    "sslmode": "disable"
  }
}
```
## Документация
**[Swagger UI](http://localhost:8080/swagger/index.html)**
