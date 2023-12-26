package config

import (
	"os"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MariaStringAkademik string = os.Getenv("MARIASTRINGAKADEMIK")

var DBUlbimongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "tes_db",
}

var Ulbimongoconn = atdb.MongoConnect(DBUlbimongoinfo)
