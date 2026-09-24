package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string             `bson:"name" json:"name"`
	Username     string             `bson:"username" json:"username"`
	Email        string             `bson:"email" json:"email"`
	PhoneNo      string             `bson:"phoneNo" json:"phoneNo"`
	JoiningDate  string             `bson:"joiningDate" json:"joiningDate"`
	DOB          string             `bson:"dob" json:"dob"`
	OfficeTimeIn string             `bson:"officeTimeIn" json:"officeTimeIn"`
	OfficeTimeOut string            `bson:"officeTimeOut" json:"officeTimeOut"`
	PhotoBase64  string             `bson:"photoBase64" json:"photoBase64"`
}