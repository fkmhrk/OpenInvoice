interface IUserRepository {
    getAll(): Promise<IUser[]>;
}

interface IUser {
    id: string;
    login_name: string;
    display_name: string;
    tel: string;
}
