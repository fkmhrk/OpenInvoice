interface IModels {
    account: IAccountRepository;
    environment: IEnvironmentRepository;
    trading: ITradingRepository;
    user: IUserRepository;
    company: ICompanyRepository;
    token: IAccessToken;
}
