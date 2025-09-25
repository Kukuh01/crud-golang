package controllers

import (
	"strconv"

	"github.com/kukuh01/crud-golang/models"
	"github.com/kukuh01/crud-golang/services"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	products, err := services.GetAllProducts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		// return c.Status(500).SendString(err.Error())
	}
	return c.JSON(products)
	// return c.Render("index", fiber.Map{
	// 	"Tilte":    "Daftar Produk",
	// 	"Products": products,
	// })
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product, err := services.GetProductByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := services.CreateProduct(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product, err := services.GetProductByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := services.UpdateProduct(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product, err := services.GetProductByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}

	if err := services.DeleteProduct(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
