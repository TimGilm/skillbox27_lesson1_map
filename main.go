package main

import "fmt"

// map
func main() {
	m := make(map[int]string)
	fmt.Println(m) //map[]
	//если не инициализируем мар, и будем добавлять в нее данные получим панику в рантайм
	m[1] = "first"
	m[2] = "second"
	fmt.Println(m) //map[1:first 2;second]
	//посмотрим как можем обратиться к ключам в мапе
	fmt.Println(m[2]) //second
	//можно удалять данные карты встроенной функцией delete
	delete(m, 2)
	fmt.Println(m) //map[1:first]
	//рассмотрим мар с другими типами ключа и значения
	mString := make(map[string]float64)
	mString["first"] = 1.0
	mString["second"] = 2.0
	fmt.Println(mString) //map[first:1 second:2]
	//ключами в мар могут быть любые сравниваемые типы, а значениями любые типы, в т.ч. и др. мар
	//рассмотрим способ инициализации карты, в котором в нее сразу можно добавить значения
	//такой способ назывется композитный литерал
	mString = map[string]float64{
		"third":  3.0,
		"fourth": 4.0,
	}
	fmt.Println(mString) //map[third:3 fourth:4]
	//что будет, если попробуем объявить переменную не инициализируя ее?
	var mNil map[int]string
	//при попытке чтения или записи получим панику, рассмотрим пример
	mNil[1] = "first" // panic: assignment to entry in nil map
	//если инициализируем
	var mNil map[int]string = make(map[int]string)
	mNil[1] = "first" // паники не будет
	//встроенная функция len, которая позволяет узнать сколько ключей записано в карту
	mLen := map[int]string{
		1: "first",
		2: "second",
		3: "third",
	}
	fmt.Println(len(mLen)) //3
	//если обратимся к мар по несуществующему ключу, то получим зеро вэлью
	fmt.Println(mLen[4]) //пустая строка
	//иногда мы не хотим получать зеро вэлью, тогда можем воспользоваться следующим синтаксисом
	v, ok := mLen[4] //в такмо случае в v попадет зеро вэлью, а в ok попадет false в случае, если ключ не существует
	//и true, если ключ существует
	fmt.Println(v, ok) //   false - пустая строка и false
	//как это может быть полезно на практике? Мы можем сделать проверку на bool, чтобы не работать с несуществующим ключем
	//для этого можно вывести только ok без значения
	_, ok = mLen[4]
	fmt.Println(ok) //false
}
