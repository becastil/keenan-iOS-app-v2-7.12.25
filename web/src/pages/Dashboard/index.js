import React from 'react';
import {styled} from 'fusion-plugin-styletron-react';
import {useHistory} from 'fusion-plugin-react-router';
import Card from '../../components/Card';
import DeductibleTracker from '../../components/DeductibleTracker';
import RecentClaims from '../../components/RecentClaims';
import QuickActions from '../../components/QuickActions';

const PageTitle = styled('h1', {
  fontSize: '32px',
  fontWeight: '600',
  color: '#333',
  marginBottom: '8px',
});

const Subtitle = styled('p', {
  fontSize: '16px',
  color: '#666',
  marginBottom: '32px',
});

const Grid = styled('div', {
  display: 'grid',
  gridTemplateColumns: 'repeat(auto-fit, minmax(300px, 1fr))',
  gap: '20px',
  marginBottom: '32px',
});

const Row = styled('div', {
  display: 'grid',
  gridTemplateColumns: '2fr 1fr',
  gap: '20px',
  marginBottom: '32px',
});

const Dashboard = () => {
  const history = useHistory();
  
  // Mock data
  const coverageSummary = {
    medical: { status: 'Active', type: 'PPO' },
    dental: { status: 'Active', type: 'Basic' },
    vision: { status: 'Active', type: 'Standard' },
    pharmacy: { status: 'Active', type: 'Formulary' },
  };

  return (
    <>
      <PageTitle>Welcome back, John</PageTitle>
      <Subtitle>Here's your health plan overview</Subtitle>
      
      <Grid>
        <Card
          title="Medical Coverage"
          value={coverageSummary.medical.status}
          subtitle={`${coverageSummary.medical.type} Plan`}
          onClick={() => history.push('/benefits')}
          color="#276ef1"
        />
        <Card
          title="Dental Coverage"
          value={coverageSummary.dental.status}
          subtitle={`${coverageSummary.dental.type} Plan`}
          onClick={() => history.push('/benefits')}
          color="#00a862"
        />
        <Card
          title="Vision Coverage"
          value={coverageSummary.vision.status}
          subtitle={`${coverageSummary.vision.type} Plan`}
          onClick={() => history.push('/benefits')}
          color="#8b572a"
        />
        <Card
          title="Pharmacy Coverage"
          value={coverageSummary.pharmacy.status}
          subtitle={`${coverageSummary.pharmacy.type} Plan`}
          onClick={() => history.push('/benefits')}
          color="#e54b4b"
        />
      </Grid>
      
      <Row>
        <DeductibleTracker />
        <QuickActions />
      </Row>
      
      <RecentClaims />
    </>
  );
};

export default Dashboard;