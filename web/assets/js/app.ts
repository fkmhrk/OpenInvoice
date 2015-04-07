/// <reference path="./ClientImpl.ts"/>
/// <reference path="./MockClient.ts"/>
/// <reference path="./ractive.d.ts"/>
var $;
var _;
var Backbone;

var TopApp = {
    onCreate : () => {
        console.log(this);
    },
    login : (username : string, password : string, r : any) => {
        r.set('loginInProgress', true);
        app.client.login(username, password, {
            success : (token : string) => {
                r.set('loginInProgress', false);
                app.token = token;
                app.router.navigate('tradings', {trigger:true})
            },
            error : (msg : string) => {
                r.set('loginInProgress', false);
                console.log('error ' + msg);
            }
        });
    }
}

var TradingApp = {
    loadTradings : (token : string) => {
        app.client.getTradings(token, {
            success : (list : Array<Invoice.Trading>) => {
                app.tradings = list;
                app.tradingMap = {};
                _.each(list, (item) => {
                    app.tradingMap[item.id] = item;
                });
                TradingApp.loadUsers(token);
            },
            error : (statuc : number, msg : string) => {
                console.log('error ' + msg);
            }
        });
    },
    loadUsers : (token : string) => {
        app.client.getUsers(token, {
            success : (list : Array<Invoice.User>) => {
                app.users = list;
                TradingApp.loadCompanies(token);
            },
            error : (msg : string) => {
                console.log('error getUsers ' + msg);
            }
        });
    },
    loadCompanies : (token : string) => {
        app.client.getCompanies(token, {
            success : (list : Array<Invoice.Company>) => {
                app.companies = list;
                TradingApp.show();
            },
            error : (msg : string) => {
                console.log('error getCompanies ' + msg);
            }
        });
    },    
    show : () => {
        app.router.r = new Ractive({
            el : '#container',
            template : '#tradingTemplate',
            data : {
                tradings : app.tradings,
                token : app.token
            }
        });
        app.router.r.on('itemClick', (e : any, i : any) => {
            TradingApp.edit(i);
        });
        app.router.r.on('printQuotation', (e : any, i : any) => {
            TradingApp.printQuotation(i);
        });
        app.router.r.on('printBill', (e : any, i : any) => {
            TradingApp.printBill(i);
        });
        app.router.r.on('newTrading', (e : any) => {
            TradingApp.newTrading(app.router.r.get('newId'));
        });
        app.router.r.on('company', (e : any) => {
            app.router.navigate('companies', {trigger:true})            
        });        
    },
    newTrading : (id : string) => {
        if (id == null || id.length == 0) {
            return;
        }
        app.trading = {
            id : null,
            date : id,
            company_id : '',
            title_type : 0,
            subject : '',
            work_from : new Date().getTime(),
            work_to : new Date().getTime(),
            quotation_date : new Date().getTime(),
            bill_date : new Date().getTime(),
            tax_rate : 8,
            assignee : '',
            product : '',
        };
        app.tradingMap['new'] = app.trading;
        app.router.navigate('tradings/new', {trigger:true})
    },
    edit : (i : any) => {
        console.log(app.tradings[i]);
        app.trading = app.tradings[i];
        app.router.navigate('tradings/' + app.tradings[i].id, {trigger:true})
    },
    printQuotation : (i : any) => {
        var trading = app.tradings[i];
        window.location.href = "/php/quotation.php?access_token=" + app.token + "&trading_id=" + trading.id;
    },
    printBill : (i : any) => {
        var trading = app.tradings[i];
        window.location.href = "/php/bill.php?access_token=" + app.token + "&trading_id=" + trading.id;        
    }
}

