///<reference path="./Dialog.ts"/>
///<reference path="./Client.ts"/>

class App {
    router : any;
    client : Client;
    ractive : Ractive;
    dialogs : Ractive;
    page : Page;

    accessToken : string;
    tradings : Array<Trading>;
    tradingsMap : any;    
    trading : any;
    companies : Array<Company>;
    companyMap : any;
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
    start() {
        this.client = createClient();
        this.initDialog();
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
    loadData(callback : LoadCallback) {
        this.loadTradings(callback);
    }
    private loadTradings(callback : LoadCallback) {
        if (this.tradings != null) {
            this.loadCompanies(callback);
            return;
        }
        this.client.getTradings(this.accessToken, {
            success : (list : Array<Trading>) => {
                this.tradings = list;
                this.tradingsMap = {};
                _.each(this.tradings, (item : Trading) => {
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
        this.companies = companyList;
        this.companyMap = {};
        _.each(this.companies, (item : Company) => {
            this.companyMap[item.id] = item;
        });
        callback.done();
    }
}

interface LoadCallback {
    done : () => void;
    error : () => void;
}