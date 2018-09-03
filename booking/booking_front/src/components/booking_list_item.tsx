import * as React from "react";
import {Booking} from "../model/event";

export interface BookingListItemProps {
    booking: Booking;
}

export class BookingListItem extends React.Component<BookingListItemProps, {}> {
    render() {
        console.log('BookingListItem userid=',this.props.booking.EventID);
        const start = new Date(this.props.booking.Date * 1000);

        return <tr>
            <td>{this.props.booking.Date.toString()}</td>
            <td>{this.props.booking.Name}</td>
            <td>{this.props.booking.Seats}</td>

        </tr>
    }
}