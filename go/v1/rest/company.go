package rest

import (
	"net/http"

	rj "github.com/fkmhrk-go/rawjson"
	s "github.com/fkmhrk/OpenInvoice/v1/service"
	"github.com/mokelab-go/hop"
)

func getCompanies(company s.CompanyService) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		return company.GetList()
	})
}

func createCompany(company s.CompanyService) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		// read input
		json := rj.RawJsonObject(hop.BodyJSON(req.Context()))

		name, _ := json.String("name")
		zip, _ := json.String("zip")
		address, _ := json.String("address")
		phone, _ := json.String("phone")
		unit, _ := json.String("unit")

		return company.Create(name, zip, address, phone, unit)
	})
}

func updateCompany(company s.CompanyService) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		c := req.Context()
		id := hop.PathString(c, "companyId")
		json := rj.RawJsonObject(hop.BodyJSON(c))

		// read input
		name, _ := json.String("name")
		zip, _ := json.String("zip")
		address, _ := json.String("address")
		phone, _ := json.String("phone")
		unit, _ := json.String("unit")

		return company.Update(id, name, zip,
			address, phone, unit)
	})
}

func deleteCompany(services s.Services) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		id := hop.PathString(req.Context(), "companyId")

		return services.Company.Delete(id)
	})
}

func getNextNumber(services s.Services) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		c := req.Context()
		seqType := hop.PathString(c, "seqType")
		json := rj.RawJsonObject(hop.BodyJSON(c))

		// read input
		date, _ := json.Long("date")

		return services.Trading.GetNextNumber(seqType, date)
	})
}
