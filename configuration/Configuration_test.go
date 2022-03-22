package configuration_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/alcor67/repo-Go-level-2-home-work/configuration"
	"github.com/stretchr/testify/assert"
)

//todo тесты не работают пакетом

func TestConfSmoke(t *testing.T) {
	//устанавливаем переменные окружения
	os.Setenv("MY_DIR", "testfolder")

	received, err := configuration.Load()
	if err != nil {
		t.Errorf("configuration.Load() expected to return no error, but received: %s", err)
	}
	expected := fmt.Sprintf("file1.txt\nfile1.txt\nfile1.txt\n")
	if received.MyDubFile != expected {
		t.Errorf("configuration.Load() expected to return:\n %v\n but received:\n %v\n", expected, received.MyDubFile)
	}
}

func Example_confDub() {
	//устанавливаем переменные окружения
	os.Setenv("MY_DIR", "testfolder")

	received, _ := configuration.Load()
	//	received2 :=  strings.Replace(received.MyDubFile, "\n", ",", 3)
	//	fmt.Println(received2)
	fmt.Println(received.MyDubFile)
	// Output: os.Getenv(MY_DIR):  testfolder
	//file1.txt
	//file1.txt
	//file1.txt
}

func Example_confEmpt() {
	//устанавливаем переменные окружения
	os.Setenv("MY_DIR", "testfolderempty")

	received, _ := configuration.Load()
	received2 := received.MyDubFile
	fmt.Println(received2)
	// Output: os.Getenv(MY_DIR):  testfolderempty
	//дубликаты файлов не найдены

}

func TestTestifyConf(t *testing.T) {

	//устанавливаем переменные окружения
	os.Setenv("MY_DIR", "testfolder")

	received, err := configuration.Load()
	if err != nil {
		t.Errorf("configuration.Load() expected to return no error, but received: %s", err)
	}
	received1 := received.MyDubFile
	expected := fmt.Sprintf("file1.txt\nfile1.txt\nfile1.txt\n")
	assert.Equal(t, expected, received1, "they should be equal")
}

func TestConfFailureDir(t *testing.T) {
	//устанавливаем неправильное имя каталога
	invalidFolderName := "testfolder1"
	//устанавливаем переменные окружения
	os.Setenv("MY_DIR", "testfolder1")
	_, err := configuration.Load()
	if err == nil {
		t.Errorf("invalid folder: %s\n expected to return error, but received no error", invalidFolderName)
	}
}

func TestConfFile(t *testing.T) {
	//устанавливаем переменные окружения
	os.Setenv("MY_DIR", "testfolder")

	received, err := configuration.Load()
	if err != nil {
		t.Errorf("configuration.Load() expected to return no error, but received: %s", err)
	}
	received1 := string(received.MyDubFile)
	expected := fmt.Sprintf("file1.txt\nfile1.txt\nfile1.txt\n")
	if received1 != expected {
		t.Errorf("configuration.Load() expected to return error but received no error:\n expected:\n%v\n received:\n%v\n", expected, received1)
	}
}
func TestConfFailureFile(t *testing.T) {
	//устанавливаем переменные окружения
	os.Setenv("MY_DIR", "testfolder")

	received, err := configuration.Load()
	if err != nil {
		t.Errorf("configuration.Load() expected to return no error, but received: %s", err)
	}
	received1 := string(received.MyDubFile)
	expected := fmt.Sprintf("file1.txt\nfile1.txt\nfile2.txt\n")
	if received1 == expected {
		t.Errorf("configuration.Load() expected to return error but received no error:\n expected:\n%v\n received:\n%v\n", expected, received1)
	}
}
