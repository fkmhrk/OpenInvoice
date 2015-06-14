
interface Client {
    /**
     * Sets refresh token
     */
    setRefreshToken(refreshToken : string);

    /**
     * Gets access token
     */
    getAccessToken() : string;
    
    /**
     * Logs in with username and password.
     */
    login(username : string, password : string, callback : ItemCallback<string>);

   /**
     * Gets all users
     */
    getUsers(callback : ItemListCallback<User>);

    /**
     * Gets all companies
     */
    getCompanies(callback : ItemListCallback<Company>);

    /**
     * Saves company
     * @return item is Company ID
     */
    saveCompany(item : Company, callback : ItemCallback<string>);    
    
    /**
     * Gets Tradings
     */
    getTradings(callback : ItemListCallback<Trading>);

    /**
     * Gets trading items of specified Trading
     */
    getTradingItems(tradingId : string, callback : ItemListCallback<TradingItem>);

    /**
     * Saves Trading
     * @return item is trading ID
     */
    saveTrading(item : Trading, callback : ItemCallback<string>);
    
    /**
     * Saves Trading item of specified Trading
     * @return item is item ID
     */
    saveTradingItem(tradingId : string, item : TradingItem, callback : ItemCallback<string>);

    /**
     * Deltes Trading item of specified Trading
     */
    deleteTradingItem(tradingId : string, itemId : string, callback : ItemCallback<string>);

    /**
     * Gets Environment
     */
    getEnvironment(callback : ItemCallback<Environment>);
    
    /**
     * Saves Environment
     */
    saveEnvironment(env : Environment, callback : Callback);

    /**
     * Gets my company name
     */
    getMyCompanyName(callback : ItemCallback<string>);

    /**
     * Gets next number
     */
    getNextNumber(type : string, date : number, callback : ItemCallback<number>);
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