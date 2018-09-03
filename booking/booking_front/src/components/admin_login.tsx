import * as React from "react";
import { Link } from 'react-router-dom';

import { Button, FormGroup, FormControl, ControlLabel, FormControlProps } from "react-bootstrap";

export interface AdminLoginState {
	email: string;
	password: string;

}
export interface AdminLoginProps {
	history: any;
}

export class AdminLogin extends React.Component<AdminLoginProps, AdminLoginState> {
	constructor(props: AdminLoginProps) {
		super(props);
		this.handleSubmit = this.handleSubmit.bind(this);
		this.handleChange = this.handleChange.bind(this);

		this.state = {
			email: "",
			password: "",
		};
	}

	validateForm() {
		if (this.state.email.length > 0 && this.state.password.length > 0) {
			return true;
		}
		return false;

	}

	handleChange(e: React.FormEvent<FormControlProps>) {
		let name = e.currentTarget.name;
		let val = e.currentTarget.value;
		this.setState({ ...this.state, [name]: val })
	}

	handleSubmit(e: React.FormEvent<HTMLFormElement>) {
		e.preventDefault();

		fetch("http://localhost:8181" + "/admin/verifyAdminUser/" + this.state.email + "/" + this.state.password, { method: "GET" })
			.then(r =>  {
				if (!r.ok)
				{
					this.props.history.push('/adminError')
				}
				else {
					console.log('ok from server');
					r.json()
					.then(data => ({status:r.status,body:data}))
					.then(obj => {
						const loc1 = {
							pathname:'/admin/event',
							state:{USERID:obj.body.ID,first:obj.body.First}
						}
						if(obj.status === 200) {
							//console.log('pushing to event')
							console.log('admin log userid=',loc1.state.USERID)
							this.props.history.push(loc1);
						}		
						})
					}
			})
			
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
					<form className="formCustom" onSubmit={e => this.handleSubmit(e)}>
						<h3> Admin Events Board</h3>
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
						<Button className="customButton"
							block
							bsSize="small"
							disabled={!this.validateForm()}
							type="submit"
						>
							Login
						</Button>
					</form>

				</div>
			</div>
		);
	}
}


