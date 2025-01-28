/* eslint-disable react-hooks/rules-of-hooks */
/* eslint-disable @typescript-eslint/no-unused-vars */
"use client";

import { FC, useState, useEffect } from "react";
import { Calendar, dateFnsLocalizer, Event } from "react-big-calendar";
import withDragAndDrop, {
  withDragAndDropProps,
} from "react-big-calendar/lib/addons/dragAndDrop";
import {
  format,
  parse,
  startOfWeek,
  getDay,
  addHours,
  startOfHour,
  startOfDecade,
} from "date-fns";
import { enUS, ro } from "date-fns/locale";
import { useEventsStore, IEvent } from "@/store/EventsStore";
import { useUserStore } from "@/store/UserStore";
import { useRouter } from "next/navigation";

import "react-big-calendar/lib/addons/dragAndDrop/styles.css";
import "react-big-calendar/lib/css/react-big-calendar.css";

const page = () => {
  
  const { isLoading, events, fetchEvents , EventNotification } = useEventsStore();
  const { user } = useUserStore();

  const [showModal, setShowModal] = useState(false);
  const [selectedEvent, setSelectedEvent] = useState<IEvent | null>(null);

  const router = useRouter();

  if (!user) {
    router.push("/");
  }

  const handleSelectEvent = (event: Event) => {

    events?.forEach((sevent) => {
      if (sevent.title == event.title) {
        setSelectedEvent(sevent);
        setShowModal(true);
      }
    });
  };

  const handleClose = () => setShowModal(false);

  console.log("User" , user)


  const handleSetupNotification = () => {
    EventNotification(selectedEvent as IEvent , user?.email as string);
    setShowModal(false);
  };

  console.log(events);

  const onEventResize: withDragAndDropProps["onEventResize"] = (data) => {
    const { start, end } = data;
  };

  const onEventDrop: withDragAndDropProps["onEventDrop"] = (data) => {
    console.log(data);
  };




  useEffect(() => {
    fetchEvents(user?.id as string);
  }, []);

  if (isLoading) {
    return (
      <div>
        <h1>Loading events...</h1>
      </div>
    );
  }

  return (
    <>
      <DnDCalendar
        defaultView="agenda"
        onSelectEvent={handleSelectEvent}
        events={events ? events : []}
        localizer={localizer}
        onEventDrop={onEventDrop}
        onEventResize={onEventResize}
        resizable
        style={{ height: "80vh" }}
      />
      {showModal && (
        <div className="fixed inset-0 flex items-center justify-center z-50">
          <div className="bg-black opacity-50 absolute inset-0"></div>
          <div className="bg-white rounded-lg shadow-lg z-10 p-6 w-96">
            <h2 className="text-xl font-semibold mb-4">Setup Notification</h2>
            <p className="mb-4">
              Would you like to setup a notification for the event{" "}
              {selectedEvent?.title}?
            </p>
            <div className="flex justify-end">
              <button
                className="bg-gray-500 text-white px-4 py-2 rounded mr-2"
                onClick={handleClose}
              >
                Close
              </button>
              <button
                className="bg-blue-500 text-white px-4 py-2 rounded"
                onClick={handleSetupNotification}
              >
                Setup Notification
              </button>
            </div>
          </div>
        </div>
      )}
    </>
  );
};

const locales = {
  "en-US": enUS,
};
const endOfHour = (date: Date): Date => addHours(startOfHour(date), 1);
const now = new Date();
const start = endOfHour(now);
const end = addHours(start, 2);
const localizer = dateFnsLocalizer({
  format,
  parse,
  startOfWeek,
  getDay,
  locales,
});

const DnDCalendar = withDragAndDrop(Calendar);

export default page;
