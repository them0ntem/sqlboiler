package main

import (
	"github.com/themontem/sqlboiler/v4/drivers"
	"github.com/themontem/sqlboiler/v4/drivers/sqlboiler-mssql/driver"
)

func main() {
	drivers.DriverMain(&driver.MSSQLDriver{})
}
