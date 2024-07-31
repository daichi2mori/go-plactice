package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type A struct{}

// バックスラッシュでJSONで使う名前を変更できる
type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       A         `json:"A"`
}

// MarshalJSON()という関数名を作成するとオーバーライドする
// func (u User) MarshalJSON() ([]byte, error) {
// 	v, err := json.Marshal(&struct {
// 		Name string
// 	}{
// 		Name: "Mr " + u.Name,
// 	})
// 	return v, err
// }

// UnmarshalJSONでもオーバーライドできる
func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		log.Fatalln(err)
	}
	u.Name = u2.Name + "!"
	return err
}

func main() {
	user1 := new(User)
	user1.ID = 1
	user1.Name = "test"
	user1.Email = "example@example.com"
	user1.Created = time.Now()

	//JSONに変換
	byteSlice, err := json.Marshal(user1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(byteSlice))

	//データに変換
	user2 := new(User)
	if err := json.Unmarshal(byteSlice, user2); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(*user2)
}
