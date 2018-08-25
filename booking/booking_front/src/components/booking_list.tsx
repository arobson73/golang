import * as React from "react";
import { EventListItem } from "./event_list_item";
import { BookingListItem } from "./booking_list_item";
import { Event } from "../model/event";
import { Link } from 'react-router-dom';
import {Booking} from "../model/event";
export interface BookingListProps {
    userID: string;
    bookings: Booking[];
}

//<td><Link to={`/bookings/${this.props.event.ID}/${this.props.userID}/bookings`}
export class BookingList extends React.Component<BookingListProps, {}> {
    public render() {
       // console.log('EventList userid=', this.props.userID)
        let i = 0;
        const items = this.props.bookings.map(booking =>
            <BookingListItem key={i++} booking={booking} />
        );

        return <div>
            <nav className="navbar navbar-expand-md navbar-dark bg-dark">
                <div className="navbar-text"> <strong>Your Bookings</strong></div>
                <div className="nav navbar-nav ml-auto">
                    <Link to="/list">Back To Events</Link>
                 </div>
            </nav>
            <table className="table">
                <thead>
                    <tr>
                        <th>Date</th>
                        <th>Event</th>
                        <th>Seats</th>
                    </tr>
                </thead>
                <tbody>
                    {items}
                </tbody>
            </table>
        </div>
    }
}