package geveechogo

import "github.com/labstack/echo/v4"

type GeveEcho struct {
	*echo.Echo
}

func (g *GeveEcho) GetOne(readOne func(collection, id string) (any, error)) {
	g.GET("/:collection/:id", func(c echo.Context) error {
		collection := c.Param("collection")
		id := c.Param("id")

		item, err := readOne(collection, id)
		if err != nil {
			return err
		}

		return c.JSON(200, item)
	})
}

func (g *GeveEcho) GetMany(readMany func(collection string) ([]any, error)) {
	g.GET("/:collection", func(c echo.Context) error {
		collection := c.Param("collection")

		items, err := readMany(collection)
		if err != nil {
			return err
		}

		return c.JSON(200, items)
	})
}
