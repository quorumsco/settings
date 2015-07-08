package settings

type Sqlite3 struct {
	Path string
}

var DefaultSqlite3 = Sqlite3{
	Path: "/tmp/users.sqlite",
}

func (s Sqlite3) String() string {
	return s.Path
}

func (config TOMLConfig) Sqlite3() Sqlite3 {
	var sq = DefaultSqlite3

	path, ok := config.Database["path"].(string)
	if path != "" && ok {
		sq.Path = path
	}
	return sq
}
