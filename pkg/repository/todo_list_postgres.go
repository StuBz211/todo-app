package repository

import (
	"fmt"
	"strings"

	"github.com/StuBz211/todo-app"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(UserId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuert := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListTable)
	row := tx.QueryRow(createListQuert, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersList := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListTable)
	_, err = tx.Exec(createUsersList, UserId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()

}

func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, title, description FROM %s tl INNER JOIN %s ul ON tl.id=ul.list_id WHERE ul.user_id=$1", todoListTable, usersListTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf(`SELECT tl.id, title, description FROM %s tl INNER JOIN %s ul ON tl.id=ul.list_id 
	WHERE ul.user_id=$1 AND ul.list_id=$2`, todoListTable, usersListTable)
	err := r.db.Get(&list, query, userId, listId)
	return list, err

}

func (r *TodoListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf(`
	DELETE FROM %s tl USING %s ul 
	WHERE tl.id=ul.list_id AND ul.user_id=$1 AND ul.list_id=$2
	`, todoListTable, usersListTable)

	_, err := r.db.Exec(query, userId, listId)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoListPostgres) Update(userId, listId int, input todo.UpdateListInput) error {

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s tl SET %s FROM %s ul 
		WHERE tl.id=ul.list_id AND ul.user_id=$%v AND ul.list_id=$%v`,
		todoListTable, setQuery, usersListTable, argId, argId+1)

	args = append(args, userId, listId)

	_, err := r.db.Exec(query, args...)
	return err
}
