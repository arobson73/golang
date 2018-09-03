import * as React from "react";
import { EventListItem } from "./event_list_item";
import { BookingListItem } from "./booking_list_item";
import { Event } from "../model/event";
import { Link } from 'react-router-dom';
import {Booking} from "../model/event";
import { AdminHallItem,AdminHallItemState } from "./admin_hall_item";
import { FormControlProps } from "react-bootstrap";

export interface AdminHallListProps {
   // userID: string;
    //bookings: Booking[];
	//
//    indexs: number[];
    max:number;
    //cb:(ah:AdminHallItemState,i: number) => any;
    cb:(e:React.FormEvent<FormControlProps>,i: number) => any;
}

//<td><Link to={`/bookings/${this.props.event.ID}/${this.props.userID}/bookings`}
export class AdminHallList extends React.Component<AdminHallListProps, {}> {
    public render() {
       // console.log('EventList userid=', this.props.userID)
		let items = [];
        for(let i=0;i < this.props.max; i++){
            items.push(<AdminHallItem key={i} index={i} cb={this.props.cb} />)
        }
         
		/*
        const items = this.props.indexs.map(ix =>
            <AdminHallItem index={ix} cb={this.props.cb} />
        );*/

        return (<div>
                    {items}
                </div>
        )
    }
}
