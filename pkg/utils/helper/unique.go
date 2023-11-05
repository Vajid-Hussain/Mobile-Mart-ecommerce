package helper

import "github.com/google/uuid"

func GenerateUUID() string{
	newUUID:=uuid.New()
	
	uuidString:=newUUID.String()
	return uuidString
}