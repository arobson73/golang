import * as React from "react";
import {EventBookingForm} from "./event_booking_form";
import {Event} from "../model/event";

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
            return <div>Loading...</div>;
        }

        if (!this.state.event) {
            return <div>Unknown error</div>;
        }

        if (this.state.state === "done") {
            return <div className="alert alert-success">Booking successfully completed!</div>
        }

        return <EventBookingForm event={this.state.event} onSubmit={amount => this.handleSubmit(amount)}/>
    }

    handleSubmit(seats: number) {
        const url = this.props.bookingServiceURL + "/events/" + this.props.eventID + "/" + this.props.userID +  "/bookings";
        const payload = {Seats: seats};

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