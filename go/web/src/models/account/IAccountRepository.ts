interface IAccountRepository {
    signIn(username: string, password: string): Promise<string>;
}
