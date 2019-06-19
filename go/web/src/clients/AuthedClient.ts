import AccessToken from "../models/token/AccessToken";

export default class AuthedClient implements IAuthedClient {
    private client: HTTPClient;
    private token: AccessToken;

    constructor(client: HTTPClient, token: AccessToken) {
        this.client = client;
        this.token = token;
    }

    send(
        method: Method,
        url: string,
        header: any,
        body: any
    ): Promise<HTTPResponse> {
        if (header == null) {
            header = {};
        }
        if (this.token.token.length > 0) {
            header["Authorization"] = `bearer ${this.token.token}`;
        }

        return this.client
            .send(method, url, header, body)
            .then((r: HTTPResponse) => {
                if (r.status != 401) {
                    return r;
                }
                return this.refresh(method, url, header, body, r);
            });
    }

    private refresh(
        method: Method,
        url: string,
        header: any,
        body: any,
        originalResp: HTTPResponse
    ): Promise<HTTPResponse> {
        if (this.token.refresh.length == 0) {
            return Promise.resolve(originalResp);
        }
        const refreshURL = "/api/v1/token/refresh";
        const refreshParams = {
            token: this.token.refresh,
        };
        return this.client
            .send(Method.POST, refreshURL, {}, JSON.stringify(refreshParams))
            .then((r: HTTPResponse) => {
                if (r.status != 200) {
                    return originalResp;
                }
                const nextToken = r.body["access_token"];
                const isAdmin = r.body["is_admin"];
                this.token.save(nextToken, this.token.refresh, isAdmin);
                return this.send(method, url, header, body);
            });
    }
}
