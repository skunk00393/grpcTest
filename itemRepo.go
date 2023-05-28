package main

import (
	"time"
	"toDoApp/protos/toDoItem"
)

func DeleteItem(id int32) string {
	sqlStatement := `delete from to_do_item where id = $1`
	_, e := db.Exec(sqlStatement, id)
	if e != nil {
		return "Error while deleting"
	}
	return "ToDo item deleted"
}

func UpdateItem(id int32, name string) string {
	sqlStatement := `update to_do_item set name = $1 where id = $2`
	_, e := db.Exec(sqlStatement, name, id)
	if e != nil {
		return "Error while altering"
	}
	return "ToDo item name altered"
}

func MarkItem(id int32) string {
	sqlStatement := `update to_do_item set is_finished = not is_finished where id = $1`
	_, e := db.Exec(sqlStatement, id)
	if e != nil {
		return "Error while altering"
	}
	return "ToDo item name altered"
}

func CreateItem(id int32, item *toDoItem.ListItem) string {
	sqlStatement := `insert into to_do_item ("name","description","created_on","to_do_list_id") values($1,$2,$3,$4)`
	_, e := db.Exec(sqlStatement, item.Name, item.Description, time.Now(), id)
	if e != nil {
		return e.Error()
	}

	return "ToDo item created"
}
