// search for duplicate files project doc.go
/*
search for duplicate files document
Программа поиска дубликатов файлов получает на вход путь до целевой директории.
Программа выводит в стандартный поток вывода список дублирующихся файлов,
которые находятся как в целевой директории, переданной через аргумент командной строки, так и в ее поддиректориях.

Флаги программы:

dir - наименование целевого каталога
по умолчанию текущая директория является целевой
допустимо вводить как абсолютные, так и относительные пути к каталогу:

$ go run main.go -dir C:/Source/Golang/Go-level-2/lesson-8/Homework-8, 
$ go run main.go -dir folder, 
$ go run main.go -dir ./folder, 
$ go run main.go -dir ./folder/tmp1

h - вызов справки по программе

d - признак удаления дубликатов найденных файлов с сохранением оригинала
*/
package doc

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

//объявляем структуру описания файла

type myConfig struct {
	MyDir     string // -dir ./ или ./...  folder ./tmp1 ./tmp2
	MyHelp    bool   // -h false true
	MyDel     bool   //-del
	MyDubFile string // список дубликатов файлов
}

//создаем глобальные переменные для парсинга флагов

var (
	//первый параметр флаг, второй параметр значение по дефолту, третий параметр описание
	//библиотека flag всегда возвращает не nil а указатель на дефолтные значения
	DirFlag  = flag.String("dir", "", "directory name")
	HelpFlag = flag.Bool("h", false, "help information")
	DelFlag  = flag.Bool("d", false, "deleting detected duplicate files")
)

