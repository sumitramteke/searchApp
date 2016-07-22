package controllers

import (
	"strings"
	//"net/http"
	//"encoding/json"
	//"bytes"
	"fmt"
	//"io"

	"github.com/revel/revel"
	//"github.com/shwoodard/jsonapi"
	//"gopkg.in/mgo.v2/bson"

	//"searchApp/app/database"
	//"searchApp/app/models"
)

type Company struct {

	*revel.Controller
}

type CompanyVO struct {

	category, companyname  string
}

var m = map[string]CompanyVO{

	"IT": {"Bell Labs","Infosys"},
	"SEARCH ENGINE" : {"Google","AOL"},
}



func (c Company) List() revel.Result {

	var search string
	

	c.Params.Query = c.Request.URL.Query()
	c.Params.Bind(&search, "q")
       
        search = strings.TrimSpace(search)

	if search == "" {
		fmt.Println("No search parameters passed")
		
	} else {
		search = strings.ToUpper(search)
		fmt.Println(m[search])
	}

	return c.RenderJson(m[search])
}


