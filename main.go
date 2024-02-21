// Вам предстоит написать простую версию кеша. Не пугайтесь это самая простая реализация какую можно придумать. Самое
// главное это запомнить зачем нам нужны интерфейсы. Если вы загляните в код, то увидите несколько структур, вам нужно
// доделать реализацию cacheImpl. Она должна хранить все значения, которые вы захотите разместить в этой структуре по
// ключ и значению. Используйте структуры данных, которые знаете.
// В функции main мы создаем некий аналог базы данных куда передаем наш кеш. Вам нужно создать структуру кеша и
// наполнить ее данными, которые добавляются в структуру dbsImpl. Далее проверить, что данные возвращаются с вашего кеша
// *********** Задача со звездочкой ***********
// Задача не сложная, постарайтесь выполнить!
// Реализуйте удаление старых ключей по времени. Для этого вам возможно понадобятся функции работы с временем,а также
// возможно вы захотите создать структуру в которой будут хранится значения и временный метки.
// Удалять ключи можно например при операциях Get. Проверять временную метку и удалять, если ключ старый.


package main

import (
	"fmt"
	"time"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k,v string)}

type cacheDate struct {
	cacheVal string
	date time.Time
}

var _ Cache = (*cacheImpl)(nil)

func newCacheImpl() *cacheImpl {
	return &cacheImpl{
		cacheMap: make(map[string]cacheDate),
	}
}

type cacheImpl struct {
	cacheMap map[string]cacheDate
}

func (c *cacheImpl) Get(key string) (string, bool) {
	for k,v := range c.cacheMap {
		if k == key {
			if time.Since(v.date) > 5 * time.Second {
				delete(c.cacheMap, k)
			}
		}
	}
	v, ok := c.cacheMap[key]
	if ok {
		return fmt.Sprintf("value: %s for key: %s.", v.cacheVal, key), ok
	}
	return fmt.Sprintf("no value for key: %s.", key), ok
}

func (c *cacheImpl) Set(key, value string) {
	c.cacheMap[key] = cacheDate{cacheVal: value, date: time.Now() }
}

func main() {
	myCache := newCacheImpl()
	myCache.Set("hello", "world")
	myCache.Set("test", "my test")
	fmt.Println(myCache.Get("hello"))
	fmt.Println(myCache.Get("test"))
	time.Sleep(6 * time.Second)
	fmt.Println("After 6 second")
	fmt.Println(myCache.Get("hello"))
	fmt.Println(myCache.Get("test"))
}