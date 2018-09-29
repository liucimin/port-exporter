package cache

import(
	"reflect"
	cmap "github.com/streamrail/concurrent-map"

)


type Cache struct{

	db cmap.ConcurrentMap

}

func NewCache() *Cache{
	return &Cache{db:cmap.New()}
}


func (self *Cache) SetFromList(values []interface{}){


	//
	for value := range values{

		t := reflect.ValueOf(value)
		fieldSturct := t.FieldByName("Name")
		if !fieldSturct.IsValid() || !fieldSturct.IsNil(){
			continue
		}
		key := fieldSturct.String()
		self.db.Set(key,value)

	}
}


func (self *Cache) Set(id string, value interface{}){


	self.db.Set(id, value)


}

func (self *Cache) Get(id string) interface{}{


	containerInfo, ok := self.db.Get(id)
	if ok {
		return containerInfo
	}

}



