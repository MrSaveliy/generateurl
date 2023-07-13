package repository

type Repository interface {
	Get(short string) (string, error)
	Set(url, short string) error
	GetLastID() (int64, error)
}
