/*
package main

import (

	"fmt"

)

	func main() {
		// в цикле спрашиваем ввод транзакций
		// добавлять каждую транзакцию в массив
		// вывести массив
		var trn []float64
		var quation bool
		var str string
		for {
			trn = transactions(trn)
			fmt.Println("Хотите продолжить? [y/n]:")
			fmt.Scan(&str)
			switch str {
			case "y", "Y":
				quation = true
			case "n", "N":
				quation = false
			}
			if !quation {
				break
			}
		}
		fmt.Println(trn)
		summ := resultTransactions(trn)
		fmt.Printf("Сумма всех транзакций: %.2f", summ)

}

	func transactions(transactionsUser []float64) []float64 {
		var trn float64
		fmt.Println("Введите транзакцию:")
		fmt.Scan(&trn)
		transactionsUser = append(transactionsUser, trn)

		return transactionsUser
	}

	func resultTransactions(trn []float64) float64 {
		var summ float64
		for i := range trn {
			summ += trn[i]
		}
		return summ
	}
*/
// package main

// import "fmt"

// // **Описание**: Создайте программу для добавления элементов в слайс с помощью функции append
// //
// // **Входные данные**: Используйте готовый слайс чисел []int{5, 15, 25}
// //
// // **Выходные данные**: Выведите исходный слайс, затем слайс после добавления новых элементов
// //
// // **Ограничения**: Добавьте в слайс числа 35 и 45 за одну операцию append
// //
// // **Примеры**:
// // Input: слайс []int{5, 15, 25}
// // Output: исходный слайс: [5 15 25]
// //
// // Входные данные: добавление 35 и 45
// // Output: новый слайс: [5 15 25 35 45]

// func main() {
// 	// Ваш код здесь
// 	sl := []int{5, 15, 25}
// 	res := appending(sl)
// 	fmt.Println(sl)
// 	fmt.Println(res)
// }

//	func appending(sl []int) []int {
//		res := append(sl, 35, 45)
//		return res
//	}
// package main

// import "fmt"

// // **Описание**: Создайте программу для удаления элемента из слайса по указанному индексу
// //
// // **Входные данные**: Используйте готовый слайс строк []string{"apple", "banana", "cherry", "date"} и индекс элемента для удаления: 2
// //
// // **Выходные данные**: Выведите исходный слайс, затем слайс после удаления элемента
// //
// // **Ограничения**: Удалите элемент с индексом 2 ("cherry") путем объединения частей слайса до и после удаляемого элемента
// //
// // **Примеры**:
// // Input: слайс []string{"apple", "banana", "cherry", "date"}, индекс 2
// // Output: исходный слайс: [apple banana cherry date]
// //
// // Входные данные: удаление элемента с индексом 2
// // Output: новый слайс: [apple banana date]

// func main() {
// 	// Ваш код здесь
// 	sl := []string{"apple", "banana", "cherry", "date"}
// 	res := deletedEl(sl)
// 	fmt.Println(sl)
// 	fmt.Println(res)
// }

// func deletedEl(sl []string) []string {
// 	res := append(sl[:2], sl[3:]...)
// 	return res
// }

// package main

// import "fmt"

// // **Описание**: Создайте программу для работы с capacity слайса, которая демонстрирует изменение ёмкости при добавлении элементов
// //
// // **Входные данные**: Используйте готовый слайс чисел []int{1, 2} с начальной capacity 2
// //
// // **Выходные данные**: Выведите длину и capacity слайса до и после добавления элементов
// //
// // **Ограничения**: Добавьте в слайс числа 3, 4, 5 по одному элементу за раз и отслеживайте изменения capacity
// //
// // **Примеры**:
// // Input: слайс []int{1, 2} с cap=2
// // Output: len=2, cap=2
// //
// // Входные данные: добавление элемента 3
// // Output: len=3, cap=4

// func main() {
// 	// Ваш код здесь
// 	sl := []int{1, 2}
// 	apSl := []int{3, 4, 5}
// 	fmt.Printf("len=%v, cap=%v\n", len(sl), cap(sl))
// 	lenCap(sl, apSl)

// }

// func lenCap(sl, apSl []int) {
// 	for i := 0; i <= len(apSl)-1; i++ {
// 		sl = append(sl, apSl[i])
// 		fmt.Printf("len=%v, cap=%v\n", len(sl), cap(sl))
// 	}
// }

