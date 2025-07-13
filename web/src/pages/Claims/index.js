import React from 'react';
import {styled} from 'fusion-plugin-styletron-react';

const PageTitle = styled('h1', {
  fontSize: '32px',
  fontWeight: '600',
  color: '#333',
  marginBottom: '32px',
});

const Claims = () => {
  return (
    <>
      <PageTitle>Claims</PageTitle>
      <p>Claims management page - Coming soon</p>
    </>
  );
};

export default Claims;