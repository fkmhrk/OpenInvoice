
interface Client {
    /**
     * Logs in with username and password.
     */
    login(username : string, password : string, callback : ItemCallback<string>);

   /**
     * Gets all users
     */
    getUsers(token : string, callback : ItemListCallback<User>);

    /**
     * Gets all companies
     */
    getCompanies(token : string, callback : ItemListCallback<Company>);

    /**
     * Saves company
     * @return item is Company ID
     */
    saveCompany(token : string, item : Company, callback : ItemCallback<string>);    
    
    /**
     * Gets Tradings
     */
    getTradings(token : string, callback : ItemListCallback<Trading>);

    /**
     * Gets trading items of specified Trading
     */
    getTradingItems(token : string, tradingId : string, callback : ItemListCallback<TradingItem>);

    /**
     * Saves Trading
     * @return item is trading ID
     */
    saveTrading(token : string, item : Trading, callback : ItemCallback<string>);
    
    /**
     * Saves Trading item of specified Trading
     * @return item is item ID
     */
    saveTradingItem(token : string, tradingId : string,
                    item : TradingItem, callback : ItemCallback<string>);

    /**
     * Deltes Trading item of specified Trading
     */
    deleteTradingItem(token : string, tradingId : string,
                      itemId : string, callback : ItemCallback<string>);
}

interface Callback {
    success : () => void;
    error : (status : number, msg : string) => void;
}

interface ItemCallback<T> {
    success : (item : T) => void;
    error : (status : number, msg : string) => void;
}

interface ItemListCallback<T> {
    success : (item : Array<T>) => void;
    error : (status : number, msg : string) => void;
}