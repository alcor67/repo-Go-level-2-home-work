// printStruct_4 project main.go
//https://go.dev/play/p/1p1BKh7aw5m
package printStruct_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func printStruct(in interface{}, values map[string]interface{}) error {
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

			if typeField.Type.Kind() == reflect.Struct {
				fmt.Printf("\tnested field вложенное поле: %v", typeField.Name)
				//если это вложенное поле, то вызываем рекурсию нашей функции printStruct()
				err := printStruct(val.Field(i).Interface(), values)
				if err != nil {
					return err
				}
				continue
			}

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

func Test_TestifyprintStruct(t *testing.T) {
	v := struct {
		FieldString string //`json:"field_string"`
		FieldInt    int
		Key         string //`structure field: Key`
		Slice       []int
	}{
		FieldString: "My string",
		FieldInt:    107,
		Key:         "old value",
		Slice:       []int{112, 107, 207},
	}
	values := make(map[string]interface{}) //создаем map
	values["value"] = "new value"          //присваиваем полю value значение "new value"
	var expected error = nil
	received := printStruct(&v, values) //передадим структуру по указателю
	assert.Equal(t, expected, received, "they should be equal")
}

func Test_TestifyprintStructFailureStructKey(t *testing.T) {
	v := struct {
		FieldString string //`json:"field_string"`
		FieldInt    int
		Key1        string //`structure field: Key`
		Slice       []int
	}{
		FieldString: "My string",
		FieldInt:    107,
		Key1:        "old value",
		Slice:       []int{112, 107, 207},
	}
	values := make(map[string]interface{}) //создаем map
	values["value"] = "new value"          //присваиваем полю value значение "new value"
	var expected error = nil
	received := printStruct(&v, values) //передадим структуру по указателю
	if received != nil {
		received = nil
	} else {
		t.Errorf("%s expected to return error, but received no error", received)
	}
	assert.Equal(t, expected, received, "they should not be equal")
}

func Test_TestifyprintStructFailureMapValue(t *testing.T) {
	v := struct {
		FieldString string //`json:"field_string"`
		FieldInt    int
		Key         string //`structure field: Key`
		Slice       []int
	}{
		FieldString: "My string",
		FieldInt:    107,
		Key:         "old value",
		Slice:       []int{112, 107, 207},
	}
	values := make(map[string]interface{}) //создаем map
	values["value1"] = "new value"         //присваиваем полю value значение "new value"
	var expected error = nil
	received := printStruct(&v, values) //передадим структуру по указателю
	if received != nil {
		received = nil
	} else {
		t.Errorf("%s expected to return error, but received no error", received)
	}
	assert.Equal(t, expected, received, "they should not be equal")
}

func Example_printStruct() {
	v := struct {
		FieldString string //`json:"field_string"`
		FieldInt    int
		Key         string //`structure field: Key`
		Slice       []int
	}{
		FieldString: "My string",
		FieldInt:    107,
		Key:         "old value",
		Slice:       []int{112, 107, 207},
	}
	values := make(map[string]interface{}) //создаем map
	values["value"] = "new value"          //присваиваем полю value значение "new value"
	received := printStruct(&v, values)    //передадим структуру по указателю
	fmt.Println(received)
	// Output: <nil>
}

func Example_printStructValueFailure() {
	v := struct {
		FieldString string //`json:"field_string"`
		FieldInt    int
		Key         string //`structure field: Key`
		Slice       []int
	}{
		FieldString: "My string",
		FieldInt:    107,
		Key:         "old value",
		Slice:       []int{112, 107, 207},
	}
	values := make(map[string]interface{}) //создаем map
	values["value1"] = "new value"         //присваиваем полю value значение "new value"
	received := printStruct(&v, values)    //передадим структуру по указателю
	fmt.Println(received)
	// Output: error: в map отсутствует ключ value
}

func Example_printStructKeyFailure() {
	v := struct {
		FieldString string //`json:"field_string"`
		FieldInt    int
		Key1        string //`structure field: Key`
		Slice       []int
	}{
		FieldString: "My string",
		FieldInt:    107,
		Key1:        "old value",
		Slice:       []int{112, 107, 207},
	}
	values := make(map[string]interface{}) //создаем map
	values["value"] = "new value"          //присваиваем полю value значение "new value"
	received := printStruct(&v, values)    //передадим структуру по указателю
	fmt.Println(received)
	// Output: error: в структуре отсутствует поле Key
}
