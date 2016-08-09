package controllers

import (
	"strings"
	"net/http"
	"fmt"
	

	"github.com/revel/revel"
	"github.com/shwoodard/jsonapi"
	
	
	"searchApp/app/models"
	"searchApp/app/helpers"
)

type Company struct {

	*revel.Controller
}




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
	

	var(

		 search,types string
		 page,per_page int
	)


	companies := []models.Company{}

	c.Params.Query = c.Request.URL.Query()

	c.Params.Bind(&search, "q")
	c.Params.Bind(&types, "category")

	c.Params.Bind(&page, "page")
	c.Params.Bind(&per_page, "per_page")
	
    search = strings.TrimSpace(search)
    types = strings.TrimSpace(types)



	
	 if page != 1{

     	page = ((page - 1 ) * per_page) + 1 
     }	


	if search == ""{
		
		companies = helpers.ListAllCompany(companies,page,per_page)

	}else if types == ""  {  
		
		fmt.Println("No type parameters passed")

		search = strings.ToUpper(search)

		companies = helpers.GenericSearch(companies,search,page,per_page)	
	}else {	
		
		search = strings.ToUpper(search)
	
		companies = helpers.FilterSearch(companies,search,types,page,per_page)

	}

	company := helpers.CompanyJsonSpec(companies)

	return jsonApiRespCompany(company)
}






