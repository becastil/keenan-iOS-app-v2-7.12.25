import React from 'react';
import {styled} from 'fusion-plugin-styletron-react';
import Navigation from '../Navigation';
import Header from '../Header';

const LayoutContainer = styled('div', {
  display: 'flex',
  minHeight: '100vh',
});

const Sidebar = styled('aside', {
  width: '240px',
  backgroundColor: '#000',
  color: '#fff',
  position: 'fixed',
  height: '100vh',
  overflowY: 'auto',
});

const MainContent = styled('main', {
  flex: 1,
  marginLeft: '240px',
  backgroundColor: '#f5f5f5',
  minHeight: '100vh',
});

const ContentWrapper = styled('div', {
  padding: '80px 20px 20px',
  maxWidth: '1200px',
  margin: '0 auto',
});

const Layout = ({children}) => {
  return (
    <LayoutContainer>
      <Sidebar>
        <Navigation />
      </Sidebar>
      <MainContent>
        <Header />
        <ContentWrapper>
          {children}
        </ContentWrapper>
      </MainContent>
    </LayoutContainer>
  );
};

export default Layout;