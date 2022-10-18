package main

import (
	"github.com/themontem/sqlboiler/v4/drivers"
	"github.com/themontem/sqlboiler/v4/drivers/sqlboiler-psql/driver"
)

func main() {
	drivers.DriverMain(&driver.PostgresDriver{})
}