var EditTradingApp = {
    loadTrading : (token : string, id : string) => {
        if (id == 'new') {
            app.tradingItems = [];
            EditTradingApp.show(id);
            return;
        }
        app.client.getTradingItems(token, id, {
            success : (list : Array<Invoice.TradingItem>) => {
                app.tradingItems = _.map(list, (item) => {
                    item.unit_price = util.numToCurrency(item.unit_price);
                    return item;
                });
                EditTradingApp.show(id);
            },
            error : (msg : string) => {
                console.log('error ' + msg);
            }
        });
    },
    show : (id) => {
        var es = (node : any) => {
            $(node).easySelectBox({speed: 200});
            return {
                teardown : () => {
                    // nop?
                }
            }
        };
        var toDateStr = (date : any) => {
            var m = date.getMonth() + 1;
            var d = date.getDate();
            if (m < 10) { m = "0" + m; }
            if (d < 10) { d = "0" + d; }
            return date.getFullYear() + "-" + m + "-" + d;
        };
        var item = app.tradingMap[id];
        var workFrom = toDateStr(new Date(item.work_from));
        var workTo = toDateStr(new Date(item.work_to));
        var quotationDate = toDateStr(new Date(item.quotation_date));
        var billDate = toDateStr(new Date(item.bill_date));
        app.router.r = new Ractive({
            el : '#container',
            template : '#editTradingTemplate',
            decorators: {
                easyselect: es
            },
            data : {
                trading : item,
                tradingItems : app.tradingItems,
                users : app.users,
                companies : app.companies,
                workFrom : workFrom,
                workTo : workTo,
                quotationDate : quotationDate,
                billDate : billDate,
                deleteList : [],
                numToCurrency : (val : any) => {
                    return util.numToCurrency(val);
                }
            }
        });
        app.router.r.on('numFocus', (e : any, val : any) => {
            e.node.value = util.currencyToNum(val);
            app.router.r.updateModel();
        });
        app.router.r.on('sumBlur', (e : any, val : any, index : any) => {
            e.node.value = util.numToCurrency(val);
            app.router.r.updateModel();
            var item = e.context;
            item.sum = util.currencyToNum(item.unit_price) * item.amount;
            app.router.r.update();
        });
        app.router.r.on('amountBlur', (e : any) => {
            var item = e.context;
            item.sum = util.currencyToNum(item.unit_price) * item.amount;
            app.router.r.update();
        });
        app.router.r.on('deleteItem', (e :any, index : any) => {
            if (!confirm('この項目を削除しますか？')) {
                return;
            }
            var tradings : Array<any> = app.router.r.get('tradingItems');
            var trading = tradings[index];
            if (trading.id != null) {
                var list : Array<string> = app.router.r.get('deleteList');
                list.push(trading.id);
                app.router.r.set('deleteList', list);
            }
            tradings.splice(index, 1);
            app.router.r.set('tradingItems', tradings);
            app.router.r.update(); 
        });
        app.router.r.on('addItem', (e : any) => {
            var list = app.router.r.get('tradingItems');
            list.push({
                id : null,
                subject : "",
                unit_price : 0,
                amount : 0,
                degree : "人月",
                tax_type : 1,
                memo : "",
                sum : 0,
            });
            app.router.r.set('tradingItems', list);
            app.router.r.update();
        });
        app.router.r.on('save', () => {
            var r = app.router.r;
            var company = r.get('companies')[$('#company').prop('selectedIndex')];
            var assignee = r.get('users')[$('#assignee').prop('selectedIndex')];
            var trading = r.get('trading');
            trading.company_id = company.id;
            trading.title_type = $('#titleType').prop('selectedIndex');
            trading.assignee = assignee.id;
            trading.work_from = new Date(r.get('workFrom')).getTime();
            trading.work_to = new Date(r.get('workTo')).getTime();
            trading.quotation_date = new Date(r.get('quotationDate')).getTime();
            trading.bill_date = new Date(r.get('billDate')).getTime();
            trading.tax_rate = parseFloat(r.get('trading.tax_rate'));

            var items = r.get('tradingItems');
            var list = [];
            for (var i = 0 ; i < items.length ; ++i) {
                var item = items[i];
                item.unit_price = util.currencyToNum(item.unit_price);
                item.amount = parseInt(item.amount)
                item.tax_type = parseInt(item.tax_type)
                
                list.push(item);
            }
            var deleteList : Array<string> = r.get('deleteList');

            EditTradingApp.save(trading, list, deleteList);
        });
    },
    save : (trading : any, items : Array<any>, deleteList : Array<string>) => {
        EditTradingApp.deleteItems(trading, items, deleteList);
        app.client.saveTrading(app.token, trading, {
            success : (id : string) => {
                console.log('ok');
                EditTradingApp.saveItems(id, items);
            },
            error : (msg : string) => {
                console.log('failed to save ' + msg);
            }
        });
    },
    deleteItems : (trading : any, items : Array<any>, deleteList : Array<string>) => {
        if (deleteList.length == 0) {
            EditTradingApp.saveTrading(trading, items);
            return;
        }
        app.client.deleteTradingItem(app.token, trading.id, deleteList[0], {
            success : (id : string) => {
                deleteList.shift();
                EditTradingApp.deleteItems(trading, items, deleteList);
            },
            error : (msg : string) => {
                console.log('failed to delete ' + msg);
            }            
        });
    },
    saveTrading : (trading : any, items : Array<any>) => {
    },
    saveItems : (tradingId : string, items : Array<any>) => {
        if (items.length == 0) {
            window.history.back();
            return;
        }
        app.client.saveTradingItem(app.token, tradingId, items[0], {
            success : (id : string) => {
                console.log('ok');
                items.shift();
                EditTradingApp.saveItems(tradingId, items);
            },
            error : (msg : string) => {
                console.log('failed to save ' + msg);
            }            
        });

    }
}

