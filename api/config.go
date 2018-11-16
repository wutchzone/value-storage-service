package api

// Config definition
type Config struct {
	// Port to the database
	Port int `json:"port"`
	//DbUrl to the database
	DbURL string `json:"DBURL"`
	// TableName of root DB set
	TableName string
}
