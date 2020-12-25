package main

// User define the user fields
type User struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Address   string
	Birthday  int64
}

// getUsers1 simulate to get a list of users information
func getUsers1() ([]User, error) {
	users := []User{
		{
			FirstName: "FirstName 200",
			LastName:  "LastName 200",
			Email:     "Email 200",
			Phone:     "Phone 200",
			Address:   "Address 200",
			Birthday:  1609459200,
		},
		{
			FirstName: "FirstName 201",
			LastName:  "LastName 201",
			Email:     "Email 201",
			Phone:     "Phone 201",
			Address:   "Address 201",
			Birthday:  1609459201,
		},
		{
			FirstName: "FirstName 202",
			LastName:  "LastName 202",
			Email:     "Email 202",
			Phone:     "Phone 202",
			Address:   "Address 202",
			Birthday:  1609459202,
		},
		{
			FirstName: "FirstName 203",
			LastName:  "LastName 203",
			Email:     "Email 203",
			Phone:     "Phone 203",
			Address:   "Address 203",
			Birthday:  1609459203,
		},
		{
			FirstName: "FirstName 204",
			LastName:  "LastName 204",
			Email:     "Email 204",
			Phone:     "Phone 204",
			Address:   "Address 204",
			Birthday:  1609459204,
		},
		{
			FirstName: "FirstName 205",
			LastName:  "LastName 205",
			Email:     "Email 205",
			Phone:     "Phone 205",
			Address:   "Address 205",
			Birthday:  1609459205,
		},
		{
			FirstName: "FirstName 206",
			LastName:  "LastName 206",
			Email:     "Email 206",
			Phone:     "Phone 206",
			Address:   "Address 206",
			Birthday:  1609459206,
		},
		{
			FirstName: "FirstName 207",
			LastName:  "LastName 207",
			Email:     "Email 207",
			Phone:     "Phone 207",
			Address:   "Address 207",
			Birthday:  1609459207,
		},
		{
			FirstName: "FirstName 208",
			LastName:  "LastName 208",
			Email:     "Email 208",
			Phone:     "Phone 208",
			Address:   "Address 208",
			Birthday:  1609459208,
		},
		{
			FirstName: "FirstName 209",
			LastName:  "LastName 209",
			Email:     "Email 209",
			Phone:     "Phone 209",
			Address:   "Address 209",
			Birthday:  1609459209,
		},
		{
			FirstName: "FirstName 210",
			LastName:  "LastName 210",
			Email:     "Email 210",
			Phone:     "Phone 210",
			Address:   "Address 210",
			Birthday:  1609459210,
		},
		{
			FirstName: "FirstName 211",
			LastName:  "LastName 211",
			Email:     "Email 211",
			Phone:     "Phone 211",
			Address:   "Address 211",
			Birthday:  1609459211,
		},
		{
			FirstName: "FirstName 212",
			LastName:  "LastName 212",
			Email:     "Email 212",
			Phone:     "Phone 212",
			Address:   "Address 212",
			Birthday:  1609459212,
		},
		{
			FirstName: "FirstName 213",
			LastName:  "LastName 213",
			Email:     "Email 213",
			Phone:     "Phone 213",
			Address:   "Address 213",
			Birthday:  1609459213,
		},
		{
			FirstName: "FirstName 214",
			LastName:  "LastName 214",
			Email:     "Email 214",
			Phone:     "Phone 214",
			Address:   "Address 214",
			Birthday:  1609459214,
		},
		{
			FirstName: "FirstName 215",
			LastName:  "LastName 215",
			Email:     "Email 215",
			Phone:     "Phone 215",
			Address:   "Address 215",
			Birthday:  1609459215,
		},
		{
			FirstName: "FirstName 216",
			LastName:  "LastName 216",
			Email:     "Email 216",
			Phone:     "Phone 216",
			Address:   "Address 216",
			Birthday:  1609459216,
		},
	}
	return users, nil
}

