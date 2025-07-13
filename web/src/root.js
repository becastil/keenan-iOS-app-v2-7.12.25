import React from 'react';
import {Route, Switch} from 'fusion-plugin-react-router';
import {styled} from 'fusion-plugin-styletron-react';

import Layout from './components/Layout';
import Dashboard from './pages/Dashboard';
import Benefits from './pages/Benefits';
import Claims from './pages/Claims';
import Providers from './pages/Providers';
import MemberCard from './pages/MemberCard';
import Messages from './pages/Messages';
import Login from './pages/Login';
import PrivateRoute from './components/PrivateRoute';

const AppContainer = styled('div', {
  minHeight: '100vh',
  backgroundColor: '#f5f5f5',
  fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif',
});

const Root = () => {
  return (
    <AppContainer>
      <Switch>
        <Route path="/login" component={Login} />
        <PrivateRoute path="/benefits" component={Benefits} />
        <PrivateRoute path="/claims" component={Claims} />
        <PrivateRoute path="/providers" component={Providers} />
        <PrivateRoute path="/member-card" component={MemberCard} />
        <PrivateRoute path="/messages" component={Messages} />
        <PrivateRoute path="/" exact component={Dashboard} />
      </Switch>
    </AppContainer>
  );
};

export default Root;