/// <reference path="./Client.ts"/>
var $;
var _;
var Invoice;
(function (Invoice) {
    var AppClientImpl = (function () {
        function AppClientImpl(url) {
            this.url = url;
        }
        AppClientImpl.prototype.login = function (username, password, callback) {
            var params = {
                username: username,
                password: password
            };
            this.exec(this.url + '/api/v1/token', 'POST', null, params, {
                success: function (json) {
                    callback.success(json.access_token);
                },
                error: function (status, body) {
                    callback.error(body.msg);
                }
            });
        };
        AppClientImpl.prototype.getTradings = function (token, callback) {
            this.exec(this.url + '/api/v1/tradings', 'GET', token, null, {
                success: function (json) {
                    callback.success(_.map(json.tradings, function (item) {
                        item.date = item.id;
                        return item;
                    }));
                },
                error: function (status, body) {
                    callback.error(body.msg);
                }
            });
        };

        AppClientImpl.prototype.getTradingItems = function (token, tradingId, callback) {
            var url = this.url + '/api/v1/tradings/' + tradingId + '/items';
            this.exec(url, 'GET', token, null, {
                success: function (json) {
                    callback.success(_.map(json.items, function (item) {
                        item.sum = item.unit_price * item.amount;
                        return item;
                    }));
                },
                error: function (status, body) {
                    callback.error(body.msg);
                }
            });
        };

        AppClientImpl.prototype.getUsers = function (token, callback) {
            var url = this.url + '/api/v1/users';
            this.exec(url, 'GET', token, null, {
                success: function (json) {
                    callback.success(json.users);
                },
                error: function (status, body) {
                    callback.error(body.msg);
                }
            });
        };

        AppClientImpl.prototype.getCompanies = function (token, callback) {
            var url = this.url + '/api/v1/companies';
            this.exec(url, 'GET', token, null, {
                success: function (json) {
                    callback.success(json.companies);
                },
                error: function (status, body) {
                    callback.error(body.msg);
                }
            });
            var list = [];
            list.push({
                id: "dummy0001",
                name: "仮会社会社",
                zip: "111-2222",
                address: "日本",
                phone: "03-1111-2222",
                unit: "開発部"
            });
            callback.success(list);
        };

        AppClientImpl.prototype.saveTrading = function (token, item, callback) {
            if (item.id === null) {
                this.createTrading(token, item, callback);
            } else {
                this.updateTrading(token, item, callback);
            }
        };

        AppClientImpl.prototype.saveTradingItem = function (token, tradingId, item, callback) {
            if (item.id === null) {
                this.createTradingItem(token, tradingId, item, callback);
            } else {
                this.updateTradingItem(token, tradingId, item, callback);
            }
        };
        AppClientImpl.prototype.deleteTradingItem = function (token, tradingId, itemId, callback) {
            var url = this.url + '/api/v1/tradings/' + tradingId + '/items/' + itemId;
            this.exec(url, 'DELETE', token, null, {
                success: function (json) {
                    callback.success(itemId);
                },
                error: function (status, body) {
                    if (status == 404) {
                        callback.success(itemId);
                    } else {
                        callback.error(body.msg);
                    }
                }
            });
        };

        AppClientImpl.prototype.saveCompany = function (token, item, callback) {
            if (item.id === null || item.id.length == 0) {
                this.createCompany(token, item, callback);
            } else {
                this.updateCompany(token, item, callback);
            }
        };

        AppClientImpl.prototype.createTrading = function (token, item, callback) {
            var url = this.url + '/api/v1/tradings';
            this.exec(url, 'POST', token, item, {
                success: function (json) {
                    callback.success(json.id);
                },
                error: function (status, body) {
                    callback.error(body.msg);
                }
            });
        };

        AppClientImpl.prototype.updateTrading = function (token, item, callback) {
            var url = this.url + '/api/v1/tradings/' + item.id;
            this.exec(url, 'PUT', token, item, {
                success: function (json) {
                    callback.success(item.id);
                },
                error: function (status, body) {
                    callback.error(body.msg);
                }
            });
        };

        AppClientImpl.prototype.createTradingItem = function (token, tradingId, item, callback) {
            var url = this.url + '/api/v1/tradings/' + tradingId + '/items';
            this.exec(url, 'POST', token, item, {
                success: function (json) {
                    callback.success(json.id);
                },
                error: function (status, body) {
                    callback.error(body.msg);
                }
            });
        };

        AppClientImpl.prototype.updateTradingItem = function (token, tradingId, item, callback) {
            var url = this.url + '/api/v1/tradings/' + tradingId + '/items/' + item.id;
            this.exec(url, 'PUT', token, item, {
                success: function (json) {
                    callback.success(item.id);
                },
                error: function (status, body) {
                    callback.error(body.msg);
                }
            });
        };

        AppClientImpl.prototype.createCompany = function (token, item, callback) {
            var url = this.url + '/api/v1/companies';
            this.exec(url, 'POST', token, item, {
                success: function (json) {
                    callback.success(json.id);
                },
                error: function (status, body) {
                    callback.error(body.msg);
                }
            });
        };

        AppClientImpl.prototype.updateCompany = function (token, item, callback) {
            var url = this.url + '/api/v1/companies/' + item.id;
            this.exec(url, 'PUT', token, item, {
                success: function (json) {
                    callback.success(json.id);
                },
                error: function (status, body) {
                    callback.error(body.msg);
                }
            });
        };

        AppClientImpl.prototype.exec = function (url, method, token, params, callback) {
            var data = {
                url: url,
                type: method,
                dataType: 'json',
                scriptCharset: 'utf-8',
                processData: false
            };
            if (token != null) {
                data.headers = {
                    authorization: 'bearer ' + token
                };
            }
            if (params != null) {
                data.data = JSON.stringify(params);
            }
            $.ajax(data).done(function (data_, status, data) {
                if (data.status == 204) {
                    callback.success({});
                } else {
                    callback.success(JSON.parse(data.responseText));
                }
            }).fail(function (data) {
                callback.error(data.status, JSON.parse(data.responseText));
            });
        };
        return AppClientImpl;
    })();
    Invoice.AppClientImpl = AppClientImpl;
})(Invoice || (Invoice = {}));