// getUsers2 simulate to get a list of users information
func getUsers2() ([]*User, error) {
	users := []User{
		{
			FirstName: "FirstName 300",
			LastName:  "LastName 300",
			Email:     "Email 300",
			Phone:     "Phone 300",
			Address:   "Address 300",
			Birthday:  1609459300,
		},
		{
			FirstName: "FirstName 301",
			LastName:  "LastName 301",
			Email:     "Email 301",
			Phone:     "Phone 301",
			Address:   "Address 301",
			Birthday:  1609459301,
		},
		{
			FirstName: "FirstName 302",
			LastName:  "LastName 302",
			Email:     "Email 302",
			Phone:     "Phone 302",
			Address:   "Address 302",
			Birthday:  1609459302,
		},
		{
			FirstName: "FirstName 303",
			LastName:  "LastName 303",
			Email:     "Email 303",
			Phone:     "Phone 303",
			Address:   "Address 303",
			Birthday:  1609459303,
		},
		{
			FirstName: "FirstName 304",
			LastName:  "LastName 304",
			Email:     "Email 304",
			Phone:     "Phone 304",
			Address:   "Address 304",
			Birthday:  1609459304,
		},
		{
			FirstName: "FirstName 305",
			LastName:  "LastName 305",
			Email:     "Email 305",
			Phone:     "Phone 305",
			Address:   "Address 305",
			Birthday:  1609459305,
		},
		{
			FirstName: "FirstName 306",
			LastName:  "LastName 306",
			Email:     "Email 306",
			Phone:     "Phone 306",
			Address:   "Address 306",
			Birthday:  1609459306,
		},
		{
			FirstName: "FirstName 307",
			LastName:  "LastName 307",
			Email:     "Email 307",
			Phone:     "Phone 307",
			Address:   "Address 307",
			Birthday:  1609459307,
		},
		{
			FirstName: "FirstName 308",
			LastName:  "LastName 308",
			Email:     "Email 308",
			Phone:     "Phone 308",
			Address:   "Address 308",
			Birthday:  1609459308,
		},
		{
			FirstName: "FirstName 309",
			LastName:  "LastName 309",
			Email:     "Email 309",
			Phone:     "Phone 309",
			Address:   "Address 309",
			Birthday:  1609459309,
		},
		{
			FirstName: "FirstName 310",
			LastName:  "LastName 310",
			Email:     "Email 310",
			Phone:     "Phone 310",
			Address:   "Address 310",
			Birthday:  1609459310,
		},
		{
			FirstName: "FirstName 311",
			LastName:  "LastName 311",
			Email:     "Email 311",
			Phone:     "Phone 311",
			Address:   "Address 311",
			Birthday:  1609459311,
		},
		{
			FirstName: "FirstName 312",
			LastName:  "LastName 312",
			Email:     "Email 312",
			Phone:     "Phone 312",
			Address:   "Address 312",
			Birthday:  1609459312,
		},
		{
			FirstName: "FirstName 313",
			LastName:  "LastName 313",
			Email:     "Email 313",
			Phone:     "Phone 313",
			Address:   "Address 313",
			Birthday:  1609459313,
		},
		{
			FirstName: "FirstName 314",
			LastName:  "LastName 314",
			Email:     "Email 314",
			Phone:     "Phone 314",
			Address:   "Address 314",
			Birthday:  1609459314,
		},
		{
			FirstName: "FirstName 315",
			LastName:  "LastName 315",
			Email:     "Email 315",
			Phone:     "Phone 315",
			Address:   "Address 315",
			Birthday:  1609459315,
		},
		{
			FirstName: "FirstName 316",
			LastName:  "LastName 316",
			Email:     "Email 316",
			Phone:     "Phone 316",
			Address:   "Address 316",
			Birthday:  1609459316,
		},
	}
	var usersPtr []*User
	for i := range users {
		usersPtr = append(usersPtr, &users[i])
	}
	return usersPtr, nil
}
