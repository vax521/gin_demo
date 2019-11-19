package main

import (
	"errors"
	"fmt"
	"reflect"
)

//实现有限状态机

//状态接口
type State interface {
	//获取状态名字
	Name() string
	//该状态是否允许同状态转移
	EnableSameTransit() bool
	//响应状态开始时
	OnBegin()
	//响应状态结束时
	OnEnd()
	//判断能否转移到某个状态
	CanTransitTo(name string) bool
}

//从状态实例获取状态名
func StateName(s State) string {
	if s == nil {
		return "none"
	}
	//通过反射获取状态名称
	return reflect.TypeOf(s).Elem().Name()
}

type StateInfo struct {
	name string
}

//获取状态名
func (s *StateInfo) Name() string {
	return s.name
}
func (s *StateInfo) setName(name string) {
	s.name = name
}

//允许同种状态转移
func (s *StateInfo) EnableSameTransit() bool {
	return false
}

//默认将状态开启实现
func (s *StateInfo) OnBegin() {

}
func (s *StateInfo) OnEnd() {

}
func (s *StateInfo) CanTransitTo() bool {
	return true
}

type IdleState struct {
	StateInfo //使用StateInfo实现基础接口
}

func (i *IdleState) OnBegin() {
	fmt.Println("Idle State begin")
}
func (i *IdleState) OnEnd() {
	fmt.Println("Idle State end")
}
func (i *IdleState) CanTransitTo(name string) bool {
	return true
}

type MoveState struct {
	StateInfo
}

func (m *MoveState) OnBegin() {
	fmt.Println("Move State Begin")
}
func (m *MoveState) EnableSameTransit() bool {
	return true
}
func (m *MoveState) CanTransitTo(name string) bool {
	return true
}

type JumpState struct {
	StateInfo
}

func (j *JumpState) OnBegin() {
	fmt.Println("Jump State begin")
}
func (j *JumpState) CanTransitTo(name string) bool {
	return name != "MoveState"
}

type StateManager struct {
	//应添加的状态
	stateByName map[string]State
	//状态改变时的回调
	onChange func(from, to State)
	//当前状态
	curr State
}

func (sm *StateManager) Get(name string) State {
	if v, ok := sm.stateByName[name]; ok {
		return v
	}
	return nil
}

//添加一个状态到管理器
func (sm *StateManager) Add(s State) {
	//获取状态的名称
	name := StateName(s)
	//将s转换为能设置名字的接口，然后调用
	s.(interface {
		setName(name string)
	}).setName(name)
	//检查状态是否已存在
	if sm.Get(name) != nil {
		panic("duplicate state:" + name)
	}
	//根据名字保存在map中
	sm.stateByName[name] = s
}

//初始化状态管理器
func NewStateManager() *StateManager {
	return &StateManager{
		stateByName: make(map[string]State),
	}
}

//状态没有找到的错误
var ErrStateNotFound = errors.New("state not found!")

//禁止在同状态间转移
var ErrFobidSameStateTransit = errors.New("cannot transit to state")

//不能转移到指定状态
var ErrCanotTransitToState = errors.New("cannot transit to state")

//当前状态能否转移到目标状态
func (sm *StateManager) CanCurrTransitTo(name string) bool {
	if sm.curr == nil {
		return true
	}
	//相同状态不用转换
	if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
		return false
	}
	return sm.curr.CanTransitTo(name)
}

//转移到指定状态
func (sm *StateManager) Transit(name string) error {
	next := sm.Get(name)
	if next == nil {
		return ErrStateNotFound
	}
	//记录转移前的状态
	pre := sm.curr

	//当前有状态
	if sm.curr != nil {
		//相同状态不用转换
		if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
			return ErrFobidSameStateTransit
		}
		//不能转移到目标状态
		if !sm.curr.CanTransitTo(name) {
			return ErrCanotTransitToState
		}
		//结束当前状态
		sm.curr.OnEnd()
	}
	//将当前状态切换为要转移到的状态
	sm.curr = next
	//新状态开始
	sm.curr.OnBegin()
	//通知回调
	if sm.onChange != nil {
		sm.onChange(pre, sm.curr)
	}
	return nil
}

func main() {
	sm := NewStateManager()
	//响应状态转移的通知
	sm.onChange = func(from, to State) {
		fmt.Printf("%s ----> %s\n", StateName(from), StateName(to))
	}
	//添加三个状态
	sm.Add(new(IdleState))
	sm.Add(new(MoveState))
	sm.Add(new(JumpState))

	//在不同状态间转移
	transitAndReport(sm, "IdleState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "IdleState")

}

//封装转移状态和输出日志
func transitAndReport(sm *StateManager, target string) {
	if err := sm.Transit(target); err != nil {
		fmt.Printf("Faild! %s-->%s, %s\n", sm.curr.Name(), target, err.Error())
	}
}
