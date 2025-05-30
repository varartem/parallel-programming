# Параллельное программирование

Этот репозиторий предназначен для изучения и реализации задач по параллельному программированию на языке **Go**.

## О проекте

Репозиторий содержит практические задания для освоения концепций параллельного и конкурентного программирования, включая работу с горутинами, каналами, синхронизацией и другими механизмами параллельности в Go.

## Требования к реализации

### Язык программирования
Все задачи должны быть реализованы на языке **Go**.

### Тестирование
- Тесты должны размещаться рядом с реализацией в той же директории
- Файлы тестов должны иметь суффикс `_test.go`

### Важно для написания тестов
**Обработка ошибок в негативных тест-кейсах**: Если вы пишете тесты для отрицательных сценариев (когда ожидается ошибка), необходимо **обрабатывать ошибки**, а не просто их возвращать. В противном случае тесты могут не пройти.

**Правильно:**
```go
func TestNegativeCase(t *testing.T) {
    result, err := SomeFunction(invalidInput)
    if err != nil {
        // Обрабатываем ошибку, а не возвращаем её
        t.Logf("Ожидаемая ошибка: %v", err)
        return
    }
    // Остальная логика теста
}
```

**Неправильно:**
```go
func TestNegativeCase(t *testing.T) {
    result, err := SomeFunction(invalidInput)
    if err != nil {
        return err // ❌ Не возвращайте ошибку из теста
    }
}
```

## Запуск тестов локально

Для запуска всех тестов в проекте:
```bash
go test ./...
```

Для запуска тестов в конкретной директории:
```bash
go test ./task/
```

## Запуск тестов GitHub

Тесты срабатывают автоматически, после вызова push и попадания последних изменений на удаленный репозиторий 