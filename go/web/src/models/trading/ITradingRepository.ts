interface ITradingRepository {
    getAll(): Promise<ITrading[]>;
    getById(id: string): Promise<ITrading>;
    save(item: ITrading): Promise<ITrading>;
    deleteTrading(id: string): Promise<boolean>;
    deleteItems(id: string, items: ITradingItem[]): Promise<void>;
    getNextNumber(type: string, date: number): Promise<string>;
}

interface ITrading {
    id: string;
    date: string;
    company_id: string;
    company_name: string;
    title_type: number;
    subject: string;
    work_from: number;
    work_to: number;
    quotation_date: number;
    quotation_number: string;
    bill_date: number;
    bill_number: string;
    delivery_date: number;
    delivery_number: string;
    tax_rate: number;
    assignee: string;
    product: string;
    memo: string;
    total: number;
    modified_time: number;

    items: ITradingItem[];
}

interface ITradingItem {
    id: string;
    sort_order: number;
    subject: string;
    unit_price: number;
    amount: number;
    degree: string;
    tax_type: number;
    memo: string;
    sum: number;
}
