package routes

import (
	"fmt"

	"github.com/felipeospina21/todo_fiber_api/database"
	"github.com/felipeospina21/todo_fiber_api/model"
	"github.com/felipeospina21/todo_fiber_api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAll(c *fiber.Ctx) error {
	var items []model.TodoItem
	res := database.DB.Instance.Find(&items)
	if res.Error != nil {
		return utils.ThrowNotFoundError(res.Error)
	}

	return c.Status(200).JSON(items)
}

func Create(c *fiber.Ctx) error {
	todo := new(model.TodoItem)
	itemId, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	todo.ItemId = itemId

	if err := c.BodyParser(todo); err != nil {
		return utils.ThrowBodyParserError(err)
	}

	res := database.DB.Instance.Create(&todo)

	if res.Error != nil {
		return utils.ThrowNotFoundError(res.Error)
	}

	return c.Status(200).JSON(todo)
}

func Update(c *fiber.Ctx) error {
	todo := new(model.TodoItem)

	if err := c.BodyParser(todo); err != nil {
		return utils.ThrowBodyParserError(err)
	}

	res := database.DB.Instance.Model(&model.TodoItem{}).Where("item_id = ?", c.Params("id")).Update("done", todo.Done)

	if res.Error != nil {
		return utils.ThrowNotFoundError(res.Error)
	}

	return c.Status(200).JSON(todo)
}

func Delete(c *fiber.Ctx) error {
	todo := new(model.TodoItem)
	itemId := c.Params("id")

	res := database.DB.Instance.Where("item_id = ?", itemId).Delete(&todo)

	if res.Error != nil {
		return utils.ThrowNotFoundError(res.Error)
	}

	r := make(map[string]string)
	r["message"] = fmt.Sprintf("item %v deleted", itemId)
	return c.Status(200).JSON(r)
}
