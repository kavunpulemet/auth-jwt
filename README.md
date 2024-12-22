## Тесты
`go test -v ./...`

## Запуск приложения с использованием Docker
`docker-compose up --build`

## API Endpoints

### Получение пары Access, Refresh токенов

**Endpoint:** `POST /api/get-tokens/{userID}/`

**Описание:** Выдает пару Access, Refresh токенов для пользователя с идентификатором (GUID) указанным в параметре запроса.

### Обновление пары Access, Refresh токенов

**Endpoint:** `GET /api/refresh/`

**Описание:** Выполняет Refresh операцию на пару Access, Refresh токенов.

**Пример запроса:**
```json
{
  "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJyZWZyZXNoX3Rva2VuX2lkIjoiZmVlYzE1OGItYjUzZC00ZDBjLTgyNjAtY2Q3ZTBhMTM3ZjFhIiwidXNlcl9pZCI6IjUzMTAwMGQ4LTUyMWYtNGY5Ny1iYjE5LTEzZmRkOWYzYTQzOSIsInVzZXJfaXAiOiIxNzIuMTguMC4xIiwiZXhwIjoxNzM0ODgwOTUyLCJpYXQiOjE3MzQ4ODAwNTJ9.3xxY0mvWCkfMmQSYu_x6MzV9bNp5D6DE95EU6R-fc7eve7vFXxiITFhZFuo5dsv5EbjpvVYuVCwEVi3mvQYnsQ",
  "refresh_token": "2wL9E3hX3/onX8k/F/fsPAFF1ziDTPuni3YYQuvVAD0="
}
```