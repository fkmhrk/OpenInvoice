import {
    isStatus200,
    getBody,
    isStatus201,
    isStatus204,
} from "../../clients/Functions";

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

    save(user: IUser, password: string): Promise<IUser> {
        if (user.id.length == 0) {
            return this.create(user, password);
        } else {
            return this.update(user, password);
        }
    }

    deleteUser(user: IUser): Promise<boolean> {
        const url = `/api/v1/users/${user.id}`;
        return this.client
            .send(Method.DELETE, url, {}, null)
            .then(isStatus204)
            .then((r: HTTPResponse) => {
                this.cache = this.cache.filter((u: IUser) => u.id != user.id);
                return true;
            });
    }

    private create(user: IUser, password: string): Promise<IUser> {
        const url = "/api/v1/users";
        const params = {
            login_name: user.login_name,
            display_name: user.display_name,
            tel: user.tel,
            password: password,
        };
        return this.client
            .send(Method.POST, url, {}, JSON.stringify(params))
            .then(isStatus201)
            .then(getBody)
            .then((body: any) => {
                this.cache.push(body);
                return <IUser>body;
            });
    }

    private update(user: IUser, password: string): Promise<IUser> {
        const url = `/api/v1/users/${user.id}`;
        const params = {
            id: user.id,
            login_name: user.login_name,
            display_name: user.display_name,
            tel: user.tel,
            password: password,
        };
        return this.client
            .send(Method.PUT, url, {}, JSON.stringify(params))
            .then(isStatus200)
            .then(getBody)
            .then((body: any) => {
                const index = this.cache.findIndex(
                    (u: IUser) => u.id == user.id
                );
                if (index != -1) {
                    this.cache[index] = user;
                }
                return user;
            });
    }
}
