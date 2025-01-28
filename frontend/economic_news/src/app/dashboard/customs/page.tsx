"use client";

import React, { useEffect } from 'react';

import { useEventsStore } from '@/store/EventsStore';

import { useUserStore } from '@/store/UserStore';


const Dashboard: React.FC = () => {

    const { FetchEventNotifications , RemoveNotification , eventNotifications } = useEventsStore();

    const { user } = useUserStore();

    useEffect(() => {

        FetchEventNotifications(user?.email as string);
    }, []);

    const handleCancel = async (calendarId: string) => {
        await RemoveNotification(calendarId , () => FetchEventNotifications(user?.email as string));
    };

    return (
        <div className="p-4">
            <h1 className="text-2xl font-bold mb-4">Scheduled Notifications</h1>
            <ul className="space-y-4">
                {eventNotifications?.map(notification => (
                    <li key={notification.CalendarId} className="p-4 border rounded shadow">
                        <p><strong>Calendar ID:</strong> {notification.CalendarId}</p>
                        <p><strong>Importance:</strong> {notification.Importance}</p>
                        <p><strong>Event:</strong> {notification.Event}</p>
                        <p><strong>Date:</strong> {notification.Date}</p>
                        <button 
                            onClick={() => handleCancel(notification.CalendarId)} 
                            className="mt-2 px-4 py-2 bg-red-500 text-white rounded hover:bg-red-700"
                        >
                            Cancel Notifiation
                        </button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default Dashboard;