import AccountRepository from "./account/AccountRepository";
import EnvironmentRepository from "./environment/EnvironmentRepository";
import AccessToken from "./token/AccessToken";
import TradingRepository from "./trading/TradingRepository";
import UserRepository from "./user/UserRepository";
import CompanyRepository from "./company/CompanyRepository";

export class Models implements IModels {
    account: IAccountRepository;
    environment: IEnvironmentRepository;
    trading: ITradingRepository;
    user: IUserRepository;
    company: ICompanyRepository;
    token: IAccessToken;
    constructor(
        client: HTTPClient,
        authedClient: IAuthedClient,
        token: AccessToken
    ) {
        this.account = new AccountRepository(client, token);
        this.environment = new EnvironmentRepository(authedClient);
        this.trading = new TradingRepository(authedClient);
        this.user = new UserRepository(authedClient);
        this.company = new CompanyRepository(authedClient);
        this.token = token;
    }
}
