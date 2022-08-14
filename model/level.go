//Done
package model

type Level struct {
	LevelID    string  `csv:"level_id"`
	LevelIndex float64 `csv:"level_index"`
	LevelName  string  `csv:"level_name"`
}

func (l Level) TableFileName() string {
	return "levels.txt"
}
