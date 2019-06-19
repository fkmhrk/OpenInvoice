import { isStatus200, getBody } from "../../clients/Functions";
import AccessToken from "../token/AccessToken";

export default class AccountRepository implements IAccountRepository {
    private client: HTTPClient;
    private token: AccessToken;

    constructor(client: HTTPClient, token: AccessToken) {
        this.client = client;
        this.token = token;
    }

    signIn(username: string, password: string): Promise<string> {
        const url = "/api/v1/token";
        const params = {
            username: username,
            password: password,
        };
        return this.client
            .send(Method.POST, url, {}, JSON.stringify(params))
            .then(isStatus200)
            .then(getBody)
            .then((json: any) => {
                const accessToken = json.access_token;
                const refreshToken = json.refresh_token;
                const isAdmin = json.is_admin;
                this.token.save(accessToken, refreshToken, isAdmin);
                return refreshToken;
            });
    }

    refresh(): Promise<void> {
        const url = "/api/v1/token/refresh";
        const params = {
            token: this.token.refresh,
        };
        return this.client
            .send(Method.POST, url, {}, JSON.stringify(params))
            .then(isStatus200)
            .then(getBody)
            .then((body: any) => {
                const nextToken = body["access_token"];
                const isAdmin = body["is_admin"];
                this.token.save(nextToken, this.token.refresh, isAdmin);
            });
    }
}