// package main

// import "fmt"

// // **Описание**: Создайте программу для работы с опущенными границами слайса, которая демонстрирует создание слайсов с различными вариантами границ
// //
// // **Входные данные**: Используйте готовый массив строк [5]string{"red", "green", "blue", "yellow", "purple"}
// //
// // **Выходные данные**: Выведите исходный массив и три созданных слайса с разными границами
// //
// // **Ограничения**: Создайте три слайса: первые два элемента ([:2]), с третьего до конца ([2:]), и копию всех элементов ([:])
// //
// // **Примеры**:
// // Input: массив ["red", "green", "blue", "yellow", "purple"]
// // Output: исходный массив: [red green blue yellow purple]
// //
// // Входные данные: создание слайсов с опущенными границами
// // Output: первые два: [red green], с третьего: [blue yellow purple], все элементы: [red green blue yellow purple]

// func main() {
// 	// Ваш код здесь
// 	arr := [5]string{"red", "green", "blue", "yellow", "purple"}
// 	fmt.Println(arr)
// 	createSlices(arr)
// }

//	func createSlices(arr [5]string) {
//		// Ваша реализация
//		sl1 := arr[:2]
//		sl2 := arr[2:]
//		sl3 := arr[:]
//		fmt.Println(sl1)
//		fmt.Println(sl2)
//		fmt.Println(sl3)
//	}
// package main

// import "fmt"

// func main() {
// 	/*Создать приложение, которое сначала выдаёт меню:

// 	Посмотреть закладки
// 	Добавить закладку
// 	Удалить закладку
// 	Выход
// 	При 1 - Выводит закладки

// 	При 2 - 2 поля ввода названия и адреса и после добавление

// 	При 3 - Ввод названия и удаление по нему

// 	При 4 - Завершение*/

// 	m := map[string]string{}
// 	for {
// 		num := menu()
// 		b := choice(num, m)
// 		if b {
// 			break
// 		}
// 	}

// }

// func menu() (i int) {
// 	fmt.Println("Выберите пункт из меню 1-4:")
// 	fmt.Println("1)Посмотреть закладки")
// 	fmt.Println("2)Добавить закладку")
// 	fmt.Println("3)Удалить закладку")
// 	fmt.Println("4)Выход")
// 	fmt.Scan(&i)
// 	return i
// }

// func choice(num int, m map[string]string) (str bool) {
// 	switch num {
// 	case 1:
// 		//Выводим закладки
// 		for key := range m {
// 			fmt.Println(key)
// 		}
// 		str = false
// 	case 2:
// 		//2 поля ввода названия и адреса и после добавление
// 		var bookmark, adress string
// 		fmt.Println("Введите название закладки:")
// 		fmt.Scan(&bookmark)
// 		fmt.Println("Введите адрес закладки:")
// 		fmt.Scan(&adress)
// 		m[bookmark] = adress
// 		str = false
// 	case 3:
// 		//Ввод названия и удаление по нему
// 		var bookmark string
// 		fmt.Println("Введите название закладки для удаления:")
// 		fmt.Scan(&bookmark)
// 		delete(m, bookmark)
// 		str = false
// 	case 4:
// 		str = true
// 	}
// 	return str
// }

// package main

// import "fmt"

// type mapStringInt = map[string]int

// func main() {
// 	// **Описание**: Создайте map для хранения информации о товарах и их ценах, добавьте новый товар и выведите обновленную длину map.
// 	// **Входные данные**: Нет (данные встроены в код)
// 	// **Выходные данные**: Длина map до добавления, длина map после добавления
// 	// **Ограничения**:
// 	// - Используйте тип map[string]int
// 	// - Создайте map с 2 парами ключ-значение
// 	// - Добавьте один новый товар
// 	// - Выведите длину до и после добавления

// 	// Ваш код здесь
// 	m := mapStringInt{"Сахар": 25, "Хлеб": 54}
// 	iMap := len(m)
// 	m["Молоко"] = 108
// 	fmt.Printf("len before=%d after=%d\n", iMap, len(m))

// }

// package main

// import "fmt"

