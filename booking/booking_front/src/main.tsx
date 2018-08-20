import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { Switch, Route,BrowserRouter} from 'react-router-dom';
import { Login } from './components/Login';
import { Register } from './components/Register';
import { EventBookingFormContainer } from './components/event_booking_form_container';
import { EventListContainer } from "./components/event_list_container";
import {ErrorLogin} from './components/error'

class App extends React.Component<{}, {}> {
  render() {
    const eventList = (props) => <EventListContainer eventServiceURL="http://localhost:8181" {...props}/>;
    const eventBooking = ({ match }: any) => <EventBookingFormContainer userID={match.params.userid} eventID={match.params.id} eventServiceURL="http://localhost:8181"
      bookingServiceURL="http://localhost:8182" />;

    return <BrowserRouter>
      <Switch>
        <Route exact={true} path="/" component={Login} />
        <Route path="/register" component={Register} />
        <Route path="/list" component={eventList} />
        <Route path="/error" component={ErrorLogin}/>
        <Route path="/events/:id/:userid/bookings" component={eventBooking}/>

      </Switch>
    </BrowserRouter>
  }
}

ReactDOM.render(<App />, document.getElementById('root'));
/*
 function eventList() {
   return <EventListContainer eventServiceURL="http://localhost:8181"/>;

 };
 function eventBooking(m:any) {
  return <EventBookingFormContainer eventID={m.params.id} eventServiceURL="http://localhost:8181" 
bookingServiceURL="http://localhost:8182" />;
 }
*/
/*
<HashRouter>
  <Switch>
    <Route exact={true} path="/" component={Login} />
    <Route path="/register" component={Register} />
    <Route path="/list" component={eventList}/>

    </Switch>
</HashRouter>
  , document.getElementById('root')
);
*/
/*
import * as React from "react";
import * as ReactDOM from "react-dom";
//import {HashRouter as Router, Route} from "react-router-dom";
import {Route,Switch,BrowserRouter} from 'react-router-dom';
//import {BrowserRouter as Router, Route} from "react-router-dom";
import {EventListContainer} from "./components/event_list_container";
import {Navigation} from "./components/navigation";
import {EventBookingFormContainer} from "./components/event_booking_form_container";
import {Register} from './components/Register';
import {Login} from './components/Login';

 export default class App extends React.Component {
    render() {
        return <h1> Hello A</h1>;
        
        const eventList = () => <EventListContainer eventServiceURL="http://localhost:8181"/>;
        const eventBooking = ({match}:any) => <EventBookingFormContainer eventID={match.params.id}
                                                                         eventServiceURL="http://localhost:8181"
                                                                         bookingServiceURL="http://localhost:8182"/>;
       // const login = () => <Login history></Login>
        const reg = () => <Register/>;

        return <BrowserRouter>
				<Switch>
					<Route exact={true} path="/" component={reg}/>
				</Switch>
        </BrowserRouter>
        
    }
}

ReactDOM.render(
    <App/>,
    document.getElementById('myevents-app')
);*/
