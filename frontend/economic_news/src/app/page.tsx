"use client";

import React , {useEffect , useState} from 'react';

import { useUserStore } from '@/store/UserStore';

import { useRouter } from 'next/navigation';

const LoginPage: React.FC = () => {

    const router = useRouter();

    const [email , setEmail] = useState('');
    const [password , setPassword] = useState('');

    const {user , Login} = useUserStore();

    const handleSubmit = (event: React.FormEvent) => {
        event.preventDefault();
        Login({email , password});
    };

    useEffect(() => {
        if (user) {
            router.push('/dashboard/events');
        }
    },[user])

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <div className="w-full max-w-md p-8 space-y-8 bg-white rounded-lg shadow-md">
            <h2 className="text-2xl font-bold text-center text-gray-900">Login</h2>
            <form className="mt-8 space-y-6">
                <div className="rounded-md shadow-sm space-y-12">
                <div>
                    <label htmlFor="email-address" className="sr-only">
                    Email address
                    </label>
                    <input
                    id="email-address"
                    name="email"
                    type="email"
                    onChange={(email) => setEmail(email.target.value)}
                    autoComplete="email"
                    required
                    className="w-full px-3 py-2 mt-1 border text-black border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                    placeholder="Email address"
                    />
                </div>
                <div>
                    <label htmlFor="password" className="sr-only">
                    Password
                    </label>
                    <input
                    id="password"
                    onChange={(password) => setPassword(password.target.value)}
                    name="password"
                    type="password"
                    autoComplete="current-password"
                    required
                    className="w-full px-3 py-2 mt-1 text-black border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                    placeholder="Password"
                    />
                </div>
                </div>
                <div>
                <button
                    onClick={handleSubmit}
                    type="submit"
                    className="relative flex justify-center w-full px-4 py-2 text-sm font-medium text-white bg-indigo-600 border border-transparent rounded-md group hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                >
                    Sign in
                </button>
                </div>
            </form>
            <p className="text-center text-sm text-gray-600">
                    Dont have an account?{' '}
                    <a onClick={() => {router.push("/register")}} className="font-medium text-indigo-600 hover:text-indigo-500 cursor-pointer">
                        Register here
                    </a>
                </p>
            </div>
        </div>
    );
};

export default LoginPage;