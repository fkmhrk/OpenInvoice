/// <reference path="./Client.ts"/>
var $;
module Invoice {
    export class MockClient implements AppClient {
        login(username : string, password : string, callback : LoginCallback) {
            callback.success('token1122');
        }
        getTradings(token : string, callback : TradingCallback) {
            var tradings : Array<Trading> = [];
            for (var i = 0 ; i < 10 ; ++i) {
                tradings.push({
                    id : 'trade1122' + i,
                    date : 'trade1122' + i,
                    company_id : "会社ID" + i,
                    title_type : 0,
                    subject : "件名" + i,  
                    work_from : 1122,
                    work_to : 2233,
                    quotation_date : 1423502769379,
                    bill_date : 5555,
                    tax_rate : 8,
                    assignee : "担当者ID" + i,   
                    product : "成果物" + i,
                });
            }
            callback.success(tradings);
        }

        getTradingItems(token : string, tradingId : string, callback : TradingItemCallback) {
            var tradings : Array<TradingItem> = [];
            for (var i = 0 ; i < 10 ; ++i) {
                tradings.push({
                    id : 'item111' + i,
                    subject : "件名" + i,
                    unit_price : i * 100 + 200,
                    amount : i * 3 + 2,
                    degree : "人月",
                    tax_type : 1,
                    memo : "備考" + i,
                    sum : 1000,
               });
            }
            callback.success(tradings);
        }
        
        getUsers(token : string, callback : UserListCallback) {
            var list : Array<User> = [];
            for (var i = 0 ; i < 10 ; ++i) {
                list.push({
                    id : "担当者ID" + i,
                    display_name : '担当' + i,
                });
            }
            callback.success(list);
        }

        getCompanies(token : string, callback : CompanyListCallback) {
            var list : Array<Company> = [];
            for (var i = 0 ; i < 10 ; ++i) {
                list.push({
                    id : "company" + i,
                    name : "会社" + i,
                    zip : "111-222" + i,
                    address : "日本の" + i,
                    phone : "03-1111-222" + i,
                    unit : "第" + i + "開発部",
                });
            }
            callback.success(list);
        }
        saveTrading(token : string, item : Trading, callback : SaveCallback) {
            callback.success('id1122');
        }
        saveTradingItem(token : string, tradingId : string,
                        item : TradingItem, callback : SaveCallback) {
            callback.success('item1122');
        }
        saveCompany(token : string, item : Company, callback : SaveCallback) {
            callback.success('company1122');
        }
    }
}