import axios, { AxiosInstance } from "axios";

export default function CreateAxios(url:string, headers={ "Content-type": "application/json"}) : AxiosInstance {
    const apiClient: AxiosInstance = axios.create({
        baseURL: "https://"+import.meta.env.VITE_ADDRES+url,
        headers
    });
    return apiClient
}
