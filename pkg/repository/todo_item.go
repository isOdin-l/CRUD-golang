package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/types/databaseTypes"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoItemRepository struct {
	db *pgxpool.Pool
}

func NewTodoItemRepository(db *pgxpool.Pool) *TodoItemRepository {
	return &TodoItemRepository{db: db}
}

func (r *TodoItemRepository) CreateItem(listId uuid.UUID, item databaseTypes.TodoItem) (uuid.UUID, error) {
	tx, errTx := r.db.Begin(context.Background())
	if errTx != nil {
		return uuid.Nil, errTx
	}

	// T1 -> craete item
	var itemId uuid.UUID
	queryCreateItem := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", databaseTypes.TableTodoItems)
	errCreateItem := tx.QueryRow(context.Background(), queryCreateItem, item.Title, item.Description).Scan(&itemId)
	if errCreateItem != nil {
		tx.Rollback(context.Background())
		return uuid.Nil, errCreateItem
	}

	// T2 -> create item-list relation
	queryCreateListItemRelation := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", databaseTypes.TableListsItems)
	_, errCreateRelation := tx.Exec(context.Background(), queryCreateListItemRelation, listId, itemId)
	if errCreateRelation != nil {
		tx.Rollback(context.Background())
		return uuid.Nil, errCreateRelation
	}

	return itemId, tx.Commit(context.Background())
}

func (r *TodoItemRepository) GetAllItems(userId uuid.UUID) (*[]databaseTypes.TodoItem, error) {
	var items []databaseTypes.TodoItem
	queryGetAllItems := fmt.Sprintf("SELECT i.* FROM %s i INNER JOIN %s il ON i.id = il.item_id INNER JOIN %s l ON il.list_id = l.id INNER JOIN %s ul ON l.id = ul.list_id WHERE ul.user_id=$1",
		databaseTypes.TableTodoItems, databaseTypes.TableListsItems, databaseTypes.TableTodoLists, databaseTypes.TableUsersLists)

	rowGetAllItems, err := r.db.Query(context.Background(), queryGetAllItems, userId)
	if err != nil {
		return &items, err
	}

	items, err = pgx.CollectRows(rowGetAllItems, pgx.RowToStructByName[databaseTypes.TodoItem])

	return &items, err
}
func (r *TodoItemRepository) GetItemById(userId, itemId uuid.UUID) (*databaseTypes.TodoItem, error) {
	var itemById databaseTypes.TodoItem

	queryGetItemById := fmt.Sprintf("SELECT i.id, i.title, i.description, i.done FROM %s i INNER JOIN %s il ON i.id = il.item_id INNER JOIN %s l ON il.list_id = l.id INNER JOIN %s ul ON l.id = ul.list_id WHERE ul.user_id=$1 AND i.id = $2",
		databaseTypes.TableTodoItems, databaseTypes.TableListsItems, databaseTypes.TableTodoLists, databaseTypes.TableUsersLists)

	err := r.db.QueryRow(context.Background(), queryGetItemById, userId, itemId).Scan(&itemById.Id, &itemById.Title, &itemById.Description, &itemById.Done)

	return &itemById, err
}
func (r *TodoItemRepository) DeleteItem(userId, itemId uuid.UUID) error {
	queryDeleteItemById := fmt.Sprintf(`
		DELETE FROM %s i 
		USING %s il 
		INNER JOIN %s l ON il.list_id = l.id
		INNER JOIN %s ul ON l.id = ul.list_id 
		WHERE i.id = il.item_id
	  	AND ul.user_id = $1
		AND i.id = $2`,
		databaseTypes.TableTodoItems, databaseTypes.TableListsItems, databaseTypes.TableTodoLists, databaseTypes.TableUsersLists)

	_, err := r.db.Exec(context.Background(), queryDeleteItemById, userId, itemId)

	return err
}

func (r *TodoItemRepository) UpdateItem(setArgs *[]interface{}, setValuesQuery *string, argId int) error {
	queryUpdateItem := fmt.Sprintf(`
		UPDATE %s tl SET %s FROM %s li, %s ul
		WHERE tl.id = li.item_id AND li.list_id = ul.list_id AND tl.id = $%d AND ul.user_id = $%d`,
		databaseTypes.TableTodoItems, *setValuesQuery, databaseTypes.TableListsItems, databaseTypes.TableUsersLists, argId, argId+1)

	_, err := r.db.Exec(context.Background(), queryUpdateItem, *setArgs...)

	return err
}
