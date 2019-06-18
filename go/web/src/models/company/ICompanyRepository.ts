interface ICompanyRepository {
    getAll(): Promise<ICompany[]>;
    save(company: ICompany): Promise<ICompany>;
    deleteCompany(company: ICompany): Promise<boolean>;
}

interface ICompany {
    id: string;
    name: string;
    zip: string;
    address: string;
    phone: string;
    fax: string;
    unit: string;
    assignee: string;
}
