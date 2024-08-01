package game

type Items struct {
	Id        uint32 `db:"id"`
	Name      string `db:"name"`
	Price1    int    `db:"price1"`
	Price10   int    `db:"price10"`
	Price100  int    `db:"price100"`
	Timestamp string `db:"timestamp"`
}
