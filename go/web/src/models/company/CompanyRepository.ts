import { isStatus200, getBody } from "../../clients/Functions";

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
}
