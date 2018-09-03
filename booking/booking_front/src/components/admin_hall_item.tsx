import * as React from "react";
import {Booking} from "../model/event";
import { Button, FormGroup, FormControl, ControlLabel, FormControlProps, HelpBlock, FormControlFeedback } from "react-bootstrap";
import { createBrowserHistory } from "history";

export interface AdminHallItemProps {
    index:number;
  //  handler:(e : React.FormEvent<FormControlProps>) => any
    //valid:(c:number) => boolean;
    cb:(e:React.FormEvent<FormControlProps>,i: number) => any;

}
export interface AdminHallItemState {
    name:string;
    location:string;
    capacity:number;
}

export class AdminHallItem extends React.Component<AdminHallItemProps, AdminHallItemState> {
    constructor(props: AdminHallItemProps){
        super(props);
        this.handleChange = this.handleChange.bind(this);
        this.handleValid = this.handleValid.bind(this);
        
        this.state={
            name:'',
            location:'',
            capacity:0,
        };
    }
    handleChange(e: React.FormEvent<FormControlProps>) {
        /*
	    let name = e.currentTarget.name;
		let val = e.currentTarget.value;
		let newState = Object.assign({},this.state);
		newState[name] = val;
        this.setState(newState);
        */
        this.props.cb(e,this.props.index)

    }
    handleValid(v:number) {
        v > 0 ? true : false;
    }
    render() {
        return (
            <div>
				<FormGroup controlId="nameHall" bsSize="small">
							<ControlLabel>Hall name</ControlLabel>
							<FormControl
								value={this.state.name}
								onChange={e => this.handleChange(e)}
								type="string"
								name="name"
							/>
						</FormGroup>
						<FormGroup controlId="locHall" bsSize="small">
							<ControlLabel>Hall Location</ControlLabel>
							<FormControl
								value={this.state.location}
								onChange={e => this.handleChange(e)}
								type="string"
								name="location"
							/>
						</FormGroup>
						<FormGroup controlId="nameHall" bsSize="small">
                            <ControlLabel> Capacity </ControlLabel>
							<FormControl
								value={this.state.capacity}
								onChange={e => this.handleChange(e)}
								type="string"
								name="capacity"
								placeholder='1'
							/>
							<FormControl.Feedback />
							<HelpBlock> {this.handleValid(this.state.capacity) ? "" : "must have greater than 0 seats"}</HelpBlock>

						</FormGroup>
        </div>
        )     
    }
}