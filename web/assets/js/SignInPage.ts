///<reference path="./app.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./UserListDialog.ts"/>
///<reference path="./CompanyListDialog.ts"/>
///<reference path="./SettingsDialog.ts"/>
declare class PasswordCredential{
    constructor(data : any);
}

class SignInPage implements Page {
    onCreate(app : App) {
        if ((<any>navigator).credentials !== undefined) { 
            (<any>navigator).credentials.get({
                password: true,
            }).then((c : any) => {
                this.signIn(app, c.id, c.password);
            }).catch((e : any) => {
            });
        }        
        // Racriveオブジェクトを作る
        var r = app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el : '#container',
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template : '#signInTemplate',
            // データを設定。テンプレートで使います。
            data : {
                myCompanyName : app.myCompanyName,
                inProgress : false,
            }
        });

        app.ractive.on({
            'signIn' : (e : any, item : Trading) => {
                var username = r.get('username');
                var password = r.get('password');
                this.signIn(app, username, password);
            },
        });        
    }
    private signIn(app : App, username : string, password : string) {
        app.ractive.set('inProgress', true);
        app.ractive.update();
        app.client.login(username, password, {
            success : (token : string) => {
                localStorage.setItem('refreshToken', token);
                if ((<any>navigator).credentials === undefined) {
                    app.router.navigate('top', {trigger:true});
                    return; 
                }
                (<any>navigator).credentials.store(new PasswordCredential({
                    id : username,
                    password : password,
                })).then((c : any) => {
                    app.router.navigate('top', {trigger:true});
                });
                
            },
            error : (status : number, msg : string) => {
                app.ractive.set('inProgress', false);
                app.ractive.update();
                switch (status) {
                case 1000: app.addSnack('ユーザー名を入力してください'); break;
                case 1001: app.addSnack('パスワードを入力してください'); break;
                }
                console.log('failed to login status=' + status);
            }
        });

    }
}