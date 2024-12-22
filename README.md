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

## Проверка различных исходов

### get-tokens с корректными данными (Ok)
![image](https://github.com/user-attachments/assets/f6bc58ff-d77d-483c-a357-16facdd8ec84)

### get-tokens с некорректным userID (Bad Request)
![image](https://github.com/user-attachments/assets/dc0568de-1aff-4762-882b-0d45f32e071c)

### refresh с корректными данными (Ok)
![image](https://github.com/user-attachments/assets/1e9c976f-2835-4607-8231-8a156f17a51c)

### refresh при измененном ip (Warning email на почту)
![image](https://github.com/user-attachments/assets/a6592232-4c5d-4b8c-b6ae-849b060615a0)

### refresh при измененном с некорректным Access токеном (Unauthorized)
![image](https://github.com/user-attachments/assets/b71a562e-c79f-4745-bdd5-334bdeaae4b3)

### refresh при измененном с некорректным Refresh токеном (Unauthorized)
![image](https://github.com/user-attachments/assets/3b7bd7f8-5169-4317-bc05-e2c1847d4bff)

### refresh без одного из токенов (Bad Request)
![image](https://github.com/user-attachments/assets/46acf247-20dc-4041-b2cf-fae51789c27a)




