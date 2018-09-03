import * as React from "react";
import {EventList} from "./event_list";
import {Loader} from "./loader";
import {Booking} from "../model/event";
import {RouteComponentProps} from 'react-router-dom'
import {BookingList} from './booking_list' 
import { Button, FormGroup, FormControl, ControlLabel, FormControlProps, HelpBlock, FormControlFeedback } from "react-bootstrap";
import { AdminHallList } from "./admin_hall_list";
import {AdminHallItemState} from "./admin_hall_item"


//export interface AdminHallContainerProps extends RouteComponentProps<any> {
export interface AdminHallContainerProps{
//    cb:(ah:AdminHallItemState,i: number) => void;
   cb:(e:React.FormEvent<FormControlProps>,i: number) => any;

}

export interface AdminHallContainerState {
    total:number;
}

export class AdminHallContainer extends React.Component<AdminHallContainerProps, AdminHallContainerState> {
    constructor(p: AdminHallContainerProps) {
        super(p);
        this.handleChange = this.handleChange.bind(this);
        this.validate = this.validate.bind(this);
        this.state ={
            total:0,
        }
    };
    
    handleChange(e: React.FormEvent<FormControlProps>) {
		let name = e.currentTarget.name;
		let val = e.currentTarget.value;
		this.setState({...this.state,[name]:val});
	}
	validate(d):boolean {
		if (d > 0) return true;
		else return false;

	}

    render() {
        return (
            <div>
				<FormGroup controlId="totalhalls" bsSize="small">
							<ControlLabel>Duration in mins</ControlLabel>
							<FormControl
								autoFocus
								type="number"
								value={this.state.total}
								onChange={e => this.handleChange(e)}
								name="total"
								placeholder="1"
							/>
							<FormControl.Feedback />
							<HelpBlock> {this.validate(this.state.total) ? "" : "Duration must be number and greater than 0"}</HelpBlock>
						</FormGroup>
                        <AdminHallList max={this.state.total} cb={this.props.cb} />
                        </div>

        )
        
    }
}