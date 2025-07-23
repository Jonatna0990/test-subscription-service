#!/bin/bash

# Выполняем миграции
echo "⏳ Running migrations..."
./app migrate --db-host=postgres --db-name=postgres --mode=up --db-user=postgres --db-pass=example --migrations-path=./migrations

# Проверяем статус миграций
if [ $? -ne 0 ]; then
  echo "Migration failed"
  exit 1
fi

# Запускаем основное приложение
echo "🚀 Starting main app..."
exec ./app http --config=config/configd.yaml