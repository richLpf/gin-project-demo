package products

import "errors"

//Product 实现一个抽象产品
type Product interface {
	SetName(name string)
	GetName() string
}

//Product1 实现具体的产品1
type Product1 struct {
	name string
}

//SetName 设置名称
func (p1 *Product1) SetName(name string) {
	p1.name = name
}

//GetName 获取名称
func (p1 *Product1) GetName() string {
	return "产品2的name为" + p1.name
}

//Product2 实现具体的产品2
type Product2 struct {
	name string
}

//SetName 设置名称
func (p2 *Product2) SetName(name string) {
	p2.name = name
}

//GetName 设置名称
func (p2 *Product2) GetName() string {
	return "产品2的name为" + p2.name
}

type productType int

const (
	p1 productType = iota //0
	p2                    // 1
)

//productFactory 实现简单的工厂类
type productFactory struct {
}

func (pf productFactory) Create(productType productType) Product {
	if productType == p1 {
		return &Product1{}
	}
	if productType == p2 {
		return &Product2{}
	}
	return nil
}

//Cache 定义一个Cache接口，作为父类
type Cache interface {
	Set(key, value string)
	Get(key string) string
}

//RedisCache 实现具体的Cache: RedisCache
type RedisCache struct {
	data map[string]string
}

//Set 方法
func (r *RedisCache) Set(key, value string) {
	r.data[key] = value
}

//Get 方法
func (r *RedisCache) Get(key string) string {
	return r.data[key]
}

//MemCache 实现具体的Cache: MemCache
type MemCache struct {
	data map[string]string
}

//Set memChche方法
func (m *MemCache) Set(key, value string) {
	m.data[key] = value
}

//Get memChche方法
func (m *MemCache) Get(key string) string {
	return m.data[key]
}

type cacheType int

const (
	redis cacheType = iota
	mem
)

//CacheFactory 实现简单工厂
type CacheFactory struct {
}

//Create 工厂创建
func (factory *CacheFactory) Create(cacheType cacheType) (Cache, error) {
	if cacheType == redis {
		return &RedisCache{
			data: map[string]string{},
		}, nil
	}
	if cacheType == mem {
		return &MemCache{
			data: map[string]string{},
		}, nil
	}
	return nil, errors.New("create fail")
}
