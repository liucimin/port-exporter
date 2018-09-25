package cache

import(

	cmap "github.com/streamrail/concurrent-map"

)


type Cache struct{

	db cmap.ConcurrentMap

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



