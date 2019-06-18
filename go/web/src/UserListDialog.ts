///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>

/*
class UserListDialog implements Dialog {
    private ractive!: Ractive;

    callback?: (result: any) => void;

    attach(app: App, el: HTMLElement) {
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: "#userListTemplate",
            data: {
                userList: app.users,
            },
        });
        this.ractive.on({
            windowClicked: () => {
                return false;
            },
            close: () => {
                app.closeDialog();
                return false;
            },
            showEdit: (e: any, item: User) => {
                this.showEditDialog(app, item);
                return false;
            },
            delete: (e: any, item: User) => {
                this.deleteUser(app, item);
                return false;
            },
            create: () => {
                this.createUser(app);
                return false;
            },
        });
        //dialog内だけスクロールするように調整
        // var listUserHeight = $(".listTemplate").height();
        // $(".listTemplate .list").css("height", listUserHeight - 330);
    }

    private showEditDialog(app: App, item: User) {
        app.showDialog(
            new AddUserDialog(item, (result: User) => {
                this.ractive!.update();
            })
        );
    }

    private createUser(app: App) {
        var loginName = this.ractive.get("loginName");
        var displayName = this.ractive.get("displayName");
        var tel = this.ractive.get("tel");
        var password = this.ractive.get("password");

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
    }

    private deleteUser(app: App, user: User) {
        if (!window.confirm("この担当者を削除しますか？")) {
            return;
        }
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
    }

    private clear() {
        this.ractive.set("loginName", "");
        this.ractive.set("displayName", "");
        this.ractive.set("tel", "");
        this.ractive.set("password", "");
    }
}

*/
