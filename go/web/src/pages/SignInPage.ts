import Ractive from "../ractive";
import { handleError } from "./ErrorHandler";

export class SignInPage implements IPage {
    private app: IApplication;
    private ractive!: Ractive;

    constructor(app: IApplication) {
        this.app = app;
    }

    async onCreate() {
        this.ractive = new Ractive({
            el: "#container",
            template: "#signInTemplate",
            data: {
                // myCompanyName: app.myCompanyName,
                inProgress: false,
            },
            on: {
                signIn: () => this.signIn(),
            },
        });
    }

    private async signIn() {
        const username = this.ractive.get("username");
        const password = this.ractive.get("password");

        this.ractive.set("inProgress", true);
        try {
            const token = await this.app.models.account.signIn(
                username,
                password
            );
            // TODO save credential
            this.app.navigate("/top");
        } catch (e) {
            handleError(this.app, e, "サインインに失敗しました");
        }
    }
}
