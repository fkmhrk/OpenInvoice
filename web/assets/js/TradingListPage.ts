/// <reference path="./Page.ts"/>

class TradingListPage implements Page {
    app : Application
    onCreate(app : Application) {
        this.app = app;
        var r = app.ractive = new Ractive({
            el : '#container',
            template : '#tradingTemplate',
            data : {
                tradings : app.tradings,
                token : app.token
            }
        });
        r.on({
            itemClick : (e : any, i : any) => {
                this.edit(i);
            },
            printQuotation : (e : any, i : any) => {
                this.printQuotation(i);
            },
            printBill :  (e : any, i : any) => {
                this.printBill(i);
            },
            newTrading : (e : any) => {
                this.newTrading(r.get('newId'));
            },
            company : (e : any) => {
                this.app.router.navigate('companies', {trigger:true});
            }
        });
        this.loadTradings();
    }
    loadTradings(){
        this.app.client.getTradings(this.app.token, {
            success : (list : Array<Invoice.Trading>) => {
                this.app.tradings = list;
                this.app.tradingMap = {};
                _.each(list, (item) => {
                    this.app.tradingMap[item.id] = item;
                });
                this.loadUsers();
            },
            error : (statuc : number, msg : string) => {
                console.log('error ' + msg);
            }
        });
    }
    loadUsers() {
        this.app.client.getUsers(this.app.token, {
            success : (list : Array<Invoice.User>) => {
                this.app.users = list;
                this.loadCompanies();
            },
            error : (msg : string) => {
                console.log('error getUsers ' + msg);
            }
        });
    }
    loadCompanies() {
        this.app.client.getCompanies(this.app.token, {
            success : (list : Array<Invoice.Company>) => {
                this.app.companies = list;
                this.app.companyMap = {};
                _.each(list, (item) => {
                    this.app.companyMap[item.id] = item;
                });
                // set company name
                _.each(this.app.tradings, (item) => {
                    var company = this.app.companyMap[item.company_id];
                    item.company_name = (company === undefined) ? '' : company.name;
                });
                this.show();
            },
            error : (msg : string) => {
                console.log('error getCompanies ' + msg);
            }
        });
    }
    show() {
        this.app.ractive.set('tradings', this.app.tradings);
        this.app.ractive.set('token', this.app.token);
        this.app.ractive.update();
    }
    newTrading(id : string) {
        if (id == null || id.length == 0) {
            return;
        }
        app.trading = {
            id : null,
            date : id,
            modified_time : 0,
            company_id : '',
            company_name : '',
            title_type : 0,
            subject : '',
            work_from : new Date().getTime(),
            work_to : new Date().getTime(),
            quotation_date : new Date().getTime(),
            bill_date : new Date().getTime(),
            tax_rate : 8,
            assignee : '',
            product : '',
            total : 0,
        };
        app.tradingMap['new'] = app.trading;
        app.router.navigate('tradings/new', {trigger:true})
    }
    edit(i : any) {
        console.log(app.tradings[i]);
        app.trading = app.tradings[i];
        app.router.navigate('tradings/' + app.tradings[i].id, {trigger:true})
    }
    printQuotation(i : any) {
        var trading = app.tradings[i];
        window.location.href = "/php/quotation.php?access_token=" + app.token + "&trading_id=" + trading.id;
    }
    printBill (i : any) {
        var trading = app.tradings[i];
        window.location.href = "/php/bill.php?access_token=" + app.token + "&trading_id=" + trading.id;        
    }
}