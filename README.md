# accruals-service

Микросервис для создания скидок и формирования цен из реализованных скидок

| Branch        | Pipeline          | Code coverage  |
| ------------- |:-----------------:| --------------:|
| main          | [![pipeline status](https://penahub.gitlab.yandexcloud.net/pena-services/discount-service/badges/main/pipeline.svg)](https://penahub.gitlab.yandexcloud.net/pena-services/discount-service/-/pipelines) | [![coverage report](https://penahub.gitlab.yandexcloud.net/pena-services/discount-service/badges/main/coverage.svg?job=test)](https://penahub.gitlab.yandexcloud.net/pena-services/discount-service/-/pipelines) |
| staging           | [![pipeline status](https://penahub.gitlab.yandexcloud.net/pena-services/discount-service/badges/staging/pipeline.svg)](https://penahub.gitlab.yandexcloud.net/pena-services/discount-service/-/pipelines) | [![coverage report](https://penahub.gitlab.yandexcloud.net/pena-services/discount-service/badges/staging/coverage.svg?job=test)](https://penahub.gitlab.yandexcloud.net/pena-services/discount-service/-/pipelines) |
| dev           | [![pipeline status](https://penahub.gitlab.yandexcloud.net/pena-services/discount-service/badges/dev/pipeline.svg)](https://penahub.gitlab.yandexcloud.net/pena-services/discount-service/-/pipelines) | [![coverage report](https://penahub.gitlab.yandexcloud.net/pena-services/discount-service/badges/dev/coverage.svg?job=test)](https://penahub.gitlab.yandexcloud.net/pena-services/discount-service/-/pipelines) |

## Переменные окружения приложения

```
GRPC_HOST                   - хост прослушивания приложения (gRPC)
GPRC_PORT                   - порт прослушивания приложения (gRPC)
HTTP_PORT                   - порт прослушивания приложения (HTTP)

MONGO_HOST                  - хост базы данных
MONGO_PORT                  - порт базы данных
MONGO_USER                  - имя пользователя базы данных для авторизации
MONGO_PASSWORD              - пароль пользователя базы данных для авторизации
MONGO_AUTH                  - название базы данных для авторизации
MONGO_DATABASE_NAME         - название базы данных, к которой будет идти подключение
```

## Пример переменных окружения:

```
GRPC_HOST=0.0.0.0
GPRC_PORT=9001
HTTP_PORT=8001

MONGO_HOST=localhost
MONGO_PORT=27017
MONGO_USER=test
MONGO_PASSWORD=test
MONGO_AUTH=admin
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
