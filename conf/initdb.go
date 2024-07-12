package conf

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// 创建postgresql连接池
func (m *PostgreSQL) GetConnPool() (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		m.Username, m.Password, m.Host, m.Port, m.DB)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("conn postgresql<%s> error, reason: %s", dsn, err.Error())
	}
	db.SetMaxOpenConns(int(m.MaxOpenConn))
	db.SetMaxIdleConns(int(m.MaxIdleConn))
	if m.MaxLifeTime != 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	}
	if m.MaxIdleTime != 0 {
		db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping postgresql<%s> error, reason: %s", dsn, err.Error())
	}
	return db, nil
}

// 从postgresql连接池获取连接
func (m *PostgreSQL) GetDbConn() (*sql.DB, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.db == nil {
		db, err := m.GetConnPool()
		if err != nil {
			return nil, err
		}
		m.db = db
	}
	return m.db, nil
}
