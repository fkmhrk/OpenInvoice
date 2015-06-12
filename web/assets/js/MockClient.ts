/// <reference path="./Client.ts"/>

var globalSeq = {
    'company' : 1,
    'trading' : 1,
    'tradingItem' : 1,
};

var userList : Array<User> = [{
    id : "user1",
    display_name : "担当者1",
}];

var companyList = {
    "company1" : {
        id : "company1",
        name : "会社1",
        zip : "111-2222",
        address : "東京都のどっか",
        phone : "090-1111-2222",
        fax : "090-1111-3333",
        unit : "開発部",
        assignee : "担当さん",
    }
};

var sheetList = {
    "trading1" : {
        id : "trading1",
        date : "2015-6-12",
        company_id : "company1",
        company_name : "会社1",
        title_type : 1,
        work_from : 10,
        work_to : 20,
        quotation_date : 10,
        bill_date : 20,
        tax_rate : 8,
        subject : "プロジェクト1",
        assignee : "user1",
        product : "製品1",
        total : 10,
        modified_time : 10,
    }
};

var tradingItemList = {
    "trading1" : {
        "item1-1" :{
            id : "item1-1",
            sort_order : 1,
            subject : "品目1-1",
            unit_price : 10000,
            amount : 2,
            degree : "人月",
            tax_type : 1,
            memo : "メモ1-1",
            sum : 20000,
        }
    }
};

class MockClient implements Client {
    /**
     * Logs in with username and password.
     */
    login(username : string, password : string, callback : ItemCallback<string>) {
        callback.success('token1122');
    }

   /**
     * Gets all users
     */
    getUsers(token : string, callback : ItemListCallback<User>) {
        callback.success(userList);
    }

    /**
     * Gets all companies
     */
    getCompanies(token : string, callback : ItemListCallback<Company>) {
        var list = [];
        for (var k in companyList) {
            list.push(companyList[k]);
        }
        callback.success(list);
    }

    /**
     * Saves company
     * @return item is Company ID
     */
    saveCompany(token : string, item : Company, callback : ItemCallback<string>) {
        if (item.id == null) {
            item.id = "compnay" + (++globalSeq['company']);
        }
        companyList[item.id] = item;
        callback.success(item.id);
    }    
    
    /**
     * Gets Tradings
     */
    getTradings(token : string, callback : ItemListCallback<Trading>) {
        var list = [];
        for (var k in sheetList) {
            list.push(sheetList[k]);
        }
        callback.success(list);
    }

    /**
     * Gets trading items of specified Trading
     */
    getTradingItems(token : string, tradingId : string, callback : ItemListCallback<TradingItem>) {
        var items = tradingItemList[tradingId];
        var list = [];
        for (var k in items) {
            list.push(items[k]);
        }
        callback.success(list);
    }

    /**
     * Saves Trading
     * @return item is trading ID
     */
    saveTrading(token : string, item : Trading, callback : ItemCallback<string>) {
        if (item.id == null) {
            item.id = "trading" + (++globalSeq['trading']);
        }
        item.modified_time = new Date().getTime();
        sheetList[item.id] = item;
        callback.success(item.id);            
    }
    
    /**
     * Saves Trading item of specified Trading
     * @return item is item ID
     */
    saveTradingItem(token : string, tradingId : string,
                    item : TradingItem, callback : ItemCallback<string>) {
        var items = tradingItemList[tradingId];
        if (items === undefined) { items = {}; }

        if (item.id == null) {
            item.id = "item" + (++globalSeq['tradingItem']);
        } 
        items[item.id] = item;
        callback.success(item.id);
    }

    /**
     * Deltes Trading item of specified Trading
     */
    deleteTradingItem(token : string, tradingId : string,
                      itemId : string, callback : ItemCallback<string>) { }

    /**
     * Gets Environment
     */
    getEnvironment(token : string, callback : ItemCallback<Environment>) {
        callback.success(new Environment());
    }

    /**
     * Saves Environment
     */
    saveEnvironment(token : string, env : Environment, callback : Callback) {
        callback.success();
    };

    /**
     * Gets my company name
     */
    getMyCompanyName(callback : ItemCallback<string>) {
        callback.success('テスト用株式会社');
    }
}

function createClient() {
    return new MockClient();
}

