package myawesomegopackage

import "strconv"

type Entity struct {
	id int32
}

func New(id int32) Entity {
	return Entity{id: id}
}

func (e Entity) ProcessEntity() {
	println("processEntity " + strconv.Itoa(int(e.id)))
}
