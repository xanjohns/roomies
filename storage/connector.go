package storage

import "database/sql"

type Connector struct {
	scanner sql.Scanner
}

func NewConnector() *Connector {
	return &Connector{}
}

func (c *Connector) Initialize(file *ConfigFile) {

}

type Query struct {
	Db    string
	Query string
}

func (c *Connector) QueryDB(query Query) {

}