// // **Описание**: Создайте map для хранения информации о студентах и их оценках, удалите одного студента и выведите содержимое обновленного map.
// //
// // **Входные данные**: Нет (данные встроены в код)
// //
// // **Выходные данные**: Содержимое map после удаления в формате: имя оценка для каждой пары
// //
// // **Ограничения**:
// // - Используйте тип map[string]int
// // - Создайте map с 3 парами ключ-значение
// // - Удалите одного студента с помощью delete
// // - Имена студентов на английском языке
// // - Оценки от 1 до 100
// //
// // **Примеры**:
// // Output:
// // Alice 85
// // Charlie 92

// type stringIntMap = map[string]int

// func main() {
// 	// Ваш код здесь
// 	m := stringIntMap{"Serge": 24, "Alice": 25, "Max": 38}
// 	delete(m, "Serge")
// 	fmt.Println(m)
// }

// package main

// import "fmt"

// // **Описание**: Создайте map для хранения информации о книгах и их авторах, найдите и выведите автора конкретной книги.
// // **Входные данные**: Нет (данные встроены в код)
// // **Выходные данные**: Автор найденной книги
// // **Ограничения**:
// // - Используйте тип map[string]string
// // - Создайте map с 3 парами ключ-значение (название книги - автор)
// // - Найдите автора книги "1984"
// // - Названия книг и авторы на английском языке

// type mapAutorsStr = map[string]string

// func main() {
// 	// Ваш код здесь
// 	m := mapAutorsStr{"Harry Potter": "J.K. Rowling", "The lord of the rings": "Tolkien", "1984": "George Orwell"}
// 	fmt.Println(m["1984"])
// }

// package main

// import "fmt"

// **Описание**: Создайте map для хранения информации о языках программирования и их годах создания, проверьте существование конкретного языка и выведите результат проверки.
//
// **Входные данные**: Нет (данные встроены в код)
//
// **Выходные данные**: Результат проверки существования языка в формате: язык exists: true/false
//
// **Ограничения**:
// - Используйте тип map[string]int
// - Создайте map с 3 парами ключ-значение (язык - год создания)
// - Проверьте существование языка "Python"
// - Названия языков на английском языке
// - Годы создания реальные
//
// **Примеры**:
// Output:
// Python exists: true

// type mapStrInt = map[string]int

//	func main() {
//		// Ваш код здесь
//		m := mapStrInt{"C#": 2000, "Python": 1991, "Go": 2009}
//		if _, ok := m["Python"]; ok {
//			fmt.Println("Python exists: true")
//		} else {
//			fmt.Println("Python exists: false")
//		}
//	}
// package main

// import "fmt"

// func main() {
// 	a := [4]int{1, 2, 3, 4}
// 	revers(&a)
// 	fmt.Println(a)
// }

// func revers(arr *[4]int) {
// 	for i, value := range *arr {
// 		(*arr)[len(*arr)-1-i] = value
// 	}
// }

// package main

// import "fmt"

// // **Описание**: Создайте программу, которая демонстрирует создание указателя на переменную и вывод её адреса в памяти.
// //
// // **Входные данные**: Программа работает с заранее определенной переменной типа int со значением 42
// //
// // **Выходные данные**: Адрес переменной в памяти в формате "address: 0x..."
// //
// // **Ограничения**: Используйте только базовые операции с указателями
// //
// // **Примеры**:
// // Input: переменная a = 42
// // Output: address: 0xc0000140a8
// //
// // Входные данные: переменная value = 100
// // Output: address: 0xc000014098

// func main() {
// 	a := 42

//		// Ваш код здесь
//		fmt.Printf("address: %p", &a)
//	}
// package main

// import "fmt"

// // **Описание**: Реализуйте программу, которая принимает указатель на переменную типа int и изменяет её значение, удваивая его.
// // **Входные данные**: Программа работает с заранее определенной переменной типа int со значением 15
// // **Выходные данные**: Значение переменной до и после изменения в формате "before: X, after: Y"
// // **Ограничения**: Используйте функцию, которая принимает указатель и изменяет значение через разыменование

// func main() {
// 	num := 15

// 	// Ваш код здесь
// 	double(&num)
// 	fmt.Println(num)
// }

// // Ваша реализация функции здесь
// func double(a *int) {
// 	*a = *a * 2
// }

