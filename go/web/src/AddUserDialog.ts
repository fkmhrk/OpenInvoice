///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./Functions.ts"/>

import { Ractive } from "./ractive";
import { handleError } from "./pages/ErrorHandler";

export class AddUserDialog implements IDialog {
    dialogId: number = 0;
    private app: IApplication;
    private ractive!: Ractive;
    private user: IUser;
    callback: (result: IUser) => void;

    constructor(
        app: IApplication,
        user: IUser,
        callback: (result: User) => void
    ) {
        this.app = app;
        this.user = user;
        this.callback = callback;
    }

    async onCreate(elem: HTMLElement) {
        this.ractive = new Ractive({
            el: elem,
            template: "#addUserTemplate",
            data: {
                isNew: this.user.id.length == 0,
                user: this.user,
            },
            on: {
                windowClicked: () => false,
                close: () => {
                    this.app.closeDialog(this);
                    return false;
                },
                save: () => {
                    this.save();
                    return false;
                },
            },
        });
    }

    private async save() {
        const user = this.ractive.get("user");
        const password = this.ractive.get("password");

        try {
            const saved = await this.app.models.user.save(user, password);
            this.app.addSnack("作成しました");
            this.app.closeDialog(this);
            this.callback(saved);
        } catch (e) {
            handleError(this.app, e, "保存に失敗しました。");
        }
    }
}
