# test-subscription-service --- REST-сервис для агрегации данных об онлайн-подписках пользователей.

## Структура проекта

```markdown
└── cmd
│   └── app             # Запуск сервера
│   └── migrator        # Запуск миграций
└── config              # Файл конфигурации(заменить на свой)
└── docker              # Запуск и работа сервиса
├── internal            # Доменная логика
│   └── app             # Главное приложение
│   └── config          # Конфигурация и загрузка
│   └── domain          # Домены приложения
│   └── http            # Хэндлеры и middleware для http
│   └── lib             # Конфигурация и загрузка
└── services        # Слой сервисов
│   └── storage         # Работа с БД
├── tests               # Функциональные тесты
       
```
## Настройка окружения

Конфигурации лежат по адресу `/config/config.yaml`:

## Команда запуска приложения
```
go run . http --config=...config\local.yaml
```

## Команда миграции
```
go run .\cmd\migrator\main.go --db-host=localhost --db-name=postgres --mode=down --db-user=postgres --db-pass=example --migrations-path=...config/migrations
```

## Proto файлы брать из проекта https://github.com/Jonatna0990/protos

## Необходимо доделать
1. add token fail validation cases
2. удалена работа с RabbitMQ, перейти на Kafka


