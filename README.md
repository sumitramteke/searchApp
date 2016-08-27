#Search App

##Introduction
This search app is developed in revel golang framework and it uses mongodb database.The search app is developed to perform search operations on mongodb collections.Basically it uses company collection in mongodb.The Rest api's developed in search app have the capability to perform <b>Generic & Faceted Search</b> operations.It also has the capability to retreive limited number of documents per page based on the match.

####Prerequisite to run this app on you local machine
 1. You will need golang to be install on your machine.
 2. You will need mongodb installed.

####Installation
    git clone https://github.com/rishirajdev/searchApp.git

####Loading data in MongoDB
    mongoimport --db workype --collection company --drop --file \location of the company json file\company.json
    
####Creating text index in mongodb
    db.company.createIndex({requirements: "text",specialties: "text",company_name: "text",language: "text"},{language_override: "dummy"})
    
####Running the app
     revel run searchApp
     
####To perform generic search hit the below URL
      localhost:9000/company?q=cloud&page=1&per_page=5
      
  The value of "q" should be any keywords from specialties,requirements,language & company_name 
  
####To perform Faceted Search
  1. To perform search on company names
          `localhost:9000/company?q=mynn&category=company_name&page=1&per_page=5`
  2. To perform search on requirements
          `localhost:9000/company?q=consulting&category=requirements&page=1&per_page=5`
  3. To perform search on specialties
          `localhost:9000/company?q=cloud&category=specialties&page=1&per_page=5`
  4. To perform search on language
           `localhost:9000/company?q=french&category=language&page=1&per_page=5`

####To list all companies
      `localhost:9000/company?page=1&per_page=5`
