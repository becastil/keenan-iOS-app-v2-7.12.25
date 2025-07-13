import React from 'react';
import {styled} from 'fusion-plugin-styletron-react';
import {useHistory} from 'fusion-plugin-react-router';
import Card from '../Card';

const ClaimsTable = styled('table', {
  width: '100%',
  borderCollapse: 'collapse',
});

const TableHeader = styled('thead', {
  borderBottom: '2px solid #f0f0f0',
});

const TableRow = styled('tr', ({$clickable}) => ({
  cursor: $clickable ? 'pointer' : 'default',
  transition: 'background-color 0.2s',
  ':hover': $clickable ? {
    backgroundColor: '#f8f8f8',
  } : {},
}));

const TableHeaderCell = styled('th', {
  padding: '12px',
  textAlign: 'left',
  fontSize: '12px',
  fontWeight: '600',
  color: '#666',
  textTransform: 'uppercase',
  letterSpacing: '0.5px',
});

const TableCell = styled('td', {
  padding: '16px 12px',
  fontSize: '14px',
  color: '#333',
  borderBottom: '1px solid #f0f0f0',
});

const StatusBadge = styled('span', ({$status}) => ({
  display: 'inline-block',
  padding: '4px 12px',
  borderRadius: '12px',
  fontSize: '12px',
  fontWeight: '500',
  backgroundColor: 
    $status === 'Paid' ? '#e6f4ea' :
    $status === 'Pending' ? '#fef3e2' :
    $status === 'Denied' ? '#fce8e6' : '#f0f0f0',
  color:
    $status === 'Paid' ? '#1e7e34' :
    $status === 'Pending' ? '#f9a825' :
    $status === 'Denied' ? '#d32f2f' : '#666',
}));

const ViewAllLink = styled('div', {
  marginTop: '16px',
  textAlign: 'center',
});

const Link = styled('a', {
  color: '#276ef1',
  textDecoration: 'none',
  fontSize: '14px',
  fontWeight: '500',
  cursor: 'pointer',
  ':hover': {
    textDecoration: 'underline',
  },
});

const RecentClaims = () => {
  const history = useHistory();

  // Mock recent claims data
  const recentClaims = [
    {
      id: 'CLM001234',
      date: '2024-01-15',
      provider: 'Bay Area Medical Center',
      service: 'Office Visit',
      amount: 250.00,
      yourCost: 25.00,
      status: 'Paid',
    },
    {
      id: 'CLM001235',
      date: '2024-01-10',
      provider: 'City Diagnostics Lab',
      service: 'Lab Work',
      amount: 450.00,
      yourCost: 0.00,
      status: 'Paid',
    },
    {
      id: 'CLM001236',
      date: '2024-01-08',
      provider: 'Dr. Sarah Johnson',
      service: 'Specialist Consultation',
      amount: 350.00,
      yourCost: 40.00,
      status: 'Pending',
    },
    {
      id: 'CLM001237',
      date: '2024-01-02',
      provider: 'Regional Pharmacy',
      service: 'Prescription',
      amount: 120.00,
      yourCost: 20.00,
      status: 'Paid',
    },
  ];

  const formatCurrency = (amount) => {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: 'USD',
    }).format(amount);
  };

  const formatDate = (dateString) => {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric',
    });
  };

  return (
    <Card>
      <h3 style={{ margin: '0 0 20px 0', fontSize: '18px', fontWeight: '600' }}>
        Recent Claims
      </h3>
      <ClaimsTable>
        <TableHeader>
          <TableRow>
            <TableHeaderCell>Date</TableHeaderCell>
            <TableHeaderCell>Provider</TableHeaderCell>
            <TableHeaderCell>Service</TableHeaderCell>
            <TableHeaderCell>Total</TableHeaderCell>
            <TableHeaderCell>Your Cost</TableHeaderCell>
            <TableHeaderCell>Status</TableHeaderCell>
          </TableRow>
        </TableHeader>
        <tbody>
          {recentClaims.map((claim) => (
            <TableRow
              key={claim.id}
              $clickable
              onClick={() => history.push(`/claims?id=${claim.id}`)}
            >
              <TableCell>{formatDate(claim.date)}</TableCell>
              <TableCell>{claim.provider}</TableCell>
              <TableCell>{claim.service}</TableCell>
              <TableCell>{formatCurrency(claim.amount)}</TableCell>
              <TableCell>{formatCurrency(claim.yourCost)}</TableCell>
              <TableCell>
                <StatusBadge $status={claim.status}>
                  {claim.status}
                </StatusBadge>
              </TableCell>
            </TableRow>
          ))}
        </tbody>
      </ClaimsTable>
      <ViewAllLink>
        <Link onClick={() => history.push('/claims')}>
          View All Claims →
        </Link>
      </ViewAllLink>
    </Card>
  );
};

export default RecentClaims;