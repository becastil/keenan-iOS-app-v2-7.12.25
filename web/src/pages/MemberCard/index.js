import React from 'react';
import {styled} from 'fusion-plugin-styletron-react';

const PageTitle = styled('h1', {
  fontSize: '32px',
  fontWeight: '600',
  color: '#333',
  marginBottom: '32px',
});

const MemberCard = () => {
  return (
    <>
      <PageTitle>Member ID Card</PageTitle>
      <p>Digital member ID card - Coming soon</p>
    </>
  );
};

export default MemberCard;