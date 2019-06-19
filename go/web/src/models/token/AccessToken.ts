export default class AccessToken implements IAccessToken {
    token: string;
    refresh: string;
    isAdmin_: boolean;

    constructor() {
        this.isAdmin_ = false;
        try {
            this.token = getFromStorage("token");
            this.refresh = getFromStorage("refresh");
        } catch (err) {
            this.token = "";
            this.refresh = "";
        }
    }

    save(token: string, refresh: string, isAdmin: boolean) {
        this.token = token;
        this.refresh = refresh;
        this.isAdmin_ = isAdmin;
        try {
            localStorage.setItem("token", token);
            localStorage.setItem("refresh", refresh);
        } catch (err) {
            // nop
        }
    }

    isLoggedIn(): boolean {
        return this.token.length > 0;
    }

    isAdmin(): boolean {
        return this.isAdmin_;
    }
}

const getFromStorage = (key: string) => {
    const v = localStorage.getItem(key) as string;
    if (v == null) {
        return "";
    } else {
        return v;
    }
};
