import { isStatus200, getBody } from "../../clients/Functions";

export default class EnvironmentRepository implements IEnvironmentRepository {
    private client: IAuthedClient;

    constructor(client: IAuthedClient) {
        this.client = client;
    }

    get(): Promise<IEnvironment> {
        const url = "/api/v1/environments";
        return this.client
            .send(Method.GET, url, {}, null)
            .then(isStatus200)
            .then(getBody)
            .then((json: any) => <IEnvironment>json);
    }

    save(env: IEnvironment): Promise<IEnvironment> {
        const url = "/api/v1/environments";
        return this.client
            .send(Method.PUT, url, {}, JSON.stringify(env))
            .then(isStatus200)
            .then((r: HTTPResponse) => env);
    }
}
