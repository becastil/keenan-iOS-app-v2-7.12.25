import React, {useState} from 'react';
import {styled} from 'fusion-plugin-styletron-react';
import Card from '../../components/Card';

const PageTitle = styled('h1', {
  fontSize: '32px',
  fontWeight: '600',
  color: '#333',
  marginBottom: '8px',
});

const TabContainer = styled('div', {
  display: 'flex',
  gap: '4px',
  marginBottom: '32px',
  borderBottom: '2px solid #e0e0e0',
});

const Tab = styled('button', ({$active}) => ({
  padding: '12px 24px',
  backgroundColor: 'transparent',
  border: 'none',
  borderBottom: $active ? '2px solid #276ef1' : '2px solid transparent',
  fontSize: '16px',
  fontWeight: $active ? '600' : '400',
  color: $active ? '#276ef1' : '#666',
  cursor: 'pointer',
  transition: 'all 0.2s',
  marginBottom: '-2px',
  ':hover': {
    color: '#276ef1',
  },
}));

const BenefitsList = styled('div', {
  display: 'grid',
  gap: '16px',
});

const BenefitItem = styled('div', {
  backgroundColor: '#fff',
  borderRadius: '8px',
  padding: '20px',
  boxShadow: '0 2px 4px rgba(0,0,0,0.08)',
  display: 'grid',
  gridTemplateColumns: '1fr auto',
  gap: '16px',
  alignItems: 'start',
});

const BenefitDetails = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  gap: '8px',
});

const BenefitName = styled('h3', {
  fontSize: '18px',
  fontWeight: '600',
  color: '#333',
  margin: 0,
});

const BenefitDescription = styled('p', {
  fontSize: '14px',
  color: '#666',
  margin: 0,
});

const CoverageDetails = styled('div', {
  display: 'flex',
  gap: '24px',
  marginTop: '12px',
});

const CoverageItem = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  gap: '4px',
});

const CoverageLabel = styled('span', {
  fontSize: '12px',
  color: '#999',
  textTransform: 'uppercase',
});

const CoverageValue = styled('span', {
  fontSize: '16px',
  fontWeight: '600',
  color: '#333',
});

