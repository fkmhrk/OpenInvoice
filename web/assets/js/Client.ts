module Invoice {
    export interface AppClient {
        login(username : string, password : string, callback : LoginCallback);
        getTradings(token : string, callback : TradingCallback);
        getTradingItems(token : string, tradingId : string, callback : TradingItemCallback)
        getUsers(token : string, callback : UserListCallback);
        getCompanies(token : string, callback : CompanyListCallback);

        saveTrading(token : string, item : Trading, callback : SaveCallback);
        saveTradingItem(token : string, tradingId : string,
                        item : TradingItem, callback : SaveCallback);

        saveCompany(token : string, item : Company, callback : SaveCallback);
    }

    export interface User {
        id : string;
        display_name : string;
    }

    export interface Company {
        id : string;
        name : string;
        zip : string;
        address : string;
        phone : string;
        unit : string;
    }

    export interface Trading {
        id : string;
        date : string;
        company_id : string;
        title_type : number;
        subject : string;
        work_from : number;
        work_to : number;
        quotation_date : number;
        bill_date : number;
        tax_rate : number;
        assignee : string;
        product : string;
    }

    export interface TradingItem {
        id : string;
        subject : string;
        unit_price : number;
        amount : number;
        degree : string;
        tax_type : number;
        memo : string;
        sum : number;
    }    

    export interface UserListCallback {
        success : (list : Array<User>) => void;
        error : (msg : string) => void;        
    }

    export interface LoginCallback {
        success : (token : string) => void;
        error : (msg : string) => void;
    }

    export interface CompanyListCallback {
        success : (list : Array<Company>) => void;
        error : (msg : string) => void;                
    }

    export interface TradingCallback {
        success : (list : Array<Trading>) => void;
        error : (msg : string) => void;        
    }

    export interface TradingItemCallback {
        success : (list : Array<TradingItem>) => void;
        error : (msg : string) => void;        
    }

   export interface SaveCallback {
        success : (id : string) => void;
        error : (msg : string) => void;        
    }        
}