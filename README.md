1. Create a postgres database named toDoApp
2. Execute the follwoing queries:
```
    create table to_do_list (
	id serial primary key,
	name varchar(50) not null,
	created_on timestamp not null
    )
```
```
    create table to_do_item (
        id serial primary key,
        name varchar(50) not null,
        description varchar(255) not null,
        created_on timestamp not null,
        is_finished boolean not null default false,
        to_do_list_id serial not null references to_do_list(id)
    )
```
3. Run go run . while in the project toDoApp folder
 
NOTE:  I tested it using Postman because I couldn't get it working with grpcurl (searched for the issue and some answers said the issue is connected to Windows)
