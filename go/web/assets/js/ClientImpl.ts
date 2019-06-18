/// <reference path="./Client.ts"/>
/// <reference path="./ClientValidator.ts"/>
var $;
var _;
var baseURL = '';
class AppClientImpl implements Client {
    url : string;
    accessToken : string;
    refreshToken : string;
    is_admin : boolean;
    isRetry : boolean;
    
    constructor(url : string) {
        this.url = url;
        this.is_admin = false;
        this.isRetry = false;
    }

    setRefreshToken(refreshToken : string) {
        if (refreshToken == null) { return; }
        this.accessToken = '';
        this.refreshToken = refreshToken;
    }

    isAdmin() : boolean {
        return this.is_admin;
    }

    getAccessToken() : string {
        return this.accessToken;
    }
    
    login(username : string, password : string, callback : ItemCallback<string>) {
        if (!ClientValidator.isValidLogin(username, password, callback)) { return; }
        var params : any = {
            username : username,
            password : password,
        };
        this.exec(this.url + '/api/v1/token', 'POST', null, params, {
            success : (json : any) => {
                this.accessToken = json.access_token;
                this.refreshToken = json.refresh_token;
                this.is_admin = json.is_admin;
                callback.success(json.refresh_token);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }
    getTradings(callback : ItemListCallback<Trading>) {
        this.exec(this.url + '/api/v1/tradings', 'GET', this.accessToken, null, {
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
    
    getTradingItems(tradingId : string, callback : ItemListCallback<TradingItem>) {
        var url = this.url + '/api/v1/tradings/' + tradingId + '/items';
        this.exec(url, 'GET', this.accessToken, null, {
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

    createUser(loginName : string, displayName : string, tel : string,
               password : string, callback : ItemCallback<User>) {
        if (!ClientValidator.isValidCreateUser(loginName, displayName, tel, password, callback)) { return; }
        var url = this.url + '/api/v1/users';
        var params = {
            login_name : loginName,
            display_name : displayName,
            tel : tel,
            password : password,
        };
        this.exec(url, 'POST', this.accessToken, params, {
            success : (json : any) => {
                callback.success(json);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });        
    }    
    
    getUsers(callback : ItemListCallback<User>) {
        var url = this.url + '/api/v1/users';
        this.exec(url, 'GET', this.accessToken, null, {
            success : (json : any) => {
                callback.success(json.users);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });            
    }

    saveUser(user : User, password : string, callback : ItemCallback<User>) {
        if (!ClientValidator.isValidSaveUser(user, password, callback)) { return; }
        var url = this.url + '/api/v1/users/' + user.id;
        var params = {
            id : user.id,
            login_name : user.login_name,
            display_name : user.display_name,
            tel : user.tel,
            password : password,
        };
        this.exec(url, 'PUT', this.accessToken, params, {
            success : (json : any) => {
                callback.success(params);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }

    deleteUser(id : string, callback : Callback) {
        var url = this.url + '/api/v1/users/' + id;
        this.exec(url, 'DELETE', this.accessToken, null, {
            success : (json : any) => {
                callback.success();
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });        
    }
    
    getCompanies(callback : ItemListCallback<Company>) {
        var url = this.url + '/api/v1/companies';
        this.exec(url, 'GET', this.accessToken, null, {
            success : (json : any) => {
                callback.success(json.companies);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });                        
    }
    
    saveTrading(item : Trading, callback : ItemCallback<string>) {
        if (!ClientValidator.isValidSaveTrading(item, callback)) { return; }
        if (item.id === null) {
            this.createTrading(item, callback);
        } else {
            this.updateTrading(item, callback);
        }
    }
    
    saveTradingItem(tradingId : string, item : TradingItem, callback : ItemCallback<string>) {
        if (!ClientValidator.isValidSaveTradingItem(tradingId, item, callback)) { return; }
        if (item.id === null) {
            this.createTradingItem(tradingId, item, callback);
        } else {
            this.updateTradingItem(tradingId, item, callback);
        }            
    }

    deleteTrading(tradingId : string, callback : Callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId;
        this.exec(url, 'DELETE', this.accessToken, null, {
            success : (json : any) => {
                callback.success();
            },
            error : (status : any, body : any) => {
                if (status == 404) {
                    callback.success();
                } else {
                    callback.error(status, body.msg);
                }
            }
        });                    
    }
    
    deleteTradingItem(tradingId : string, itemId : string, callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/tradings/' + tradingId +
            '/items/' + itemId;
        this.exec(url, 'DELETE', this.accessToken, null, {
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
    
    saveCompany(item : Company, callback : ItemCallback<string>) {
        if (!ClientValidator.isValidSaveCompany(item, callback)) { return; }
        if (item.id === null || item.id.length == 0 ) {
            this.createCompany(item, callback);
        } else {
            this.updateCompany(item, callback);
        }
    }

    deleteCompany(id : string, callback : Callback) {
        var url = this.url + '/api/v1/companies/' + id;
        this.exec(url, 'DELETE', this.accessToken, null, {
            success : (json : any) => {
                callback.success();
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });                
    }

    getEnvironment(callback : ItemCallback<Environment>) {
        var url = this.url + '/api/v1/environments';
        this.exec(url, 'GET', this.accessToken, null, {
            success : (json : any) => {
                callback.success(json);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });        
    }

    saveEnvironment(env : Environment, callback : Callback) {
        if (!ClientValidator.isValidSaveEnvironment(env, callback)) { return; }
        var url = this.url + '/api/v1/environments';
        this.exec(url, 'PUT', this.accessToken, env, {
            success : (json : any) => {
                callback.success();
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }

    getMyCompanyName(callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/myCompany/name';
        this.exec(url, 'GET', null, null, {
            success : (json : any) => {
                callback.success(json['name']);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });        
    }

    getNextNumber(type : string, date : number, callback : ItemCallback<number>) {
        var url = this.url + '/api/v1/sequences/' + type;
        var params = {
            date : date,
        }
        this.exec(url, 'POST', this.accessToken, params, {
            success : (json : any) => {
                callback.success(json['number']);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });                
    }

    createInvoice(items : Array<any>, callback : ItemCallback<ArrayBuffer>) {
        var params = {
            access_token : this.accessToken,
            customer : {
                name : "サンプル会社",
                address : "サンプル住所",
            },
            myCompany : {
                name : "サンプル会社",
                address : "住所\n\n担当",
            },
            item_title : 'ご請求書在中',
            date : new Date().getTime(),
            items : items,
        };
        var xhr = new XMLHttpRequest();
        xhr.open('POST', this.url + '/php/invoice.php', true);
        xhr.responseType = 'arraybuffer';
        xhr.onload = function(e : any){
            if (this.status == 200) {
                callback.success(this.response);
            } else {
                callback.error(this.status, this.response);
            }
        };
        xhr.send(JSON.stringify(params));
    }
    
    private createTrading(item : Trading, callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/tradings';
        this.exec(url, 'POST', this.accessToken, item, {
            success : (json : any) => {
                callback.success(json.id);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }
    
    private updateTrading(item : Trading, callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/tradings/' + item.id;
        this.exec(url, 'PUT', this.accessToken, item, {
            success : (json : any) => {
                callback.success(item.id);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }
    
    private createTradingItem(tradingId : string, item : TradingItem, callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/tradings/' + tradingId + '/items';
        this.exec(url, 'POST', this.accessToken, item, {
            success : (json : any) => {
                callback.success(json.id);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }
    
    private updateTradingItem(tradingId : string, item : TradingItem, callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/tradings/' + tradingId +
            '/items/' + item.id;
        this.exec(url, 'PUT', this.accessToken, item, {
            success : (json : any) => {
                callback.success(item.id);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }
    
    private createCompany(item : Company, callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/companies';
        this.exec(url, 'POST', this.accessToken, item, {
            success : (json : any) => {
                callback.success(json.id);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }
    
    private updateCompany(item : Company, callback : ItemCallback<string>) {
        var url = this.url + '/api/v1/companies/' + item.id;
        this.exec(url, 'PUT', this.accessToken, item, {
            success : (json : any) => {
                callback.success(json.id);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });
    }

    private tokenRefresh(url : string, method : string, params : any, callback : HttpCallback) {
        var refreshURL = this.url + '/api/v1/token/refresh';
        var refreshParams = {
            token : this.refreshToken
        };
        this.exec(refreshURL, 'POST', null, refreshParams, {
            success : (json : any) => {
                this.accessToken = json.access_token;
                this.is_admin = json.is_admin;
                this.isRetry = true;
                this.exec(url, method, this.accessToken, params, callback);
            },
            error : (status : any, body : any) => {
                callback.error(status, body.msg);
            }
        });        
    }
    
    private exec(url : string, method : string, token : string, params : any, callback : HttpCallback) {
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
        $.ajax(data).done((data_ : any, status : any, data : any) => {
            this.isRetry = false;
	    if (data.status == 204) {
		callback.success({});
	    } else {
		callback.success(JSON.parse(data.responseText));
	    }
	}).fail((data : any) => {
            if (data.status == 204) {
                this.isRetry = false;
                callback.success({});
            } else if (data.status == 401) {
                if (this.isRetry) {
                    this.isRetry = false;
                    callback.error(data.status, JSON.parse(data.responseText));
                } else {
                    this.isRetry = true;
                    this.tokenRefresh(url, method, params, callback);
                }
            } else {
                this.isRetry = false;
		callback.error(data.status, JSON.parse(data.responseText));
            }
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