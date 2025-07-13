import React from 'react';
import {styled} from 'fusion-plugin-styletron-react';
import Card from '../Card';

const TrackerContainer = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  gap: '16px',
});

const TrackerItem = styled('div', {
  marginBottom: '20px',
});

const TrackerHeader = styled('div', {
  display: 'flex',
  justifyContent: 'space-between',
  marginBottom: '8px',
});

const TrackerLabel = styled('div', {
  fontSize: '14px',
  color: '#666',
});

const TrackerValue = styled('div', {
  fontSize: '14px',
  fontWeight: '500',
  color: '#333',
});

const ProgressBar = styled('div', {
  height: '8px',
  backgroundColor: '#f0f0f0',
  borderRadius: '4px',
  overflow: 'hidden',
});

const ProgressFill = styled('div', ({$percentage, $color}) => ({
  height: '100%',
  backgroundColor: $color || '#276ef1',
  width: `${$percentage}%`,
  transition: 'width 0.3s ease',
}));

const DeductibleTracker = () => {
  // Mock deductible data
  const deductibles = {
    individual: {
      met: 750,
      total: 1500,
      percentage: 50,
    },
    family: {
      met: 1200,
      total: 3000,
      percentage: 40,
    },
    outOfPocket: {
      spent: 2500,
      max: 6000,
      percentage: 42,
    },
  };

  const formatCurrency = (amount) => {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: 'USD',
      minimumFractionDigits: 0,
    }).format(amount);
  };

  return (
    <Card>
      <TrackerContainer>
        <h3 style={{ margin: 0, fontSize: '18px', fontWeight: '600' }}>
          2024 Deductible & Out-of-Pocket
        </h3>
        
        <TrackerItem>
          <TrackerHeader>
            <TrackerLabel>Individual Deductible</TrackerLabel>
            <TrackerValue>
              {formatCurrency(deductibles.individual.met)} / {formatCurrency(deductibles.individual.total)}
            </TrackerValue>
          </TrackerHeader>
          <ProgressBar>
            <ProgressFill
              $percentage={deductibles.individual.percentage}
              $color="#276ef1"
            />
          </ProgressBar>
        </TrackerItem>
        
        <TrackerItem>
          <TrackerHeader>
            <TrackerLabel>Family Deductible</TrackerLabel>
            <TrackerValue>
              {formatCurrency(deductibles.family.met)} / {formatCurrency(deductibles.family.total)}
            </TrackerValue>
          </TrackerHeader>
          <ProgressBar>
            <ProgressFill
              $percentage={deductibles.family.percentage}
              $color="#00a862"
            />
          </ProgressBar>
        </TrackerItem>
        
        <TrackerItem>
          <TrackerHeader>
            <TrackerLabel>Out-of-Pocket Maximum</TrackerLabel>
            <TrackerValue>
              {formatCurrency(deductibles.outOfPocket.spent)} / {formatCurrency(deductibles.outOfPocket.max)}
            </TrackerValue>
          </TrackerHeader>
          <ProgressBar>
            <ProgressFill
              $percentage={deductibles.outOfPocket.percentage}
              $color="#e54b4b"
            />
          </ProgressBar>
        </TrackerItem>
      </TrackerContainer>
    </Card>
  );
};

export default DeductibleTracker;