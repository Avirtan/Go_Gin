import CreateAxios from "./HttpCommon";
import { AxiosInstance } from "axios";

class ProfileService {
    client: AxiosInstance;
    file: AxiosInstance;

    constructor() {
        this.client = CreateAxios("/api/v1/admin/profile");
        this.file = CreateAxios("/api/v1/admin/profile", { "Content-type": "multipart/form-data" });
    }
    GetUsers(): Promise<any> {
        return this.client.get("/users");
    }
    GetUserById(id: string): Promise<any> {
        return this.client.get("/user/" + id);
    }
    UpdateUser(data: any): Promise<any> {
        return this.client.patch("/user", data);
    }
    GetRoles(): Promise<any> {
        return this.client.get("/roles");
    }
    CreateRole(data: any): Promise<any> {
        return this.client.post("/role", data);
    }
    DeleteRole(data: any): Promise<any> {
        return this.client.delete("/role", { data: data });
    }
    // SetRole(data: any): Promise<any> {
    //     return this.client.post("/setRole", data);
    // }
    UploadThuings(data: any, name: string): Promise<any> {
        return this.file.post("/upload/" + name, data);
    }
    UploadThing(data: any, name: string): Promise<any> {
        return this.file.post("/upload/" + name, data);
    }
    GetItems(): Promise<any> {
        return this.file.get("/items");
    }
    DeleteThing(data: any): Promise<any> {
        return this.file.delete("/items", { data: data });
    }
    GetCharacters(): Promise<any> {
        return this.file.get("/characters");
    }
}

export default new ProfileService();
