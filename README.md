# recipe-service

Сервис рецептов

## Переменные окружения приложения

```
HTTP_HOST                   - хост прослушивания приложения (HTTP)
HTTP_PORT                   - порт прослушивания приложения (HTTP)

FILES_DIST                  - путь от корневой папки, куда будут сохранятся файлы

POSTGRES_DB_HOST            - хост базы данных
POSTGRES_DB_PORT            - порт базы данных
POSTGRES_DB_USER            - имя пользователя базы данных для авторизации
POSTGRES_DB_PASSWORD        - пароль пользователя базы данных для авторизации
POSTGRES_DB_NAME            - название базы данных, к которой будет идти подключение
```

## Пример переменных окружения:

```
HTTP_HOST=0.0.0.0
HTTP_PORT=8000

FILES_DIST=./dist

POSTGRES_DB_HOST=localhost
POSTGRES_DB_PORT=5432
POSTGRES_DB_USER=test
POSTGRES_DB_PASSWORD=test
POSTGRES_DB_NAME=recipes

```

## Команды для работы с приложением:

```
make help                   - вывести список доступных команд с описанием
make install                - устанавливает все необходимые зависимости и инструменты
make generate               - генерирует proto файлы gRPC сервиса
make test                   - запускает unit и интеграционные тесты
make test.unit              - запуск unit тестов
make test.integration       - запуск интеграционных тестов (поднятие и завершение окружения)
make test.integration.up    - поднятие тестового окружения
make test.integration.start - запуск интеграционных тестов
make test.integration.down  - завершение тестового окружения
make run                    - запуск приложения
```
