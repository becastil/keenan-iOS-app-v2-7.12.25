import React from 'react';
import {styled} from 'fusion-plugin-styletron-react';

const PageTitle = styled('h1', {
  fontSize: '32px',
  fontWeight: '600',
  color: '#333',
  marginBottom: '32px',
});

const Messages = () => {
  return (
    <>
      <PageTitle>Messages</PageTitle>
      <p>Secure messaging - Coming soon</p>
    </>
  );
};

export default Messages;