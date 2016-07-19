package models

import "gopkg.in/mgo.v2/bson"

type (

	
	Address struct{
		Id	string	`jsonapi:"primary,address"`
		Country string  `jsonapi:"attr,country"`  
		State   string  `jsonapi:"attr,state"`
		City    string  `jsonapi:"attr,city"`
		Street  string  `jsonapi:"attr,street"`
		Zipcode string  `jsonapi:"attr,zipcode"`
	}

	
	Company struct {
		ID        		bson.ObjectId 	`json:"id" bson:"_id"`
		CName      		string        	`json:"companyName" bson:"company_name"`
		CURL        	        string       	`json:"url" bson:"url"`
		Address   		[]*Address         `json:"address" bson:"address"`
		CreatedBy 		bson.ObjectId  	`json:"createdBy" bson:"created_by"`
		Contact  		string  	`json:"contact" bson:"contact"`
		Email 			string 		`json:"email" bson:"email"`
		Language  		[]string        `json:"language" bson:"language"`
		EmployeeCount 	        string        	`json:"employeeCount" bson:"employee_count"`
		Specialities 	        []string     	`json:"specialties" bson:"specialties"`
		Requirements 	        []string     	`json:"requirements" bson:"requirements"`
		VideoURL  		string          `json:"videoURL" bson:"video_url"`
		FileURL   		string          `json:"filesURL" bson:"files_url"`
	}

	CompanyVO struct {
		ID        		string   	`jsonapi:"primary,company" `
		CName      		string        	`jsonapi:"attr,companyName" `
		CURL        	        string       	`jsonapi:"attr,url" `
		Address   		[]*Address      `jsonapi:"relation,address" `
		CreatedBy 		string  	`jsonapi:"relation,createdBy" `
		Contact  		string  	`jsonapi:"attr,contact" `
		Email 			string 		`jsonapi:"attr,email" `
		Language  		[]string        `jsonapi:"attr,language" `
		EmployeeCount 	        string        	`jsonapi:"attr,employeeCount" `
		Specialities 	        []string     	`jsonapi:"attr,specialties" `
		Requirements 	        []string     	`jsonapi:"attr,requirements" `
		VideoURL  		string          `jsonapi:"attr,videoURL" `
		FileURL   		string          `jsonapi:"attr,filesURL" `
	}

)