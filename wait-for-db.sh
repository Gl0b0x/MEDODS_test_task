#!/bin/sh

# Скрипт ожидает доступности порта базы данных
echo "Waiting for PostgreSQL to be ready..."

# Ожидание до тех пор, пока не будет доступен указанный порт
while ! nc -z postgres 5432; do
  sleep 1
done

echo "PostgreSQL is ready. Starting the app..."

# Запускаем приложение
exec "$@"