/// <reference path="./Client.ts"/>
/// <reference path="./ClientValidator.ts"/>

var globalSeq = {
    'company' : 1,
    'trading' : 1,
    'tradingItem' : 1,
    'quotationNumber' : 1,
    'billNumber' : 1,
    'deliveryNumber' : 1,
};

var userList : Array<User> = [{
    id : "user1",
    login_name : "user1",
    display_name : "担当者1",
    tel : "08011112222",
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
    "trading1" : <Trading>{
        id : "trading1",
        date : "2015-6-12",
        company_id : "company1",
        company_name : "会社1",
        title_type : 1,
        work_from : 10,
        work_to : 20,
        quotation_date : 10,
        quotation_number : '',
        bill_date : 20,
        bill_number : '',
        delivery_date : 20,
        delivery_number : '',        
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
    setRefreshToken(refreshToken : string) {
    }

    getAccessToken() : string {
        return '';
    }
    
    /**
     * Logs in with username and password.
     */
    login(username : string, password : string, callback : ItemCallback<string>) {
        if (!ClientValidator.isValidLogin(username, password, callback)) { return; }
        callback.success('token1122');
    }

    /**
     * Creates user. This API requires Admin token.
     */
    createUser(loginName : string, displayName : string, tel : string,
               password : string, callback : ItemCallback<User>) {
        if (!ClientValidator.isValidCreateUser(loginName, displayName, tel, password, callback)) { return; }
        console.log('createUser');
        var user = new User();
        user.id = '';
        user.login_name = loginName;        
        user.display_name = displayName;
        user.tel = tel;
        callback.success(user);
    }

   /**
     * Gets all users
     */
    getUsers(callback : ItemListCallback<User>) {
        callback.success(userList);
    }

    /**
     * Gets all companies
     */
    getCompanies(callback : ItemListCallback<Company>) {
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
    saveCompany(item : Company, callback : ItemCallback<string>) {
        if (!ClientValidator.isValidSaveCompany(item, callback)) { return; }
        if (item.id == null) {
            item.id = "compnay" + (++globalSeq['company']);
        }
        companyList[item.id] = item;
        callback.success(item.id);
    }

    /**
     * Deletes comapny
     */
    deleteCompany(id : string, callback : Callback) {
        console.log('DeleteCompany id=' + id);
        callback.success();
    }
    
    /**
     * Gets Tradings
     */
    getTradings(callback : ItemListCallback<Trading>) {
        var list = [];
        for (var k in sheetList) {
            list.push(sheetList[k]);
        }
        callback.success(list);
    }

    /**
     * Gets trading items of specified Trading
     */
    getTradingItems(tradingId : string, callback : ItemListCallback<TradingItem>) {
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
    saveTrading(item : Trading, callback : ItemCallback<string>) {
        if (!ClientValidator.isValidSaveTrading(item, callback)) { return; }
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
    saveTradingItem(tradingId : string, item : TradingItem, callback : ItemCallback<string>) {
        if (!ClientValidator.isValidSaveTradingItem(tradingId, item, callback)) { return; }
        var items = tradingItemList[tradingId];
        if (items === undefined) { items = {}; }

        if (item.id == null) {
            item.id = "item" + (++globalSeq['tradingItem']);
        } 
        items[item.id] = item;
        callback.success(item.id);
    }

    /**
     * Deletes Trading
     */
    deleteTrading(tradingId : string, callback : Callback) {
        console.log('Delete traing id=' + tradingId);
        callback.success();
    }

    /**
     * Deltes Trading item of specified Trading
     */
    deleteTradingItem(tradingId : string, itemId : string, callback : ItemCallback<string>) { }

    /**
     * Gets Environment
     */
    getEnvironment(callback : ItemCallback<Environment>) {
        var env = new Environment();
        env.tax_rate = "8";
        callback.success(env);
    }

    /**
     * Saves Environment
     */
    saveEnvironment(env : Environment, callback : Callback) {
        if (!ClientValidator.isValidSaveEnvironment(env, callback)) { return; }
        callback.success();
    };

    /**
     * Gets my company name
     */
    getMyCompanyName(callback : ItemCallback<string>) {
        callback.success('テスト用株式会社');
    }

    /**
     * Gets next number
     */
    getNextNumber(type : string, date : number, callback : ItemCallback<number>) {
        callback.success(20150000 + (++globalSeq[type + 'Number']));
    }
}

function createClient() {
    return new MockClient();
}

