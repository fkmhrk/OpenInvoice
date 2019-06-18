import { isStatus200, getBody } from "../../clients/Functions";

export default class UserRepository implements IUserRepository {
    private client: IAuthedClient;
    private cache: IUser[];

    constructor(client: IAuthedClient) {
        this.client = client;
        this.cache = [];
    }

    getAll(): Promise<IUser[]> {
        if (this.cache.length > 0) {
            return Promise.resolve(this.cache);
        }
        const url = "/api/v1/users";
        return this.client
            .send(Method.GET, url, {}, null)
            .then(isStatus200)
            .then(getBody)
            .then((body: any) => {
                this.cache = <IUser[]>body["users"];
                return this.cache;
            });
    }
}
