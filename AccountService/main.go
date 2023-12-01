package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/dgrijalva/jwt-go"
)

// User struct to hold user info
type User struct {
	UserID   string `json:"userID"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Secret key for JWT
var mySigningKey = []byte("mysupersecretkey")

func main() {
	http.HandleFunc("/user", addUser)

	http.ListenAndServe(":18080", nil)
}

// Middleware to validate JWT
func validateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from the header
		tokenString := r.Header.Get("Authorization")

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
	})
}

func addUser(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `User` instance
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate user data (you can add more validation logic here)
	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	// Initialize a session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Create item in DynamoDB
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Users"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User added successfully")
}
