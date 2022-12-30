package handler

import (
	// "database/sql"

	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/prajapatiomkar/crud-api-golang-mysql/database"
	"github.com/prajapatiomkar/crud-api-golang-mysql/model"
)

// Get all Employee records from MySQL
func GetAllEmployee(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT * from employees")

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	result := model.Employees{}

	for rows.Next() {
		employee := model.Employee{}
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age); err != nil {
			return err // Exit if we get an error
		}

		// Append Employee to Employees
		result.Employees = append(result.Employees, employee)
	}
	// Return Employees in JSON format
	return c.JSON(result)
}

// Add Employee record into MySQL
func CreateEmployee(c *fiber.Ctx) error {
	u := new(model.Employee)

	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := database.DB.Query("INSERT INTO employees (ID ,NAME, SALARY, AGE) VALUES (?,?, ?, ?)", u.ID, u.Name, u.Salary, u.Age)
	if err != nil {
		return err
	}

	return c.JSON(u)
}

// Update Employee record into MySQL
func UpdateEmployee(c *fiber.Ctx) error {
	u := new(model.Employee)

	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := database.DB.Query("UPDATE employees SET name=?,salary=?,age=? WHERE id=?", u.Name, u.Salary, u.Age, u.ID)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(u)
}

// Delete Employee record from MySQL
func DeleteEmployee(c *fiber.Ctx) error {
	var err error
	p := c.Params("id")

	id, err := strconv.Atoi(p)
	if err != nil {
		log.Fatal(err.Error())
	}

	// u := new(model.Employee)

	// if err := c.BodyParser(u); err != nil {
	// 	return c.Status(400).SendString(err.Error())
	// }

	_, err = database.DB.Query("DELETE FROM employees WHERE id = ?", id)

	if err != nil {
		return err
	}
	return c.JSON("Deleted")
}

// Get Employee by Id from MySQL
func GetEmployeeById(c *fiber.Ctx) error {
	p := c.Params("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		log.Fatal(err.Error())
	}
	row, err := database.DB.Query("SELECT * from employees WHERE id=?", id)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer row.Close()

	// result := model.Employees{}

	for row.Next() {
		employee := model.Employee{}
		if err := row.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age); err != nil {
			return err
		}
		return c.JSON(employee)
	}
	return nil
}
