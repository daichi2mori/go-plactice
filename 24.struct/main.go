package main

import "fmt"

type User struct {
	Name string
	Age  int
	//X, Y int
}

type T struct {
	User User
}

//構造体のスライス
type Users []*User

func UpdateUser(user *User) {
	user.Name = "A"
	user.Age = 1
}

//User構造体に対してSayNameメソッドを実装
func (u User) SayName() {
	fmt.Println(u.Name)
}

/*
構造体と関数の紐づけが可能
func (変数名 構造体) 関数名(変数名 型) 型 {
    // do something
    return 値
}
*/

func (u *User) SetName(name string) {
	u.Name = name
}

//構造体のコンストラクタ
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	var user1 User
	fmt.Println(user1)

	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)

	user2.Name = "user2"
	user2.Age = 20
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	//&（ポイント演算子）をつけることでポイント型のインスタンスを作成できる
	user5 := &User{}
	fmt.Println(user5)

	//関数の引数をポインタ型にすることで参照渡しができる
	UpdateUser(&user1)
	fmt.Println(user1)

	user6 := User{Name: "user6"}
	user6.SayName()

	user6.SetName("A")
	user6.SayName()

	t := T{User: User{Name: "userT", Age: 50}}
	fmt.Println(t.User)

	//コンストラクタ
	user7 := NewUser("user7", 60)
	fmt.Println(*user7)

	//スライス
	users := Users{}
	users = append(users, &user1, &user2, &user3, &user4)
	for i, u := range users {
		fmt.Println(i, *u)
	}

	//map
	m := map[int]User{
		1: {Name: "mapUser1", Age: 10},
		2: {Name: "mapUser2", Age: 20},
	}
	fmt.Println(m)
}
