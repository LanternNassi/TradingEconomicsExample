import { create } from "zustand";
import api from "@/Utils/Request";
import { AxiosResponse } from "axios";
import { addMonths, addHours } from "date-fns";

export interface IEvent {
  CalendarId: string;
  Date: string;
  Category: string;
  Currency: string;
  Importance: string;
  Event : string;

  title: string;
  start: Date;
  end: Date;
}

interface IEventStore {
  isLoading: boolean;
  events: IEvent[] | null;
  eventNotifications: IEvent[] | null;

  fetchEvents: (user_id: string) => Promise<void>;
  EventNotification: (event: IEvent , user : string) => Promise<void>;
  FetchEventNotifications: (email: string) => Promise<void>;
  RemoveNotification: (CalendarId: string , onsuccess : () => void) => Promise<void>;
}

export const useEventsStore = create<IEventStore>((set) => ({
  events: null,
  isLoading: false,
  eventNotifications: null,

  fetchEvents: async (user_id: string) => {
    set({ isLoading: true });

    api
      .get("/Events/", { params: { user_id } })
      .then((response: AxiosResponse) => {
        if (response.status == 200) {
          const newEvents: IEvent[] = response.data.map((event: IEvent) => {
            return {
              ...event,
              start: addMonths(new Date(event.Date), 6),
              end: addHours(addMonths(new Date(event.Date), 6), 0),
              title: event.Category,
            };
          });

          console.log(newEvents);

          set({ events: newEvents, isLoading: false });
        }
      });
  },

    EventNotification: async (event: IEvent , email : string) => {
        set({ isLoading: true });

        const sevent = {
            CalendarId: event.CalendarId,
            Date: event.Date,
            Event: event.title,
            Importance: event.Importance,
            Email : email
        }

        api.post("/Events/Notifications", { ...sevent }).then((response: AxiosResponse) => {
            if (response.status == 200) {
                set({ isLoading: false });
                alert("Notification set successfully");
            }
        });

    },

    FetchEventNotifications: async (email: string) => {
        set({ isLoading: true });

        api.get(`/Events/Notifications/${email}`).then((response: AxiosResponse) => {
            if (response.status == 200) {
                set({ isLoading: false, eventNotifications: response.data });
            }
        });
    },

    RemoveNotification: async (CalendarId: string , onsuccess: ()=>void) => {
        set({ isLoading: true });

        api.delete(`/Events/Notifications/${CalendarId}`).then((response: AxiosResponse) => {
            if (response.status == 200) {
                set({ isLoading: false });
                alert("Notification removed successfully");
                onsuccess();
            }
        });

    },
}));
