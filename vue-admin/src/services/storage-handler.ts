

class StorageLocal{
    setData(key:string,value:string){
        localStorage.setItem(key,value)
    }

    getData(key:string):string{
        let token = localStorage.getItem(key)
        if (token == null){
            return ""
        }
        return token
    }
    delKey(key:string){
        localStorage.removeItem(key)
    }
}


export default new StorageLocal();