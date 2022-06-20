<h2> Развертывание </h2>

Запуск миграций:
```
make migrate
```
Генерация proto-файлов:
```
make proto-gen
```
Запуск приложения:
```
make docker-up
```
Запуск тестов:
```
make test
```
<h2> Примеры запросов </h2>

```
evans proto/book.proto -p 9080

call GetBooksByAuthor
call GetAuthorsByBook
```

