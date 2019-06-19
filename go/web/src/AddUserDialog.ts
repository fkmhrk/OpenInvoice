///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./Functions.ts"/>

import { Ractive } from "./ractive";

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
            // どの箱に入れるかをIDで指定
            el: elem,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
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

        const saved = await this.app.models.user.save(user, password);
        //app.addSnack("作成しました");
        this.app.closeDialog(this);
        this.callback(saved);
        /*
        if (this.isNew) {
            app.client.createUser(
                user.login_name,
                user.display_name,
                user.tel,
                password,
                {
                    success: (item: User) => {
                        app.addUser(item);
                        this.callback(item);
                        app.addSnack("作成しました");
                        app.closeDialog();
                    },
                    error: (status: number, msg: string) => {
                        switch (status) {
                            case 1000:
                                app.addSnack("ユーザーIDを入力してください");
                                break;
                            case 1001:
                                app.addSnack("担当者名を入力してください");
                                break;
                            case 1002:
                                app.addSnack("TELを入力してください");
                                break;
                            case 1003:
                            case 1004:
                                app.addSnack(
                                    "パスワードを6文字以上入力してください"
                                );
                                break;
                            default:
                                app.addSnack("保存に失敗しました。");
                        }
                    },
                }
            );
        } else {
            app.client.saveUser(user, password, {
                success: (item: User) => {
                    this.userOrg.login_name = user.login_name;
                    this.userOrg.display_name = user.display_name;
                    this.userOrg.tel = user.tel;
                    this.callback(item);
                    app.addSnack("保存しました");
                    app.closeDialog();
                },
                error: (status: number, msg: string) => {
                    switch (status) {
                        case 1001:
                            app.addSnack("ユーザーIDを入力してください");
                            break;
                        case 1002:
                            app.addSnack("担当者名を入力してください");
                            break;
                        case 1003:
                            app.addSnack("TELを入力してください");
                            break;
                        case 1004:
                            app.addSnack(
                                "パスワードを6文字以上入力してください"
                            );
                            break;
                        default:
                            app.addSnack("保存に失敗しました。");
                    }
                },
            });
        }
*/
    }
}
