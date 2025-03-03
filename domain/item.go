package domain

type Item struct {
	ID   int
	Name string
	Cost int
}

type Records struct {
	ID      int
	ItemID  int
	OwnerID int
}
