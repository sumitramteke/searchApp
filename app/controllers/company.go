package controllers

import (
	"log"
	"net/http"
	//"encoding/json"
	//"bytes"
	"fmt"
	//"io"

	"github.com/revel/revel"
	"github.com/shwoodard/jsonapi"
	"gopkg.in/mgo.v2/bson"

	"searchApp/app/database"
	"searchApp/app/models"
)

type Company struct {

	*revel.Controller
}


func (c Company) List(search string) revel.Result {
    
        search = strings.TrimSpace(search)
	var company []*models.Company

	if search == "" {
		
	} else {
		search = strings.ToLower(search)
		
	}

	return c.RenderJson()
}


