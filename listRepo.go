package main

import (
	"time"
	"toDoApp/protos/toDoList"
)

func CreateNewList(name string) string {
	sqlStatement := `insert into to_do_list("name","created_on") values($1,$2)`
	_, e := db.Exec(sqlStatement, name, time.Now())
	if e != nil {
		return "Error while inserting"
	}
	return "ToDo List created"
}

func UpdateListName(id int32, name string) string {
	sqlStatement := `update to_do_list set name = $1 where id = $2`
	_, e := db.Exec(sqlStatement, name, id)
	if e != nil {
		return "Error while altering"
	}
	return "ToDo list name altered"
}

func DeleteList(id int32) string {
	sqlStatement := `delete from to_do_list where id = $1`
	_, e := db.Exec(sqlStatement, id)
	if e != nil {
		return "Error while deleting"
	}
	return "ToDo list deleted"
}

func GetList(id int32) ([]*toDoList.Item, string) {
	sqlStatement := `select * from to_do_item where to_do_list_id = $1`
	resp, e := db.Query(sqlStatement, id)
	if e != nil {
		return nil, "Error while fetching"
	}
	defer resp.Close()
	var itemRows []*toDoList.Item
	for resp.Next() {
		row := toDoList.Item{}
		err := resp.Scan(&row.Id, &row.Name, &row.Description, &row.CreatedOn, &row.IsFinished, &row.ToDoListId)
		if err != nil {
			return nil, "Error while querying"
		}
		itemRows = append(itemRows, &row)
	}
	return itemRows, ""
}
