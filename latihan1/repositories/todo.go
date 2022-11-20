package repositories

import (
	"database/sql"
	"fmt"
	"latihan1/params/request"
	"latihan1/params/views"

	_ "github.com/lib/pq"
)

type TodoRepo struct {
	Db *sql.DB
}

func (r *TodoRepo) InitDB() error {
	db, err := sql.Open("postgres", "postgres://postgres:12345@localhost/postgres?sslmode=disable")
	if err != nil {
		return err
	}
	r.Db = db
	return nil
}

func (r *TodoRepo) GetAllTodos() (data []views.GetTodoPayload, err error) {
	r.InitDB()
	query := "select id, title, description, created_at from todo"

	rows, err := r.Db.Query(query)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		var todo views.GetTodoPayload
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.CreatedAt,
		)
		if err != nil {
			return data, err
		}
		data = append(data, todo)
	}
	return data, nil
}

func (r *TodoRepo) InsertTodo(req request.CreateTodo) (err error) {
	r.InitDB()
	query := fmt.Sprintf("insert into todo(title, description, created_at) values('%s','%s',current_timestamp)", req.Title, req.Description)
	_, err = r.Db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepo) GetTodo(id string) (data views.GetTodoPayload, err error) {
	r.InitDB()
	query := fmt.Sprintf("select id, title, description, created_at from todo where id = %s", id)

	err = r.Db.QueryRow(query).Scan(
		&data.ID,
		&data.Title,
		&data.Description,
		&data.CreatedAt,
	)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *TodoRepo) UpdateTodo(req request.CreateTodo, id string) (err error) {
	r.InitDB()
	query := fmt.Sprintf("update todo set title = '%s', description = '%s' where id = %s", req.Title, req.Description, id)
	_, err = r.Db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepo) DeleteTodo(id string) (err error) {
	r.InitDB()
	query := fmt.Sprintf("delete from todo where id = %s", id)

	_, err = r.Db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
