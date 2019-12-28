import { isStatus200, getBody, isStatus204 } from "../../clients/Functions";

export default class TradingRepository implements ITradingRepository {
    private client: IAuthedClient;

    constructor(client: IAuthedClient) {
        this.client = client;
    }

    getAll(): Promise<ITrading[]> {
        const url = "/api/v1/tradings";
        return this.client
            .send(Method.GET, url, {}, null)
            .then(isStatus200)
            .then(getBody)
            .then((json: any) => {
                return <ITrading[]>json["tradings"];
            })
            .then((list: ITrading[]) => {
                list.forEach((item: ITrading) => {
                    item.date = item.id;
                });
                return list;
            });
    }

    async getById(id: string): Promise<ITrading> {
        const vals = await Promise.all([
            this.getTradingById(id),
            this.getTradingItemsById(id),
        ]);
        vals[0].items = vals[1];
        return vals[0];
    }

    async save(item: ITrading): Promise<ITrading> {
        let saved: ITrading;
        if (item.id.length == 0) {
            saved = await this.create(item);
        } else {
            saved = await this.update(item);
        }
        await this.saveItems(item);
        return saved!;
    }

    async deleteTrading(id: string): Promise<boolean> {
        const url = `/api/v1/tradings/${id}`;
        return this.client
            .send(Method.DELETE, url, {}, null)
            .then(isStatus204)
            .then((r: HTTPResponse) => true);
    }

    async deleteItems(id: string, items: ITradingItem[]): Promise<void> {
        if (id.length == 0) {
            return Promise.resolve();
        }
        const idSet = new Set<string>();
        items.forEach((item: ITradingItem) => {
            idSet.add(item.id);
        });
        idSet.forEach(async (itemId: string) => {
            await this.deleteItem(id, itemId);
        });
        return Promise.resolve();
    }

    getNextNumber(type: string, date: number): Promise<string> {
        const url = `/api/v1/sequences/${type}`;
        const params = {
            date: new Date(date).getTime(),
        };
        return this.client
            .send(Method.POST, url, {}, JSON.stringify(params))
            .then(isStatus200)
            .then(getBody)
            .then((body: any) => body["number"]);
    }

    private getTradingById(id: string): Promise<ITrading> {
        const url = `/api/v1/tradings/${id}`;
        return this.client
            .send(Method.GET, url, {}, null)
            .then(isStatus200)
            .then(getBody)
            .then((body: any) => <ITrading>body);
    }

    private getTradingItemsById(id: string): Promise<ITradingItem[]> {
        const url = `/api/v1/tradings/${id}/items`;
        return this.client
            .send(Method.GET, url, {}, null)
            .then(isStatus200)
            .then(getBody)
            .then((body: any) => <ITradingItem[]>body["items"]);
    }

    private async create(item: ITrading): Promise<ITrading> {
        const url = "/api/v1/tradings";
        const params = {
            company_id: item.company_id,
            title_type: item.title_type,
            subject: item.subject,
            work_from: item.work_from,
            work_to: item.work_to,
            total: item.total,
            quotation_date: item.quotation_date,
            quotation_number: item.quotation_number,
            bill_date: item.bill_date,
            bill_number: item.bill_number,
            delivery_date: item.delivery_date,
            delivery_number: item.delivery_number,
            tax_rate: item.tax_rate,
            assignee: item.assignee,
            product: item.product,
            memo: item.memo,
        };
        const resp = await this.client.send(
            Method.POST,
            url,
            {},
            JSON.stringify(params)
        );
        if (resp.status != 201) {
            throw resp;
        }
        item.id = resp.body.id;
        return item;
    }

    private async update(item: ITrading): Promise<ITrading> {
        const url = `/api/v1/tradings/${item.id}`;
        const params = {
            company_id: item.company_id,
            title_type: item.title_type,
            subject: item.subject,
            work_from: item.work_from,
            work_to: item.work_to,
            total: item.total,
            quotation_date: item.quotation_date,
            quotation_number: item.quotation_number,
            bill_date: item.bill_date,
            bill_number: item.bill_number,
            delivery_date: item.delivery_date,
            delivery_number: item.delivery_number,
            tax_rate: item.tax_rate,
            assignee: item.assignee,
            product: item.product,
            memo: item.memo,
        };
        const resp = await this.client.send(
            Method.PUT,
            url,
            {},
            JSON.stringify(params)
        );
        if (resp.status != 200) {
            throw resp;
        }
        return item;
    }

    private async saveItems(trading: ITrading) {
        trading.items.forEach(async (item: ITradingItem) => {
            if (item.id.length == 0) {
                await this.createItem(trading.id, item);
            } else {
                await this.updateItem(trading.id, item);
            }
        });
    }

    private async createItem(tradingId: string, item: ITradingItem) {
        const url = `/api/v1/tradings/${tradingId}/items`;
        const params = {
            sort_order: item.sort_order,
            subject: item.subject,
            unit_price: item.unit_price,
            amount: item.amount,
            degree: item.degree,
            tax_type: item.tax_type,
            memo: item.memo,
        };
        const resp = await this.client.send(
            Method.POST,
            url,
            {},
            JSON.stringify(params)
        );
        if (resp.status != 201) {
            throw resp;
        }
        item.id = resp.body.id;
    }

    private async updateItem(tradingId: string, item: ITradingItem) {
        const url = `/api/v1/tradings/${tradingId}/items/${item.id}`;
        const params = {
            sort_order: item.sort_order,
            subject: item.subject,
            unit_price: item.unit_price,
            amount: item.amount,
            degree: item.degree,
            tax_type: item.tax_type,
            memo: item.memo,
        };
        const resp = await this.client.send(
            Method.PUT,
            url,
            {},
            JSON.stringify(params)
        );
        if (resp.status != 200) {
            throw resp;
        }
    }

    private async deleteItem(tradingId: string, itemId: string) {
        const url = `/api/v1/tradings/${tradingId}/items/${itemId}`;
        const resp = await this.client.send(Method.DELETE, url, {}, null);
        if (resp.status != 204) {
            throw resp;
        }
    }
}
