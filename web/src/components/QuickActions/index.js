import React from 'react';
import {styled} from 'fusion-plugin-styletron-react';
import {useHistory} from 'fusion-plugin-react-router';
import Card from '../Card';

const ActionsContainer = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  gap: '12px',
});

const ActionButton = styled('button', {
  display: 'flex',
  alignItems: 'center',
  width: '100%',
  padding: '12px 16px',
  backgroundColor: '#f8f8f8',
  border: '1px solid #e0e0e0',
  borderRadius: '6px',
  fontSize: '14px',
  fontWeight: '500',
  color: '#333',
  cursor: 'pointer',
  transition: 'all 0.2s',
  textAlign: 'left',
  gap: '12px',
  ':hover': {
    backgroundColor: '#f0f0f0',
    borderColor: '#ccc',
  },
});

const ActionIcon = styled('span', {
  fontSize: '20px',
  width: '24px',
  textAlign: 'center',
});

const ActionText = styled('span', {
  flex: 1,
});

const Arrow = styled('span', {
  color: '#999',
});

const QuickActions = () => {
  const history = useHistory();

  const actions = [
    {
      icon: '🔍',
      text: 'Find a Doctor',
      path: '/providers',
    },
    {
      icon: '💳',
      text: 'View Member ID',
      path: '/member-card',
    },
    {
      icon: '📄',
      text: 'Submit a Claim',
      path: '/claims',
    },
    {
      icon: '💬',
      text: 'Message Support',
      path: '/messages',
    },
    {
      icon: '💊',
      text: 'Prescription Benefits',
      path: '/benefits',
    },
  ];

  return (
    <Card>
      <h3 style={{ margin: '0 0 16px 0', fontSize: '18px', fontWeight: '600' }}>
        Quick Actions
      </h3>
      <ActionsContainer>
        {actions.map((action) => (
          <ActionButton
            key={action.path}
            onClick={() => history.push(action.path)}
          >
            <ActionIcon>{action.icon}</ActionIcon>
            <ActionText>{action.text}</ActionText>
            <Arrow>→</Arrow>
          </ActionButton>
        ))}
      </ActionsContainer>
    </Card>
  );
};

export default QuickActions;