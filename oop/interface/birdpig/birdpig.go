package main

import "fmt"

type FLyer interface {
	Fly()
}
type Walker interface {
	Walk()
}

type bird struct {
}

func (b *bird) Fly() {
	fmt.Println("bird:fly")
}
func (b *bird) Walk() {
	fmt.Println("bird:walk")
}

type pig struct {
}

func (p *pig) Walk() {
	fmt.Println("pig:walk")
}

func main() {
	//创建动物名字到实例的映射
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}
	for name, obj := range animals {
		// 类型断言
		f, isFlyer := obj.(FLyer)
		w, isWalker := obj.(Walker)
		fmt.Printf("name %s is Flyer: %v is Walker:%v\n", name, isFlyer, isWalker)
		if isWalker {
			w.Walk()
		}
		if isFlyer {
			f.Fly()
		}
	}
	p1 := new(pig)
	var a Walker = p1
	p2 := a.(*pig)
	fmt.Printf("p1=%p p2=%p", p1, p2)
	p2.Walk()
}
