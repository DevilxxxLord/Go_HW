package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var mp map[string]func(sl []int) (int, string, error)

func main() {
	// Принимает операцию (AVG - среднее, SUM - сумму, MED - медиану)
	// Принимает неограниченное число чисел через запятую (2, 10, 9)
	// Разбивает строку чисел по запятым и затем делает расчёт в зависимости от операции выводя результат
	var sl []int
	var choise int
	var str string
	mp = map[string]func([]int) (int, string, error){
		"AVG": func(sl []int) (res int, str string, err error) {
			for _, value := range sl {
				res = res + value
			}
			res = res / len(sl)

			return res, "AVG", nil
		},
		"SUM": func(sl []int) (res int, str string, err error) {
			for _, value := range sl {
				res = res + value
			}
			return res, str, nil
		},
		"MED": func(sl []int) (res int, str string, err error) {
			sort.Ints(sl)
			n := len(sl)
			if n%2 == 1 {
				//Число нечетное
				res = sl[n/2]
			} else {
				res = (sl[n/2-1] + sl[n/2]) / 2
			}

			return res, str, nil
		},
	}

	fmt.Println("Введите числа через запятую:")
	fmt.Scan(&str)
	fmt.Println("Выберите операцию по числам 1-3:")
	fmt.Println("1) AVG")
	fmt.Println("2) SUM")
	fmt.Println("3) MED")
	fmt.Scan(&choise)   //выбор операции
	sl = scanSlice(str) //приведение к слайсу
	//i, name, err := operations(choise, sl)
	nameOperation := convertString(choise)
	i, name, err := mp[nameOperation](sl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Результат операции %v: %d", name, i)
}
func convertString(i int) string {
	switch i {
	case 1:
		return "AVG"

	case 2:
		return "SUM"
	case 3:
		return "MED"
	}
	return ""
}

// func operations(i int, sl []int) (int, string, error) {
// 	var err error
// 	var res int
// 	var str string

// 	switch i {
// 	case 1:
// 		str = "AVG"

// 		for _, value := range sl {
// 			res = res + value
// 		}
// 		res = res / len(sl)

// 		return res, str, nil
// 		//AVG = сумма чисел/кол-во чисел
// 	case 2:
// 		str = "SUM"
// 		for _, value := range sl {
// 			res = res + value
// 		}
// 		return res, str, nil //SUM
// 	case 3:
// 		str = "MED"

// 		sort.Ints(sl)
// 		n := len(sl)
// 		if n%2 == 1 {
// 			//Число нечетное
// 			res = sl[n/2]
// 		} else {
// 			res = (sl[n/2-1] + sl[n/2]) / 2
// 		}

// 		return res, str, nil
// 		//MED
// 	default:
// 		err = errors.New("Нет такой операции")

// 	}
// 	return 0, str, err
// }

func scanSlice(str string) []int {
	var sl []int

	re := regexp.MustCompile(`\s+`)
	result := re.ReplaceAllString(str, "")
	s := strings.Split(result, ",")

	for i := 0; i < len(s); i++ {
		num, err := strconv.ParseInt(s[i], 10, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}
		sl = append(sl, int(num))
	}
	return sl
}
