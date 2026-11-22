package restapi

import (
	"math"
	"sort"
)

// Define the Rest API interface. You should not modify the code in this block.

type User struct {
	Name    string
	Owes    map[string]float64
	OwedBy  map[string]float64
	Balance float64
}

type GetUsersRequest struct {
	Users []string
}

type GetUsersResponse struct {
	Users []User
}

type AddUserRequest struct {
	User string
}

type AddUserResponse struct {
	User User
}

type AddIouRequest struct {
	Lender   string
	Borrower string
	Amount   float64
}

type AddIouResponse struct {
	Users []User
}

type RestApi interface {
	GetUsers(GetUsersRequest) GetUsersResponse
	AddUser(AddUserRequest) AddUserResponse
	AddIou(AddIouRequest) AddIouResponse
}

// Your code goes below here. Implement the RestApi interface.

type Api struct {
	Users map[string]User
}

func NewApi(database []User) RestApi {
	users := make(map[string]User)
	for _, U := range database {
		users[U.Name] = U
	}
	return &Api{users}
}

func (a *Api) GetUsers(req GetUsersRequest) GetUsersResponse {
	var users []User
	if len(req.Users) == 0 {
		for _, user := range a.Users {
			users = append(users, user)
		}
	} else {
		for _, name := range req.Users {
			users = append(users, a.Users[name])
		}
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].Name < users[j].Name

	})
	return GetUsersResponse{users}
}

func (a *Api) AddUser(req AddUserRequest) AddUserResponse {
	addedUser := User{
		Name:    req.User,
		Owes:    make(map[string]float64),
		OwedBy:  make(map[string]float64),
		Balance: 0,
	}
	a.Users[req.User] = addedUser
	return AddUserResponse{addedUser}
}

func (a *Api) AddIou(req AddIouRequest) AddIouResponse {
	b := a.Users[req.Borrower]
	l := a.Users[req.Lender]
	bname := b.Name
	lname := l.Name
	bnet := b.Owes[lname] - b.OwedBy[lname] + req.Amount
	lnet := l.Owes[bname] - l.OwedBy[bname] - req.Amount

	//amount
	b.Balance -= req.Amount
	l.Balance += req.Amount

	delete(b.Owes, l.Name)
	delete(b.OwedBy, l.Name)
	if bnet > 0 {
		b.Owes[lname] = bnet
	} else if bnet < 0 {
		b.OwedBy[lname] = math.Abs(bnet)
	}

	delete(l.Owes, b.Name)
	delete(l.OwedBy, b.Name)
	if lnet > 0 {
		l.Owes[bname] = lnet
	} else if lnet < 0 {
		l.OwedBy[bname] = math.Abs(lnet)
	}
	//copy? 뭐라고할지
	a.Users[b.Name] = b
	a.Users[l.Name] = l

	if l.Name < b.Name {
		return AddIouResponse{Users: []User{l, b}}
	}
	return AddIouResponse{[]User{b, l}}

}
