package rest

import (
	"net/http"

	rj "github.com/fkmhrk-go/rawjson"
	s "github.com/fkmhrk/OpenInvoice/v1/service"
	"github.com/fkmhrk/OpenInvoice/v1/service/company"
	"github.com/mokelab-go/hop"
)

func getCompanies(company company.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		resp := company.GetList()
		resp.Write(w)
	}
}

func createCompany(company company.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// read input
		json := rj.RawJsonObject(hop.BodyJSON(req.Context()))

		name, _ := json.String("name")
		zip, _ := json.String("zip")
		address, _ := json.String("address")
		phone, _ := json.String("phone")
		unit, _ := json.String("unit")

		resp := company.Create(name, zip, address, phone, unit)
		resp.Write(w)
	}
}

func updateCompany(company company.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		c := req.Context()
		id := hop.PathString(c, "companyId")
		json := rj.RawJsonObject(hop.BodyJSON(c))

		// read input
		name, _ := json.String("name")
		zip, _ := json.String("zip")
		address, _ := json.String("address")
		phone, _ := json.String("phone")
		unit, _ := json.String("unit")

		resp := company.Update(id, name, zip,
			address, phone, unit)
		resp.Write(w)

	}
}

func deleteCompany(services s.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := hop.PathString(req.Context(), "companyId")

		resp := services.Company.Delete(id)
		resp.Write(w)
	}
}

func getNextNumber(services s.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		c := req.Context()
		seqType := hop.PathString(c, "seqType")
		json := rj.RawJsonObject(hop.BodyJSON(c))

		// read input
		date, _ := json.Long("date")

		resp := services.Trading.GetNextNumber(seqType, date)
		resp.Write(w)
	}
}
