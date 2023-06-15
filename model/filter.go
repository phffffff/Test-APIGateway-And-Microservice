package userModel

type Filter struct {
	Status []int `bson:"status" json:"status" form:"status"`
}
