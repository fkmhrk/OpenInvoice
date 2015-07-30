///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>

class UserListDialog implements Dialog {
    ractive : Ractive;
    
    callback : (result : any) => void;
    
    attach(app : App, el : HTMLElement) {
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el : el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template : '#userListTemplate',
            data : {
                userList : app.users,
            }
        });
        this.ractive.on({
            'windowClicked' : () => {
                return false;
            },
            'close' : () => {
                app.closeDialog();
                return false;
            },
            'showEdit' : (e : any, item : User) => {
                this.showEditDialog(app, item);
                return false;
            },
            'create' : () => {
                this.createUser(app);
                return false;
            }
        });
        //dialog内だけスクロールするように調整
        var listUserHeight = $('.listTemplate').height();
        $('.listTemplate .list').css('height', listUserHeight-330);
    }

    private showEditDialog(app : App, item : User) {
        app.showDialog(new AddUserDialog(item, (result : User) => {
            this.ractive.update();
        }));
    }

    private createUser(app : App) {
        var loginName = this.ractive.get('loginName');
        var displayName = this.ractive.get('displayName');
        var tel = this.ractive.get('tel');
        var password = this.ractive.get('password');

        app.client.createUser(loginName, displayName, tel, password, {
            success : (user : User) => {
                app.addUser(user);
                this.clear();
                app.addSnack('ユーザーを作成しました！');
            },
            error : (status : number, msg : string) => {
                switch (status) {
                case 1000: app.addSnack('ユーザー名を入力してください'); break;                    
                case 1001: app.addSnack('担当者名を入力してください'); break;
                case 1002: app.addSnack('電話番号を入力してください'); break;
                case 1003: // same message
                case 1004: app.addSnack('パスワードを6文字以上入力してください'); break;
                default: app.addSnack('ユーザー作成に失敗しました'); break;
                }

            },
        });
    }
    private clear() {
        this.ractive.set('loginName', '');
        this.ractive.set('displayName', '');
        this.ractive.set('tel', '');
        this.ractive.set('password', '');
    }
}