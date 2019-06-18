interface ICompanyRepository {
    getAll(): Promise<ICompany[]>;
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