var CompanyApp = {
    show : () => {
        app.router.r = new Ractive({
            el : '#container',
            template : '#companyTemplate',
            data : {
                companies : app.companies,
            }
        });
        app.router.r.on('itemClick', (e : any, i : any) => {
            CompanyApp.edit(i);
        });
        app.router.r.on('newCompany', (e : any) => {
            CompanyApp.newCompany();
        });
    },
    edit : (i : any) => {
        app.company = app.companies[i];
        app.router.navigate('companies/' + app.companies[i].id, {trigger:true})
    },
    newCompany : () => {
        app.company = {
            id : null,
            name : '',
            zip : '',
            address : '',
            phone : '',
            unit : '',
        };
        app.router.navigate('companies/new', {trigger:true})
    }
}

var EditCompanyApp = {
    show : (company : any) => {
        app.router.r = new Ractive({
            el : '#container',
            template : '#editCompanyTemplate',
            data : {
                company : company
            }
        });
        app.router.r.on('save', (e : any) => {
            var company = app.router.r.get('company');
            EditCompanyApp.save(company);
        });
    },
    save : (company : any) => {
        app.client.saveCompany(app.token, company, {
            success : (id : string) => {
                window.history.back();
            },
            error : (msg : string) =>{
            }
        });
    }
}

var AppRouter = Backbone.Router.extend({
    routes : {
        "" : "top",
        "tradings" : "tradings",
        "tradings(/:id)" : "editTrading",
        "companies" : "companies",
        "companies(/:id)" : "editCompanies",
    },
    initialize : function() {
        _.bindAll(this, 'top', 'tradings', 'editTrading', 'companies', 'editCompanies');
    },
    top : () => {
        this.r = new Ractive({
            el : '#container',
            template : '#topTemplate',
            data : {
                loginInProgress : false,
            }
        });
        this.r.on('login', (e : any) => {
            TopApp.login(this.r.get('username'),
                         this.r.get('password'),
                         this.r);
        });    
    },
    tradings : () => {
        if (app.token == null) {
            app.router.navigate('', {trigger:true})           
            return;
        }
        TradingApp.loadTradings(app.token);
    },
    editTrading : (id : any) => {
        if (app.token == null) {
            app.router.navigate('', {trigger:true})           
            return;
        }        
        EditTradingApp.loadTrading(app.token, id);
    },
    companies : () => {
        if (app.token == null) {
            app.router.navigate('', {trigger:true})           
            return;
        }
        CompanyApp.show();
    },
    editCompanies : (id : any) => {
        if (app.token == null) {
            app.router.navigate('', {trigger:true})           
            return;
        }
        var company : any = null;
        if (id === 'new') {
            company = {
                id : null
            };
            EditCompanyApp.show(company);        
            return;
        }
        for (var i = 0 ; i < app.companies.length ; ++i) {
            if (app.companies[i].id === id) {
                company = app.companies[i];
                break;
            }
        }
        if (company === null) {
            app.router.navigate('', {trigger:true})           
            return;
        }        
        EditCompanyApp.show(company);        
    }
});

class App {
    token : string;
    router : any;
    client : Invoice.AppClient;
    users : Array<Invoice.User>;
    trading : Invoice.Trading;
    tradings : Array<Invoice.Trading>;
    tradingMap : any;    
    tradingItems : Array<Invoice.TradingItem>;
    company : Invoice.Company;
    companies : Array<Invoice.Company>;
    
    constructor() {
        //this.client = new Invoice.AppClientImpl('http://localhost:9001');
        this.client = new Invoice.MockClient();
    }
}

var app : App = new App();

var util : any = {
    numToCurrency : (val : any) => {
        var n = parseInt(val);
        var ret = '';
        do {
            var n1 = (n % 1000);
            var c = ("00" + n1).slice(-3);
            n = Math.floor(n / 1000);
            if (n > 0) {
                ret = c + "," + ret;
            } else {
                ret = n1 + "," + ret;
            }
        } while (n > 0);
        return ret.substring(0, ret.length - 1);
    },
    currencyToNum : (val : any) => {
        if (typeof(val) === 'number') { return val; }
        return parseInt(val.replace(",", ""));
    },    
};

(($) => {
    $(() => {
        app.router = new AppRouter();
        Backbone.history.start();
    });
})($);