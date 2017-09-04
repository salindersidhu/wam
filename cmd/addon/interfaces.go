package addon

/*
Utils represents an interface containing web and file functions for downloading,
installing, updating and removing World of Warcraft addons
*/
type Utils interface {
	GetData(string) (*Data, error)
	Install(string) error
}
