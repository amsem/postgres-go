package helper

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)
const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "gituser"
	password = "passme123"
	dbname   = "mydb"
)
type Shipment struct {
    gorm.Model
    Packages []Package
    Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"` 
}
type Package struct {
   gorm.Model
   Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

func (Shipment) TableName() string  {
    return "Shipment"
}

func (Package) TableName() string {
    return "Package"
}

/*func InitDB() (*gorm.DB, error)  {
    connectString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",user ,password, host, dbname )
    db, err := gorm.Open("postgres", connectString)
    if err != nil {
        return nil, err
    }
    db.AutoMigrate(&Shipment{}, &Package{})
    return db, nil
}
*/ 
