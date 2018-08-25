import * as React from "react";
import {EventList} from "./event_list";
import {Loader} from "./loader";
import {Event} from "../model/event";
import {RouteComponentProps} from 'react-router-dom'

export interface EventListContainerProps extends RouteComponentProps<any> {
    eventServiceURL: string;
}

export interface EventListContainerState {
    loading: boolean;
    events: Event[];
}

export class EventListContainer extends React.Component<EventListContainerProps, EventListContainerState> {
    constructor(p: EventListContainerProps) {
        super(p);
         if (typeof p.location.state != 'undefined'){
             localStorage.setItem('userid',p.location.state.USERID); //use this for refresh case
             localStorage.setItem('username',p.location.state.first);
             console.log('setting:', p.location.state.USERID )

         }
        // console.log('name=',p.location.state.first);

        this.state = {
            loading: true,
            events: []
        };

  //      console.log(this.props)
        fetch(p.eventServiceURL + "/events", {method: "GET"})
            .then<Event[]>(response => response.json())
            .then(events => {
                this.setState({
                    loading: false,
                    events: events
                })
            })
    }

    private handleEventBooked(e: Event) {
        console.log("booking event...");
    }

    render() {
        
        return <Loader loading={this.state.loading} message="Loading events...">
            <EventList userID= {localStorage.getItem('userid')} name={localStorage.getItem('username')} events={this.state.events} onEventBooked={e => this.handleEventBooked(e)}/>
        </Loader>
    }
}