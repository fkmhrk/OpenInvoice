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