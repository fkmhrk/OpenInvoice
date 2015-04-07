/// <reference path="./Page.ts"/>
class TopPage implements Page {
    app : Application
    onCreate(app : Application) {
        this.app = app;
        var r = app.ractive = new Ractive({
            el : '#container',
            template : '#topTemplate',
            data : {
                loginInProgress : false,
            }
        });
        r.on('login', (e : any) => {
            this.login(r.get('username'),
                       r.get('password'));
        });            
        console.log(this);
    }
    login(username : string, password : string){
        this.app.ractive.set('loginInProgress', true);
        this.app.client.login(username, password, {
            success : (token : string) => {
                this.app.ractive.set('loginInProgress', false);
                this.app.token = token;
                this.app.router.navigate('tradings', {trigger:true})
            },
            error : (msg : string) => {
                this.app.ractive.set('loginInProgress', false);
                console.log('error ' + msg);
            }
        });
    }
}