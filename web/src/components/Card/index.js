import React from 'react';
import {styled} from 'fusion-plugin-styletron-react';

const CardContainer = styled('div', ({$clickable, $color}) => ({
  backgroundColor: '#fff',
  borderRadius: '8px',
  padding: '24px',
  boxShadow: '0 2px 4px rgba(0,0,0,0.08)',
  transition: 'all 0.2s',
  cursor: $clickable ? 'pointer' : 'default',
  borderLeft: `4px solid ${$color || '#276ef1'}`,
  ':hover': $clickable ? {
    boxShadow: '0 4px 12px rgba(0,0,0,0.12)',
    transform: 'translateY(-2px)',
  } : {},
}));

const CardTitle = styled('h3', {
  fontSize: '14px',
  fontWeight: '500',
  color: '#666',
  marginBottom: '8px',
  textTransform: 'uppercase',
  letterSpacing: '0.5px',
});

const CardValue = styled('div', {
  fontSize: '24px',
  fontWeight: '600',
  color: '#333',
  marginBottom: '4px',
});

const CardSubtitle = styled('div', {
  fontSize: '14px',
  color: '#999',
});

const Card = ({ title, value, subtitle, onClick, color, children }) => {
  return (
    <CardContainer
      $clickable={!!onClick}
      $color={color}
      onClick={onClick}
    >
      {title && <CardTitle>{title}</CardTitle>}
      {value && <CardValue>{value}</CardValue>}
      {subtitle && <CardSubtitle>{subtitle}</CardSubtitle>}
      {children}
    </CardContainer>
  );
};

export default Card;