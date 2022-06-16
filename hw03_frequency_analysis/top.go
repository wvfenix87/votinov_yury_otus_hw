package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var regex = regexp.MustCompile(`[,."!:?]+`)

func Top10(incomeString string) []string {
	// разбиваем строку на слова
	splitString := strings.Fields(incomeString)
	// проходим по массиву и наполняем мапу слово-кол-во повторений
	frequencyWordMap := make(map[string]int)
	for _, currentWord := range splitString {
		// убираем тире
		if currentWord != "-" && len(currentWord) > 0 {
			// приводим в нижний регистр
			transformWord := strings.ToLower(currentWord)
			// выпиливаем все знаки из слов
			transformWord = regex.ReplaceAllString(transformWord, "")
			if len(transformWord) > 0 {
				frequencyWordMap[transformWord]++
			}
		}
	}
	// вынимаем ключи имена
	names := make([]string, 0, len(frequencyWordMap))
	for name := range frequencyWordMap {
		names = append(names, name)
	}
	// сортируем слайс имен со своей функцией сравнения используя подсчитанное кол-во повторений
	sort.Slice(names, func(firstWord, secondWord int) bool {
		if frequencyWordMap[names[firstWord]] == frequencyWordMap[names[secondWord]] {
			return names[firstWord] < names[secondWord]
		}
		return frequencyWordMap[names[firstWord]] > frequencyWordMap[names[secondWord]]
	})
	// возвращяем подслайс из первых 10 элементов (или меньше)
	endTopIndex := 10
	if len(names) < 10 {
		endTopIndex = len(names)
	}
	return names[:endTopIndex]
}
