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
            'create' : () => {
                this.createUser(app);
                return false;
            }
        });
        //dialog内だけスクロールするように調整
        var listUserHeight = $('.listTemplate').height();
        $('.listTemplate .list').css('height', listUserHeight-330);
    }

    private createUser(app : App) {
        var loginName = this.ractive.get('loginName');
        var displayName = this.ractive.get('displayName');
        var tel = this.ractive.get('tel');
        var password = this.ractive.get('password');

        app.client.createUser(loginName, displayName, tel, password, {
            success : (user : User) => {
                this.ractive.push('userList', user);
                this.clear();
                app.addSnack('ユーザーを作成しました！');
            },
            error : (status : number, msg : string) => {
                app.addSnack('ユーザー作成に失敗しました');
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