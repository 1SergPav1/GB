// Напишите программу, которая будет хранить ваши url. На основании созданного шаблона допишите код, который позволяет
//добавлять новые ссылки, удалять и выводить список.
//Для решения задачи используйте структуры. Обязательные поля структуры должны быть дата добавления, имя ссылки,
//теги для ссылки через запятую и сам url.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type Item struct {
	Link string
	URL  string
	Tags string
	Date string
}

func createItem(url, link, tags string) Item {
	date := time.Now().Format(time.DateTime)
	return Item{URL: url, Link: link, Tags: tags, Date: date}
}

func (i Item) ShowItem() {
	fmt.Printf("Имя: <%s>\nURL: <%s>\nТеги: <%s>\nДата: <%s>\n", i.Link, i.URL, i.Tags, i.Date)
}

type URLmap map[string]Item

func (u URLmap) AddURL(i Item) {
	u[i.Link] = i
}

func (u URLmap) DelURL(name string) {
	delete(u, name)
}

func (u URLmap) ShowAllURLs() {
	fmt.Printf("\nВсего закладок: %d\n", len(u))
	for _, v := range u {
		v.ShowItem()
	}
}

func main() {
	defer func() {
		// Завершаем работу с клавиатурой при выходе из функции
		_ = keyboard.Close()
	}()

	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для выхода и приложения нажмите Esc")

	myMap := URLmap{}

OuterLoop:
	for {
		// Подключаем отслеживание нажатия клавиш
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			// Добавление нового url в список хранения
			fmt.Println("Введите новую запись в формате <url описание теги>")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Введите правильный аргументы в формате url описание теги")
				continue OuterLoop
			}

			newURL := createItem(args[0], args[1], args[2])
			myMap.AddURL(newURL)
		case 'r':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Удаление url из списка хранения
			fmt.Println("Введите имя ссылки, которое нужно удалить")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			_ = text

			myMap.DelURL(text)
		case 'l':
			// Вывод списка добавленных url. Выведите количество добавленных url и список с данными url
			myMap.ShowAllURLs()

		default:
			// Если нажата Esc выходим из приложения
			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}
