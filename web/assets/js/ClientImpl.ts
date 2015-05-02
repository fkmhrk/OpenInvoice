/// <reference path="./Client.ts"/>
var $;
var _;
var baseURL = '';
class AppClientImpl implements Client {
    url : string;
    
    constructor(url : string) {
        this.url = url;
    }
    
    login(username : string, password : string, callback : ItemCallback<string>) {
        var params : any = {
            username : username,
            password : password,
        };
        this.exec(this.url + '/api/v1/token', 'POST', null, params, {
            success : (json : any) => {
                callback.success(json.access_token);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }
    getTradings(token : string, callback : ItemListCallback<Trading>) {
        this.exec(this.url + '/api/v1/tradings', 'GET', token, null, {
            success : (json : any) => {
                callback.success(_.map(json.tradings, (item) => {
                    item.date = item.id;
                    return item;
                }));
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });            
    }
    
    getTradingItems(token : string, tradingId : string,
                    callback : ItemListCallback<TradingItem>) {
        var url = this.url + '/api/v1/tradings/' + tradingId + '/items';
        this.exec(url, 'GET', token, null, {
            success : (json : any) => {
                callback.success(_.map(json.items, (item) => {
                    item.sum = item.unit_price * item.amount;
                    return item;
                }));
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }
    
    getUsers(token : string, callback : ItemListCallback<User>) {
        var url = this.url + '/api/v1/users';
        this.exec(url, 'GET', token, null, {
            success : (json : any) => {
                callback.success(json.users);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });            
    }
    
    getCompanies(token : string, callback : ItemListCallback<Company>) {
        var url = this.url + '/api/v1/companies';
        this.exec(url, 'GET', token, null, {
            success : (json : any) => {
                callback.success(json.companies);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });                        
    }
    
    saveTrading(token : string, item : Trading, callback : ItemCallback<string>) {
        if (item.id === null) {
            this.createTrading(token, item, callback);
        } else {
            this.updateTrading(token, item, callback);
        }
    }
    
    saveTradingItem(token : string, tradingId : string,
                    item : TradingItem, callback : ItemCallback<string>) {
        if (item.id === null) {
            this.createTradingItem(token, tradingId, item, callback);
        } else {
            this.updateTradingItem(token, tradingId, item, callback);
        }            
    }
    deleteTradingItem(token : string, tradingId : string,
                      itemId : string, callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/tradings/' + tradingId +
            '/items/' + itemId;
        this.exec(url, 'DELETE', token, null, {
            success : (json : any) => {
                callback.success(itemId);
            },
            error : (status : any, body : any) => {
                if (status == 404) {
                    callback.success(itemId);
                } else {
                    callback.error(status, body.msg);
                }
            }
        });            
    }
    
    saveCompany(token : string, item : Company, callback : ItemCallback<string>) {
        if (item.id === null || item.id.length == 0 ) {
            this.createCompany(token, item, callback);
        } else {
            this.updateCompany(token, item, callback);
        }
    }
    
    private createTrading(token : string, item : Trading, callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/tradings';
        this.exec(url, 'POST', token, item, {
            success : (json : any) => {
                callback.success(json.id);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }
    
    private updateTrading(token : string, item : Trading, callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/tradings/' + item.id;
        this.exec(url, 'PUT', token, item, {
            success : (json : any) => {
                callback.success(item.id);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }
    
    private createTradingItem(token : string, tradingId : string,
                              item : TradingItem, callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/tradings/' + tradingId + '/items';
        this.exec(url, 'POST', token, item, {
            success : (json : any) => {
                callback.success(json.id);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }
    
    private updateTradingItem(token : string, tradingId : string,
                              item : TradingItem, callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/tradings/' + tradingId +
            '/items/' + item.id;
        this.exec(url, 'PUT', token, item, {
            success : (json : any) => {
                callback.success(item.id);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }
    
    private createCompany(token : string, item : Company, callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/companies';
        this.exec(url, 'POST', token, item, {
            success : (json : any) => {
                callback.success(json.id);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }
    
    private updateCompany(token : string, item : Company, callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/companies/' + item.id;
        this.exec(url, 'PUT', token, item, {
            success : (json : any) => {
                callback.success(json.id);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }        
    
    private exec(url : string, method : string, token : string,
                 params : any, callback : HttpCallback) {
        var data : any = {
            url : url,
            type : method,
            dataType : 'json',
            scriptCharset : 'utf-8',
            processData : false,
        };
        if (token != null) {
            data.headers = {
                authorization : 'bearer ' + token
            };
        }
        if (params != null) {
            data.data = JSON.stringify(params);
        }
        $.ajax(data)
            .done(function(data_, status, data) {
		if (data.status == 204) {
		    callback.success({});
		} else {
		    callback.success(JSON.parse(data.responseText));
		}
	    }).fail(function(data) {
		callback.error(data.status, JSON.parse(data.responseText));
	    });
    }
}

interface HttpCallback {
    success : (json : any) => void;
    error : (status : any, body : any) => void;
}

function createClient() {
    return new AppClientImpl(baseURL);
}