///<reference path="./app.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./UserListDialog.ts"/>
///<reference path="./CompanyListDialog.ts"/>
///<reference path="./SettingsDialog.ts"/>

class SignInPage implements Page {
    onCreate(app : App) {
        // Racriveオブジェクトを作る
        var r = app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el : '#container',
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template : '#signInTemplate',
            // データを設定。テンプレートで使います。
            data : {
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
                app.accessToken = token;
                app.router.navigate('top', {trigger:true});
            },
            error : (status : number, msg : string) => {
                app.ractive.set('inProgress', false);
                app.ractive.update();                                
                console.log('failed to login status=' + status);
            }
        });

    }
}