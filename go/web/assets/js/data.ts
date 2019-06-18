class User {
    id : string;
    login_name : string;
    display_name : string;
    tel : string;
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
    quotation_number : string;
    bill_date : number;
    bill_number : string;
    delivery_date : number;
    delivery_number : string;
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

class Environment {
    tax_rate : string;
    quotation_limit : string;
    closing_month : string;
    pay_limit : string;
    company_name : string;
    company_zip : string;
    company_address : string;
    company_tel : string;
    company_fax : string;
    company_bankname : string;
    company_bank_type : string;
    company_bank_num : string;
    company_bank_name : string;
}