// package main

// import "fmt"

// // **Описание**: Создайте программу, которая демонстрирует проверку nil-указателя и безопасную инициализацию перед использованием.
// //
// // **Входные данные**: Программа работает с заранее объявленным nil-указателем типа *int
// //
// // **Выходные данные**: Статус указателя и значение после инициализации в формате "pointer is nil: true/false, value: X"
// //
// // **Ограничения**: Используйте проверку на nil и инициализацию указателя перед разыменованием
// //
// // **Примеры**:
// // Input: var p *int (nil-указатель)
// // Output: pointer is nil: true, value: 25
// //
// // Входные данные: var ptr *int (nil-указатель)
// // Output: pointer is nil: true, value: 50

// // Возвращает указатель на int со значением v.
// func newInt(v int) *int {
//     p := new(int)
//     *p = v
//     return p
// }

// func main() {
//     var p *int

//     // Инициализируем через возвращаемое значение
//     p = newInt(42)

//     // p уже не nil
//     fmt.Printf("pointer is nil: %t, value: %d\n", p == nil, *p)
// }

// package main

// import "fmt"

// // **Описание**: Создайте программу, которая демонстрирует передачу указателя в функцию для модификации элемента слайса по индексу.
// //
// // **Входные данные**: Программа работает с заранее определенным слайсом []int{10, 20, 30} и индексом 1
// //
// // **Выходные данные**: Слайс до и после изменения в формате "before: [10 20 30], after: [10 40 30]"
// //
// // **Ограничения**: Используйте функцию, которая принимает указатель на элемент слайса и удваивает его значение
// //
// // **Примеры**:
// // Input: слайс [10, 20, 30], индекс 1
// // Output: before: [10 20 30], after: [10 40 30]
// //
// // Входные данные: слайс [5, 15, 25], индекс 2
// // Output: before: [5 15 25], after: [5 15 50]

// // Ваша функция здесь

// func main() {
// 	// Ваш код здесь
// 	sl := []int{10, 20, 30}
// 	i := 1
// 	double(&sl[i])
// 	fmt.Println(sl)
// }

// func double(i *int) {
// 	*i = *i * 2
// }

// package main

// import "fmt"

// // **Описание**: Создайте программу, которая демонстрирует создание указателя через операцию new() и последующий вывод значения через разыменование.
// //
// // **Входные данные**: Программа работает с указателем, созданным через new(int)
// //
// // **Выходные данные**: Значение по умолчанию и значение после присвоения в формате "default: 0, assigned: 77"
// //
// // **Ограничения**: Используйте функцию new() для создания указателя и операцию * для разыменования
// //
// // **Примеры**:
// // Input: указатель, созданный через new(int)
// // Output: default: 0, assigned: 77
// //
// // Входные данные: указатель через new(int)
// // Output: default: 0, assigned: 123

// func main() {
// 	// Ваш код здесь

// 	i := new(int)
// 	fmt.Printf("default: %d, ", *i)
// 	*i = 77
// 	fmt.Printf("assigned: %d", *i)
// }

// ------------------------------------------
// Генерация пароля
// package main

// import (
// 	"errors"
// 	"fmt"
// 	"math/rand/v2"
// 	"net/url"
// )

// var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890-*!")

// type account struct {
// 	login    string
// 	password string
// 	url      string
// }

// func (acc *account) outputPassword() {
// 	fmt.Println(acc.login, acc.password, acc.url)
// }

// func (acc *account) generetionPassword(n int) {
// 	res := make([]rune, n)
// 	for i := range res {
// 		res[i] = letterRunes[rand.IntN(len(letterRunes))]
// 	}
// 	acc.password = string(res)
// }

// func newAccount(login, password, urlString string) (*account, error) {

// 	if login == "" {
// 		return nil, errors.New("INVALID_LOGIN")
// 	}
// 	_, err := url.ParseRequestURI(urlString)
// 	if err != nil {
// 		return nil, errors.New("INVALID_URL")
// 	}
// 	acc := &account{
// 		login:    login,
// 		password: password,
// 		url:      urlString,
// 	}
// 	if password == "" {
// 		acc.generetionPassword(10)
// 	}
// 	return acc, nil
// }

