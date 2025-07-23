# test-subscription-service --- REST-сервис для агрегации данных об онлайн-подписках пользователей.

## Структура проекта

```markdown
├── cmd
│   ├── http            # Запуск сервера
│   └── migrator        # Запуск миграций
├── config              # Файл конфигурации(заменить на свой)
├── internal            # Доменная логика
│   ├── app             # Главное приложение
│   ├── config          # Конфигурация и загрузка
│   ├── dto             # Модели валидации запросов и отправки ответов
│   ├── entity          # Домены приложения
│   ├── http            # Хэндлеры и middleware для http
│   ├── repository      # Слой репозитория
│   └── usecase         # Слой кейсов
├── migrations          # Миграции
├── pkg                 # Вспомогательные пакеты
└── tests               # Тесты
       
```
## Настройки окружения

Конфигурации лежат по адресу `/config/config.yaml`: (configd.yaml для докера)

### Команда запуска приложения
```
go run . http --config=...config\local.yaml
```

### Команда миграции
```
go run .\cmd\migrator\main.go --db-host=localhost --db-name=postgres --mode=down --db-user=postgres --db-pass=example --migrations-path=...config/migrations
```
### Приложение
```
http://localhost:3001
```

### Swagger
```
http://localhost:3001/swagger/index.html
```

## Необходимо доделать
1. Добавить тесты
2. DTO
3. Добавить сервисы
4. Мигратор


