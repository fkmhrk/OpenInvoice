/// <reference path="./MockClient.ts"/>
/// <reference path="./ractive.d.ts"/>
/// <reference path="./Page.ts"/>

class Application {
    token : string;
    router : any;
    ractive : Ractive;
    page : Page;
    client : Invoice.AppClient;
    users : Array<Invoice.User>;
    trading : Invoice.Trading;
    tradings : Array<Invoice.Trading>;
    tradingMap : any;    
    tradingItems : Array<Invoice.TradingItem>;
    company : Invoice.Company;
    companies : Array<Invoice.Company>;
    
    constructor() {
        this.client = new Invoice.MockClient();
    }
}