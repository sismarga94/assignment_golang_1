package utils

type ConnectionData struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

func ConnectDb()
