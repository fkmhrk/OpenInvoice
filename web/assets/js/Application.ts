///<reference path="./Dialog.ts"/>
///<reference path="./Client.ts"/>
///<reference path="./Functions.ts"/>

class App {
    router : any;
    client : Client;
    ractive : Ractive;
    dialogs : Ractive;
    snackbars : Ractive;
    page : Page;

    accessToken : string;
    myCompanyName : string;
    environment : Environment;
    users : Array<User>;
    tradingsMap : any;    
    trading : any;
    companies : Array<Company>;
    companyMap : any;

    // getter
    getTradings() : Array<Trading> {
        return Utils.toList(this.tradingsMap);
    }
    
    showDialog(dialog : Dialog) {
        (<HTMLElement>document.querySelector('#dialogs')).style.display = 'block';
        app.dialogs.push('dialogs', dialog).then(() => {
            var list = app.dialogs.get('dialogs');
            app.updateDialogs(list);
        });
        $('#container').addClass('dialogOpened');
     
    }
    updateDialogs(list) {
        for (var i = 0 ; i < list.length ; ++i) {
            var s = document.querySelector('#dialog' + i);
            list[i].attach(this, s);
        }        
    }
    closeDialog() {
        this.dialogs.pop('dialogs').then(() => {
            // hide overlay
            var list = this.dialogs.get('dialogs');
            if (list.length == 0) {
                (<HTMLElement>document.querySelector('#dialogs')).style.display = 'none';
                $('#container').removeClass('dialogOpened');
            } else {
                this.updateDialogs(list);
            }
        });
    }
    // snack bar
    addSnack(item : string) {
        this.snackbars.push('snackbars', item);
        var closeFunc = () => {
            var list = this.snackbars.get('snackbars');
            if (list.length == 0) { return; }
            this.snackbars.splice('snackbars', 0, 1);
            if (this.snackbars.get('snackbars').length > 0) {
                setTimeout(closeFunc, 3000);
            }
        }
        setTimeout(closeFunc, 3000);
    }
    start() {
        this.client = createClient();
        var refreshToken = localStorage.getItem('refreshToken');
        this.client.setRefreshToken(refreshToken);
        this.initDialog();
        this.initSnackbar();
        this.loadMyCompanyName();
    }
    private initDialog() {
        // dialogの準備
        this.dialogs = new Ractive({
            el : '#dialogs',
            template : '#dialogsTemplate',
            data : {
                dialogs : [],
            }
        });
        this.dialogs.on({
            'closeClick' : () => {
                this.closeDialog();
            }
        });            
    }
    private initSnackbar() {
        // snackbarsの準備
        this.snackbars = new Ractive({
            el : '#snacks',
            template : '#snackbarsTemplate',
            data : {
                snackbars : [],
            }
        });
        this.snackbars.on({
            'close' : (e : any, index : number) => {
                this.snackbars.splice('snackbars', index, 1);
            }
        });                    
    }
    private loadMyCompanyName() {
        if (this.myCompanyName != null && this.myCompanyName.length > 0) {
            return;
        }
        this.client.getMyCompanyName({
            success : (name : string) => {
                this.myCompanyName = name;
            },
            error : (status : number, msg : string) => {
                console.log('Failed to get my company name status=' + status);
            }
        });
    }
    
    loadData(callback : LoadCallback) {
        this.loadEnvironment(callback);
    }
    
    private loadEnvironment(callback : LoadCallback) {
        if (this.environment != null) {
            this.loadUsers(callback);
            return;
        }
        this.client.getEnvironment({
            success : (item : Environment) => {
                this.environment = item;
                this.loadUsers(callback);
            },
            error : (status : number, msg : string) => {
                console.log('Failed to get environment status=' + status);
                callback.error();
            }
        });
    }
    
    private loadUsers(callback : LoadCallback) {
        if (this.users != null) {
            this.loadTradings(callback);
            return;
        }
        this.client.getUsers({
            success : (list : Array<User>) => {
                this.users = list;
                this.loadTradings(callback);
            },
            error : (status : number, msg : string) => {
                console.log('Failed to get users status=' + status);
                callback.error();
            }
        });
    }
    private loadTradings(callback : LoadCallback) {
        if (this.tradingsMap != null) {
            this.loadCompanies(callback);
            return;
        }
        this.client.getTradings({
            success : (list : Array<Trading>) => {
                this.tradingsMap = {};
                _.each(list, (item : Trading) => {
                    this.tradingsMap[item.id] = item;
                });
                this.loadCompanies(callback);
            },
            error : (status : number, msg : string) => {
                console.log('Failed to get tradings status=' + status);
                callback.error();
            }
        });
    }
    private loadCompanies(callback : LoadCallback) {
        if (this.companies != null) {
            callback.done();
            return;
        }
        this.client.getCompanies({
            success : (list : Array<Company>) => {
                this.companies = list;
                this.companyMap = {};
                _.each(this.companies, (item : Company) => {
                    this.companyMap[item.id] = item;
                });
                callback.done();                
            },
            error : (status : number, msg : string) => {
                console.log('Failed to get companies status=' + status);
                callback.error();
            }            
        });
    }

    addCompany(c : Company) {
        this.companies.push(c);
        this.companyMap[c.id] = c;
    }
}

interface LoadCallback {
    done : () => void;
    error : () => void;
}