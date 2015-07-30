///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./Functions.ts"/>

class AddUserDialog implements Dialog {
    ractive : Ractive;
    user : User;
    userOrg : User;
    isNew : boolean;
    callback : (result : User) => void;

    constructor(user : User, callback : (result : User) => void) {
        if (user == null) {
            this.isNew = true;
            this.user = new User();
            this.userOrg = null;
        } else {
            this.isNew = false;
            this.user = Utils.clone(user);
            this.userOrg = user;
        }
        this.callback = callback;
    }
    
    attach(app : App, el : HTMLElement) {
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el : el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template : '#addUserTemplate',
            data : {
                isNew : this.isNew,
                user : this.user,
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
            'save' : () => {
                this.save(app);
                return false;
            }
        });        
    }

    private save(app : App) {
        var user = this.ractive.get('user');
        var password = this.ractive.get('password');
        if (this.isNew) {
            app.client.createUser(user.login_name, user.display_name, user.tel, password, {
                success : (item : User) => {
                    this.callback(item);
                    app.addUser(item);
                    app.addSnack('作成しました');
                    app.closeDialog();
                },
                error : (status : number, msg : string) => {
                    switch (status) {
                    case 1000: app.addSnack('ユーザーIDを入力してください'); break;
                    case 1001: app.addSnack('担当者名を入力してください'); break;            
                    case 1002: app.addSnack('TELを入力してください'); break;
                    case 1003: 
                    case 1004: app.addSnack('パスワードを6文字以上入力してください'); break;
                    default : app.addSnack('保存に失敗しました。');
                    }
                }
            });
        } else {
            app.client.saveUser(user, password, {
                success : (item : User) => {
                    this.userOrg.login_name = user.login_name;
                    this.userOrg.display_name = user.display_name;
                    this.userOrg.tel = user.tel;
                    this.callback(item);
                    app.addSnack('保存しました');
                    app.closeDialog();
                },
                error : (status : number, msg : string) => {
                    switch (status) {
                    case 1001: app.addSnack('ユーザーIDを入力してください'); break;
                    case 1002: app.addSnack('担当者名を入力してください'); break;            
                    case 1003: app.addSnack('TELを入力してください'); break;
                    case 1004: app.addSnack('パスワードを6文字以上入力してください'); break;
                    default : app.addSnack('保存に失敗しました。');
                    }
                }
            });            
        }
    }
}