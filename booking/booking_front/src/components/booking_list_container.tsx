import * as React from "react";
import {EventList} from "./event_list";
import {Loader} from "./loader";
import {Booking} from "../model/event";
import {RouteComponentProps} from 'react-router-dom'
import {BookingList} from './booking_list' 

export interface BookingListContainerProps extends RouteComponentProps<any> {
    bookingServiceURL: string;
    userid:string;
}

export interface BookingListContainerState {
    loading: boolean;
    bookings: Booking[];
}

export class BookingListContainer extends React.Component<BookingListContainerProps, BookingListContainerState> {
    constructor(p: BookingListContainerProps) {
        super(p);
        /*
         if (typeof p.location.state != 'undefined'){
             localStorage.setItem('userid',p.location.state.USERID); //use this for refresh case
             localStorage.setItem('username',p.location.state.first);
             console.log('setting:', p.location.state.USERID )

         }*/
        // console.log('name=',p.location.state.first);

        this.state = {
            loading: true,
            bookings: []
        };
        console.log('userid=',p.userid)

  //      console.log(this.props)
        fetch(p.bookingServiceURL + "/bookings/" + p.userid, {method: "GET"})
            .then<Booking[]>(response => response.json())
            .then(bookings => {
                this.setState({
                    loading: false,
                    bookings: bookings
                })
            })
    }

    private handleEventBooked(e: Event) {
        console.log("getting bookings...");
    }

    render() {
        
        return <Loader loading={this.state.loading} message="Loading bookings...">
        <BookingList userID={this.props.userid} bookings={this.state.bookings}/>
        {/*    <EventList userID= {localStorage.getItem('userid')} name={localStorage.getItem('username')} events={this.state.events} onEventBooked={e => this.handleEventBooked(e)}/> */}
        </Loader>
    }
}