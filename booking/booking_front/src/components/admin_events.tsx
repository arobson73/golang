import * as React from "react";
import { Link, RouteComponentProps } from 'react-router-dom';
import { ObjectID } from 'bson';

import { Button, FormGroup, FormControl, ControlLabel, FormControlProps, HelpBlock, FormControlFeedback } from "react-bootstrap";
import { AdminHallContainer } from "./admin_hall_container";
import {AdminHallItemState} from "./admin_hall_item"
import {Hall} from "../model/event";

export interface AdminEventState {
	id: ObjectID;
	name: string;
	duration: number;
	startDate: number;
	startString: string;
	startTime: string;
	endDate: number;
	endString: string;
	error: string;
	
	Location:{
		id:ObjectID;
		name:string;
		address:string;
		country:string;
		openTime:string;
		closeTime:string;
		Halls:Hall[],
		/*
		Halls:{
			name:string;
			location:string;
			capacity:number;
		}*/
	}

}
export interface AdminEventProps extends RouteComponentProps<any> {
//	history: any;
}

export class AdminEvent extends React.Component<AdminEventProps, AdminEventState> {
	constructor(props: AdminEventProps) {
		super(props);
		console.log('props',props)
		this.handleSubmit = this.handleSubmit.bind(this);
		this.handleChange = this.handleChange.bind(this);
		this.handleLocChange = this.handleLocChange.bind(this);
		console.log('in constructor')
		this.state = {
			id: new ObjectID(),
			name: "wimbledon",
			duration: 1,
			startDate: 0,
			startString: "01/01/2018",
			startTime: "00:00",
			endDate: 0,
			endString: "01/01/2018",
			error: "",
			
			Location:{
				id: new ObjectID(),
				name:'palace',
				address:'downing st',
				country:'uk',
				openTime:'00:00',
				closeTime:'00:00',
				Halls:[],
/*i
				Halls:{
					name:'',
					location:'',
					capacity:1,
				}*/
			}	
		};
	}

	validateForm() {
		let ln = this.state.name.length > 0;
		if(!ln) return false;
		console.log('h1');
		let dur = this.validateDuration(this.state.duration);
		if(!dur) return false;
		console.log('h2');
		let dat = this.validateDate(this.state.startString);
		if(!dat) return false;
		console.log('h3');
		let dat2 = this.validateDate(this.state.endString);
		if(!dat2) return false;
		console.log('h4');
		if(!(this.state.Location.name.length > 1)) return false;
		console.log('h5');
		if(!(this.state.Location.address.length > 1)) return false;
		console.log('h6');
		if(!(this.state.Location.country.length > 1)) return false;
		console.log('h7');
		if(!this.validateTime(this.state.Location.openTime)) return false;
		console.log('h8');
		if(!this.validateTime(this.state.Location.openTime)) return false;
		console.log('h9');
		if(!this.validateTime(this.state.Location.closeTime)) return false;
		console.log('h10');
	//	if(!this.validateDuration(this.state.Location.Halls.capacity)) return false;
		console.log('h11');


		return true;

	}
	validateDuration(d):boolean {
		if (d > 0) return true;
		else return false;

	}
	validateDate(d:string):boolean {
		let f = /^(3[01]|[12][0-9]|0[1-9])\/(1[0-2]|0[1-9])\/[0-9]{4}$/.test(d);
		return f;
	}
	validateTime(time:string):boolean {
		let f = /^(2[0-3]|[01]?[0-9]):([0-5]?[0-9])$/.test(time);
		return f;
	}

	handleChange(e: React.FormEvent<FormControlProps>) {
		let name = e.currentTarget.name;
		let val = e.currentTarget.value;
		this.setState({...this.state,[name]:val});
	}
	handleLocChange(e: React.FormEvent<FormControlProps>) {
		let name = e.currentTarget.name;
		let val = e.currentTarget.value;
		this.setState(state => ({Location:{...state.Location,[name]:val}}));	

	}
	handleHallChange(e: React.FormEvent<FormControlProps>,idx:number) {
		let name = e.currentTarget.name;
		let val = e.currentTarget.value;
		console.log('name=',name,',val=',val,',idx=',idx)
		/*
		let newState = Object.assign({},this.state);
		newState.Location.Halls[name] = val;
		this.setState(newState);*/
	}