// func main() {
// 	login := promptData("Введите логин:")
// 	password := promptData("Введите пароль:")
// 	url := promptData("Введите адрес:")
// 	myAccount, err := newAccount(login, password, url)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	myAccount.generetionPassword(10)
// 	//fmt.Println(m)

// 	myAccount.outputPassword()
// }

//	func promptData(str string) string {
//		var res string
//		fmt.Println(str)
//		fmt.Scanln(&res)
//		return res
//	}
// package main

// import "fmt"

// // **Описание**: Создайте структуру User с полями Name (строка), Age (целое число) и Email (строка), затем создайте экземпляр этой структуры и выведите все её поля
// //
// // **Входные данные**: Данные встроены в код - не требуется внешний ввод
// //
// // **Выходные данные**: Вывод всех полей структуры в консоль в формате: Name: [имя], Age: [возраст], Email: [email]
// //
// // **Ограничения**:
// // - Используйте только базовые типы данных Go
// // - Не используйте внешние пакеты
// // - Создайте структуру с именно тремя указанными полями
// //
// // **Примеры**:
// // Входные данные: встроенные значения Name: "Alice", Age: 25, Email: "alice@example.com"
// // Output: Name: Alice, Age: 25, Email: alice@example.com
// //
// // Входные данные: встроенные значения Name: "Bob", Age: 30, Email: "bob@test.com"
// // Output: Name: Bob, Age: 30, Email: bob@test.com

// // Ваш код здесь

// type user struct {
// 	name  string
// 	age   int
// 	email string
// }

// func main() {
// 	// Ваша реализация
// 	p := user{"Andrey", 27, "a@mail.ru"}
// 	fmt.Printf("Name: %v, Age: %d, Email: %v", p.name, p.age, p.email)
// }

// package main

// import (
// 	"fmt"
// )

// // **Описание**: Создайте функцию-конструктор NewProduct, которая принимает название товара, цену и категорию, выполняет валидацию (название и категория не должны быть пустыми строками), и возвращает указатель на структуру Product или строку с ошибкой
// //
// // **Входные данные**: Параметры функции - name (string), price (float64), category (string)
// //
// // **Выходные данные**: Указатель на Product и пустая строка при успехе, или nil и строка с ошибкой при неудаче
// //
// // **Ограничения**:
// // - Используйте только базовые типы данных Go
// // - Валидация должна проверять только пустые строки
// // - Возвращайте строку "empty field" при ошибке валидации
// //
// // **Примеры**:
// // Входные данные: name="Laptop", price=999.99, category="Electronics"
// // Output: указатель на Product{Name: "Laptop", Price: 999.99, Category: "Electronics"}, ""
// //
// // Входные данные: name="", price=50.0, category="Books"
// // Output: nil, "empty field"

// type Product struct {
// 	// Ваш код здесь
// 	name     string
// 	price    float64
// 	category string
// }

// func NewProduct(name string, price float64, category string) (*Product, string) {
// 	// Ваш код здесь

// 	if name == "" {
// 		return nil, "empty field"
// 	}
// 	if category == "" {
// 		return nil, "empty field"
// 	}

// 	return &Product{name: name, price: price, category: category}, ""
// }

// func main() {
// 	// Ваш код здесь
// 	p, err := NewProduct("Laptop", 999.99, "Electronics")
// 	if err != "" {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(*p)
// }

// package main

// import "fmt"

// // **Описание**: Создайте структуру Book с полями Title (строка), Author (строка) и Pages (целое число), затем создайте пустой экземпляр этой структуры и присвойте значения полям через точечную нотацию
// //
// // **Входные данные**: Данные встроены в код - не требуется внешний ввод
// //
// // **Выходные данные**: Вывод всех полей структуры в консоль в формате: Title: [название], Author: [автор], Pages: [страницы]
// //
// // **Ограничения**:
// // - Используйте только базовые типы данных Go
// // - Создайте сначала пустой экземпляр структуры
// // - Присвойте значения полям через точечную нотацию (book.Title = ...)
// // - Не используйте литерал структуры для инициализации
// //
// // **Примеры**:
// // Входные данные: встроенные значения Title: "Go Programming", Author: "John Doe", Pages: 350
// // Output: Title: Go Programming, Author: John Doe, Pages: 350
// //
// // Входные данные: встроенные значения Title: "Learning Structs", Author: "Jane Smith", Pages: 200
// // Output: Title: Learning Structs, Author: Jane Smith, Pages: 200

