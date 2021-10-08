package structures

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Id              primitive.ObjectID `json:"Id,omitempty" bson:"Id,omitempty"`
	UserId          primitive.ObjectID `json:"UserId,omitempty" bson:"Id,omitempty"`
	Caption         string             `json:"Caption,omitempty" bson:"Name,omitempty"`
	ImageUrl        string             `json:"ImageUrl,omitempty" bson:"Email,omitempty"`
	PostedTimeStamp primitive.DateTime `json:"PostedTimeStamp,omitempty" bson:"Password,omitempty"`
}
