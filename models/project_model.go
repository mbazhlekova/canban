package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	Description string             `json:"description,omitempty"`
	Columns     []Column           `json:"columns,omitempty"`
}

type Column struct {
	Title string `json:"title,omitempty" validate:"required"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
