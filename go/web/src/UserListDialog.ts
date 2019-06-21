///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>

import { Ractive } from "./ractive";
import { AddUserDialog } from "./AddUserDialog";
import { handleError } from "./pages/ErrorHandler";

export class UserListDialog implements IDialog {
    dialogId: number = 0;
    private app: IApplication;
    private ractive!: Ractive;

    constructor(app: IApplication) {
        this.app = app;
    }

    async onCreate(elem: HTMLElement) {
        const userList = await this.app.models.user.getAll();

        this.ractive = new Ractive({
            el: elem,
            template: "#userListTemplate",
            data: {
                userList: userList,
            },
            on: {
                windowClicked: () => false,
                close: () => {
                    this.app.closeDialog(this);
                    return false;
                },
                showEdit: (e: any, item: User) => {
                    this.showEditDialog(item);
                    return false;
                },
                delete: (e: any, item: User) => {
                    this.deleteUser(item);
                    return false;
                },
                create: () => {
                    this.createUser();
                    return false;
                },
            },
        });
    }

    private showEditDialog(item: User) {
        this.app.showDialog(
            new AddUserDialog(this.app, item, (result: User) => {
                this.ractive!.update();
            })
        );
    }

    private async createUser() {
        const user = <IUser>{
            id: "",
            login_name: this.ractive.get("loginName"),
            display_name: this.ractive.get("displayName"),
            tel: this.ractive.get("tel"),
        };
        const password = this.ractive.get("password");

        try {
            const u = await this.app.models.user.save(user, password);
            this.clear();
            // list is cached array. so we don't need to add
            this.ractive.set("userList", await this.app.models.user.getAll());
            this.app.addSnack("ユーザーを作成しました！");
        } catch (e) {
            handleError(this.app, e, "ユーザー作成に失敗しました");
        }
    }

    private async deleteUser(user: IUser) {
        if (!window.confirm("この担当者を削除しますか？")) {
            return;
        }

        try {
            await this.app.models.user.deleteUser(user);
            this.ractive.set("userList", await this.app.models.user.getAll());
            this.app.addSnack("担当者を削除しました");
        } catch (e) {
            handleError(this.app, e, "削除に失敗しました");
        }
    }

    private clear() {
        this.ractive.set({
            loginName: "",
            displayName: "",
            tel: "",
            password: "",
        });
    }
}
