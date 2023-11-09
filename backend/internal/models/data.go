package models

type Data struct {
	ID   string `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `bson:"name" json:"name"`
	Bags []Bag  `bson:"bags" json:"bags"`
}

type Bag struct {
	Facts []Fact `bson:"facts" json:"facts"`
}

type Fact struct {
	Info              string `bson:"info" json:"info"`
	InterestingNumber int    `bson:"number" json:"number"`
}
