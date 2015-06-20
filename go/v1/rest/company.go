package rest

import (
	s "../service"
	rj "github.com/fkmhrk-go/rawjson"
	"github.com/gorilla/mux"
	"net/http"
)

func getCompanies(company s.CompanyService) handler {
	return makeHandler(func(token, tType string,
		req *http.Request) s.Result {
		return company.GetList(token)
	})
}

func createCompany(company s.CompanyService) handler {
	return makeJsonHandler(func(token, tType string,
		json rj.RawJsonObject) s.Result {
		// read input
		name, _ := json.String("name")
		zip, _ := json.String("zip")
		address, _ := json.String("address")
		phone, _ := json.String("phone")
		unit, _ := json.String("unit")

		return company.Create(token, name, zip, address, phone, unit)
	})
}

func updateCompany(company s.CompanyService) handler {
	return makeHandler(func(token, tType string,
		req *http.Request) s.Result {
		// read path param
		vars := mux.Vars(req)
		id := vars["companyId"]

		// to json
		json, _ := rj.ObjectFromString(readBody(req))

		// read input
		name, _ := json.String("name")
		zip, _ := json.String("zip")
		address, _ := json.String("address")
		phone, _ := json.String("phone")
		unit, _ := json.String("unit")

		return company.Update(token, id, name, zip,
			address, phone, unit)
	})
}

func deleteCompany(services s.Services) handler {
	return makeHandler(func(token, tType string, req *http.Request) s.Result {
		// read path param
		vars := mux.Vars(req)
		id := vars["companyId"]

		return services.Company.Delete(token, id)
	})
}

func getNextNumber(services s.Services) handler {
	return makeHandler(func(token, tType string, req *http.Request) s.Result {
		// read path param
		vars := mux.Vars(req)
		seqType := vars["seqType"]

		// to json
		json, _ := rj.ObjectFromString(readBody(req))

		// read input
		date, _ := json.Long("date")

		return services.Trading.GetNextNumber(token, seqType, date)
	})
}
