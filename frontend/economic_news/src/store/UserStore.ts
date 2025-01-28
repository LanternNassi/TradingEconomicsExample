import {create} from "zustand"
import api from "@/Utils/Request";
import { AxiosResponse } from "axios";



export interface IUser {
    id : string;
    username : string ;
    password : string;
    email : string;
}

interface UserDto {
    username : string;
    password : string;
    email : string;
}

export interface LoginUser {
    email : string ;
    password : string;
}

interface IUserStore{
    isLoading : boolean;
    user : IUser | null;
    Login : (login : LoginUser) => Promise<void>;
    Logout : () => Promise<void>;
    Register : (user : UserDto) => Promise<void>;
}

export const useUserStore = create<IUserStore>((set) => ({
    user : null,
    isLoading : false,

    Login : async(Login : LoginUser) => {
        set({isLoading : true})

        api.post("/Users/login" , Login)
            .then((response : AxiosResponse) => {
                if (response.status == 202){
                    set({user : response.data , isLoading : false})
                    alert("Login successful")
                }
            })

    },
    Register : async(user : UserDto) => {
        set({isLoading : true})

        api.post("/Users/" , user)
            .then((response : AxiosResponse) => {
                if (response.status == 200){
                    set({user : response.data , isLoading : false})
                    alert("Registation successful")
                }
            })
    },
    Logout : async() => {
        set({user : null})
    }

}))