// // Ваш код здесь

// type book struct {
// 	Title  string
// 	Author string
// 	Pages  int
// }

// func main() {
// 	// Ваш код здесь
// 	p := book{}
// 	p.Title = "Go Programming"
// 	p.Author = "John Doe"
// 	p.Pages = 350
// 	fmt.Printf("Title: %v, Author: %v, Pages: %d", p.Title, p.Author, p.Pages)
// }

// package main

// import "fmt"

// // **Описание**: Создайте структуру Car с полями Brand (строка), Model (строка) и Year (целое число), затем напишите функцию DisplayCar, которая принимает экземпляр Car по значению и выводит информацию о машине в указанном формате
// //
// // **Входные данные**: Данные встроены в код - не требуется внешний ввод
// //
// // **Выходные данные**: Вывод информации о машине в формате: Car: [Brand] [Model] ([Year])
// //
// // **Ограничения**:
// // - Используйте только базовые типы данных Go
// // - Функция должна принимать struct по значению, а не по указателю
// // - Создайте экземпляр структуры с конкретными значениями
// // - Не используйте внешние пакеты
// //
// // **Примеры**:
// // Входные данные: встроенные значения Brand: "Toyota", Model: "Camry", Year: 2020
// // Output: Car: Toyota Camry (2020)
// //
// // Входные данные: встроенные значения Brand: "BMW", Model: "X5", Year: 2022
// // Output: Car: BMW X5 (2022)

// // Ваш код здесь

// type Car struct {
// 	Brand string
// 	Model string
// 	Year  int
// }

// func DisplayCar(p Car) {
// 	fmt.Printf("Car: %v %v %d", p.Brand, p.Model, p.Year)
// }

// func main() {
// 	// Ваш код здесь
// 	p := Car{}
// 	p.Brand = "BMW"
// 	p.Model = "X5"
// 	p.Year = 2022
// 	DisplayCar(p)

// }

// package main

// import "fmt"

// // **Описание**: Создайте структуру Employee с полями Name (строка), Position (строка) и Salary (целое число), затем напишите метод GetInfo, который возвращает отформатированную строку с информацией о сотруднике
// //
// // **Входные данные**: Данные встроены в код - не требуется внешний ввод
// //
// // **Выходные данные**: Строка в формате: "Employee: [Name], Position: [Position], Salary: [Salary]"
// //
// // **Ограничения**:
// // - Используйте только базовые типы данных Go
// // - Метод должен быть привязан к структуре Employee
// // - Создайте экземпляр структуры с конкретными значениями
// // - Не используйте внешние пакеты
// //
// // **Примеры**:
// // Входные данные: встроенные значения Name: "Alice Johnson", Position: "Developer", Salary: 75000
// // Output: "Employee: Alice Johnson, Position: Developer, Salary: 75000"
// //
// // Входные данные: встроенные значения Name: "Bob Smith", Position: "Manager", Salary: 85000
// // Output: "Employee: Bob Smith, Position: Manager, Salary: 85000"

// // Ваш код здесь

// type Employee struct {
// 	Name     string
// 	Position string
// 	Salary   int
// }

// func (e Employee) GetInfo() string {
// 	return fmt.Sprintf("Employee: %s, Position: %s, Salary: %d", e.Name, e.Position, e.Salary)
// }

// func main() {
// 	// Ваш код здесь
// 	e := Employee{Name: "Bob Smith", Position: "Manager", Salary: 85000}
// 	fmt.Println(e.GetInfo())
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// // **Описание**: Создайте программу для чтения JSON-файла с данными аккаунта и вывода информации на экран
// //
// // **Входные данные**: JSON-файл "account.json" со следующей структурой:
// // {
// //   "url": "example.com",
// //   "user": "alice",
// //   "pass": "secret123"
// // }
// //
// // **Выходные данные**: Информация об аккаунте в формате:
// // URL: example.com
// // User: alice
// // Password: secret123
// //
// // **Ограничения**:
// // - Файл "account.json" гарантированно существует
// // - JSON всегда корректный
// // - Все поля обязательны
// //
// // **Примеры**:
// // Входные данные: файл account.json с содержимым {"url": "github.com", "user": "john", "pass": "mypass"}
// // Output:
// // URL: github.com
// // User: john
// // Password: mypass

