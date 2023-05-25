package commons

var (
	MongoDbDriver = "mongodb"
)
var AllowedDrivers = []string{MongoDbDriver}

func GetAllowedDbDrivers() []string {
	return AllowedDrivers
}
