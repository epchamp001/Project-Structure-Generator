# 🏗️Project Structure Generator

Генератор шаблонного проекта на Go, построенного по принципам **Go Clean Architecture**, с поддержкой:

- gRPC
- HTTP
- Redis
- Prometheus (metrics)
- Grafana (dashboard provisioning)
- K6 (нагрузочное тестирование)

## 🚀 Использование

Соберите и установите бинарник в `$GOPATH/bin` (или в директорию из `$GOBIN`):

```bash
go install github.com/epchamp001/Project-Structure-Generator/cmd/gen@v0.1.0
```

> Убедитесь, что $(go env GOPATH)/bin (или ваш GOBIN) добавлен в переменную среды PATH.

### Запуск установленного бинарника

```bash
gen -n <appName> -c <cmdName> [-d <targetDir>] [-g] [-t] [-r] [-f] [-m] [-l]
```

### Запуск без установки

Если вы не хотите устанавливать бинарник глобально, запустите команду из корня проекта:

```bash
go run cmd/gen/main.go \
  -n <appName> \
  -c <cmdName> \
  [-d <targetDir>] \
  [-g] [-t] [-r] [-f] [-m] [-l]
```

## 🧩 Поддерживаемые флаги

| Флаг              | Название                          | Тип    | Обязательный | Описание                                              |
|-------------------|-----------------------------------|--------|--------------|-------------------------------------------------------|
| `-n`, `--name`    | Название проекта                  | string | ✅            | Название директории проекта                           |
| `-c`, `--cmd`     | Название директории в `cmd/`      | string | ✅            | Имя поддиректории в `cmd/`                            |
| `-d`, `--dir`     | Целевая директория                | string | ❌            | Путь сохранения проекта                               |
| `-g`, `--grpc`    | Включить gRPC                     | bool   | ❌            | Включает генерацию gRPC слоёв                         |
| `-t`, `--http`    | Включить HTTP                     | bool   | ❌            | Включает генерацию HTTP слоёв                         |
| `-r`, `--redis`   | Включить Redis                    | bool   | ❌            | Добавляет Redis-репозиторий                           |
| `-f`, `--grafana` | Включить Grafana                  | bool   | ❌            | Добавляет Grafana dashboards                          |
| `-m`, `--metrics` | Включить Prometheus               | bool   | ❌            | Добавляет директории и файлы для настройки Prometheus |
| `-l`, `--load`    | Включить нагрузочное тестирование | bool   | ❌            | Добавляет `scripts/k6/load_test.js`                   |

> ❗ По умолчанию gRPC и HTTP включены, если не передан ни один из этих двух флагов. Все остальные опции по умолчанию отключены.

## 🔧 Первичная настройка

После генерации:

```bash
go mod tidy
```

Настрой вручную некоторые файлы.

P.S. Позже планируется автоматизировать этот процесс при помощи дополнительных флагов.

### ⚙️ Что нужно вручную настроить

1. Если используется Redis добавь:

```go
viper.BindEnv("storage.redis.host", "REDIS_HOST")
```

2. Настройка Логгера:

Всё прописано в pkg/logger. Если не нужен sampling и initial fields, то в config.yaml сделай так:

```yaml
sampling: null
initialFields: {}
```

3. Обнови shutdown в main.go. Проверь, что используется Shutdown из нужной тебе структуры:

```go
shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.PublicServer.ShutdownTimeout)*time.Second)
```

4. Проверь и поправь:
* config.yaml
* prometheus.yaml
* Dockerfile
* .env, .env.example
* docker-compose.yaml
* Makefile
* импорты во всех .go файлах

5. Проверь порты и соответствующие значения в grafana, prometheus, docker-compose.yaml

6. Проверь tests/integration. Нужно один раз настроить контейнер с БД.