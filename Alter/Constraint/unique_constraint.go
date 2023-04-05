package Constraint

import (
	"gorm.io/gorm"
)

type AlterUniqueConstraint struct {
	TableName      string
	ConstraintName string
	Columns        string
}

func AddUniqueConstraint(DB *gorm.DB, constraints ...AlterUniqueConstraint) error {
	for _, constraint := range constraints {
		alterUniqueConstraintError := DB.Exec("ALTER TABLE " + constraint.TableName + " ADD CONSTRAINT " + constraint.ConstraintName + " UNIQUE (" + constraint.Columns + ")").Error
		if alterUniqueConstraintError != nil {
			return alterUniqueConstraintError
		}
	}
	return nil
}
