/// <reference path="./Client.ts"/>

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
    getUsers(token : string, callback : ItemListCallback<User>) { }

    /**
     * Gets all companies
     */
    getCompanies(token : string, callback : ItemListCallback<Company>) {
        callback.success(companyList);
    }

    /**
     * Saves company
     * @return item is Company ID
     */
    saveCompany(token : string, item : Company, callback : ItemCallback<string>) { }    
    
    /**
     * Gets Tradings
     */
    getTradings(token : string, callback : ItemListCallback<Trading>) {
        callback.success(sheetList);
    }

    /**
     * Gets trading items of specified Trading
     */
    getTradingItems(token : string, tradingId : string, callback : ItemListCallback<TradingItem>) {
    }

    /**
     * Saves Trading
     * @return item is trading ID
     */
    saveTrading(token : string, item : Trading, callback : ItemCallback<string>) { }
    
    /**
     * Saves Trading item of specified Trading
     * @return item is item ID
     */
    saveTradingItem(token : string, tradingId : string,
                    item : TradingItem, callback : ItemCallback<string>) { }

    /**
     * Deltes Trading item of specified Trading
     */
    deleteTradingItem(token : string, tradingId : string,
                      itemId : string, callback : ItemCallback<string>) { }
}

function createClient() {
    return new MockClient();
}

