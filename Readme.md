![Go 1.24.1](https://img.shields.io/badge/go-1.24.1-blue.svg) ![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)

# Base-Go-API-Engine

Базовый демонстрационный REST‑API на Go с использованием [chi](https://github.com/go-chi/chi) и PostgreSQL. Проект служит отправной точкой для быстрого старта микросервисов и учебных примеров.

## Оглавление

- [Описание](#описание)  
- [Архитектура](#архитектура)  
- [Технологии](#технологии)  
- [Установка и запуск](#установка-и-запуск)  
  - [Предварительные требования](#предварительные-требования)  
  - [Конфигурация](#конфигурация)  
  - [Запуск миграций](#запуск-миграций)  
  - [Запуск приложения](#запуск-приложения)  
- [API Endpoints](#api-endpoints)  
- [Рекомендации по улучшению](#рекомендации-по-улучшению)  
- [Возможные области применения](#возможные-области-применения)  
- [Контакты](#контакты)  
- [Лицензия](#лицензия)  

---

## Описание

Base-Go-API-Engine — это минималистичный демонстрационный каркас HTTP‑сервера на Go. Включает:

- Организацию проекта по «чистой архитектуре» (cmd, internal, pkg).  
- Конфигурирование через переменные окружения.  
- Подключение к PostgreSQL с пулом соединений.  
- Простейший роутинг через chi и middleware (логгинг, таймаут, Recoverer).  
- SQL‑миграции в виде файлов (up/down).  

Цель — дать готовую основу для быстрого прототипирования API‑микросервисов.

---

## Архитектура

```
.
├── .air.toml              # Конфиг live‑reload (Air)
├── docker-compose.yml     # PostgreSQL для локальной разработки
├── go.mod                 # Модули Go
├── cmd
│   ├── api                # HTTP‑сервер
│   │   └── components     # Роуты и handlers
│   └── migrate            # Утилита для миграций
│       └── migrations     # SQL‑файлы миграций
├── internal
│   ├── config             # Константы и строки запросов
│   ├── db                 # Инициализация и пул соединений
│   ├── env                # Чтение переменных окружения
│   └── store              # Репозитории (CRUD для users/posts)
├── scripts
│   └── initDb.sql         # Скрипт создания БД
└── .gitignore
```

- **cmd/api** — точка входа HTTP‑сервера, создаёт `Application`, монтирует роуты и запускает `http.Server`.  
- **cmd/migrate** — утилита для применения SQL‑миграций вручную.  
- **internal/db** — обёртка над `database/sql`, пул соединений и `PingContext`.  
- **internal/env** — вспомогательные функции `GetString`, `GetInt`.  
- **internal/store** — слои доступа к данным: `UsersStore`, `PostStore`.  
- **internal/config** — хранит константы (имена переменных окружения, строки запросов).  

---

## Технологии

- [Go 1.24](https://golang.org/)  
- [chi v5](https://github.com/go-chi/chi) — лёгкий HTTP‑роутер  
- [pq v1.10](https://github.com/lib/pq) — PostgreSQL‑драйвер  
- [Air](https://github.com/cosmtrek/air) — hot‑reload во время разработки  
- Docker & Docker Compose — локальная БД  

---

## Установка и запуск

### Предварительные требования

- Go 1.24+  
- Docker & Docker Compose  
- `make` (опционально)  

### Конфигурация

Переменные окружения (смотрите `internal/config/config.go`):

| Переменная               | Описание                             | Значение по умолчанию                          |
| ------------------------ | ------------------------------------ | ----------------------------------------------- |
| `ADDR`                   | Адрес прослушивания HTTP‑сервера     | `:8080`                                         |
| `DB_ADDR`                | DSN для подключения к PostgreSQL     | `postgres://postgres:password@localhost/appdb?sslmode=disable` |
| `DB_MAX_OPEN_CONNS`      | Максимум открытых соединений         | `100`                                           |
| `DB_MAX_IDLE_CONNS`      | Максимум простаивающих соединений    | `100`                                           |
| `DB_MAX_IDLE_TIME_MINS`  | Время простоя соединения (например `15m`) | `15m`                                      |

### Запуск миграций

1. Убедитесь, что БД создана (скрипт `scripts/initDb.sql` или `psql`).  
2. Запустите вручную:
   ```bash
   psql "$DB_ADDR" -f cmd/migrate/migrations/000001_create_users.up.sql
   psql "$DB_ADDR" -f cmd/migrate/migrations/000002_create_posts.up.sql
   ```

*(Можно подключить `github.com/golang-migrate/migrate` для автоматизации.)*

### Запуск приложения

#### Через Docker Compose

```bash
docker-compose up -d db
go run cmd/api/main.go
```

#### Локально с Air

```bash
air
```

#### Без live‑reload

```bash
go build -o bin/main.exe ./cmd/api
./bin/main.exe
```

---

## API Endpoints

| Метод | Путь            | Описание               | Ответ             |
| ----- | --------------- | ---------------------- | ----------------- |
| GET   | `/v1/status`    | Проверка статуса API   | `OK: it works`    |

> *Демонстрационный статус-чек. Дополните CRUD-эндпоинтами для `users` и `posts`.*

---

## Рекомендации по улучшению

1. **CRUD для пользователей и постов**  
   – Реализовать `GET`, `POST`, `PUT`, `DELETE` для ресурсов.  
2. **Аутентификация и авторизация**  
   – JWT, OAuth2, RBAC.  
3. **Валидация входящих данных**  
   – `go-playground/validator`, middleware.  
4. **Документация API**  
   – Swagger / OpenAPI, `swaggo/swag`.  
5. **Логирование и трассировка**  
   – `zap`/`logrus`, OpenTelemetry.  
6. **Тесты**  
   – Unit и интеграционные тесты, mock-репозитории.  
7. **CI/CD**  
   – GitHub Actions, Docker Hub, Helm-чарты.  
8. **Кэширование**  
   – Redis для горячих данных.  
9. **Мониторинг и метрики**  
   – Prometheus, Grafana.  

---

## Возможные области применения

- Быстрый старт микросервисов на Go.  
- Учебные проекты и демонстрации Go‑архитектуры.  
- Внутренние инструменты и прототипы API.  
- Бэкенд для SPA / мобильных приложений.  

---

## Контакты

📧 dim4dmi7rij@yandex.ru
📧 dimaprihodko180@gmail.com
📱 +7 904 926‒57‒29  

---

## Лицензия

MIT

```