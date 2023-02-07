package cache

import (
	"sync"
	"time"
)

type container struct {
	Map  		map[string]map[string]interface{}
	MapAddTime 	map[string]map[string]time.Time
	Mu 			*sync.RWMutex
}

var Container = &container{
	Map:map[string]map[string]interface{}{},
	MapAddTime:map[string]map[string]time.Time{},
	Mu:&sync.RWMutex{},
}

func (c *container) Set(mapKey string, key string, item interface{}) { c.Mu.Lock(); defer c.Mu.Unlock()
	if c.Map[mapKey]== nil {
		c.Map[mapKey] = map[string]interface{}{}
		c.MapAddTime[mapKey] = map[string]time.Time{}
	}
	// add item
	c.Map[mapKey][key] = item
	c.MapAddTime[mapKey][key] = time.Now()
}

func (c *container) Get(mapKey string, key string) (i interface{}, ok bool) { c.Mu.RLock(); defer c.Mu.RUnlock()
	// get item
	if _, ok = c.Map[mapKey]; !ok {return nil, false}
	if _, ok = c.Map[mapKey][key]; !ok {return nil, false}
	i = c.Map[mapKey][key]
	return i, true
}

func (c *container) Delete(mapKey string, key string) { c.Mu.Lock(); defer c.Mu.Unlock()
	// delete keys
	delete(c.Map[mapKey], key)
	delete(c.MapAddTime[mapKey], key)
}
