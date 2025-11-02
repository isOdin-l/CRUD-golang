package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/types/databaseTypes"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoListRepository struct {
	db *pgxpool.Pool
}

func NewTodoListRepository(db *pgxpool.Pool) *TodoListRepository {
	return &TodoListRepository{db: db}
}

func (r *TodoListRepository) CreateList(userId uuid.UUID, list databaseTypes.TodoList) (uuid.UUID, error) {
	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return uuid.Nil, err
	}

	var todoListid uuid.UUID
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", databaseTypes.TableTodoLists)

	rowCreateList := tx.QueryRow(context.Background(), createListQuery, list.Title, list.Description)
	if err := rowCreateList.Scan(&todoListid); err != nil {
		tx.Rollback(context.Background())
		return uuid.Nil, err
	}

	createUserListRelationQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", databaseTypes.TableUsersLists)
	_, err = tx.Exec(context.Background(), createUserListRelationQuery, userId, todoListid)
	if err != nil {
		tx.Rollback(context.Background())
		return uuid.Nil, err
	}

	return todoListid, tx.Commit(context.Background())
}

func (r *TodoListRepository) GetAllLists(userId uuid.UUID) (*[]databaseTypes.TodoList, error) {
	var lists []databaseTypes.TodoList

	getAllListsQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", databaseTypes.TableTodoLists, databaseTypes.TableUsersLists)
	rowsGetAllLists, err := r.db.Query(context.Background(), getAllListsQuery, userId)
	if err != nil {
		return &lists, err
	}

	lists, err = pgx.CollectRows(rowsGetAllLists, pgx.RowToStructByName[databaseTypes.TodoList])

	return &lists, err
}

func (r *TodoListRepository) GetListById(userId, listId uuid.UUID) (*databaseTypes.TodoList, error) {
	var list databaseTypes.TodoList

	getListByIdQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", databaseTypes.TableTodoLists, databaseTypes.TableUsersLists)
	err := r.db.QueryRow(context.Background(), getListByIdQuery, userId, listId).Scan(&list.Id, &list.Title, &list.Description)

	return &list, err
}

func (r *TodoListRepository) DeleteList(userId, listId uuid.UUID) error {
	queryDeleteList := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2", databaseTypes.TableTodoLists, databaseTypes.TableUsersLists)
	_, err := r.db.Exec(context.Background(), queryDeleteList, userId, listId)

	return err
}

func (r *TodoListRepository) UpdateList(setArgs *[]interface{}, argId int, setValuesQuery *string) error {
	queryUpdateList := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d", databaseTypes.TableTodoLists, *setValuesQuery, databaseTypes.TableUsersLists, argId, argId+1)

	_, err := r.db.Exec(context.Background(), queryUpdateList, *setArgs...)

	return err
}
