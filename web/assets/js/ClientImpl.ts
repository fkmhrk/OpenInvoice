/// <reference path="./Client.ts"/>
var $;
var _;
module Invoice {
    export class AppClientImpl implements AppClient {
        url : string;

        constructor(url : string) {
            this.url = url;
        }
        
        login(username : string, password : string, callback : LoginCallback) {
            var params : any = {
                username : username,
                password : password,
            };
            this.exec(this.url + '/api/v1/token', 'POST', null, params, {
                success : (json : any) => {
                    callback.success(json.access_token);
                },
                error : (status : any, body : any) => {
                    callback.error(body.msg);
                }
            });
        }
        getTradings(token : string, callback : TradingCallback) {
            this.exec(this.url + '/api/v1/tradings', 'GET', token, null, {
                success : (json : any) => {
                    callback.success(_.map(json.tradings, (item) => {
                        item.date = item.id;
                        return item;
                    }));
                },
                error : (status : any, body : any) => {
                    callback.error(body.msg);
                }
            });            
        }

        getTradingItems(token : string, tradingId : string,
                        callback : TradingItemCallback) {
            var url = this.url + '/api/v1/tradings/' + tradingId + '/items';
            this.exec(url, 'GET', token, null, {
                success : (json : any) => {
                    callback.success(_.map(json.items, (item) => {
                        item.sum = item.unit_price * item.amount;
                        return item;
                    }));
                },
                error : (status : any, body : any) => {
                    callback.error(body.msg);
                }
            });
        }

        getUsers(token : string, callback : UserListCallback) {
            var url = this.url + '/api/v1/users';
            this.exec(url, 'GET', token, null, {
                success : (json : any) => {
                    callback.success(json.users);
                },
                error : (status : any, body : any) => {
                    callback.error(body.msg);
                }
            });            
        }

        getCompanies(token : string, callback : CompanyListCallback) {
            var list : Array<Company> = [];
            list.push({
                id : "dummy0001",
                name : "仮会社会社",
                zip : "111-2222",
                address : "日本",
                phone : "03-1111-2222",
                unit : "開発部",
            });
            callback.success(list);            
        }

        saveTrading(token : string, item : Trading, callback : SaveCallback) {
            if (item.id === null) {
                this.createTrading(token, item, callback);
            } else {
                this.updateTrading(token, item, callback);
            }
        }

        saveTradingItem(token : string, tradingId : string,
                        item : TradingItem, callback : SaveCallback) {
            if (item.id === null) {
                this.createTradingItem(token, tradingId, item, callback);
            } else {
                this.updateTradingItem(token, tradingId, item, callback);
            }            
        }

        private createTrading(token : string, item : Trading, callback : SaveCallback) {
            var url = this.url + '/api/v1/tradings';
            this.exec(url, 'POST', token, item, {
                success : (json : any) => {
                    callback.success(json.id);
                },
                error : (status : any, body : any) => {
                    callback.error(body.msg);
                }
            });
        }

        private updateTrading(token : string, item : Trading, callback : SaveCallback) {
            var url = this.url + '/api/v1/tradings/' + item.id;
            this.exec(url, 'PUT', token, item, {
                success : (json : any) => {
                    callback.success(item.id);
                },
                error : (status : any, body : any) => {
                    callback.error(body.msg);
                }
            });
        }

        private createTradingItem(token : string, tradingId : string,
                                  item : TradingItem, callback : SaveCallback) {
            var url = this.url + '/api/v1/tradings/' + tradingId + '/items';
            this.exec(url, 'POST', token, item, {
                success : (json : any) => {
                    callback.success(json.id);
                },
                error : (status : any, body : any) => {
                    callback.error(body.msg);
                }
            });
        }

        private updateTradingItem(token : string, tradingId : string,
                                  item : TradingItem, callback : SaveCallback) {
            var url = this.url + '/api/v1/tradings/' + tradingId +
                '/items/' + item.id;
            this.exec(url, 'PUT', token, item, {
                success : (json : any) => {
                    callback.success(item.id);
                },
                error : (status : any, body : any) => {
                    callback.error(body.msg);
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
}