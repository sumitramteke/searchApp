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
	

	var search,types string

	companies := []models.Company{}
	c.Params.Query = c.Request.URL.Query()
	c.Params.Bind(&search, "q")
	c.Params.Bind(&types, "category")
	
        search = strings.TrimSpace(search)
        types = strings.TrimSpace(types)

	

	if search == ""{
		
		companies = helpers.ListAllCompany(companies)

	}else if types == ""  {  
		
		fmt.Println("No type parameters passed")

		search = strings.ToUpper(search)

		companies = helpers.GenericSearch(companies,search)	
	}else {	
		
		search = strings.ToUpper(search)
	
		companies = helpers.FilterSearch(companies,search,types)

	}

	company := helpers.CompanyJsonSpec(companies)

	return jsonApiRespCompany(company)
}

/*func (c Company)GetQueryParameters(p ParamValues) ParamValues{
	
	c.Params.Query = c.Request.URL.Query()
	c.Params.Bind(&p.search, "q")
	c.Params.Bind(&p.category, "category")	
	
	p.search = strings.TrimSpace(p.search)
        p.category = strings.TrimSpace(p.category)
	
	return p
}*/




