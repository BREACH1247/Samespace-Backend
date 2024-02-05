package scylladb

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

func SetupScyllaDB() (*gocql.Session, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "todo"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to ScyllaDB: %v", err)
	}
	defer session.Close()

	err = session.Query(`
	CREATE KEYSPACE todo
	WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}
`).Exec()
	if err != nil {
		log.Fatalf("Failed to create keyspace: %v", err)
	}

	err = session.Query(`
        CREATE TABLE todos (
            id UUID,
            user_id UUID,
            title TEXT,
            description TEXT,
            status TEXT,
            created TEXT,
            updated TEXT,
			PRIMARY KEY (id, user_id)
        )
    `).Exec()
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	fmt.Println("Keyspace 'todo' and table 'todos' created successfully")
	return session, nil
}
