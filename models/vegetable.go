package models

type Vegetable struct {
    Name  string
    Price int32
    Image *string
}

// Utils
func StrPtr(str string) *string {
    return &str
}