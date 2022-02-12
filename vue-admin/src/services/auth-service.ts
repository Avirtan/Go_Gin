import CreateAxios from "./HttpCommon"
import login from "@/types/login"
import  { AxiosInstance } from "axios";

class Authentication {
    client:AxiosInstance

    constructor() {
        this.client = CreateAxios("/api/v1/auth")

    }
    login(data:login):Promise<any>{
        return this.client.post("/login", data);
    }
}

export default new Authentication();