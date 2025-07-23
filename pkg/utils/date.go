package utils

import (
	"fmt"
	"strings"
	"time"
)

// ParseMonthYear принимает строку формата "MM-YYYY" и возвращает time.Time (начало месяца)
func ParseMonthYear(input string) (time.Time, error) {
	// Проверим формат: должно быть ровно два числа, разделённые "-"
	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		return time.Time{}, fmt.Errorf("invalid date format: expected MM-YYYY")
	}

	// Преобразуем строку в дату
	t, err := time.Parse("01-2006", input)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date: %w", err)
	}

	return t, nil
}

// FormatMonthYear принимает time.Time и возвращает строку формата "MM-YYYY"
func FormatMonthYear(t time.Time) string {
	// Обнуляем дату до первого числа месяца, без времени
	t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	return t.Format("01-2006")
}
