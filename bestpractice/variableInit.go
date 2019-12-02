package main

import "fmt"

type Foo struct {
	key    string
	option Option
}

type Option struct {
	num int
	str string
}

func New(option Option) *Foo {
	return &Foo{
		option: option,
	}
}

type ModOption func(option *Option)

func NewMod(modOption ModOption) *Foo {
	option := Option{
		num: 100,
		str: "heelo",
	}
	modOption(&option)
	return &Foo{
		option: option,
	}
}

func NewModKey(key string, modOption ModOption) *Foo {
	option := Option{
		num: 100,
		str: "hello",
	}
	modOption(&option)
	return &Foo{
		key:    key,
		option: option,
	}
}

func NewModWith(key string, modOptions ...ModOption) *Foo {
	option := Option{
		num: 100,
		str: "hello",
	}
	for _, fn := range modOptions {
		fn(&option)
	}
	return &Foo{
		key:    key,
		option: option,
	}
}

func withNum(num int) ModOption {
	return func(option *Option) {
		option.num = num
	}
}
func withStr(str string) ModOption {
	return func(option *Option) {
		option.str = str
	}
}
func main() {
	foo := New(Option{
		str: "world",
	})
	fmt.Println(foo)

	foo1 := NewMod(func(option *Option) {
		option.num = 200
	})
	fmt.Println(foo1)

	fooKey := NewModKey("key", func(option *Option) {
		option.num = 300
	})
	fmt.Println(fooKey)

	fooWith := NewModWith("hello", withNum(3), withStr("world"))
	fmt.Println(fooWith)

}
