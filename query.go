package point

import (
	"log"
	"heiro/pkg/point/cache"
)

var DEBUG = false

func (p *ptr) Query(i interface{}) error {
	if item, ok := p.GetCached(); ok { 
		i = item
		return nil 
	}
	_, err := Pointer(p.Endpoint, p.DataIn, i)
	if err!=nil {
		return err
	}
	if DEBUG {
		log.Println("------------------->>>")
		log.Println("------------------->>>")
		log.Println("------------------->>>")
		log.Println(i, err)
		log.Println("------------------->>>")
		log.Println("------------------->>>")
		log.Println("------------------->>>")
	}
	p.SetCached(i)
	return err
}

func (p *ptr) GetCached() (i interface{}, ok bool) {
	if p.CacheKey == "" || p.ItemKey == "" { return nil, false }
	i, ok = cache.Container.Get(p.CacheKey, p.ItemKey)
	return i, ok
}

func (p *ptr) SetCached(i interface{}) {
	if p.CacheKey == "" || p.ItemKey == "" { return }
	cache.Container.Set(p.CacheKey, p.ItemKey, i)
} 