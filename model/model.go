package model

// Employee struct
type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    int     `json:"age"`
}

// Employees struct
type Employees struct {
	Employees []Employee `json:"employees"`
}
