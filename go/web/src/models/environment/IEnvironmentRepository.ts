interface IEnvironmentRepository {
    get(): Promise<IEnvironment>;
}

interface IEnvironment {
    tax_rate: string;
    quotation_limit: string;
    closing_month: string;
    pay_limit: string;
    company_name: string;
    company_zip: string;
    company_address: string;
    company_tel: string;
    company_fax: string;
    company_bankname: string;
    company_bank_type: string;
    company_bank_num: string;
    company_bank_name: string;
}
