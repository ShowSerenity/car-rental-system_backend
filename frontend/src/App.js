import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Home from './components/Home';
import Login from './components/Login';
import Register from './components/Register';
import CarList from './components/CarList';
import RentHistory from './components/RentHistory';

function App() {
    return (
        <Router>
            <div className="App">
                <Switch>
                    <Route path="/" exact component={Home} />
                    <Route path="/login" component={Login} />
                    <Route path="/register" component={Register} />
                    <Route path="/cars" component={CarList} />
                    <Route path="/rent-history" component={RentHistory} />
                </Switch>
            </div>
        </Router>
    );
}

export default App;
