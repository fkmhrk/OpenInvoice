/// <reference path="./ractive.d.ts"/>
/// <reference path="./TopPage.ts"/>
/// <reference path="./TradingListPage.ts"/>
/// <reference path="./EditTradingPage.ts"/>

var $;
var _;
var Backbone;

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
        app.page = new TopPage();
        app.page.onCreate(app);
    },
    tradings : () => {
        if (app.token == null) {
            app.router.navigate('', {trigger:true})           
            return;
        }
        app.page = new TradingListPage();
        app.page.onCreate(app);
    },
    editTrading : (id : any) => {
        if (app.token == null) {
            app.router.navigate('', {trigger:true})           
            return;
        }
        app.page = new EditTradingPage(id);
        app.page.onCreate(app);        
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

var app : Application = new Application();

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