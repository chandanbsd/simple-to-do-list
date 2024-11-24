package models

import (
	"github.com/google/uuid"
)

type UserEntity struct {
	id string `binding:"required"`
	guid string `binding:"required"`
	firstName string `binding:"required"`
	lastName  string `binding:"required"`
	email     string `binding:"required"`
	password  string `binding:"required"`
	isActive bool `binding:"required"`
}

func (u *UserEntity) Create(
	firstName string,
	lastName string,
	email string,
	password string,
) {
	u.guid = uuid.New().String()
	u.firstName = firstName
	u.lastName = lastName
	u.email = email
	u.password = password
	u.isActive = true
} 

func (u *UserEntity) Update(
	firstName string,
	lastName string,
	email string,
	password string,
) {
	u.guid = uuid.New().String()
	u.firstName = firstName
	u.lastName = lastName
	u.email = email
	u.password = password
}

func (u *UserEntity) Delete() {
	u.isActive = false
}


