//const React = require('react')
//import * as ReactDOM from 'react-dom';
import * as React from "react";
import { Link } from 'react-router-dom';


export class ErrorLogin extends React.Component<{}, {}> {
	render() {
		return <div>
			<div><h1> Login Error </h1></div>
			<div>
				<Link className="customLink" to="/">
					Retry
					</Link>
			</div>
		</div>
	}
}
//module.exports = Register