	handleSubmit(e: React.FormEvent<HTMLFormElement>) {
		e.preventDefault();
		function toUnixDate(dateStr){
			let parts = dateStr.split("/");
			let date = new Date(parts[2],parts[1]-1,parts[0])
			return date.getTime() /1000;
		}
		/*
		console.log('id',this.state.id);
		console.log('name',this.state.name);
		console.log('duration',this.state.duration);
		console.log('Startdate',this.state.startDate);
		console.log('Enddate',this.state.endString);
		console.log('Location.id',this.state.Location.id);
		console.log('Location.name',this.state.Location.name);
		console.log('Location.',this.state.Location.name);
*/



		
				const payload = JSON.stringify({
					ID:this.state.id,
					Name:this.state.name,
					Duration:this.state.duration,
					StartDate:toUnixDate(this.state.startString),
					EndDate : toUnixDate(this.state.endString),
					Location:{
						ID:this.state.Location.id,
						Name:this.state.Location.name,
						Address:this.state.Location.address,
						Country:this.state.Location.country,
						OpenTime:this.state.Location.openTime,
						CloseTime:this.state.Location.closeTime,
						/*
						Halls:[{
							Name:this.state.Location.Halls.name,
							Location:this.state.Location.Halls.location,
							Capacity:this.state.Location.Halls.capacity,
						},]*/
						Halls:this.state.Location.Halls,


					}

				});
		
		
				console.log(payload);
		
				fetch("http://localhost:8181/admin/addEventForUser/" +this.props.location.state.USERID , {method: "POST", body: payload} )
					.then(response => {
						//console.log(response.json)
						if (response.ok) {
							//this.props.history.push('/admin/eventsCreated');
							console.log('ok added event for user')
						}
						else {
							this.props.history.push('/error');
						}
					}
					).catch(e => console.log('Error',e))
		


	}

