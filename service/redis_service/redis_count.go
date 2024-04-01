package redis_service

import (
	"fmt"
	"gvd_server/global"
	"strconv"
)

const (
	docDiggIndex = "docDiggIndex"
	docLookIndex = "docLookIndex"
)

func NewDocDigg() CountDB {
	return CountDB{
		Index: docDiggIndex,
	}
}

func NewDocLook() CountDB {
	return CountDB{
		Index: docLookIndex,
	}
}

type CountDB struct {
	Index string
}

func (c CountDB) SetById(id uint) error {
	return c.Set(fmt.Sprintf("%d", id))
}

// Set 给某一个key设置数字 调用一次+1
func (c CountDB) Set(key string) error {
	return c.SetCount(key, 1)
}

// SetCount 给某一个key设置数字,调用一次 + num
func (c CountDB) SetCount(key string, num int) error {
	oldNum, _ := global.Redis.HGet(c.Index, key).Int()
	newNum := oldNum + num
	err := global.Redis.HSet(c.Index, key, newNum).Err()
	return err
}

// Get 返回某一个key对应的值
func (c CountDB) Get(key string) int {
	num, _ := global.Redis.HGet(c.Index, key).Int()
	return num
}

// GetById 返回某一个key对应的值
func (c CountDB) GetById(id uint) int {
	return c.Get(fmt.Sprintf("%d", id))
}

// GetAll 返回这个索引下的全部数据
func (c CountDB) GetAll() map[string]int {
	var countMap = map[string]int{}
	maps := global.Redis.HGetAll(c.Index).Val()
	for key, val := range maps {
		num, _ := strconv.Atoi(val)
		countMap[key] = num
	}
	return countMap
}

// Clear 清空这个索引里面的值
func (c CountDB) Clear() {
	global.Redis.Del(c.Index)
}
