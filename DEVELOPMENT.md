# AmoCRM Go Client - Руководство по разработке

## Структура проекта

```
amocrm-go/
├── amocrm/              # Основной пакет библиотеки
│   ├── client.go        # Главный клиент и HTTP-логика
│   ├── auth.go          # OAuth2 авторизация
│   ├── contacts.go      # Работа с контактами
│   ├── companies.go     # Работа с компаниями
│   ├── leads.go         # Работа со сделками
│   ├── tasks.go         # Работа с задачами
│   ├── notes.go         # Работа с примечаниями
│   ├── webhooks.go      # Работа с вебхуками
│   ├── catalogs.go      # Работа с каталогами
│   ├── account.go       # Информация об аккаунте
│   ├── types.go         # Общие типы данных
│   ├── errors.go        # Типы ошибок
│   ├── storage.go       # Интерфейс хранилища токенов
│   └── storage/         # Реализации хранилищ
│       └── file_storage.go
├── examples/            # Примеры использования
│   ├── basic/          # Базовый пример
│   ├── oauth2/         # OAuth2 авторизация
│   └── batch/          # Пакетные операции
├── go.mod
├── go.sum
├── README.md
└── LICENSE
```

## Основные компоненты

### 1. Client (client.go)
- HTTP-клиент с rate limiting
- Автоматическое обновление OAuth2 токенов
- Поддержка контекстов
- Логирование запросов

### 2. Модели сущностей
- Contact - контакты
- Company - компании
- Lead - сделки
- Task - задачи
- Note - примечания
- Webhook - вебхуки
- Catalog - каталоги

### 3. Сервисы
Каждая сущность имеет свой сервис с методами:
- List() - получение списка
- GetByID() - получение по ID
- Create() - создание
- CreateBatch() - пакетное создание
- Update() - обновление
- UpdateBatch() - пакетное обновление

## Добавление новой функциональности

### Добавление нового сервиса

1. Создайте файл `amocrm/entity_name.go`
2. Определите структуру модели
3. Создайте сервис:

```go
type EntityService struct {
    client *Client
}

func (s *EntityService) List(ctx context.Context, filter *Filter) ([]Entity, error) {
    // Реализация
}
```

4. Добавьте сервис в Client:

```go
// В client.go
type Client struct {
    // ...
    EntityName *EntityService
}

// В NewClient()
client.EntityName = &EntityService{client: client}
```

### Добавление нового метода API

```go
func (s *Service) MethodName(ctx context.Context, params *Params) (*Result, error) {
    path := "/endpoint"
    
    var result Result
    if err := s.client.GetJSON(ctx, path, &result); err != nil {
        return nil, err
    }
    
    return &result, nil
}
```

## Тестирование

### Запуск тестов

```bash
go test ./...
```

### Запуск с покрытием

```bash
go test -cover ./...
```

### Создание мок-сервера для тестов

```go
server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"result": "ok"}`))
}))
defer server.Close()
```

## Рекомендации

### 1. Обработка ошибок
Всегда возвращайте информативные ошибки:

```go
if err != nil {
    return nil, fmt.Errorf("failed to create contact: %w", err)
}
```

### 2. Использование контекстов
Всегда передавайте контекст для поддержки таймаутов:

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

result, err := client.Service.Method(ctx, params)
```

### 3. Rate Limiting
Клиент автоматически ограничивает запросы (7 req/s по умолчанию).
Не нужно добавлять дополнительную логику.

### 4. Логирование
Используйте встроенный logger:

```go
client := amocrm.NewClient(
    amocrm.WithDebug(true), // включить подробное логирование
)
```

## Публикация новой версии

1. Обновите версию в документации
2. Создайте git tag:

```bash
git tag v1.0.0
git push origin v1.0.0
```

3. Go автоматически подхватит новую версию

## Совместимость с PHP-библиотекой

Библиотека спроектирована для совместимости с PHP-версией:

| PHP | Go |
|-----|-----|
| `AmoAPI::oAuth2()` | `client.Auth.ExchangeCode()` |
| `AmoAPI::permanentToken()` | `WithPermanentToken()` |
| `AmoContact` | `Contact` |
| `AmoLead` | `Lead` |
| `AmoAPI::getContacts()` | `client.Contacts.List()` |
| `$contact->save()` | `client.Contacts.Create()` |

## Поддержка

При возникновении проблем:
1. Проверьте примеры в `examples/`
2. Изучите документацию AmoCRM API
3. Создайте issue на GitHub
