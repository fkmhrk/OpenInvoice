import {
    isStatus200,
    getBody,
    isStatus201,
    isStatus204,
} from "../../clients/Functions";

export default class CompanyRepository implements ICompanyRepository {
    private client: IAuthedClient;
    private cache: ICompany[];

    constructor(client: IAuthedClient) {
        this.client = client;
        this.cache = [];
    }

    getAll(): Promise<ICompany[]> {
        if (this.cache.length > 0) {
            return Promise.resolve(this.cache);
        }
        const url = "/api/v1/companies";
        return this.client
            .send(Method.GET, url, {}, null)
            .then(isStatus200)
            .then(getBody)
            .then((body: any) => {
                this.cache = body["companies"];
                return this.cache;
            });
    }

    save(company: ICompany): Promise<ICompany> {
        if (company.id.length == 0) {
            return this.createCompany(company);
        } else {
            return this.updateCompany(company);
        }
    }

    deleteCompany(company: ICompany): Promise<boolean> {
        const url = `/api/v1/companies/${company.id}`;
        return this.client
            .send(Method.DELETE, url, {}, null)
            .then(isStatus204)
            .then((r: HTTPResponse) => {
                this.cache = this.cache.filter(
                    (c: ICompany) => c.id != company.id
                );
                return true;
            });
    }

    private createCompany(company: ICompany): Promise<ICompany> {
        const url = "/api/v1/companies";
        return this.client
            .send(Method.POST, url, {}, JSON.stringify(company))
            .then(isStatus201)
            .then(getBody)
            .then((body: any) => {
                company.id = body["id"];
                this.cache.push(company);
                return company;
            });
    }

    private updateCompany(company: ICompany): Promise<ICompany> {
        const url = `/api/v1/companies/${company.id}`;
        return this.client
            .send(Method.PUT, url, {}, JSON.stringify(company))
            .then(isStatus200)
            .then((r: HTTPResponse) => company);
    }
}
