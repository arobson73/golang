import * as React from "react";
import {EventBookingForm} from "./event_booking_form";
import {Event} from "../model/event";
import {Link} from 'react-router-dom';

export interface EventBookingFormContainerProps {
    userID:string;
    eventID: string;
    eventServiceURL: string;
    bookingServiceURL: string;
}

export interface EventBookingFormState {
    state: "loading"|"ready"|"saving"|"done"|"error";
    event?: Event;
}

export class EventBookingFormContainer extends React.Component<EventBookingFormContainerProps, EventBookingFormState> {
    constructor(p: EventBookingFormContainerProps) {
        super(p);

	this.handleSubmit = this.handleSubmit.bind(this);
        this.state = {
            state: "loading"
        };

        console.log('EventBookingContainer booking url=',this.props.bookingServiceURL);

        fetch(p.eventServiceURL + "/events/" + p.eventID)
            .then<Event>(resp => resp.json())
            .then(event => {
                this.setState({
                    state: "ready",
                    event: event
                });
            })
    }

    render() {
        if (this.state.state === "loading") {
            console.log("state:loading")
            return <div>Loading...</div>;
        }

        if (!this.state.event) {
            console.log('state:unknown')
            return <div>Unknown error</div>;
        }

        if (this.state.state === "done") {
            return <div>
            <nav className="navbar navbar-expand-md navbar-dark bg-dark">
                <div className="navbar-text"> <strong>Booking successfully completed</strong></div>
                <div className="nav navbar-nav ml-auto">
                    <Link to="/list">Back To Events</Link>
                 </div>
            </nav>
            </div>

//            return <div className="alert alert-success">Booking successfully completed!</div>
        }

        return <EventBookingForm event={this.state.event} onSubmit={amount => this.handleSubmit(amount)}/>
    }

    handleSubmit(seats: number) {
        const url = this.props.bookingServiceURL + "/bookings/" + this.props.eventID + "/" + this.props.userID ;
       // const url = this.props.bookingServiceURL + "/" + this.props.eventID + "/" + this.props.userID ;
        const payload = {Seats: seats,Name:this.state.event.Name};
        console.log("state :saving")
        this.setState({
            event: this.state.event,
            state: "saving"
        });
        console.log('EventBookingFormContainer url=', url);
        console.log('payload=',JSON.stringify(payload));
        fetch(url, {method: "POST", body: JSON.stringify(payload)})
            .then(response => {
                this.setState({
                    event: this.state.event,
                    state: response.ok ? "done" : "error"
                });
            })
    }
}