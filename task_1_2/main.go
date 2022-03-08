// PrintStruct_4 project main.go
//https://go.dev/play/p/1p1BKh7aw5m
package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

func PrintStruct(in interface{}, values map[string]interface{}) error {
	if in == nil { //Проверка на то, что входящий параметр не nil.
		return errors.New("error: входящий параметр nil")
	}

	val := reflect.ValueOf(in) //получим значение интерфейса (возвращает значение входящего параметра)

	if val.Kind() == reflect.Ptr { //val проверяем на то, что это указатель.
		val = val.Elem() //если это указатель, то используем функцию Elem
	}

	if val.Kind() != reflect.Struct { //val проверяем на то, что его первоначальный тип это структура
		//иначе ниже идущая функция reflect.NumField вернет панику, если элемент входящий не является структурой
		return errors.New("error: первоначальный тип это не структура")
	}

	if valMap, ok := values["value"]; ok {

		//функция reflect.NumField возвращает количество полей входящей структуры
		//функця reflect.NumField вернет панику, если элемент входящий не является структурой
		for i := 0; i < val.NumField(); i++ {
			typeField := val.Type().Field(i) //получаем тип i-го поля структуры val
			//typeField проверяем на то, что его первоначальный тип i-го поля это структура
			/*
				if typeField.Type.Kind() == reflect.Struct {
					fmt.Printf("\tnested field вложенное поле: %v", typeField.Name)
					//если это вложенное поле, то вызываем рекурсию нашей функции PrintStruct()
					err := PrintStruct(val.Field(i).Interface(), values)
					if err != nil {
						return err
					}
					continue
				}
			*/
			if typeField.Name == "Key" {

				strValMap, ok := valMap.(string)
				if !ok {
					return errors.New("error: could not assert value to string")
				}
				val.Field(i).SetString(strValMap)
				return nil
			}

		}
		return errors.New("error: в структуре отсутствует поле Key")
	}
	return errors.New("error: в map отсутствует ключ value")
}

func main() {
	v := struct {
		FieldString string //`json:"field_string"`
		FieldInt    int
		Key         string //`structure field: Key`
		Slice       []int
		Object      struct {
			NestedField int
			Key1        string //`structure nested field: Key`
		}
	}{
		FieldString: "My string",
		FieldInt:    107,
		Key:         "old value",
		Slice:       []int{112, 107, 207},
		Object: struct {
			NestedField int
			Key1        string
		}{
			NestedField: 302,
			Key1:        "nested field old value",
		},
	}

	values := make(map[string]interface{}) //создаем map
	values["value"] = "new value"          //присваиваем полю value значение "new value"
	/*
		var values map[string]interface{} //объявляем map
		values = map[string]interface{}{} //инициализируем map
		values["value"]="new value" //присваиваем полю value значение "new value"
	*/
	err := PrintStruct(&v, values) //передадим структуру по указателю
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println("цикл по полям структуры и чтение тегов полей структуры прошли успешно.")
	fmt.Printf("Новое значение поля Key: %+v\n", v.Key)
	//fmt.Printf("Новые значение полей структуры: %+v\n", v)
}
