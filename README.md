
# Calculator HTTP Service

## Описание проекта

Этот проект представляет собой HTTP-сервис, который вычисляет математические выражения, отправленные пользователем в формате JSON. Сервис принимает POST-запросы с телом, содержащим математическое выражение, и возвращает результат вычисления или сообщение об ошибке.

### Основные возможности:

- Вычисление арифметических выражений с поддержкой приоритетов операций и скобок.
- Обработка ошибок: неверные выражения, деление на ноль и некорректный формат запроса.
- HTTP-ответы с кодами:
  - **200 OK**: успешное вычисление выражения.
  - **422 Unprocessable Entity**: ошибки валидации входных данных.
  - **500 Internal Server Error**: внутренние ошибки сервера.

---

## Инструкция по запуску проекта

1. **Склонировать репозиторий:**

   ```bash
   git clone https://github.com/antalkon/http_calc_YAL_sprint_1.git
   cd http_calc_YAL_sprint_1
   ```

2. **Запустить сервер:**

   ```bash
   go run ./cmd/app/...
   ```
   (or Makefile)
    ```bash
   Make run
   ```
   Сервер запустится на порту `8080` и будет доступен по адресу `http://localhost:8080`.

---

## Примеры использования

### Успешный запрос

**Входные данные:**

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expression":"2+2*2"}'
```

**Ответ:**

```json
{
    "result": 6
}
```

---

### Ошибка: неверное выражение (422)

**Входные данные:**

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expression":"5/0"}'
```

**Ответ:**

```json
{
    "error": "Expression is not valid"
}
```

---

### Ошибка: некорректный формат JSON (422)

**Входные данные:**

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expr":"2+2"}'
```

**Ответ:**

```json
{
    "error": "Invalid input format"
}
```

---

### Ошибка: неверный метод запроса (405)

**Запрос с использованием GET:**

```bash
curl -X GET http://localhost:8080/api/v1/calculate
```

**Ответ:**

```json
{
    "error": "Method not allowed"
}
```

---

## Структура проекта

```
calculator-service/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── handlers/
|   |   ├── handler_test.go
│   │   └── handler.go
│   ├── router/
│   │   └── router.go
│   ├── services/
|   |   ├── calc_test.go
│   │   └── calc.go
│   └── models/
│       └── models.go
├── tests/
│   ├── handler_test.go
│   └── calc_test.go
└── go.mod
```

---

## Тестирование

Запустить тесты можно следующей командой:

```bash
go test ./internal/...
```
(or Makefile)
```bash
   Make test
```

Все модули будут протестированы на корректную обработку и ошибки.

---

