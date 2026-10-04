# Disk Analyzer

[![hexlet-check](https://github.com/k1ll4reall/go-from-scratch-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/k1ll4reall/go-from-scratch-project-242/actions)

Утилита на Go для подсчёта размера файлов и каталогов.

## Запуск

```sh
go run ./cmd/hexlet-path-size testdata/test.txt
```

## Флаги

- `--human`, `-H` — удобный формат размера: `B`, `KB`, `MB` и далее.
- `--all`, `-a` — учитывать скрытые файлы и каталоги с именами, начинающимися с точки.
- `--recursive`, `-r` — учитывать содержимое вложенных каталогов.

Без флагов размер выводится в байтах, скрытые элементы и содержимое вложенных каталогов не учитываются.

## Примеры

```sh
go run ./cmd/hexlet-path-size --human testdata
go run ./cmd/hexlet-path-size --all testdata
go run ./cmd/hexlet-path-size --recursive --all testdata
```

## Использование как библиотеки

Пакет `code` экспортирует функцию:

```go
GetPathSize(path string, recursive, human, all bool) (string, error)
```

Пример вызова:

```go
size, err := code.GetPathSize("testdata", true, true, false)
```

## Проверки

```sh
go test ./...
golangci-lint run
```
<details>
<summary>Автоматические тесты Хекслета</summary>

Тесты запускаются на каждый коммит. За запуск отвечает файл `.github/workflows/hexlet-check.yml` — не удаляйте и не переименовывайте ни его, ни репозиторий.

</details>

## О Хекслете

[Хекслет](https://ru.hexlet.io/) — школа программирования: авторские программы обучения с практикой, поддержкой наставников и реальными проектами, которые остаются в резюме. Этот репозиторий — один из таких проектов.
