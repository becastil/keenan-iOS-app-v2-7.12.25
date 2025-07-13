import React from 'react';
import {NavLink} from 'fusion-plugin-react-router';
import {styled} from 'fusion-plugin-styletron-react';

const Nav = styled('nav', {
  padding: '20px 0',
});

const Logo = styled('div', {
  fontSize: '24px',
  fontWeight: 'bold',
  padding: '0 20px 30px',
  borderBottom: '1px solid #333',
  marginBottom: '20px',
});

const NavList = styled('ul', {
  listStyle: 'none',
  padding: 0,
  margin: 0,
});

const NavItem = styled('li', {
  margin: 0,
});

const StyledNavLink = styled(NavLink, {
  display: 'flex',
  alignItems: 'center',
  padding: '12px 20px',
  color: '#ccc',
  textDecoration: 'none',
  transition: 'all 0.2s',
  ':hover': {
    backgroundColor: '#111',
    color: '#fff',
  },
});

const activeStyle = {
  backgroundColor: '#276ef1',
  color: '#fff',
};

const Icon = styled('span', {
  marginRight: '12px',
  fontSize: '20px',
});

const Navigation = () => {
  const navItems = [
    { path: '/', label: 'Dashboard', icon: '🏠' },
    { path: '/benefits', label: 'Benefits', icon: '📋' },
    { path: '/claims', label: 'Claims', icon: '📄' },
    { path: '/providers', label: 'Find Care', icon: '🏥' },
    { path: '/member-card', label: 'Member ID', icon: '💳' },
    { path: '/messages', label: 'Messages', icon: '💬' },
  ];

  return (
    <Nav>
      <Logo>Sydney Health</Logo>
      <NavList>
        {navItems.map((item) => (
          <NavItem key={item.path}>
            <StyledNavLink
              to={item.path}
              exact={item.path === '/'}
              activeStyle={activeStyle}
            >
              <Icon>{item.icon}</Icon>
              {item.label}
            </StyledNavLink>
          </NavItem>
        ))}
      </NavList>
    </Nav>
  );
};

export default Navigation;