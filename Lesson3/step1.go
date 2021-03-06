package Lesson3

import "fmt"

// TODO На стандартный ввод подается 10 целых чисел, разделенных пробелами (числа могут повторяться).
//  Для чтения из стандартного ввода импортирован пакет fmt. Вам необходимо вычислить результат
//  выполнения функции work для каждого из полученных чисел. Функция work имеет следующий вид:
//  work(x int) int
//  Результаты вычислений , разделенные пробелами, должны быть напечатаны в строку.
//  Однако работа функции work занимает слишком много времени. Выделенного вам времени выполнения не хватит
//  на последовательную обработку каждого числа, поэтому необходимо реализовать кэширование уже готовых результатов
//  и использовать их в работе.
//  После завершения работы программы результат выполнения будет дополнен информацией о соблюдении установленного лимита времени выполнения.
func cachedWork() {
	results := make(map[int]int)
	for i := 0; i < 10; i++ {
		var element int
		fmt.Scan(&element)
		if _, ok := results[element]; !ok {
			results[element] = work(element)
			fmt.Printf("%d: Value added; ", results[element])
		} else {
			fmt.Printf("%d: Value getted;", results[element])
		}
	}
}

// Just a mocking func
func work(x int) int {
	x++
	return x
}

// TODO В ходе анализа результатов переписи населения информация была сохранена в объекте типа map:
//  groupCity := map[int][]string{
// 		10:   []string{...}, // города с населением 10-99 тыс. человек
//  	100:  []string{...}, // города с населением 100-999 тыс. человек
//		1000: []string{...}, // города с населением 1000 тыс. человек и более
//  }
//  При подготовке важного отчета о городах с населением 100-999 тыс. человек был подготовлен другой объект типа map:
//  cityPopulation := map[string]int{
//		город: население города в тысячах человек,
//  }
//  Однако база данных с информацией о точной численности населения содержала ошибки,
// 	поэтому в cityPopulation в т.ч. была сохранена информация о городах, которые входят в другие группы из groupCity.
//  Ваша программа имеет доступ к обоим указанным отображениям, требуется исправить cityPopulation,
//  чтобы в ней была сохранена информация только о городах из группы groupCity[100].
func mapRepair(groupCity *map[int][]string, cityPopulation *map[string]int) {
	for i := 10; i <= 1000; i *= 100 {
		for _, cityStr := range (*groupCity)[i] {
			if _, ok := (*cityPopulation)[cityStr]; ok {
				delete(*cityPopulation, cityStr)
			}
		}
	}
}