// Load
//функция загрузки флагов поиска и удаления дубликатов файлов
func Load() (*myConfig, error) {

	conf := &myConfig{}

	//вызываем обработку флагов
	//если флаг не указан, то под этим флагом будет первоначальное значение по умолчанию - пустая строка
	// библиотека flag всегда возвращает не nil а указатель на дефолтные значения
	flag.Parse()

	//инициализация конфигурации значениями флагов
	conf.MyDir = *DirFlag
	conf.MyHelp = *HelpFlag
	conf.MyDel = *DelFlag
	conf.MyDubFile = ""

	if conf.MyDel {
		var nInp string
		fmt.Printf("Для подтверждения удаление дубликатов файлов нажмите y: \n")
		_, err := fmt.Scanln(&nInp)
		if err != nil {
			return conf, err
		}
		if nInp != "y" {
			conf.MyDel = false
		}
		//fmt.Println("nInp: ", nInp)
	}
	/*
		fmt.Printf("Загружаемые значения флагов: \n")
		fmt.Printf("        Dir: %+v \n", conf.MyDir)
		fmt.Printf("       Help: %+v \n", conf.MyHelp)
		fmt.Printf("        Del: %+v \n", conf.MyDel)
	*/
	//вывод help
	if conf.MyHelp {
		_, err := fmt.Printf(" Программа поиска дубликатов файлов получает на вход путь до целевой директории.\n Программа выводит в стандартный поток вывода список дублирующихся файлов,\n которые находятся как в целевой директории, переданной через аргумент командной строки, так и в ее поддиректориях.\n\n Флаги программы:\n dir - наименование целевого каталога\n по умолчанию текущая директория является целевой\n допустимо вводить как абсолютные, так и относительные пути к каталогу\n $ go run main.go -dir C:/Source/Golang/Go-level-2/lesson-8/Homework-8\n $ go run main.go -dir folder\n $ go run main.go -dir ./folder\n $ go run main.go -dir ./folder/tmp1\n h - вызов справки по программе\n d - признак удаления дубликатов найденных файлов с сохранением оригинала\n")
		return conf, err
	}

	var dirFile string

	if !strings.Contains(conf.MyDir, "/") {
		dirFile = "./" + conf.MyDir
	} else {

		dirFile = conf.MyDir
	}

	err := os.Chdir(dirFile)
	if err != nil {
		return conf, err
	}

	//todo как получить список файлов в каталоге?
	//объявляем структуру описания файла
	/*
			type FileInfo interface {
		    Name() string       // base name of the file
		    Size() int64        // length in bytes for regular files;
	*/

	//производим перебор всех объектов в каталоге и подкаталогов
	// описываем функцию, которая будет анализировать каждый элемент файловой системы, на который наткнется
	// мы вставим эту функцию в качестве аргумента для filepath.Walk.
	// То есть filepath.Walk просматривает каталог (в данном случае - текущий каталог ".") и его подкаталоги
	// а функцния printFiles уже будет делать нужные вещи с найденными файлами

	type fileStruct struct {
		filInf    fs.FileInfo //описание файла
		filePath  string      //путь к файлу
		dub       bool        //признак дубликата
		checkFile int         //счетчик проверяемых файлов
	}

	//объявляем переменную структуры описания файла
	var currentFileStruct fileStruct
	//создаем слайс структур описаний файлов
	var foundFiles = make([]fileStruct, 0, 100)
	//создаем глобальную переменную для счетчика найденных файлов
	iGlob := 0

	//инициализация структуры описания файла
	currentFileStruct.dub = false
	currentFileStruct.checkFile = iGlob

	//============================================

	// описываем функцию, которая будет анализировать каждый элемент файловой системы, на который наткнется
	// мы вставим эту функцию в качестве аргумента для filepath.Walk.

	printFiles := func(path string, info fs.FileInfo, err error) error {

		//получаем путь к файлу

		// проверяем, не является ли текущий элемент каталогом (а нам нужен файл)

		if !info.IsDir() {
			currentFileStruct.filInf = info
			currentFileStruct.checkFile = iGlob
			//fmt.Printf("currentFileStruct.checkFile: %v\n", currentFileStruct.checkFile)
			iGlob++
			//todo=====================================
			//получаем путь к файлу
			currentFileStruct.filePath = path
			//формируем слайс структур со списком и информацией о найденных файлах
			//fmt.Printf("currentFileStruct.filePath: %s\n", currentFileStruct.filePath)
			foundFiles = append(foundFiles, currentFileStruct)

			var wg = sync.WaitGroup{} //создаем переменную для группы горутин

			var mutex sync.Mutex // определяем мьютекс

			//производим поиск дубликатов файлов
			for i := 0; i < len(foundFiles); i++ {

				value1 := foundFiles[i]

				for j := 0; j < len(foundFiles); j++ {

					value2 := foundFiles[j]

					//добавляем горутину в группу wg
					wg.Add(1)
					go func(i, j int, mutex *sync.Mutex) {
						//отложенное закрытие горутины в группе wg
						defer wg.Done()
						mutex.Lock() // блокируем доступ
						if i != j && value1.filInf.Name() == value2.filInf.Name() && value1.filInf.Size() == value2.filInf.Size() {

							var fileStructPointer1 *fileStruct = &foundFiles[i]
							var fileStructPointer2 *fileStruct = &foundFiles[j]
							*&fileStructPointer1.dub = true //признак дубликата файла
							*&fileStructPointer2.dub = true //признак дубликата файла

							//удаление дубликатов с сохранением оригинала при наличии соответствующего флага
							if conf.MyDel {
								if value1.checkFile < value2.checkFile {
									os.Remove(value2.filePath)
								} else {
									os.Remove(value1.filePath)
								}
							}
						}
						mutex.Unlock() // деблокируем доступ
					}(i, j, &mutex) //передаем мьютекс в горутину
				}
			}
			//ожидание завершения работы всех горутин в группе wg
			wg.Wait()
		}
		return nil
	}
	filepath.Walk(".", printFiles)

	//создание списка дублированных файлов

	for _, value1 := range foundFiles {
		if value1.dub {
			//conf.MyDubFile += fmt.Sprintf(" file name: %v size: %v  dub: %v\n", value1.filInf.Name(), value1.filInf.Size(), value1.dub)
			conf.MyDubFile += fmt.Sprintf(" %v\n", value1.filInf.Name())
		}
	}

	if conf.MyDubFile == "" {
		fmt.Println("дубликаты файлов не найдены")
		return conf, err
	}

	return conf, err
}
