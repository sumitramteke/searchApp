package helpers

import (
	
	"log"
	"gopkg.in/mgo.v2/bson"
	"searchApp/app/database"
	"searchApp/app/models"
)



func ListAllCompany(c []models.Company) []models.Company{

	if err := database.Company.Find(bson.M{}).All(&c); err != nil {
			// Internal Server Error
			log.Fatal(err)
	}
	
	return c

}

func GenericSearch(c []models.Company,search string) []models.Company{

	if err := database.Company.Find(bson.M{"$text" : bson.M{"$search": search}}).All(&c); err != nil {
			// Internal Server Error
			log.Fatal(err)
	}
	
	return c
}

func FilterSearch(c []models.Company,search string,types string) []models.Company{

	
	if err := database.Company.Find(bson.M{types : &bson.RegEx{Pattern : search,Options : "i"}}).All(&c); err != nil {
			// Internal Server Error
			log.Fatal(err)
	}
	
	return c

}


func CompanyJsonSpec(c []models.Company) []interface{}{

	company := make([]interface{}, 0, len(c))

	for i := 0; i < len(c); i += 1 {

		company = append(company, parseToCompany(&c[i]))
	}

	return company
}


func parseToCompany(company *models.Company) *models.CompanyVO {

	CompanyVO := models.CompanyVO{}

	CompanyVO.ID = company.ID.Hex()
	CompanyVO.CName = company.CName
	CompanyVO.CURL = company.CURL
	CompanyVO.Address = company.Address
//	CompanyVO.Createdby = company.Createdby.Hex()
	CompanyVO.Contact = company.Contact
	CompanyVO.Email = company.Email
	CompanyVO.Language = company.Language
	CompanyVO.EmployeeCount = company.EmployeeCount
	CompanyVO.Specialities = company.Specialities
	CompanyVO.Requirements = company.Requirements
	CompanyVO.VideoURL = company.VideoURL
	CompanyVO.FileURL = company.FileURL


	return &CompanyVO
}