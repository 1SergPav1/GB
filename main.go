// Напишите программу, которая будет хранить ваши url. На основании созданного шаблона допишите код, который позволяет
//добавлять новые ссылки, удалять и выводить список.
//Для решения задачи используйте структуры. Обязательные поля структуры должны быть дата добавления, имя ссылки,
//теги для ссылки через запятую и сам url.

package main

import (
	"fmt"
	"time"
)

type Item struct {
	Name, URL, Tags, Data string
}

func newItem(url, name, tags string) Item {
	return Item{URL: url, Name: name, Tags: tags, Data: time.Now().Format(time.DateTime)}
}

func (i Item) ShowItem() string{
	return fmt.Sprintf("Имя: <%s>\nURL: <%s>\nТеги: <%s>\nДата: <%s>\n", i.Name, i.URL, i.Tags, i.Data)
}

type URLmap map[string]Item

func (u URLmap) Add(i Item) {
	u[i.Name] = i
}

func (u URLmap) Del(name string) {
	delete(u, name)
}

func (u URLmap) ShowAllURLs() {
	fmt.Printf("\nВсего закладок: %d\n", len(u))
	for _, v := range u {
		fmt.Println(v.ShowItem())
	}
}

func main() {
	myMap := URLmap{}
	myMap.ShowAllURLs()

	it1 := newItem("vk.com", "vk", "social_network")
	it2 := newItem("gitlab.com", "gitlab", "version_control_system")

	myMap.Add(it1)
	myMap.Add(it2)

	myMap.ShowAllURLs()

	myMap.Del("vk")
	myMap.ShowAllURLs()
}
