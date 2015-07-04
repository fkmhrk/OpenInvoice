/// <reference path="./Functions.ts"/>
module ClientValidator {
    export function isValidLogin(username : string, password : string, callback : ItemCallback<string>) : boolean {
        if (Utils.isEmpty(username)) {
            callback.error(1000, "Username must not be empty.");
            return false;
        }
        if (Utils.isEmpty(password)) {
            callback.error(1001, "Password must not be empty.");
            return false;
        }
        return true;
    }

    export function isValidCreateUser(loginName : string, displayName : string, tel : string,
               password : string, callback : ItemCallback<User>) {
        if (Utils.isEmpty(loginName)) {
            callback.error(1000, "LoginName must not be empty.");
            return false;
        }
        if (Utils.isEmpty(displayName)) {
            callback.error(1001, "DisplayName must not be empty.");
            return false;
        }
        if (Utils.isEmpty(tel)) {
            callback.error(1002, "Tel must not be empty.");
            return false;
        }
        if (Utils.isEmpty(password)) {
            callback.error(1003, "Password must not be empty.");
            return false;
        }
        if (password.length < 6) {
            callback.error(1004, "Password must be more than 6 characters.");
            return false;            
        }
        return true;
    }

    export function isValidSaveCompany(item : Company, callback : ItemCallback<string>) : boolean {
        if (item == null) {
            callback.error(1000, "Item must not be empty.");
            return false;
        }
        if (Utils.isEmpty(item.name)) {
            callback.error(1001, "Name must not be empty.");
            return false;
        }
        if (item.zip == null) { item.zip = ''; }
        if (item.address == null) { item.address = ''; }
        if (item.phone == null) { item.phone = ''; }
        if (item.fax == null) { item.fax = ''; } 
        if (item.unit == null) { item.unit = ''; }       
        return true;
    }

    export function isValidSaveTrading(item : Trading, callback : ItemCallback<string>) : boolean {
        if (item == null) {
            callback.error(1000, "Item must not be empty.");
            return false;
        }
        if (Utils.isEmpty(item.subject)) {
            callback.error(1001, "Subject must not be empty.");
            return false;            
        }
        if (item.work_to < item.work_from) {
            callback.error(1002, "Invalid work_from and work_to.");
            return false;
        }
        if (item.tax_rate < 0) {
            callback.error(1003, "tax_rate must be positive(tax_rate > 0).");
            return false;            
        }
        if (item.product == null) { item.product = ''; }
        return true;
    }

    export function isValidSaveTradingItem(tradingId : string, item : TradingItem, callback : ItemCallback<string>) : boolean {
        if (Utils.isEmpty(tradingId)) {
            callback.error(1000, "trading ID must not be empty.");
            return false;
        }
        if (item == null) {
            callback.error(1001, "item must not be empty.");
            return false;
        }
        if (Utils.isEmpty(item.subject)) {
            callback.error(1002, "subject must not be empty.");
            return false;
        }        
        if (item.tax_type < 0 || item.tax_type > 2) {
            callback.error(1003, "tax_type must be 0, 1, 2.");
            return false;
        }
        if (item.degree == null) { item.degree = ''; }
        if (item.memo == null) { item.memo = ''; }
        return true;
    }

    export function isValidSaveEnvironment(env : Environment, callback : Callback) : boolean {
        if (env == null) {
            callback.error(1000, "item must not be empty.");
            return false;
        }
        if (env.tax_rate == null) { env.tax_rate = ''; }
        if (env.quotation_limit == null) { env.quotation_limit = ''; }
        if (env.closing_month == null) { env.closing_month = ''; }
        if (env.pay_limit == null) { env.pay_limit = ''; }
        if (env.company_name == null) { env.company_name = ''; }
        if (env.company_zip == null) { env.company_zip = ''; }
        if (env.company_address == null) { env.company_address = ''; }
        if (env.company_tel == null) { env.company_tel = ''; }
        if (env.company_fax == null) { env.company_fax = ''; }
        if (env.company_bankname == null) { env.company_bankname = ''; }
        if (env.company_bank_type == null) { env.company_bank_type = ''; }
        if (env.company_bank_num == null) { env.company_bank_num = ''; }
        if (env.company_bank_name == null) { env.company_bank_name = ''; }
        return true;
    }
}