const Benefits = () => {
  const [activeTab, setActiveTab] = useState('medical');

  const benefitsData = {
    medical: [
      {
        name: 'Primary Care Visit',
        description: 'Regular check-ups and preventive care with your primary care physician',
        inNetwork: { copay: '$20', coinsurance: '0%' },
        outOfNetwork: { copay: 'N/A', coinsurance: '40%' },
      },
      {
        name: 'Specialist Visit',
        description: 'Consultations with medical specialists',
        inNetwork: { copay: '$40', coinsurance: '0%' },
        outOfNetwork: { copay: 'N/A', coinsurance: '40%' },
      },
      {
        name: 'Emergency Room',
        description: 'Emergency medical services',
        inNetwork: { copay: '$150', coinsurance: '20%' },
        outOfNetwork: { copay: '$150', coinsurance: '20%' },
      },
      {
        name: 'Preventive Care',
        description: 'Annual wellness visits, immunizations, and screenings',
        inNetwork: { copay: '$0', coinsurance: '0%' },
        outOfNetwork: { copay: 'N/A', coinsurance: '40%' },
      },
    ],
    dental: [
      {
        name: 'Preventive Dental',
        description: 'Cleanings, exams, and X-rays (2 per year)',
        inNetwork: { copay: '$0', coinsurance: '0%' },
        outOfNetwork: { copay: 'N/A', coinsurance: '20%' },
      },
      {
        name: 'Basic Procedures',
        description: 'Fillings, extractions, and root canals',
        inNetwork: { copay: 'N/A', coinsurance: '20%' },
        outOfNetwork: { copay: 'N/A', coinsurance: '50%' },
      },
      {
        name: 'Major Procedures',
        description: 'Crowns, bridges, and dentures',
        inNetwork: { copay: 'N/A', coinsurance: '50%' },
        outOfNetwork: { copay: 'N/A', coinsurance: '70%' },
      },
    ],
    vision: [
      {
        name: 'Eye Exam',
        description: 'Comprehensive eye examination (1 per year)',
        inNetwork: { copay: '$10', coinsurance: '0%' },
        outOfNetwork: { copay: 'N/A', coinsurance: 'Not covered' },
      },
      {
        name: 'Eyeglasses',
        description: 'Frames and lenses',
        inNetwork: { copay: '$25', allowance: '$150' },
        outOfNetwork: { copay: 'N/A', allowance: '$75' },
      },
      {
        name: 'Contact Lenses',
        description: 'Contact lenses in lieu of eyeglasses',
        inNetwork: { copay: '$25', allowance: '$150' },
        outOfNetwork: { copay: 'N/A', allowance: '$75' },
      },
    ],
    pharmacy: [
      {
        name: 'Generic Drugs',
        description: 'FDA-approved generic medications',
        retail: { copay: '$10' },
        mailOrder: { copay: '$20 (90-day supply)' },
      },
      {
        name: 'Preferred Brand',
        description: 'Brand-name drugs on the formulary',
        retail: { copay: '$35' },
        mailOrder: { copay: '$70 (90-day supply)' },
      },
      {
        name: 'Non-Preferred Brand',
        description: 'Brand-name drugs not on the formulary',
        retail: { copay: '$60' },
        mailOrder: { copay: '$120 (90-day supply)' },
      },
      {
        name: 'Specialty Drugs',
        description: 'Complex medications for chronic conditions',
        retail: { copay: '30% up to $250' },
        mailOrder: { copay: 'Not available' },
      },
    ],
  };

  const tabs = [
    { id: 'medical', label: 'Medical' },
    { id: 'dental', label: 'Dental' },
    { id: 'vision', label: 'Vision' },
    { id: 'pharmacy', label: 'Pharmacy' },
  ];

  const currentBenefits = benefitsData[activeTab];

  return (
    <>
      <PageTitle>Your Benefits</PageTitle>
      
      <TabContainer>
        {tabs.map((tab) => (
          <Tab
            key={tab.id}
            $active={activeTab === tab.id}
            onClick={() => setActiveTab(tab.id)}
          >
            {tab.label}
          </Tab>
        ))}
      </TabContainer>
      
      <BenefitsList>
        {currentBenefits.map((benefit, index) => (
          <BenefitItem key={index}>
            <BenefitDetails>
              <BenefitName>{benefit.name}</BenefitName>
              <BenefitDescription>{benefit.description}</BenefitDescription>
              
              <CoverageDetails>
                {benefit.inNetwork && (
                  <CoverageItem>
                    <CoverageLabel>In-Network</CoverageLabel>
                    {benefit.inNetwork.copay && (
                      <CoverageValue>{benefit.inNetwork.copay}</CoverageValue>
                    )}
                    {benefit.inNetwork.coinsurance && (
                      <CoverageValue>
                        {benefit.inNetwork.coinsurance} coinsurance
                      </CoverageValue>
                    )}
                    {benefit.inNetwork.allowance && (
                      <CoverageValue>
                        {benefit.inNetwork.allowance} allowance
                      </CoverageValue>
                    )}
                  </CoverageItem>
                )}
                
                {benefit.outOfNetwork && (
                  <CoverageItem>
                    <CoverageLabel>Out-of-Network</CoverageLabel>
                    {benefit.outOfNetwork.copay && (
                      <CoverageValue>{benefit.outOfNetwork.copay}</CoverageValue>
                    )}
                    {benefit.outOfNetwork.coinsurance && (
                      <CoverageValue>
                        {benefit.outOfNetwork.coinsurance}
                        {benefit.outOfNetwork.coinsurance !== 'Not covered' && ' coinsurance'}
                      </CoverageValue>
                    )}
                    {benefit.outOfNetwork.allowance && (
                      <CoverageValue>
                        {benefit.outOfNetwork.allowance} allowance
                      </CoverageValue>
                    )}
                  </CoverageItem>
                )}
                
                {benefit.retail && (
                  <CoverageItem>
                    <CoverageLabel>Retail (30-day)</CoverageLabel>
                    <CoverageValue>{benefit.retail.copay}</CoverageValue>
                  </CoverageItem>
                )}
                
                {benefit.mailOrder && (
                  <CoverageItem>
                    <CoverageLabel>Mail Order</CoverageLabel>
                    <CoverageValue>{benefit.mailOrder.copay}</CoverageValue>
                  </CoverageItem>
                )}
              </CoverageDetails>
            </BenefitDetails>
          </BenefitItem>
        ))}
      </BenefitsList>
    </>
  );
};

export default Benefits;