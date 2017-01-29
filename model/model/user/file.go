package user

import "fmt"

type FileSchema struct {
	ID string `bson:"_id"`
}

type File struct {}

func (f *File) Collection() string {
	return "UserFile"
}

func (f *File) Find() {
	fmt.Println("col:" + f.Collection())
}
