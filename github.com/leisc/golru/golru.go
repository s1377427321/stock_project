/*
 * @Description:golang lru
 * @Author: leisc
 * @Date: 2018-11-17 10:57:22
 * @LastEditTime: 2018-11-17 19:04:42
 */

package golru

import (
	"container/list"
	"errors"
	"sync"
)

type GOLRUSync struct {
	cap      int
	lock     sync.Mutex
	dataList *list.List
	dataMap  map[interface{}]*list.Element
}

type LRUNode struct {
	Lkey   interface{}
	Lvalue interface{}
}

func NewLRUNode(k, v interface{}) *LRUNode {
	return &LRUNode{
		k,
		v,
	}
}

func New(cap int) (*GOLRUSync, error) {
	if cap <= 0 {
		return nil, errors.New("capacity must greater than 0")
	}

	golru := &GOLRUSync{
		cap:      cap,
		dataList: list.New(),
		dataMap:  make(map[interface{}]*list.Element, cap),
	}

	return golru, nil
}

func (golru *GOLRUSync) Clear() {
	golru.lock.Lock()
	defer golru.lock.Unlock()

	golru.dataList = list.New()
	golru.dataMap = make(map[interface{}]*list.Element, golru.cap)
}

func (golru *GOLRUSync) ClearWithCapacity(cap int) {
	golru.lock.Lock()
	defer golru.lock.Unlock()

	golru.dataList = list.New()
	golru.dataMap = make(map[interface{}]*list.Element, cap)
}

func (golru *GOLRUSync) AddNode(node *LRUNode) error {
	if node == nil {
		return errors.New("node error, node is nil")
	}

	golru.lock.Lock()
	defer golru.lock.Unlock()
	//check datamap and datalist
	if golru.dataList == nil {
		golru.dataList = list.New()
	}

	if golru.dataMap == nil {
		golru.dataMap = make(map[interface{}]*list.Element, golru.cap)
	}

	//check exist node
	if dElement, ok := golru.dataMap[node.Lkey]; ok {

		golru.dataList.MoveToBack(dElement)
		dElement.Value.(*LRUNode).Lvalue = dElement

		return nil
	}

	//add new node
	newNode := &LRUNode{
		Lkey:   node.Lkey,
		Lvalue: node.Lvalue,
	}
	//back insert
	newElement := golru.dataList.PushBack(newNode)
	golru.dataMap[node.Lkey] = newElement

	//check size
	if golru.dataList.Len() > golru.cap {
		//front pop
		oldElement := golru.dataList.Front()
		if oldElement == nil {
			return nil
		}

		golru.removeElement(oldElement)
	}

	return nil
}

func (golru *GOLRUSync) RemoveNode(key interface{}) error {
	golru.lock.Lock()
	defer golru.lock.Unlock()

	if golru.dataMap == nil {
		return nil
	}

	if element, ok := golru.dataMap[key]; ok {
		golru.removeElement(element)
	}

	return nil
}

func (golru *GOLRUSync) GetNode(key interface{}) (node *LRUNode) {
	golru.lock.Lock()
	defer golru.lock.Unlock()

	if golru.dataMap == nil {
		return nil
	}

	if ele, ok := golru.dataMap[key]; ok {
		return &LRUNode{
			Lkey:   key,
			Lvalue: ele.Value,
		}
	}
	return nil
}

func (golru *GOLRUSync) removeElement(element *list.Element) {

	node := element.Value.(*LRUNode)

	golru.dataList.Remove(element)

	delete(golru.dataMap, node.Lkey)
}

func (golru *GOLRUSync) Keys() []interface{} {
	golru.lock.Lock()
	defer golru.lock.Unlock()

	keys := make([]interface{}, len(golru.dataMap))
	i := 0
	for k := range golru.dataMap {
		keys[i] = k
		i++
	}

	return keys
}

func (golru *GOLRUSync) Size() int {
	return golru.dataList.Len()
}

func (golru *GOLRUSync) Cap() int {
	return golru.cap
}

func (golru *GOLRUSync) ResetCap(newcap int) {
	golru.cap = newcap
}
