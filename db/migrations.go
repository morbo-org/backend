package db

type migration struct {
	version int
	sql     string
}

var migrations = []migration{
	{
		version: 1,
		sql: `
			CREATE TABLE schema_version (
				version INT PRIMARY KEY
			);

			INSERT INTO schema_version (version) VALUES (1);
	    `,
	},
	{
		version: 2,
		sql: `
			CREATE TABLE users(
				id INTEGER NOT NULL PRIMARY KEY,
				username TEXT NOT NULL UNIQUE,
				password TEXT NOT NULL
			);

			CREATE TABLE sessions(
				session_token TEXT NOT NULL PRIMARY KEY,
				user_id INTEGER NOT NULL
			);

			UPDATE schema_version SET version = 2;
		`,
	},
	{
		version: 3,
		sql: `
			INSERT INTO users
				(id, username, password)
			VALUES
				(1, 'admin', '$2a$10$X0W3DOiy9dUP0F9xOX5o.uxckTsdpnzMJLiMYqE2kHnRIDYfWDfqC');
		`,
	},
}
