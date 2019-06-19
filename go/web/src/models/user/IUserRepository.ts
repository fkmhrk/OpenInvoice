interface IUserRepository {
    getAll(): Promise<IUser[]>;
    save(user: IUser, password: string): Promise<IUser>;
    deleteUser(user: IUser): Promise<boolean>;
}

interface IUser {
    id: string;
    login_name: string;
    display_name: string;
    tel: string;
}
