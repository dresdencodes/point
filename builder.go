package point
// ptr builder

type ptr struct {
	JobKey 			string
	Endpoint 		string
	DataIn 			map[string]interface{}
	CacheKey		string
	ItemKey			string	

	// internal only
	cacheChecked    bool
}

func Ptr(jobKey string) *ptr {
	return &ptr{
		Endpoint:"/p/"+jobKey, 
		JobKey:jobKey,
		DataIn:map[string]interface{}{},
	} 
}

func (p *ptr) ClientCache(cacheKey string, itemKey string) *ptr {
	p.CacheKey = cacheKey
	p.ItemKey = itemKey
	return p
}

func (p *ptr) PostKV(k string, v interface{}) *ptr {
	p.DataIn[k] = v
	return p
}