// type Account struct {
// 	// Ваш код здесь
// 	URL  string `json:"url"`
// 	User string `json:"user"`
// 	Pass string `json:"pass"`
// }

// func main() {
// 	// Ваш код здесь
// 	err := parseJson(&Account{})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func parseJson(acc *Account) error {
// 	jsonData, err := os.ReadFile("account.json")
// 	if err != nil {
// 		return err
// 	}
// 	err = json.Unmarshal(jsonData, &acc)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("URL:", acc.URL)
// 	fmt.Println("User:", acc.User)
// 	fmt.Println("Password:", acc.Pass)
// 	return nil
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// // **Описание**: Создайте программу для записи структуры аккаунта в JSON-файл
// //
// // **Входные данные**: Структура Account с полями URL, User, Pass (значения задаются в коде)
// //
// // **Выходные данные**: JSON-файл "output.json" с данными аккаунта
// //
// // **Ограничения**:
// // - Используйте структурные теги для JSON
// // - Обязательно закройте файл после записи
// // - Обработайте ошибки создания файла и записи
// //
// // **Примеры**:
// // Входные данные: Account{URL: "google.com", User: "bob", Pass: "password123"}
// // Output: файл output.json с содержимым {"url":"google.com","user":"bob","pass":"password123"}

// type Account struct {
// 	// Ваш код здесь - добавьте поля с JSON тегами
// 	URL  string `json:"url"`
// 	User string `json:"user"`
// 	Pass string `json:"pass"`
// }

// func main() {
// 	// Создайте экземпляр Account с данными
// 	acc := Account{
// 		// Ваш код здесь
// 		URL:  "Andrey.com",
// 		User: "Andrey",
// 		Pass: "1w1wd12rfc",
// 	}
// 	file, err := acc.ToBits()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	WriteFile(file, "output.json")
// 	// Ваш код здесь - сериализуйте в JSON и запишите в файл
// 	fmt.Println("Данные записаны в output.json")
// }

// // Преобразование структуры в массив байтов
// func (acc *Account) ToBits() ([]byte, error) {
// 	file, err := json.Marshal(acc)
// 	//defer
// 	if err != nil {
// 		return nil, err
// 	}
// 	return file, nil
// }

// // Запись и создание файла json
// func WriteFile(content []byte, name string) {
// 	file, err := os.Create(name)
// 	if err != nil {
// 		fmt.Println("Ошибка создания файла")
// 		return
// 	}
// 	defer file.Close()
// 	file.Write(content)
// }

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// **Описание**: Создайте программу для добавления нового аккаунта в существующий слайс аккаунтов внутри структуры Vault
//
// **Входные данные**: Структура Vault с полем Accounts (слайс аккаунтов) и новый аккаунт для добавления (значения задаются в коде)
//
// **Выходные данные**: Обновленная структура Vault с добавленным аккаунтом, выведенная в JSON формате
//
// **Ограничения**:
// - Используйте структурные теги для JSON
// - Новый аккаунт добавляется в конец слайса
// - Выведите итоговый JSON на экран
//
// **Примеры**:
// Входные данные: Vault{Accounts: []Account{{URL: "site1.com", User: "user1", Pass: "pass1"}}} и новый Account{URL: "site2.com", User: "user2", Pass: "pass2"}
// Output: {"accounts":[{"url":"site1.com","user":"user1","pass":"pass1"},{"url":"site2.com","user":"user2","pass":"pass2"}]}

type Account struct {
	// Ваш код здесь
	URL  string `json:"url"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

type Vault struct {
	// Ваш код здесь
	Accounts []Account `json:"accounts"`
}

// Переводим в байты
func (acc *Account) ToBits() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func main() {
	// Ваш код здесь
	acc := Account{
		URL:  "site2.com",
		User: "user2",
		Pass: "pass2",
	}
	file, err := acc.ToBits()
	if err != nil {
		fmt.Println(err)
	}
	CreatedFiles(file, "account.json")
}

// Создание или добавление в файл
func CreatedFiles(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Файл записан")
}
