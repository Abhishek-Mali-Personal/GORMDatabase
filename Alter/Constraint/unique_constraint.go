package Constraint

import (
	"gorm.io/gorm"
	"log"
)

type AlterUniqueConstraint struct {
	TableName      string
	ConstraintName string
	Columns        string
}

func AddUniqueConstraint(DB *gorm.DB, constraints ...AlterUniqueConstraint) {
	for _, constraint := range constraints {
		alterUniqueConstraintError := DB.Exec("ALTER TABLE " + constraint.TableName + " ADD CONSTRAINT " + constraint.ConstraintName + " UNIQUE (" + constraint.Columns + ")").Error
		if alterUniqueConstraintError != nil {
			log.Fatal(alterUniqueConstraintError)
		}
	}
}
