class User {
    id : string;
    display_name : string;
}

class Company {
    id : string;
    name : string;
    zip : string;
    address : string;
    phone : string;
    fax : string;
    unit : string;
    assignee : string;
}

class Trading {
    id : string;
    date : string;
    company_id : string;
    company_name : string;
    title_type : number;
    subject : string;
    work_from : number;
    work_to : number;
    quotation_date : number;
    bill_date : number;
    tax_rate : number;
    assignee : string;
    product : string;
    total : number;
    modified_time : number;
}

class TradingItem {
    id : string;
    sort_order : number;
    subject : string;
    unit_price : number;
    amount : number;
    degree : string;
    tax_type : number;
    memo : string;
    sum : number;
}

var userList : Array<User> = [];
userList.push({
    'id' : 'user1',
    'display_name' : '秋葉 秀樹'
});
userList.push({
    'id' : 'user2',
    'display_name' : '秋葉 ちひろ'
});

var companyList : Array<Company> = [];
var company = new Company();
company.id = 'company1';
company.name = '株式会社NRI';
company.unit = '生産革新部';
company.zip = '111-2222';
company.address = '東京都新宿区新宿2-2-4';
company.phone = '090-1111-2222';
company.assignee = '東山 正二';
companyList.push(company);

company = new Company();
company.id = 'company2';
company.name = '株式会社カルチャー';
company.unit = '';
company.zip = '111-2222';
company.address = '東京都渋谷区神宮前4-8-10';
company.phone = '090-3333-4444';
company.assignee = '';
companyList.push(company);

company = new Company();
company.id = 'company3';
company.name = '株式会社忍者';
company.unit = '';
company.zip = '111-2222';
company.address = '東京都三鷹市一丁目3-5-95';
company.phone = '090-3333-4444';
company.assignee = '';
companyList.push(company);

company = new Company();
company.id = 'company4';
company.name = 'ツクロウ';
company.unit = '';
company.zip = '111-2222';
company.address = '広島県福山市鞆の浦4154-5';
company.phone = '090-3333-4444';
company.assignee = '';
companyList.push(company);

var sheetList : Array<Trading> = [];
for (var i = 0 ; i < 10 ; ++i ){
    sheetList.push(
    {
        'id' : 'idA' + i,
        'date' : '1432542408000',
        'company_id' : 'company1',
        'company_name' : '株式会社ABC',
        'title_type' : 0,
        'subject' : '画面作成',
        'work_from' : 1432542408000,
        'work_to' : 1432542408000,
        'quotation_date' : 1432542408000,
        'bill_date' : 1432542408000,
        'tax_rate' : 8,
        'assignee' : 'user1',
        'product' : '成果物A',
        'total' : 650000,
        'modified_time' : 1432542408000,
        'quotation_number' : 'E0124',
        'bill_number' : ''
    });
    sheetList.push(
    {
        'id' : 'idB' + i,
        'date' : '1432542408000',
        'company_id' : 'company2',
        'company_name' : '株式会社NRI',
        'title_type' : 1,
        'subject' : '【コンサルツールモック】デザイン画面作成',
        'work_from' : 1431505608,
        'work_to' : 1431505608,
        'quotation_date' : 1431505608,
        'bill_date' : 1431505608,
        'tax_rate' : 10,
        'assignee' : 'user2',
        'product' : '成果物B',
        'total' : 1030875,        
        'modified_time' : 1431505608,
        'quotation_number' : 'E0123',
        'bill_number' : 'V0238'
    });
}

var tradingItemList : Array<TradingItem> = [];