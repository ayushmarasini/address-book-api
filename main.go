package main

import (
	"net/http"
	"strconv"

	"github.com/ayushmarasini/address-book-api/models"
	"github.com/docker/docker/api/server/middleware"
	echov3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/echo/v3"
	"github.com/gofiber/fiber/middleware"
)

var addressBook = []models.AddressBookEntry{}

func main() {
	e := echov3.New()
	e.Use(middleware.logger())
	e.Use(middleware.Recover())

	e.GET("/address book", getEntries)
	e.GET("/address-book/:id", getEntryById)
	e.POST("/address-book", addEntry)
	e.PUT("/address-book/:id", updateEntry)
	e.DELETE("/address-book/:id", deleteEntry)
	e.POST("/address-book/import", importFromCSV)
	e.GET("/address-book/export", exportToCSV)

	e.Logger.Fatal(e.Start(":8080"))

}

func getEntries(c echov3.Context) error {
	return c.JSON(http.StatusOK, addressBook)
}

func getEntryById(c echov3.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, entry := range addressBook {
		if entry.ID == id {
			return c.JSON(http.StatusOK, addressBook)
		}
	}
	return c.JSON(http.StatusNotFound, "Entry Not Found :")
}

func addEntry(c echov3.Context) error {
	entry := models.AddressBookEntry{}
	if err := c.Bind(&entry); err != nil {
		return err
	}
	entry.ID = len(addressBook) + 1
	addressBook = append(addressBook, entry)
	return c.JSON(http.StatusCreated, entry)

}
func updateEntry(c echov3.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	entry := models.AddressBookEntry{}
	if err := c.Bind(&entry); err != nil {
		return err
	}
	for i, e := range addressBook {
		if e.ID == id {
			entry.ID = id
			addressBook[i] = entry
			return c.JSON(http.StatusOK, entry)
		}
	}
	return c.JSON(http.StatusNotFound, "Entry Not Found")

}
