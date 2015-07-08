package settings

type Sqlite3 struct {
	Path string
}

func (s Sqlite3) String() string {
	return s.Path
}

func (config TOMLConfig) Sqlite3() Sqlite3 {
	var sDefault = Default.Components["sqlite3"].(Sqlite3)

	sConfig, ok := config.Components["sqlite3"].(Sqlite3)
	if sConfig.Path == "" || !ok {
		sConfig.Path = sDefault.Path
	}
	return sConfig
}
