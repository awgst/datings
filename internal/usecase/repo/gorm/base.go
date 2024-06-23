package gorm

import (
	"fmt"
	"sync"

	"gorm.io/gorm"
)

type base struct {
	db  *gorm.DB
	mtx sync.Mutex
}

// Transaction
func (r *base) DBTransaction(fn func(tx *gorm.DB) error) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	tx := r.db.Session(&gorm.Session{SkipDefaultTransaction: true}).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := fn(tx)
	if err != nil {
		if rbErr := tx.Rollback().Error; rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit().Error
}
