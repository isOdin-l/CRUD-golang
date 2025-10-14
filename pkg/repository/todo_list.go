package repository

import (
	"context"
	"fmt"

	"github.com/isOdin/RestApi/internal/types/databaseTypes"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoListRepository struct {
	db *pgxpool.Pool
}

func NewTodoListRepository(db *pgxpool.Pool) *TodoListRepository {
	return &TodoListRepository{db: db}
}

func (r *TodoListRepository) CreateList(userId int, list databaseTypes.TodoList) (int, error) {
	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return -1, err
	}

	var todoListid int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", databaseTypes.TableTodoLists)

	rowCreateList := tx.QueryRow(context.Background(), createListQuery, list.Title, list.Description)
	if err := rowCreateList.Scan(&todoListid); err != nil {
		tx.Rollback(context.Background())
		return -1, err
	}

	createUserListRelationQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", databaseTypes.TableUsersLists)
	_, err = tx.Exec(context.Background(), createUserListRelationQuery, userId, todoListid)
	if err != nil {
		tx.Rollback(context.Background())
		return -1, err
	}

	return todoListid, tx.Commit(context.Background())
}