	render() {
		//console.log('in render');
		return (
			<div>
				<style type="text/css">{`
    					.formCustom {
						width:40%;
						position: relative;
						top:40%;
						left:40%;
						}
						.customButton {
        				background-color: purple;
						color: white;
						}
						.customLink {
							position:absolute;
							top:0%;
							left:80%;
						}
    				`}</style>

				<div>
					<Link className="customLink" to="/eventsCreated">
						Show created events
					</Link>

				</div>
				<div>
					<form className="formCustom" onSubmit={e => this.handleSubmit(e)}>
						<h3> Submit new event</h3>
						<FormGroup controlId="name" bsSize="small">
							<ControlLabel>Event Name</ControlLabel>
							<FormControl
								autoFocus
								type="string"
								value={this.state.name}
								onChange={e => this.handleChange(e)}
								name="name"
							/>

						</FormGroup>
						<FormGroup controlId="duration" bsSize="small">
							<ControlLabel>Duration in mins</ControlLabel>
							<FormControl
								autoFocus
								type="number"
								value={this.state.duration}
								onChange={e => this.handleChange(e)}
								name="duration"
								placeholder="1"
							/>
							<FormControl.Feedback />
							<HelpBlock> {this.validateDuration(this.state.duration) ? "" : "Duration must be number and greater than 0"}</HelpBlock>
						</FormGroup>

						<FormGroup controlId="start" bsSize="small">
							<ControlLabel>Start Date</ControlLabel>
							<FormControl
								autoFocus
								type="string"
								value={this.state.startString}
								onChange={e => this.handleChange(e)}
								name="startString"
								placeholder="20/01/2018"
							/>
							<FormControl.Feedback />
								<HelpBlock> {this.validateDate(this.state.startString) ? "" : "invalid format: use dd/mm/yyyy"}</HelpBlock>
						</FormGroup>
						<FormGroup controlId="time" bsSize="small">
							<ControlLabel>Start Time</ControlLabel>
							<FormControl
								value={this.state.startTime}
								onChange={e => this.handleChange(e)}
								type="string"
								name="startTime"
								placeholder="00:00"
							/>
							<FormControl.Feedback/>
							<HelpBlock> {this.validateTime(this.state.startTime) ? "" : "Time nas to be in 24 hr format"}</HelpBlock>
						</FormGroup>
						<FormGroup controlId="end" bsSize="small">
							<ControlLabel>End Date</ControlLabel>
							<FormControl
								value={this.state.endString}
								onChange={e => this.handleChange(e)}
								type="string"
								name="endString"
								placeholder="20/01/2018"
							/>
							<FormControl.Feedback />
							<HelpBlock>{this.validateDate(this.state.endString) ? "" : "invalid format: use dd/mm/yyyy"}</HelpBlock>
						</FormGroup>
						<FormGroup controlId="locaName" bsSize="small">
							<ControlLabel>Location</ControlLabel>
							<FormControl
								value={this.state.Location.name}
								onChange={e => this.handleLocChange(e)}
								type="string"
								name="name"
								placeholder="London"
							/>
						</FormGroup>
						<FormGroup controlId="locaAdd" bsSize="small">
							<ControlLabel>Address</ControlLabel>
							<FormControl
								value={this.state.Location.address}
								onChange={e => this.handleLocChange(e)}
								type="string"
								name="address"
							/>
						</FormGroup>
						<FormGroup controlId="locaC" bsSize="small">
							<ControlLabel>Country</ControlLabel>
							<FormControl
								value={this.state.Location.country}
								onChange={e => this.handleLocChange(e)}
								type="string"
								name="country"
							/>
						</FormGroup>
						<FormGroup controlId="locaO" bsSize="small">
							<ControlLabel>Opening Time</ControlLabel>
							<FormControl
								value={this.state.Location.openTime}
								onChange={e => this.handleLocChange(e)}
								type="string"
								name="openTime"
								placeholder='00:00'
							/>
						<FormControl.Feedback/>
							<HelpBlock> {this.validateTime(this.state.Location.openTime) ? "" : "Time nas to be in 24 hr format"}</HelpBlock>

						</FormGroup>
						<FormGroup controlId="locaCl" bsSize="small">
							<ControlLabel>Closing Time</ControlLabel>
							<FormControl
								value={this.state.Location.closeTime}
								onChange={e => this.handleLocChange(e)}
								type="string"
								name="closeTime"
								placeholder='00:00'
							/>
						<FormControl.Feedback/>
							<HelpBlock> {this.validateTime(this.state.Location.closeTime) ? "" : "Time nas to be in 24 hr format"}</HelpBlock>
						</FormGroup>
						{/*
						<FormGroup controlId="nameHall" bsSize="small">
							<ControlLabel>Hall name</ControlLabel>
							<FormControl
								value={this.state.Location.Halls.name}
								onChange={e => this.handleHallChange(e)}
								type="string"
								name="name"
							/>
						</FormGroup>
						<FormGroup controlId="locHall" bsSize="small">
							<ControlLabel>Hall Location</ControlLabel>
							<FormControl
								value={this.state.Location.Halls.location}
								onChange={e => this.handleHallChange(e)}
								type="string"
								name="location"
							/>
						</FormGroup>
						<FormGroup controlId="nameHall" bsSize="small">
							<ControlLabel>Capacity</ControlLabel>
							<FormControl
								value={this.state.Location.Halls.capacity}
								onChange={e => this.handleHallChange(e)}
								type="string"
								name="capacity"
								placeholder='1'
							/>
							<FormControl.Feedback />
							<HelpBlock> {this.validateDuration(this.state.Location.Halls.capacity) ? "" : "must have greater than 0 seats"}</HelpBlock>

						</FormGroup>
						*/}
						<AdminHallContainer cb={this.handleHallChange}/>


						<Button className="customButton"
							block
							bsSize="small"
							disabled={!this.validateForm()}
							type="submit"
						>
							Add Event
						</Button>
					</form>
					<div><span>{this.state.error}</span></div>

				</div>
			</div>
		);
	}
}


