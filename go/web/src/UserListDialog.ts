///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>

import { Ractive } from "./ractive";
import { AddUserDialog } from "./AddUserDialog";

export class UserListDialog implements IDialog {
    dialogId: number = 0;
    private app: IApplication;
    private ractive!: Ractive;

    private callback: (result: IUser) => void;

    constructor(app: IApplication, callback: (result: IUser) => void) {
        this.app = app;
        this.callback = callback;
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

        const u = await this.app.models.user.save(user, password);
        this.clear();
        // list is cached array. so we don't need to add
        this.ractive.set("userList", await this.app.models.user.getAll());
        // app.addSnack("ユーザーを作成しました！");
        /*
        app.client.createUser(loginName, displayName, tel, password, {
            success: (user: User) => {
                app.addUser(user);
                this.clear();
                app.addSnack("ユーザーを作成しました！");
            },
            error: (status: number, msg: string) => {
                switch (status) {
                    case 1000:
                        app.addSnack("ユーザー名を入力してください");
                        break;
                    case 1001:
                        app.addSnack("担当者名を入力してください");
                        break;
                    case 1002:
                        app.addSnack("電話番号を入力してください");
                        break;
                    case 1003: // same message
                    case 1004:
                        app.addSnack("パスワードを6文字以上入力してください");
                        break;
                    default:
                        app.addSnack("ユーザー作成に失敗しました");
                        break;
                }
            },
        });
*/
    }

    private async deleteUser(user: IUser) {
        if (!window.confirm("この担当者を削除しますか？")) {
            return;
        }

        await this.app.models.user.deleteUser(user);
        this.ractive.set("userList", await this.app.models.user.getAll());
        //app.addSnack("担当者を削除しました");

        /*        
        app.client.deleteUser(user.id, {
            success: () => {
                app.deleteUser(user);
                this.ractive.set("userList", app.users);
                app.addSnack("担当者を削除しました");
            },
            error: (status: number, msg: string) => {
                switch (status) {
                    default:
                        app.addSnack("削除に失敗しました");
                }
            },
        });
*/
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
