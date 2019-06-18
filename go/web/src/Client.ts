interface Client {
    /**
     * Sets refresh token
     */
    setRefreshToken(refreshToken: string): void;

    /**
     * Determines whether this account is admin.
     */
    isAdmin(): boolean;

    /**
     * Gets access token
     */
    getAccessToken(): string;

    /**
     * Logs in with username and password.
     */
    login(
        username: string,
        password: string,
        callback: ItemCallback<string>
    ): void;

    /**
     * Creates user. This API requires Admin token.
     */
    createUser(
        loginName: string,
        displayName: string,
        tel: string,
        password: string,
        callback: ItemCallback<User>
    ): void;

    /**
     * Gets all users
     */
    getUsers(callback: ItemListCallback<User>): void;

    /**
     * Saves user
     */
    saveUser(user: User, password: string, callback: ItemCallback<User>): void;

    /**
     * Deletes user
     */
    deleteUser(id: string, callback: Callback): void;

    /**
     * Gets all companies
     */
    getCompanies(callback: ItemListCallback<Company>): void;

    /**
     * Saves company
     * @return item is Company ID
     */
    saveCompany(item: Company, callback: ItemCallback<string>): void;

    /**
     * Deletes comapny
     */
    deleteCompany(id: string, callback: Callback): void;

    /**
     * Gets Tradings
     */
    getTradings(callback: ItemListCallback<Trading>): void;

    /**
     * Gets trading items of specified Trading
     */
    getTradingItems(
        tradingId: string,
        callback: ItemListCallback<TradingItem>
    ): void;

    /**
     * Saves Trading
     * @return item is trading ID
     */
    saveTrading(item: Trading, callback: ItemCallback<string>): void;

    /**
     * Saves Trading item of specified Trading
     * @return item is item ID
     */
    saveTradingItem(
        tradingId: string,
        item: TradingItem,
        callback: ItemCallback<string>
    ): void;

    /**
     * Deletes Trading
     */
    deleteTrading(tradingId: string, callback: Callback): void;

    /**
     * Deltes Trading item of specified Trading
     */
    deleteTradingItem(
        tradingId: string,
        itemId: string,
        callback: ItemCallback<string>
    ): void;

    /**
     * Gets Environment
     */
    getEnvironment(callback: ItemCallback<Environment>): void;

    /**
     * Saves Environment
     */
    saveEnvironment(env: Environment, callback: Callback): void;

    /**
     * Gets my company name
     */
    getMyCompanyName(callback: ItemCallback<string>): void;

    /**
     * Gets next number
     */
    getNextNumber(
        type: string,
        date: number,
        callback: ItemCallback<number>
    ): void;

    /**
     * Creates invoice
     */
    createInvoice(items: Array<any>, callback: ItemCallback<ArrayBuffer>): void;
}

interface Callback {
    success: () => void;
    error: (status: number, msg: string) => void;
}

interface ItemCallback<T> {
    success: (item: T) => void;
    error: (status: number, msg: string) => void;
}

interface ItemListCallback<T> {
    success: (item: Array<T>) => void;
    error: (status: number, msg: string) => void;
}
