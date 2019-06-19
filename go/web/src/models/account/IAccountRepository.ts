interface IAccountRepository {
    signIn(username: string, password: string): Promise<string>;
    refresh(): Promise<void>;
}
