package structures

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Id              primitive.ObjectID `json:"Id,omitempty" bson:"Id,omitempty"`
	UserId          primitive.ObjectID `json:"UserId,omitempty" bson:"UserId,omitempty"`
	Caption         string             `json:"Caption,omitempty" bson:"Caption,omitempty"`
	ImageUrl        string             `json:"ImageUrl,omitempty" bson:"ImageUrl,omitempty"`
	PostedTimeStamp primitive.DateTime `json:"PostedTimeStamp,omitempty" bson:"PostedTimeStamp,omitempty"`
}
