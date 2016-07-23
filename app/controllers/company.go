package controllers

import (
	"strings"
	"net/http"
	//"encoding/json"
	//"bytes"
	"fmt"
	"log"

	"github.com/revel/revel"
	"github.com/shwoodard/jsonapi"
	"gopkg.in/mgo.v2/bson"

	"searchApp/app/database"
	"searchApp/app/models"
)

type Company struct {

	*revel.Controller
}

/*type CompanyVO struct {

	category, companyname  string
}

/ar m = map[string]CompanyVO{

	"IT": {"Bell Labs","Infosys"},
	"SEARCH ENGINE" : {"Google","AOL"},
}*/

type jsonApiRespCompany []interface{}

/*--------------------Custome Response-------------------------------------*/
func (r jsonApiRespCompany) Apply(req *revel.Request, resp *revel.Response) {

	//jsonapiRuntime := jsonapi.NewRuntime().Instrument("blogs.list")

	resp.WriteHeader(http.StatusOK, "application/json")

	if err := jsonapi.MarshalManyPayload(resp.Out, r); err != nil {
		http.Error(resp.Out, err.Error(), 500)

	}

	//resp.Out.Write(r)

}


func (c Company) List() revel.Result {

	var search string
	
	companies := []models.Company{}
	c.Params.Query = c.Request.URL.Query()
	c.Params.Bind(&search, "q")
       
        search = strings.TrimSpace(search)

	if search == "" {
		fmt.Println("No search parameters passed")
		
	} else {
		search = strings.ToUpper(search)
		
		if err := database.Company.Find(bson.M{"specialties" : &bson.RegEx{Pattern : search,Options : "i"}}).All(&companies); err != nil {
			// Internal Server Error
			log.Fatal(err)
		}
	}

	company := make([]interface{}, 0, len(companies))

	for i := 0; i < len(companies); i += 1 {

		company = append(company, parseToCompany(&companies[i]))
	}

	return jsonApiRespCompany(company)
	//return c.RenderJson(m[search])
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

