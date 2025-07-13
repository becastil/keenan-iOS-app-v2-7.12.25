import React from 'react';
import {Route, Redirect} from 'fusion-plugin-react-router';
import Layout from '../Layout';

const PrivateRoute = ({ component: Component, ...rest }) => {
  // Check if user is authenticated
  const isAuthenticated = localStorage.getItem('authToken');
  
  return (
    <Route
      {...rest}
      render={(props) =>
        isAuthenticated ? (
          <Layout>
            <Component {...props} />
          </Layout>
        ) : (
          <Redirect to="/login" />
        )
      }
    />
  );
};

export default PrivateRoute;