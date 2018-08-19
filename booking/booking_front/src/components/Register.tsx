//const React = require('react')
//import * as ReactDOM from 'react-dom';
/*
import * as React from "react";

export class Register extends React.Component<{},{}> {
	render() {
		return <div><h1> Register </h1></div>
	}
}
*/
import * as React from "react";
import { Link } from 'react-router-dom';
import {ObjectID} from 'bson';

import { Button, FormGroup, FormControl, ControlLabel, FormControlProps } from "react-bootstrap";

export interface RegisterState {
	id:ObjectID;
	firstname:string;
	secondname:string;
	email:string;
	password:string;
	age:number;//replace with dob !!!TODO
	
}
export interface RegisterProps {
	history: any;
}

export class Register extends React.Component<RegisterProps, RegisterState> {
	constructor(props: RegisterProps) {
		super(props);
		this.handleSubmit = this.handleSubmit.bind(this);
		this.handleChange = this.handleChange.bind(this);

		this.state = {
			id: new ObjectID(),
			firstname:"",
			secondname:"",
			email: "",
			password: "",
			age:0
		};
	}

	validateForm() {
		//quite basic validation - make better TODO
		let f = /^[a-zA-Z]+$/.test(this.state.firstname);
		let s = /^[a-zA-Z]+$/.test(this.state.secondname);

		if (this.state.email.length > 0 && this.state.password.length > 0 && f && s) {
			return true;
		}
		return false;

	}

	handleChange(e: React.FormEvent<FormControlProps>) {
		let name = e.currentTarget.name;
		let val = e.currentTarget.value;

			this.setState({...this.state,[name]:val})
/*
		if (name === "firstname" || name === "secondname")
		{
			let val = e.currentTarget.value.toString().toLowerCase();	
			this.setState({...this.state,[name]:val})
		}
		else if(name === "age") {
			this.setState({...this.state,[name]:parseInt(e.currentTarget.value)})
		}
		else{

			this.setState({...this.state,[name]:e.currentTarget.value})
		}*/
	}

	handleSubmit(e: React.FormEvent<HTMLFormElement>) {
		e.preventDefault();
//		console.log(this.state.email);
		//console.log(this.state.password);
		const payload = JSON.stringify({
			ID:this.state.id.toString(),
			First:this.state.firstname,
			Last:this.state.secondname,
			Email:this.state.email,
			Password:this.state.password,
			Age:parseInt(this.state.age.toString()),
			Bookings:[]

		});
		//console.log(payload.ID);
	//	console.log(payload.First);
		//console.log(payload.Last);
		//console.log(payload.Email);
		//console.log(payload.Password);
		//console.log(payload.Age);
		//console.log(payload.Bookings);
		//console.log(payload);	
	
	

		console.log(payload);
		

		fetch("http://localhost:8181/users", {method: "POST", body: payload} )
			.then(response => {
				console.log(response.json)
				if (response.ok) {
					this.props.history.push('/list');
					//	this.setState({...this.state,verified:true,msg:""})
				}
				else {
					this.props.history.push('/error');
				}
			}
			).catch(e => console.log('Error',e))


           
		/*

        fetch("http://localhost:8181" + "/users/findUserEmailPass/" + this.state.email + "/" + this.state.password, {method: "GET"})
			.then(response => {
				if (response.ok) {
					this.props.history.push('/list');
				//	this.setState({...this.state,verified:true,msg:""})
				}
				else{
					this.props.history.push('/error');
				}
			}
			)
	*/			

	}

	render() {
		return (
			<div>
				<style type="text/css">{`
    					.formCustom {
						width:20%;
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
					<Link className="customLink" to="/register">
						Register
					</Link>

				</div>
				<div>
					<form className="formCustom" onSubmit={e => this.handleSubmit(e)}>
						<h3> Events Board Add New User</h3>
						<FormGroup controlId="firstname" bsSize="small">
							<ControlLabel>FirstName</ControlLabel>
							<FormControl
								autoFocus
								type="string"
								value={this.state.firstname}
								onChange={e => this.handleChange(e)}
								name="firstname"
							/>
						</FormGroup>
						<FormGroup controlId="secondname" bsSize="small">
							<ControlLabel>Surname</ControlLabel>
							<FormControl
								autoFocus
								type="string"
								value={this.state.secondname}
								onChange={e => this.handleChange(e)}
								name="secondname"
							/>
						</FormGroup>

						<FormGroup controlId="email" bsSize="small">
							<ControlLabel>Email</ControlLabel>
							<FormControl
								autoFocus
								type="string"
								value={this.state.email}
								onChange={e => this.handleChange(e)}
								name="email"
							/>
						</FormGroup>
						<FormGroup controlId="password" bsSize="small">
							<ControlLabel>Password</ControlLabel>
							<FormControl
								value={this.state.password}
								onChange={e => this.handleChange(e)}
								type="string"
								name="password"
							/>
						</FormGroup>
						<FormGroup controlId="age" bsSize="small">
							<ControlLabel>Age</ControlLabel>
							<FormControl
								value={this.state.age}
								onChange={e => this.handleChange(e)}
								type="number"
								name="age"
							/>
						</FormGroup>

						<Button className="customButton"
							block
							bsSize="small"
							disabled={!this.validateForm()}
							type="submit"
						>
							Register
						</Button>
					</form>
					
				</div>
			</div>
		);
	}
}


