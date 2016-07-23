package models

import "gopkg.in/mgo.v2/bson"

type (

	
	Address struct{
		Id	string	    `jsonapi:"primary,address"`
		Country string  `jsonapi:"attr,country"`  
		State   string  `jsonapi:"attr,state"`
		City    string  `jsonapi:"attr,city"`
		Street  string  `jsonapi:"attr,street"`
		Zipcode string  `jsonapi:"attr,zipcode"`
	}

	
	Company struct {
		ID        		bson.ObjectId 	`jsonapi:"primary,company" bson:"_id"`
		CName      		string        	`jsonapi:"attr,companyname" bson:"company_name"`
		CURL        	string       	`jsonapi:"attr,url" bson:"url"`
		Address   		[]*Address      `jsonapi:"relation,address" bson:"address"`
		//Createdby 		bson.ObjectId   `jsonapi:"attr,createdby" bson:"created_by"`
		Contact  		string  		`jsonapi:"attr,contact" bson:"contact"`
		Email 			string 			`jsonapi:"attr,email" bson:"email"`
		Language  		[]string        `jsonapi:"attr,language" bson:"language"`
		EmployeeCount 	string        	`jsonapi:"attr,employeecount" bson:"employee_count"`
		Specialities 	[]string     	`jsonapi:"attr,specialties" bson:"specialties"`
		Requirements 	[]string     	`jsonapi:"attr,requirements" bson:"requirements"`
		VideoURL  		string          `jsonapi:"attr,videoURL" bson:"video_url"`
		FileURL   		string          `jsonapi:"attr,filesURL" bson:"files_url"`
	}

	CompanyVO struct {
		ID        		string      	`jsonapi:"primary,company" `
		CName      		string        	`jsonapi:"attr,companyname"`
		CURL        	string       	`jsonapi:"attr,url"`
		Address   		[]*Address      `jsonapi:"relation,address"`
		//Createdby 		string  	    `jsonapi:"attr,createdby`
		Contact  		string  		`jsonapi:"attr,contact"`
		Email 			string 			`jsonapi:"attr,email"`
		Language  		[]string        `jsonapi:"attr,language"`
		EmployeeCount 	string        	`jsonapi:"attr,employeecount"`
		Specialities 	[]string     	`jsonapi:"attr,specialties"`
		Requirements 	[]string     	`jsonapi:"attr,requirements"`
		VideoURL  		string          `jsonapi:"attr,videoURL"`
		FileURL   		string          `jsonapi:"attr,filesURL"`
	}

)