"use client";

import { useRouter } from 'next/navigation';

import React , {useState} from 'react';


const Navbar = () => {

    const router = useRouter();

    const [activeRoute, setActiveRoute] = useState("Dashboard");

    return (
        <div className="flex flex-col">
            <header className="bg-blue-500 text-white">
                <div className="container mx-auto flex justify-between items-center p-4">
                    <h1 className="text-xl font-bold">{activeRoute}</h1>
                    <div>
                        <button onClick={()=>{
                            router.push("/dashboard")
                            setActiveRoute("Dashboard")
                        }} className="mx-6">Dashboard</button>
                        <button onClick={()=>{
                            router.push("/profile")
                            setActiveRoute("Profile")
                        }} className="mx-6">UserProfile</button>
                        <button onClick={()=>{
                            router.push("/events")
                            setActiveRoute("Events")
                        }} className="mx-6">Events</button>
                        <button onClick={()=>{
                            router.push("/customs")
                            setActiveRoute("Customs")
                        }} className="mx-6">Customs</button>
                        <button  className="mx-6">Logout</button>

                    </div>
                </div>
            </header>
            {/* <main className="flex-grow container mx-auto p-4">
                {children}
            </main> */}
        </div>
    )
}

export default Navbar;