import React from 'react';
import {styled} from 'fusion-plugin-styletron-react';
import {useHistory} from 'fusion-plugin-react-router';

const HeaderContainer = styled('header', {
  position: 'fixed',
  top: 0,
  right: 0,
  left: '240px',
  height: '60px',
  backgroundColor: '#fff',
  borderBottom: '1px solid #e0e0e0',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'space-between',
  padding: '0 20px',
  zIndex: 100,
  boxShadow: '0 2px 4px rgba(0,0,0,0.08)',
});

const UserSection = styled('div', {
  display: 'flex',
  alignItems: 'center',
  gap: '20px',
});

const UserInfo = styled('div', {
  textAlign: 'right',
});

const UserName = styled('div', {
  fontWeight: '500',
  fontSize: '14px',
  color: '#333',
});

const MemberId = styled('div', {
  fontSize: '12px',
  color: '#666',
});

const LogoutButton = styled('button', {
  backgroundColor: 'transparent',
  border: '1px solid #ddd',
  borderRadius: '4px',
  padding: '8px 16px',
  fontSize: '14px',
  cursor: 'pointer',
  transition: 'all 0.2s',
  ':hover': {
    backgroundColor: '#f5f5f5',
    borderColor: '#999',
  },
});

const SearchBar = styled('div', {
  flex: '0 0 400px',
  position: 'relative',
});

const SearchInput = styled('input', {
  width: '100%',
  padding: '8px 16px 8px 40px',
  border: '1px solid #ddd',
  borderRadius: '20px',
  fontSize: '14px',
  outline: 'none',
  transition: 'border-color 0.2s',
  ':focus': {
    borderColor: '#276ef1',
  },
});

const SearchIcon = styled('span', {
  position: 'absolute',
  left: '12px',
  top: '50%',
  transform: 'translateY(-50%)',
  color: '#999',
});

const Header = () => {
  const history = useHistory();
  
  const handleLogout = () => {
    // Clear auth token
    localStorage.removeItem('authToken');
    history.push('/login');
  };

  // Mock user data
  const user = {
    name: 'John Doe',
    memberId: 'M123456',
  };

  return (
    <HeaderContainer>
      <SearchBar>
        <SearchIcon>🔍</SearchIcon>
        <SearchInput
          type="text"
          placeholder="Search benefits, claims, providers..."
        />
      </SearchBar>
      
      <UserSection>
        <UserInfo>
          <UserName>{user.name}</UserName>
          <MemberId>ID: {user.memberId}</MemberId>
        </UserInfo>
        <LogoutButton onClick={handleLogout}>
          Sign Out
        </LogoutButton>
      </UserSection>
    </HeaderContainer>
  );
};

export default Header;